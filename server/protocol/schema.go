package protocol

import (
	"gitlab.fbk168.com/gamedevjp/blackjack/server/go-modules/frame/code"
	"gitlab.fbk168.com/gamedevjp/blackjack/server/go-modules/poker"
)

type Deal struct {
	ShoeNum int        `json:"shoe_num"`
	Card    poker.Card `json:"card"`
	No      int        `json:"no"`
	Pile    int        `json:"pile"`
	Cards   poker.Set  `json:"cards"`
	Points  int        `json:"points"`
}

type Pile struct {
	Cards  poker.Set `json:"cards"`
	Bet    float64   `json:"bet"`
	Pay    float64   `json:"pay"`
	Result code.Code `json:"result"`
	Action string    `json:"action"`
}

type Seat struct {
	No        int     `json:"no"`
	Account   string  `json:"account"`   // player UID
	Player    string  `json:"player"`    // player name
	TotalBet  float64 `json:"total_bet"` // sum of pile[0] and pile[1], insurance excluded
	Insurance float64 `json:"insurance"`
	Pay       float64 `json:"pay"`
	Piles     []Pile  `json:"piles"`
}

type TableStat struct {
	ID       int           `json:"id"`
	MaxBet   float64       `json:"max_bet,omitempty"`
	MinBet   float64       `json:"min_bet,omitempty"`
	SeatsNum int           `json:"seats_num,omitempty"`
	Occupied int           `json:"occupied"`
	History  []interface{} `json:"history"`
}

type TableUpdate struct {
	Round       string    `json:"round,omitempty"`
	State       []int     `json:"state,omitempty"`
	ShoeNum     int       `json:"shoe_num,omitempty"`
	Dealer      string    `json:"dealer,omitempty"`
	DealerCards poker.Set `json:"dealer_cards,omitempty"`
	Seats       []Seat    `json:"seats,omitempty"`
}

type Table struct {
	TableStat
	TableUpdate
}
