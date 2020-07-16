package player

import "gitlab.fbk168.com/gamedevjp/blackjack/server/go-modules/frame"

// A Client represents an user client in a game.
type Client interface {

	// Close closes the underlying connection and resources.
	Close()

	// Receive returns the client's request. Bad requests and ping-pong
	// messages should not returned but handled by the handler.
	// It should check the client's balance before returning any request
	// involving bet.
	// If the client disconnected, it should return a request with
	// command code.Exited or close the channel.
	Receive() <-chan frame.Frame

	// Send returns a channel where the client can receive frames.
	Send() chan<- frame.Frame
}
