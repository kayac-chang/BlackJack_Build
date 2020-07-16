package operate

import (
	"net/url"

	"github.com/google/go-querystring/query"
)

type AuthValidateCond struct {
	Token     string `json:"token"      url:"token"`
	GameID    string `json:"game_id"    url:"game_id"`
	GameToken string `json:"game_token" url:"game_token"`
	ConnID    string `json:"conn_id"    url:"conn_id"`
}

func (c *AuthValidateCond) Body() url.Values {
	v, _ := query.Values(c)
	return v
}

type AuthCond struct {
	Token  string `json:"token"   url:"token"   form:"token"   binding:"required"`
	GameID string `json:"game_id" url:"game_id" form:"game_id" binding:"required"`
	ConnID string `json:"conn_id" url:"conn_id"`
}

func (c *AuthCond) Body() url.Values {
	v, _ := query.Values(c)
	return v
}

type ExchangeCond struct {
	Token      string  `json:"token"       url:"token"        form:"token"       binding:"required"`
	GameToken  string  `json:"game_token"  url:"game_token"   form:"game_token"  binding:"required"`
	GameID     string  `json:"game_id"     url:"game_id"      form:"game_id"     binding:"required"`
	CoinType   int     `json:"coin_type"   url:"coin_type"    form:"coin_type"   binding:"required"`
	CoinAmount float64 `json:"coin_amount" url:"coin_amount"  form:"coin_amount" binding:"required"`
}

func (c *ExchangeCond) Body() url.Values {
	v, _ := query.Values(c)
	return v
}

type CheckoutCond struct {
	Token     string  `json:"token"       url:"token"       form:"token"       binding:"required"`
	GameToken string  `json:"game_token"  url:"game_token"  form:"game_token"  binding:"required"`
	GameID    string  `json:"game_id"     url:"game_id"     form:"game_id"     binding:"required"`
	Amount    float64 `json:"amount"      url:"amount"      form:"-"           binding:"-"`
	Win       float64 `json:"win"         url:"win"         form:"-"           binding:"-"`
	Lost      float64 `json:"lost"        url:"lost"        form:"-"           binding:"-"`
}

func (c *CheckoutCond) Body() url.Values {
	v, _ := query.Values(c)
	return v
}
