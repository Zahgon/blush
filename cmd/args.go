package cmd

import (
	"github.com/arsham/blush/blush"
)

// Note that hasArgs, setFinders and setPaths methods of args are designed to
// shrink the input as they go. Therefore the order of calls matters in some
// cases.
type args struct {
	paths       []string
	matches     []string
	remaining   []string
	finders     []blush.Finder
	cut         bool
	noFilename  bool
	recursive   bool
	insensitive bool
	stdin       bool
}

// nolint:misspell // it's ok.
func newArgs(input ...string) (*args, error) { _ = "STUB: not implemented"; return nil, nil }

// hasArgs removes any occurring `args` argument.
func (a *args) hasArgs(args ...string) (found bool) { _ = "STUB: not implemented"; return false }

// setPaths starts from the end of the slice and removes any paths/globs/files
// it finds and put them in the paths property.
func (a *args) setPaths() error { _ = "STUB: not implemented"; return nil }

// going backwards from the end.

// I don't like this label, but if we replace the `switch` statement with a
// regular if-then-else clause, it gets ugly and doesn't show its
// intentions. Order of cases in this switch matters.

// In this case, the previous input was a flag argument, therefore
// it might have been a colouring command. That is why we are
// ignoring this item.

// there is already a pattern found so we stop here.

// to return back in the same order.

// to keep the original user's preference.

func (a *args) setFinders() { _ = "STUB: not implemented"; return }

func flip(s []string) []string { _ = "STUB: not implemented"; return nil }

func inStringSlice(s string, haystack []string) bool { _ = "STUB: not implemented"; return false }
