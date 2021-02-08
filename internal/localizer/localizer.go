package localizer

import (
	"fmt"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var lang = language.English
var bundle *i18n.Bundle = i18n.NewBundle(lang)
var loc = i18n.NewLocalizer(bundle, lang.String())

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

// MustLocalise returns a localized a message
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

func LoadMessageFile(paths ...string) {
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.MustLoadMessageFile(fmt.Sprintf("./internal/localizer/locales/%v/active.%v.toml", strings.Join(paths, "/"), lang.String()))
}
