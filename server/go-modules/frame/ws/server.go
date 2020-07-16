package ws

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
	"gitlab.fbk168.com/gamedevjp/blackjack/server/go-modules/frame"
	"gitlab.fbk168.com/gamedevjp/blackjack/server/go-modules/frame/code"
)

type server struct {
	mu      sync.RWMutex
	enc     Encoding
	in, out map[code.Code]Handler

	upgrader websocket.Upgrader

	// notFound handles frame with unregistered command.
	notFound Handler

	// binary handles binary messages.
	binary Handler

	onWriteError func(*Context, error)

	// onClosed is called once the connection is closed.
	onClosed func(*Context)
}

func (s *server) onEmit(c code.Code, h Handler) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exist := s.out[c]; exist {
		panic(fmt.Sprint("ws: multiple registrations of command ", c))
	}

	s.out[c] = h
}

func (s *server) onMessage(c code.Code, h Handler) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exist := s.in[c]; exist {
		panic(fmt.Sprint("ws: multiple registrations of command ", c))
	}

	s.in[c] = h
}

func (s *server) onBinary(h Handler) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.binary = h
}

func (s *server) handleClosed(h func(*Context)) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.onClosed = h
}

func (s *server) handleNotFound(h Handler) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.notFound = h
}

func (s *server) listen(conn *websocket.Conn, c *Context) {
	for {
		i, r, err := conn.NextReader()
		if err != nil {
			if s.onClosed != nil {
				s.onClosed(c)
			}
			return
		}

		// from gorilla docs, i is either TextMessage or BinaryMessage
		if i == websocket.BinaryMessage {
			if s.binary != nil {
				s.binary(frame.New(0, code.OK, r), c)
			}
			continue
		}

		f := frame.Prepare(&json.RawMessage{})
		if err = json.NewDecoder(s.enc.NewReader(r)).Decode(&f); err != nil {
			// handle err
			// ...
			continue
		}

		if h, ok := s.in[f.Command]; ok {
			h(f, c)
		} else {
			s.notFound(f, c)
		}
	}
}

func (s *server) write(conn *websocket.Conn, f frame.Frame) error {
	w, err := conn.NextWriter(websocket.TextMessage)
	if err != nil {
		return err
	}
	defer w.Close()

	w = s.enc.NewWriter(w)
	defer w.Close()

	return json.NewEncoder(w).Encode(f)
}

func (s *server) serve(conn *websocket.Conn, c *Context) {
	for {
		select {
		case <-c.ctx.Done():
			return
		case f := <-c.w:
			if h, ok := s.out[f.Command]; ok {
				h(f, c)
			}
			if err := s.write(conn, f); err != nil {
				s.onWriteError(c, err)
			}
		}
	}
}

// listenAndServe handles the given connection.
func (s *server) listenAndServe(ctx context.Context, conn *websocket.Conn) {
	if conn == nil {
		panic("ws: nil connection applied")
	}
	defer conn.Close()

	c := newContext(context.WithCancel(ctx))
	go s.listen(conn, c)

	s.serve(conn, c)
}
