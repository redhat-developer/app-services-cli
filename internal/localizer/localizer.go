package localizer

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"

	"github.com/markbates/pkger"
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

// IncludeAssetsAndLoadMessageFiles walks the /internal/locales directory
// and allows the static assets found to be embedded into the binary
// by github.com/markbates/pkger
// It also loads all files into memory
func IncludeAssetsAndLoadMessageFiles() error {
	localeFileName := fmt.Sprintf("active.%v", getLangFormat())
	return pkger.Walk("/locales", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() || localeFileName != info.Name() {
			return nil
		}

		err = loadMessageFile(path)
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
func loadMessageFile(path string) (err error) {
	// open the static i18n file
	f, err := pkger.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	b := bytes.NewBufferString("")
	// copy to contents of the file to a buffer string
	if _, err = io.Copy(b, f); err != nil {
		panic(err)
	}
	// read the contents of the file to a byte array
	out, _ := ioutil.ReadAll(b)
	// load the contents into context
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	_, err = bundle.ParseMessageFileBytes(out, "en.toml")
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
