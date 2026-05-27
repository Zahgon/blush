package blush

import (
	"io"
)

type mode int

const (
	// Separator string between name of the reader and the contents.
	Separator = ": "

	// DefaultLineCache is minimum lines to cache.
	DefaultLineCache = 50

	// DefaultCharCache is minimum characters to cache for each line. This is in
	// effect only if Read() function is used.
	DefaultCharCache = 1000

	readMode mode = iota
	writeToMode
)

// Blush reads from reader and matches against all finders. If NoCut is true,
// any unmatched lines are printed as well. If WithFileName is true, blush will
// write the filename before it writes the output. Read and WriteTo will return
// ErrReadWriteMix if both Read and WriteTo are called on the same object. See
// package docs for more details.
// nolint:govet // we are expecting lots of these objects.
type Blush struct {
	Finders      []Finder
	Reader       io.ReadCloser
	LineCache    uint
	CharCache    uint
	Drop         bool // do not cut out non-matched lines.
	WithFileName bool
	closed       bool
	readLineCh   chan []byte
	readCh       chan byte
	mode         mode
}

// Read creates a goroutine on first invocation to read from the underlying
// reader. It is considerably slower than WriteTo as it reads the bytes one by
// one in order to produce the results, therefore you should use WriteTo
// directly or use io.Copy() on blush.
func (b *Blush) Read(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

// WriteTo writes matches to w. It returns an error if the writer is nil or
// there are not paths defined or there is no files found in the Reader.
func (b *Blush) WriteTo(w io.Writer) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (b *Blush) setup(m mode) error { _ = "STUB: not implemented"; return nil }

func (b *Blush) decorate(input string) (string, bool) { _ = "STUB: not implemented"; return "", false }

func (b *Blush) readLines() { _ = "STUB: not implemented"; return }

func (b *Blush) transfer() { _ = "STUB: not implemented"; return }

// Close closes the reader and returns whatever error it returns.
func (b *Blush) Close() error { _ = "STUB: not implemented"; return nil }

// lookInto returns a new decorated line if any of the finders decorate it, or
// the given line as it is.
func lookInto(f []Finder, line string) (string, bool) { _ = "STUB: not implemented"; return "", false }

// fileName returns an empty string if it could not query the fileName from r.
func fileName(r io.Reader) string { _ = "STUB: not implemented"; return "" }
