package cmd

import (
	"github.com/arsham/blush/blush"
)

// Main reads the provided arguments from the command line and creates a
// blush.Blush instance.
func Main() { _ = "STUB: not implemented"; return }

// this return statement should be here to support tests.

// GetBlush returns an error if no arguments are provided or it can't find all
// the passed files in the input.
//
// # Note
//
// The first argument will be dropped as it will be the application's name.
func GetBlush(input []string) (*blush.Blush, error) { _ = "STUB: not implemented"; return nil, nil }
