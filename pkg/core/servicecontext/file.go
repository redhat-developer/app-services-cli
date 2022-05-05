package servicecontext

import (
	"fmt"
	"strings"

	"github.com/redhat-developer/app-services-cli/pkg/core/configutil"
)

const errorFormat = "%v: %w"
const defaultFileName = "context.json"

// NewFile creates a new context type
func NewFile(id string) IContext {
	return &File{
		ConfigFile: configutil.NewConfigFile(
			id,
			defaultFileName,
			strings.ToUpper(id)+"_CONTEXT"),
	}
}

// File is a type which describes a context file
type File struct {
	configutil.ConfigFile
}

// Load loads the profiles from the context file. If the context file doesn't exist
// it will return an empty context object.
func (c *File) Load() (*Context, error) {
	ctx := &Context{}

	if err := c.ConfigFile.Load(ctx); err != nil {
		return ctx, fmt.Errorf(errorFormat, "unable to load context file", err)
	}

	return ctx, nil
}

// Save saves the given profiles to the context file.
func (c *File) Save(cfg *Context) error {
	if err := c.ConfigFile.Save(cfg); err != nil {
		return fmt.Errorf(errorFormat, "unable to save context file", err)
	}

	return nil
}

// Remove removes the context file.
func (c *File) Remove() error {
	if err := c.ConfigFile.Remove(); err != nil {
		return fmt.Errorf(errorFormat, "unable to remove context file", err)
	}

	return nil
}

// Returns the context file location.
func (c *File) Location() (string, error) {
	path, err := c.ConfigFile.Location()
	if err != nil {
		return "", fmt.Errorf(errorFormat, "unable to determine config file location", err)
	}

	return path, nil
}
