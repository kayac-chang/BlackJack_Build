package ws

import (
	"fmt"

	"gitlab.fbk168.com/gamedevjp/blackjack/server/go-modules/frame"
	"gitlab.fbk168.com/gamedevjp/blackjack/server/go-modules/frame/code"
)

// NotFoundHandler is the default handler for unknown commands.
var NotFoundHandler = func(f frame.Frame, c *Context) {
	c.Send(frame.New(f.Command, code.NotFound, fmt.Sprint("Unknown command ", f.Command)))
}

// A Handler handles a WebSocket frame.
type Handler func(frame.Frame, *Context)

// A Middleware is a function that chains handlers.
type Middleware func(Handler) Handler
