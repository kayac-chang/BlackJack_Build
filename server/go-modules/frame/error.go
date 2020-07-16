package frame

import "gitlab.fbk168.com/gamedevjp/blackjack/server/go-modules/frame/code"

// An Err implements the error interface.
type Err struct {
	c, s code.Code
	d    string
}

func (e *Err) Error() string {
	return e.d
}

// Frame returns a frame containing the error messages.
func (e *Err) Frame() Frame {
	return New(e.c, e.s, e.d)
}

// Error returns a pointer to the Err with given status and message.
func Error(command, status code.Code, message string) *Err {
	return &Err{command, status, message}
}
