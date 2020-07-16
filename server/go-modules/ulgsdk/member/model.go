package member

import (
	"net/url"

	"github.com/google/go-querystring/query"
)

type WalletCond struct {
	Token     string `json:"token"      url:"token"`
	GameID    string `json:"game_id"    url:"game_id"`
	GameToken string `json:"game_token" url:"game_token"`
}

func (c *WalletCond) Body() url.Values {
	v, _ := query.Values(c)
	return v
}

type Wallet struct {
	Account string  `json:"account"`
	Balance float64 `json:"balance"`
}
