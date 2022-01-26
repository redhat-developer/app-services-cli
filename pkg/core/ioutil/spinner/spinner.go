package spinner

import (
	"fmt"
	"io"
	"time"

	"github.com/redhat-developer/app-services-cli/pkg/core/localize"

	"github.com/briandowns/spinner"
)

type Spinner struct {
	writer    io.Writer
	spinner   *spinner.Spinner
	localizer localize.Localizer
}

// New create a new Spinner instance
func New(w io.Writer, localizer localize.Localizer) *Spinner {
	return &Spinner{
		writer:    w,
		localizer: localizer,
		spinner: spinner.New(
			spinner.CharSets[11],
			100*time.Millisecond,
			spinner.WithWriter(w),
			spinner.WithColor("cyan"),
		),
	}
}

// SetSuffix sets the spinner suffix message
func (s *Spinner) SetSuffix(suffix string) {
	s.spinner.Suffix = " " + suffix
}

// SetLocalizedSuffix sets the spinner suffix from an i18n ID
func (s *Spinner) SetLocalizedSuffix(id string, entries ...*localize.TemplateEntry) {
	s.SetSuffix(s.localizer.MustLocalize(id, entries...))
}

// Start starts the spinner
func (s *Spinner) Start() {
	s.spinner.Start()
}

// Stop stops the spinner
func (s *Spinner) Stop() {
	s.spinner.Stop()
	fmt.Fprintln(s.writer)
}
