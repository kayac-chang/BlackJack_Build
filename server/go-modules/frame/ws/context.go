package ws

import (
	"context"
	"sync"

	uuid "github.com/satori/go.uuid"
	"gitlab.fbk168.com/gamedevjp/blackjack/server/go-modules/frame"
)

// A Context represents a WebSocket client.
type Context struct {
	uid   uuid.UUID
	mu    sync.RWMutex
	ctx   context.Context
	exit  context.CancelFunc
	state map[string]interface{}
	w     chan frame.Frame
}

// Exit closes the underlying connection.
func (c *Context) Exit() {
	c.exit()
}

// Exited returns if the connection has closed from the server side.
func (c *Context) Exited() bool {
	select {
	case <-c.ctx.Done():
		return true
	default:
		return false
	}
}

// UUID returns the context UUID.
func (c *Context) UUID() uuid.UUID {
	return c.uid
}

// Context returns the underlying context.Context, which retrieved from
// the http.Request.
func (c *Context) Context() context.Context {
	return c.ctx
}

// Get returns the value associated with the given key.
// If the key doesn't exist, it returns false.
func (c *Context) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	c.mu.RUnlock()

	v, ok := c.state[key]
	return v, ok
}

// Set replaces the value associated with the key, creating it
// if the key doesn't exist.
func (c *Context) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.state[key] = value
}

// Put replaces the value for the key if it exists. Else it returns false.
func (c *Context) Put(key string, value interface{}) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	_, ok := c.state[key]
	if ok {
		c.state[key] = value
	}
	return ok
}

// Del deletes the value associated with key.
func (c *Context) Del(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.state, key)
}

// Send sends the given frame to the associated client.
func (c *Context) Send(f frame.Frame) {
	select {
	case <-c.ctx.Done():
		// closed, do nothing
	case c.w <- f:
		// already sent, do nothing
	}
}

func newContext(ctx context.Context, exit context.CancelFunc) *Context {
	var err error

	return &Context{
		uid:   uuid.Must(uuid.NewV4(), err),
		ctx:   ctx,
		exit:  exit,
		state: make(map[string]interface{}),
		w:     make(chan frame.Frame, 8),
	}
}
