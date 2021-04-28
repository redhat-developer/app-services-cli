package localizer

import (
	"fmt"
	"io/fs"

	"github.com/BurntSushi/toml"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var lang = language.English
var bundle *i18n.Bundle = i18n.NewBundle(lang)
var loc = i18n.NewLocalizer(bundle, lang.String())

// file format for locale files
const format = "toml"

// Config is the basic configuration needed
// to localize a message
//
// Example usage:
//
// localizeConfig := &localizer.Config{
//	MessageID: flagi18n.InvalidValueError,
//	TemplateData: map[string]interface{}{
//		"Value": "xml",
//    "Flag":  "output",
// },
type Config struct {
	// The unique ID of the message
	MessageID string
	// Mapping of variables to their template names
	// eg:
	TemplateData interface{}
	// Indicate the number of values referenced
	// If > 1 the message will be pluralized
	PluralCount int
}

// IncludeAssetsAndLoadMessageFiles walks the locales directory
// and allows the static assets found to be embedded into the binary
// It also loads all files into memory
func IncludeAssetsAndLoadMessageFiles(embeddedFiles fs.FS) error {
	localeFileName := fmt.Sprintf("active.%v", getLangFormat())
	return fs.WalkDir(embeddedFiles, ".", func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() || localeFileName != info.Name() {
			return nil
		}

		err = loadMessageFile(embeddedFiles, path)
		if err != nil {
			return err
		}

		return nil
	})
}

// MustLocalise returns a localized a message,
// and panics if it was not found
func MustLocalize(config *Config) string {
	pluralCount := config.PluralCount
	if config.PluralCount == 0 {
		pluralCount = 1
	}

	return loc.MustLocalize(&i18n.LocalizeConfig{
		MessageID:    config.MessageID,
		PluralCount:  pluralCount,
		TemplateData: config.TemplateData,
	})
}

func MustLocalizeFromID(messageID string) string {
	return MustLocalize(&Config{
		MessageID: messageID,
	})
}

// loadMessageFile loads the message file int context
// Using github.com/nicksnyder/go-i18n/v2/i18n
// pathTree to File is an array of the parent directories
func loadMessageFile(files fs.FS, path string) (err error) {
	// open the static i18n file
	buf, err := fs.ReadFile(files, path)
	if err != nil {
		return err
	}
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	_, err = bundle.ParseMessageFileBytes(buf, "en.toml")
	if err != nil {
		return err
	}

	return nil
}

// get the file extension for the current language
// Example: "en.toml", "de.yaml"
func getLangFormat() string {
	return fmt.Sprintf("%v.%v", lang.String(), format)
}
