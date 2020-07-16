package ws

import "net/http"

// An Option configures a WS.
type Option interface {
	apply(*WS)
}

type ignoreOrigin struct{}

func (o ignoreOrigin) apply(s *WS) {
	s.server.upgrader.CheckOrigin = func(_ *http.Request) bool { return true }
}

// IgnoreOrigin accepts the WebSocket connection without checking the origin.
func IgnoreOrigin() Option {
	return ignoreOrigin{}
}

type onWriteError struct {
	f func(*Context, error)
}

func (o onWriteError) apply(s *WS) {
	s.server.onWriteError = o.f
}

// OnWriteError returns an Option with the given function executed if error
// occurred when writing data to the client.
func OnWriteError(f func(*Context, error)) Option {
	if f == nil {
		panic("ws: nil function applied")
	}
	return onWriteError{f}
}
