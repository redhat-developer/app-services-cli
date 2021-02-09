// This file is a workaround to make pkger work with the CLI package structure
// See: https://github.com/markbates/pkger/issues/44#issuecomment-620118227
// File will never be included in the CLI binary

// +build !skippkger

package cli
