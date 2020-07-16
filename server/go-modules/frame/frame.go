package frame

import "gitlab.fbk168.com/gamedevjp/blackjack/server/go-modules/frame/code"

// A Frame defines the schema of a message interchanged between the
// server and the clients.
type Frame struct {
	Command code.Code   `json:"cmd"`
	Meta    MetaData    `json:"meta"`
	Data    interface{} `json:"data"`
	Success bool        `json:"success"`
	Status  code.Code   `json:"status,omitempty"`
	From    string      `json:"-"`
}

// New returns a newly created Frame with given data, which the
// field "Success" set to true if the status is code.OK.
func New(command, status code.Code, data interface{}) Frame {
	return Frame{
		Command: command,
		Meta:    make(MetaData),
		Data:    data,
		Status:  status,
		Success: status == code.OK,
	}
}

// Prepare creates a frame with given data for later uses. For
// example, unmarshaling.
func Prepare(v interface{}) Frame {
	return Frame{
		Meta: make(MetaData),
		Data: v,
	}
}
