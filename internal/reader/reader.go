package reader

import (
	"io"

	"github.com/pkg/errors"
)

// ErrNoReader is returned if there is no reader defined.
var ErrNoReader = errors.New("no input")

// MultiReader holds one or more io.ReadCloser and reads their contents when
// Read() method is called in order. The reader is loaded lazily if it is a
// file to prevent the system going out of file descriptors.
type MultiReader struct {
	currentName string
	readers     []*container
}

// NewMultiReader creates an instance of the MultiReader and passes it to all
// input functions.
func NewMultiReader(input ...Conf) (*MultiReader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Conf is used to configure the MultiReader.
type Conf func(*MultiReader) error

// WithReader adds the {name,r} reader to the MultiReader. If name is empty, the
// key will not be written in the output. You can provide as many empty names as
// you need.
func WithReader(name string, r io.ReadCloser) Conf { _ = "STUB: not implemented"; return *new(Conf) }

// WithPaths searches through the path and adds any files it finds to the
// MultiReader. Each path will become its reader's name in the process. It
// returns an error if any of given files are not found. It ignores any files
// that cannot be read or opened.
func WithPaths(paths []string, recursive bool) Conf { _ = "STUB: not implemented"; return *new(Conf) }

// nolint:gosec // we need this.

// Read is almost the exact implementation of io.MultiReader but keeps track of
// reader names. It closes each reader once they report they are exhausted, and
// it will happen on the next read.
func (m *MultiReader) Read(b []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

// Close does nothing.
func (m *MultiReader) Close() error {
	_ = "STUB: not implemented"

	// FileName returns the current reader's name.
	return nil
}

func (m *MultiReader) FileName() string { _ = "STUB: not implemented"; return "" }

// container takes care of opening the reader on demand. This is particularly
// useful when searching in thousands of files, because we want to open them on
// demand, otherwise the system gets out of file descriptors.
type container struct {
	r    io.ReadCloser
	get  func() (io.ReadCloser, error)
	open bool
}

func (c *container) Read(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }
