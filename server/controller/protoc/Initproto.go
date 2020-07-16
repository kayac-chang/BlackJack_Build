package protoc

import "net/http"

// InitRequest ...
type InitRequest struct {
	Token    string
	PlayerID int64
}

// InitData ...
func (c *InitRequest) InitData(r *http.Request) {
	c.Token = r.Header.Get("Authorization")
}

// InitRespon ...
type InitRespon struct {
}

// InitData ...
func (c *InitRespon) InitData(r *http.Request) {
}
