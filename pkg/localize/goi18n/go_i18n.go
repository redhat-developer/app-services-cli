// Internationalization implementation using nicksnyder/go-i18n
package goi18n

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"

	"github.com/redhat-developer/app-services-cli/pkg/dump"
	"github.com/redhat-developer/app-services-cli/pkg/localize"
	"gopkg.in/yaml.v2"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

// go-i18n implementation
type Goi18n struct {
	files     fs.FS
	language  *language.Tag
	bundle    *i18n.Bundle
	localizer *i18n.Localizer
	format    string
	path      string
}

type Config struct {
	files    fs.FS
	language *language.Tag
	format   string
	path     string
}

// New creates a new nicksnyder/go-i18n client.
// You can pass nil to use the pre-configured defaults
// Or pass a partial config to override some defaults.
// Default file path: locales
// Default file format: toml (yaml and json are supported)
// Default language: English
func New(cfg *Config) (localize.Localizer, error) {
	if cfg == nil {
		cfg = &Config{}
	}
	if cfg.files == nil {
		cfg.files = localize.GetDefaultLocales()
		cfg.path = "locales"
		cfg.format = "toml"
	}
	if cfg.language == nil {
		cfg.language = localize.GetDefaultLanguage()
	}
	if cfg.format == "" {
		cfg.format = "toml"
	}

	bundle := i18n.NewBundle(*cfg.language)
	loc := &Goi18n{
		files:     cfg.files,
		language:  cfg.language,
		bundle:    bundle,
		localizer: i18n.NewLocalizer(bundle),
		format:    cfg.format,
		path:      cfg.path,
	}

	err := loc.load()
	return loc, err
}

// MustLocalize loads a i18n message from the file system
// You can pass dynamic values using template entries
// Examples:
//
// `localizer.MustLocalize("my-message-id", &localize.TemplateEntry{"Name", "Danny"})`
//
func (l *Goi18n) MustLocalize(id string, tmplEntries ...*localize.TemplateEntry) string {
	templateData := map[string]interface{}{}
	// map the template entries to the go-i18n format
	for _, t := range tmplEntries {
		templateData[t.Key] = t.Value
	}
	cfg := &i18n.LocalizeConfig{MessageID: id, PluralCount: 1, TemplateData: templateData}
	return l.localizer.MustLocalize(cfg)
}

// MustLocalizePlural loads a pluralized i18n message from the file system
// You can pass dynamic values using template entries
// Examples:
//
// `localizer.MustLocalizePlural("my-message-id", 2)`
//
func (l *Goi18n) MustLocalizePlural(id string, pluralCount int, tmplEntries ...*localize.TemplateEntry) string {
	templateData := map[string]interface{}{}
	// map the template entries to the go-i18n format
	for _, t := range tmplEntries {
		templateData[t.Key] = t.Value
	}
	cfg := &i18n.LocalizeConfig{MessageID: id, PluralCount: pluralCount, TemplateData: templateData}
	return l.localizer.MustLocalize(cfg)
}

// MustLocalizeError loads a i18n message from the file system
// and returns an error
// Examples:
//
// `err := localizer.MustLocalizeError("my-message-id", &localize.TemplateEntry{"Name", "Danny"}`
//
func (l *Goi18n) MustLocalizeError(id string, tmplEntries ...*localize.TemplateEntry) error {
	return errors.New(l.MustLocalize(id, tmplEntries...))
}

// walk the file system and load each file into memory
func (l *Goi18n) load() error {
	return fs.WalkDir(l.files, l.path, func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		return l.MustLocalizeFile(l.files, path)
	})
}

// read the message file from the file system
func (l *Goi18n) MustLocalizeFile(files fs.FS, path string) (err error) {
	// open the static i18n file
	buf, err := fs.ReadFile(files, path)
	if err != nil {
		return err
	}
	fileext := fmt.Sprintf("%v.%v", l.language.String(), l.format)
	var unmarshalFunc i18n.UnmarshalFunc
	switch l.format {
	case "toml":
		unmarshalFunc = toml.Unmarshal
	case dump.YAMLFormat, dump.YMLFormat:
		unmarshalFunc = yaml.Unmarshal
	case dump.JSONFormat:
		unmarshalFunc = json.Unmarshal
	default:
		return fmt.Errorf("unsupported format \"%v\"", l.format)
	}

	l.bundle.RegisterUnmarshalFunc(l.format, unmarshalFunc)
	_, err = l.bundle.ParseMessageFileBytes(buf, fileext)

	return
}
