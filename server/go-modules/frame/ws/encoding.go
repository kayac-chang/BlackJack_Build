package ws

import (
	"encoding/base64"
	"io"
)

var (

	// Base64Encoding is an Encoding which wraps the given readers and writers
	// with additional Base64 encoder or decoder.
	Base64Encoding Encoding = new(base64Encoding)

	// StdEncoding is an Encoding which does nothing with the given readers
	// and writers.
	StdEncoding Encoding = new(stdEncoding)
)

// An Encoding is an interface that converts the stream data into the
// custom form.
type Encoding interface {
	NewReader(io.Reader) io.Reader
	NewWriter(io.WriteCloser) io.WriteCloser
}

// A base64Encoding implements the Encoding interface.
type base64Encoding struct{}

func (*base64Encoding) NewReader(r io.Reader) io.Reader {
	return base64.NewDecoder(base64.URLEncoding, r)
}

func (*base64Encoding) NewWriter(w io.WriteCloser) io.WriteCloser {
	return base64.NewEncoder(base64.URLEncoding, w)
}

// A stdEncoding implements the Encoding interface. It does nothing
// with given readers and writers.
type stdEncoding struct{}

func (*stdEncoding) NewReader(r io.Reader) io.Reader {
	return r
}

func (*stdEncoding) NewWriter(w io.WriteCloser) io.WriteCloser {
	return w
}
