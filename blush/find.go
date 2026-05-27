package blush

import (
	"regexp"
)

var (
	isRegExp = regexp.MustCompile(`[\^\$.\{\}\[\]\*\?]`)
	// This is used for matching colour groups (b1, etc.).
	grouping = regexp.MustCompile(`^([[:alpha:]]+)(\d+)$`)
)

// Finder finds texts based on a plain text or regexp logic. If it doesn't find
// any match, it will return an empty string. It might decorate the match with a
// given instruction.
type Finder interface {
	Find(string) (string, bool)
}

// NewLocator returns a Rx object if search is a valid regexp, otherwise it
// returns Exact or Iexact. If insensitive is true, the match will be case
// insensitive. The colour argument can be in short form (b) or long form
// (blue). If it cannot find the colour, it will fall-back to DefaultColour. The
// colour also can be in hex format, which should be started with a pound sign
// (#666).
func NewLocator(colour, search string, insensitive bool) Finder {
	_ = "STUB: not implemented"
	return *new(Finder)
}

// Exact looks for the exact word in the string.
type Exact struct {
	s      string
	colour Colour
}

// NewExact returns a new instance of the Exact.
func NewExact(s string, c Colour) Exact { _ = "STUB: not implemented"; return *new(Exact) }

// Find looks for the exact string. Any strings it finds will be decorated with
// the given Colour.
func (e Exact) Find(input string) (string, bool) { _ = "STUB: not implemented"; return "", false }

func (e Exact) colourise(input string, c Colour) string { _ = "STUB: not implemented"; return "" }

// Colour returns the Colour property.
func (e Exact) Colour() Colour {
	_ = "STUB: not implemented"

	// String will returned the colourised contents.
	return *new(Colour)
}

func (e Exact) String() string { _ = "STUB: not implemented"; return "" }

// Iexact is like Exact but case insensitive.
type Iexact struct {
	s      string
	colour Colour
}

// NewIexact returns a new instance of the Iexact.
func NewIexact(s string, c Colour) Iexact { _ = "STUB: not implemented"; return *new(Iexact) }

// Find looks for the exact string. Any strings it finds will be decorated with
// the given Colour.
func (i Iexact) Find(input string) (string, bool) { _ = "STUB: not implemented"; return "", false }

func (i Iexact) colourise(input string, c Colour) string { _ = "STUB: not implemented"; return "" }

// Colour returns the Colour property.
func (i Iexact) Colour() Colour {
	_ = "STUB: not implemented"

	// String will returned the colourised contents.
	return *new(Colour)
}

func (i Iexact) String() string { _ = "STUB: not implemented"; return "" }

// Rx is the regexp implementation of the Locator.
type Rx struct {
	*regexp.Regexp
	colour Colour
}

// NewRx returns a new instance of the Rx.
func NewRx(r *regexp.Regexp, c Colour) Rx { _ = "STUB: not implemented"; return *new(Rx) }

// Find looks for the string matching `r` regular expression. Any strings it
// finds will be decorated with the given Colour.
func (r Rx) Find(input string) (string, bool) { _ = "STUB: not implemented"; return "", false }

func (r Rx) colourise(input string, c Colour) string { _ = "STUB: not implemented"; return "" }

// Colour returns the Colour property.
func (r Rx) Colour() Colour { _ = "STUB: not implemented"; return *new(Colour) }
