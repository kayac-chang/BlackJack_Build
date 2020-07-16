package order

import (
	"encoding/json"
	"net/url"
	"time"

	"github.com/google/go-querystring/query"
	"gitlab.fbk168.com/gamedevjp/blackjack/server/go-modules/ulgsdk/ulg168utils"
	"gitlab.fbk168.com/gamedevjp/blackjack/server/go-modules/utils/timetool"
)

type Basic struct {
	Token     string `json:"token"      url:"token"`
	GameToken string `json:"game_token" url:"game_token"`
	GameID    string `json:"game_id"    url:"game_id"`
}

type CreateOrder struct {
	Basic
	GameType   string             `json:"game_type"     url:"game_type"`
	GameID     string             `json:"game_id"       url:"game_id"`
	GameRoom   string             `json:"game_room"     url:"game_room"`
	GameRound  string             `json:"game_round"    url:"game_round"`
	Summary    string             `json:"summary"       url:"summary"`
	Amount     float64            `json:"amount"        url:"amount"`
	Status     int                `json:"status"        url:"status"`
	OrderItems []*CreateOrderItem `json:"order_items"   url:"order_items"`
	CreatedAt  time.Time          `json:"created_at"    url:"created_at"`
}

func NewCreateOrder(basic Basic, gameRound string, amount float64, orderItems []*CreateOrderItem) *CreateOrder {
	obj := &CreateOrder{}
	obj.Basic = basic
	obj.GameType = ulg168utils.Conf.GameType
	obj.GameID = ulg168utils.Conf.GameID
	obj.GameRound = gameRound
	obj.Status = 1
	obj.CreatedAt = timetool.GetNowByUTC()
	obj.Amount = amount
	obj.OrderItems = orderItems
	return obj
}

func NewCreateOrderV2(basic Basic, gameRound string, orderItems []*CreateOrderItem) *CreateOrder {
	obj := &CreateOrder{}
	obj.Basic = basic
	obj.GameType = ulg168utils.Conf.GameType
	obj.GameID = ulg168utils.Conf.GameID
	obj.GameRound = gameRound
	obj.Status = 1
	obj.CreatedAt = timetool.GetNowByUTC()

	var amount float64
	for _, item := range orderItems {
		amount += item.Amount
	}

	obj.Amount = amount
	obj.OrderItems = orderItems
	return obj
}

func NewCreateOrderV3(basic Basic, gameRoom, gameRound string, orderItems []*CreateOrderItem) *CreateOrder {
	obj := &CreateOrder{}
	obj.Basic = basic
	obj.GameType = ulg168utils.Conf.GameType
	obj.GameID = ulg168utils.Conf.GameID
	obj.GameRoom = gameRoom
	obj.GameRound = gameRound
	obj.Status = 1
	obj.CreatedAt = timetool.GetNowByUTC()

	var amount float64
	for _, item := range orderItems {
		amount += item.Amount
	}

	obj.Amount = amount
	obj.OrderItems = orderItems
	return obj
}

func (c *CreateOrder) Body() url.Values {
	v, _ := query.Values(c)
	v.Del("order_items")
	items, _ := json.Marshal(c.OrderItems)
	v.Add("order_items", string(items))
	return v
}

type CreateOrderItem struct {
	Basic
	PlayCode  string    `json:"play_code"     url:"play_code"`
	Summary   string    `json:"summary"       url:"summary"`
	Amount    float64   `json:"amount"        url:"amount"`
	Rate      float64   `json:"rate"          url:"rate"`
	Status    int       `json:"status"        url:"status"`
	CreatedAt time.Time `json:"created_at"    url:"created_at"`
}

func NewCreateOrderItem(basic Basic, playCode, Summary string, amount, rate float64) *CreateOrderItem {
	obj := &CreateOrderItem{}
	obj.Basic = basic
	obj.PlayCode = playCode
	obj.Summary = Summary
	obj.Rate = rate
	obj.Status = 1
	obj.Amount = amount
	obj.CreatedAt = timetool.GetNowByUTC()
	return obj
}

func (coi *CreateOrderItem) Body() url.Values {
	v, _ := query.Values(coi)
	return v
}

type PayoutOrderItem struct {
	Basic
	Rate     float64   `json:"rate"         url:"rate"`
	Win      float64   `json:"win"          url:"win"`
	Result   string    `json:"result"       url:"result"`
	PayoutAt time.Time `json:"payout_at"    url:"payout_at"`
}

func (poi *PayoutOrderItem) Body() url.Values {
	v, _ := query.Values(poi)
	return v
}

type PayoutOrder struct {
	Basic
	Result   string    `json:"result"      url:"result"`
	Summary  string    `json:"summary"     url:"summary"`
	PayoutAt time.Time `json:"payout_at"   url:"payout_at"`
}

func (c *PayoutOrder) Body() url.Values {
	v, _ := query.Values(c)
	return v
}

type ListOrderCond struct {
	Basic
	GameRound string `json:"game_round"      url:"game_round"`
	Status    string `json:"status"          url:"status"`
}

func (c *ListOrderCond) Body() url.Values {
	v, _ := query.Values(c)
	if c.GameRound == "" {
		v.Del("game_round")
	}
	if c.Status == "" {
		v.Del("status")
	}
	return v
}
