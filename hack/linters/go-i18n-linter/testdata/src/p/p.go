package p

import (
	"errors"
)

type Goi18n struct {
	format string
	path   string
}

type TemplateEntry struct {
	Key   string
	Value interface{}
}

type Localizer interface {
	MustLocalize(id string) string
	MustLocalizeError(id string) error
}

func (l *Goi18n) MustLocalize(id string) string {
	return id
}

func (l *Goi18n) MustLocalizeError(id string) error {
	return errors.New(id)
}

func testFunc() error {
	a := Goi18n{}
	a.path = a.MustLocalize("test.path.string")
	a.format = a.MustLocalize("another.message.exist")

	if 2+2 == 5 {
		return a.MustLocalizeError("test.return.error.message")
	} else {
		return a.MustLocalizeError("test.message.exist")
	}

	return nil
}
