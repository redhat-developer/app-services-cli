package locales

import (
	"embed"
	"fmt"
	"io/fs"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var (
	//go:embed files
	defaultLocales  embed.FS
	defaultLanguage = &language.English

	singleCount int = 1
)

type LocalizerConfig struct {
	PluralCount  int
	TemplateData map[string]interface{}
}

type Localizer interface {
	LoadMessage(id string) string
	LoadMessageWithConfig(id string, config *LocalizerConfig) string
}

type Goi18n struct {
	files     fs.FS
	language  *language.Tag
	bundle    *i18n.Bundle
	localizer *i18n.Localizer
}

type Config struct {
	files    fs.FS
	language *language.Tag
}

func New(cfg *Config) (Localizer, error) {
	if cfg == nil {
		cfg = &Config{}
	}
	if cfg.files == nil {
		cfg.files = defaultLocales
	}
	if cfg.language == nil {
		cfg.language = defaultLanguage
	}
	bundle := i18n.NewBundle(*cfg.language)
	loc := &Goi18n{
		files:     cfg.files,
		language:  cfg.language,
		bundle:    bundle,
		localizer: i18n.NewLocalizer(bundle),
	}

	err := loc.load()
	return loc, err
}

func (l *Goi18n) LoadMessage(id string) string {
	return l.localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: id, PluralCount: singleCount})
}

func (l *Goi18n) LoadMessageWithConfig(id string, cfg *LocalizerConfig) string {
	if cfg == nil {
		cfg = &LocalizerConfig{}
	}
	if cfg.PluralCount == 0 {
		cfg.PluralCount = 1
	}
	return l.localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID:    id,
		PluralCount:  cfg.PluralCount,
		TemplateData: cfg.TemplateData,
	})
}

// loadMessageFile loads the message file int context
// Using github.com/nicksnyder/go-i18n/v2/i18n
// pathTree to File is an array of the parent directories
func (l *Goi18n) loadMessageFile(files fs.FS, path string) (err error) {
	// open the static i18n file
	buf, err := fs.ReadFile(files, path)
	if err != nil {
		return err
	}
	l.bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	_, err = l.bundle.ParseMessageFileBytes(buf, "en.toml")
	if err != nil {
		return err
	}

	return nil
}

func (l *Goi18n) load() error {
	localeFileName := fmt.Sprintf("active.%v", getLangFormat(l.language, "toml"))
	return fs.WalkDir(l.files, ".", func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() || localeFileName != info.Name() {
			return nil
		}

		err = l.loadMessageFile(l.files, path)
		if err != nil {
			return err
		}

		return nil
	})
}

// get the file extension for the current language
// Example: "en.toml", "de.yaml"
func getLangFormat(lang *language.Tag, format string) string {
	return fmt.Sprintf("%v.%v", lang.String(), format)
}
