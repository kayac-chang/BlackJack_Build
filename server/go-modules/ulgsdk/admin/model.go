package admin

import (
	"net/url"

	"github.com/google/go-querystring/query"
)

type adminValidateCond struct {
	Token string `json:"token"  url:"token"`
}

func (c *adminValidateCond) Body() url.Values {
	v, _ := query.Values(c)
	return v
}
