package cmd

import (
	"os"
)

// WaitForSignal calls exit with code 130 if receives an SIGINT or SIGTERM, 0 if
// SIGPIPE, and 1 otherwise.
func WaitForSignal(sig chan os.Signal, exit func(int)) { _ = "STUB: not implemented"; return }

// Ctrl+c
