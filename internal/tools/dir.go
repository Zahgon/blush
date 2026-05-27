// Package tools contains common tools used throughout this application.
package tools

// Files returns all files found in paths. If recursive is false, it only
// returns the immediate files in the paths.
func Files(recursive bool, paths ...string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func unique(fileList []string) []string { _ = "STUB: not implemented"; return nil }

func nonBinary(fileList []string) []string { _ = "STUB: not implemented"; return nil }

func rfiles(location string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func files(location string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

// TODO: we should ignore the line in search stage instead.
func isPlainText(name string) bool { _ = "STUB: not implemented"; return false }

// nolint:gosec // this is required.

// nolint:errcheck,gosec // not required.
