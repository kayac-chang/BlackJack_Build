package order

import (
	"encoding/json"
	"net/url"
	"time"

	"github.com/google/go-querystring/query"
)

type Order struct {
	Basic
	Order *OrderRes `json:"order"`
}

type OrderRes struct {
	Basic
	UUID       string          `json:"uuid"         url:"uuid"`
	Account    string          `json:"account"      url:"account"`
	GameType   string          `json:"game_type"    url:"game_type"`
	GameID     string          `json:"game_id"      url:"game_id"`
	GameRoom   string          `json:"game_room"    url:"game_room"`
	GameRound  string          `json:"game_round"   url:"game_round"`
	Result     *string         `json:"result"       url:"result"`
	Summary    *string         `json:"summary"      url:"summary"`
	Amount     float64         `json:"amount"       url:"amount"`
	Win        *float64        `json:"win"          url:"win"`
	Status     int             `json:"status"       url:"status"`
	GameToken  *string         `json:"game_token"   url:"game_token"`
	Remark     *string         `json:"remark"       url:"remark"`
	CreatedAt  time.Time       `json:"created_at"   url:"created_at"`
	UpdatedAt  time.Time       `json:"updated_at"   url:"updated_at"`
	PayoutAt   *time.Time      `json:"payout_at"    url:"payout_at"`
	OrderItems []*OrderItemRes `json:"order_items"  url:"order_items"`
}

func (or *OrderRes) Body() url.Values {
	v, _ := query.Values(or)
	v.Del("order_items")
	items, _ := json.Marshal(or.OrderItems)
	v.Add("order_items", string(items))
	return v
}

type OrderItemRes struct {
	UUID      string     `json:"uuid"        url:"uuid"`
	OrderUUID string     `json:"order_uuid"  url:"order_uuid"`
	PlayCode  string     `json:"play_code"   url:"play_code"`
	Summary   string     `json:"summary"     url:"summary"`
	Amount    float64    `json:"amount"      url:"amount"`
	Rate      *float64   `json:"rate"        url:"rate"`
	Win       *float64   `json:"win"         url:"win"`
	Result    *string    `json:"result"      url:"result"`
	Status    int        `json:"status"      url:"status"`
	GameToken *string    `json:"game_token"   url:"game_token"`
	Remark    *string    `json:"remark"      url:"remark"`
	CreatedAt time.Time  `json:"created_at"  url:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"  url:"updated_at"`
	PayoutAt  *time.Time `json:"payout_at"   url:"payout_at"`
}
