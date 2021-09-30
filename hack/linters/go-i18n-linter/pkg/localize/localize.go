package localize

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v2"
	"io/fs"
	"os"
	"path/filepath"
)

type localize struct {
	fsys         fs.FS
	language     *language.Tag
	bundle       *i18n.Bundle
	translations []i18n.MessageFile
	format       string
	path         string
}

var (
	defaultLanguage = &language.English
)

type config struct {
	fsys     fs.FS
	language *language.Tag
	format   string
	path     string
}

func getDefaultLanguage() *language.Tag {
	return defaultLanguage
}

// New loads localization files from a specified directory
func New(cfg *config, d string) (*localize, error) {
	if cfg == nil {
		cfg = &config{}
	}

	if cfg.fsys == nil {
		if d != "" {
			cfg.path = filepath.Join(d)
		} else {
			cfg.path = filepath.Join(d, "pkg", "localize", "locales")
		}
		cfg.fsys = os.DirFS(cfg.path)
		cfg.format = "toml"
	}
	if cfg.language == nil {
		cfg.language = getDefaultLanguage()
	}
	if cfg.format == "" {
		cfg.format = "toml"
	}

	bundle := i18n.NewBundle(*cfg.language)
	loc := &localize{
		fsys:     cfg.fsys,
		language: cfg.language,
		bundle:   bundle,
		format:   cfg.format,
		path:     cfg.path,
	}

	err := loc.load()
	return loc, err
}

func (l *localize) load() error {
	return fs.WalkDir(l.fsys, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		return l.loadLocalizationFile(l.fsys, path)
	})
}

func (l *localize) loadLocalizationFile(fsys fs.FS, path string) (err error) {
	buf, err := fs.ReadFile(fsys, path)
	if err != nil {
		return err
	}
	fileExt := fmt.Sprintf("%v.%v", l.language.String(), l.format)
	var unmarshalFunc i18n.UnmarshalFunc
	switch l.format {
	case "toml":
		unmarshalFunc = toml.Unmarshal
	case "yaml", "yml":
		unmarshalFunc = yaml.Unmarshal
	default:
		return fmt.Errorf("unsupported format \"%v\"", l.format)
	}

	l.bundle.RegisterUnmarshalFunc(l.format, unmarshalFunc)
	file, err := l.bundle.ParseMessageFileBytes(buf, fileExt)
	if err != nil {
		return err
	}

	l.translations = append(l.translations, *file)

	return nil
}

// GetTranslations returns an array with localization files
func (l *localize) GetTranslations() []i18n.MessageFile {
	return l.translations
}
