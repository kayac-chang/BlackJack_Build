package ws

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"

	"github.com/gorilla/websocket"
	"gitlab.fbk168.com/gamedevjp/blackjack/server/go-modules/frame"
	"gitlab.fbk168.com/gamedevjp/blackjack/server/go-modules/frame/code"
)

// isClosedError checks if the given error is due to close of connection.
func isClosedError(err error) bool {
	if err == nil {
		return false
	}

	_, o := err.(*net.OpError)
	_, c := err.(*websocket.CloseError)
	return o || c || err == websocket.ErrCloseSent
}

// UnmarshalData is a helper which parses the data field of the
// incoming frames.
func UnmarshalData(f frame.Frame, v interface{}) error {
	if b, ok := f.Data.(*json.RawMessage); ok {
		if len(*b) == 0 {
			return nil
		}
		return json.Unmarshal(*b, v)
	}
	return fmt.Errorf("ws: unsupport message type %T", f.Data)
}

// A WS is a multiplexer for customized websocket frames.
type WS struct {
	middlewares []Middleware
	server      *server
}

// OnEmit registers the handler for outgoing frames of
// the given command.
func (s *WS) OnEmit(command code.Code, handler Handler) *WS {
	s.server.onEmit(command, handler)
	return s
}

// Group returns a copy of s with its middlewares appended with the
// provided m.
func (s *WS) Group(m ...Middleware) *WS {
	return &WS{
		middlewares: append(m, s.middlewares...),
		server:      s.server,
	}
}

// OnMessage registers the handler for the given command.
// If a handler already exists for the command, it panics.
func (s *WS) OnMessage(command code.Code, handler Handler) *WS {
	if handler == nil {
		panic("ws: nil handler applied")
	}

	for _, v := range s.middlewares {
		handler = v(handler)
	}

	s.server.onMessage(command, handler)
	return s
}

// OnBinary registers the handler for binary messages. The registered
// middlewares are NOT chained with the given handler.
func (s *WS) OnBinary(handler Handler) *WS {
	s.server.onBinary(handler)
	return s
}

// OnClosed registers the handler for close event.
func (s *WS) OnClosed(handler func(*Context)) *WS {
	s.server.handleClosed(handler)
	return s
}

// NotFound registers the default handler for a frame with the
// unregistered command.
func (s *WS) NotFound(handler Handler) *WS {
	s.server.handleNotFound(handler)
	return s
}

// Use applies new middlewares to the WS.
func (s *WS) Use(m ...Middleware) *WS {
	s.middlewares = append(m, s.middlewares...)
	return s
}

// ListenAndServe handles the given connection.
func (s *WS) ListenAndServe(ctx context.Context, conn *websocket.Conn) {
	s.server.listenAndServe(ctx, conn)
}

// ServeHTTP implements the http.Handler interface.
func (s *WS) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := s.server.upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	s.ListenAndServe(r.Context(), conn)
}

// With applies the given options to the WS.
func (s *WS) With(opts ...Option) *WS {
	for _, o := range opts {
		o.apply(s)
	}
	return s
}

// New returns the pointer to a newly created WS with given encoding.
func New(enc Encoding) *WS {
	return &WS{
		middlewares: make([]Middleware, 0),
		server: &server{
			enc:          enc,
			in:           make(map[code.Code]Handler),
			out:          make(map[code.Code]Handler),
			onWriteError: func(_ *Context, _ error) {},
			notFound:     NotFoundHandler,
		},
	}
}
