package protocol

import (
	"gitlab.fbk168.com/gamedevjp/blackjack/server/go-modules/poker"
	"gitlab.fbk168.com/gamedevjp/blackjack/server/protocol/action"
)

type Data struct {
	ID        int            `json:"id"`
	Round     string         `json:"round,omitempty"`
	No        int            `json:"no"`
	MaxBet    float64        `json:"max_bet,omitempty"`
	MinBet    float64        `json:"min_bet,omitempty"`
	TotalBet  float64        `json:"total_bet,omitempty"`
	Bets      []BetData      `json:"bets,omitempty"`
	Pile      int            `json:"pile"`
	Expire    int            `json:"expire,omitempty"`
	Action    string         `json:"action,omitempty"`
	Options   action.Options `json:"options,omitempty"`
	Result    []Pile         `json:"result,omitempty"`
	Insurance float64        `json:"insurance,omitempty"`
}

type BetData struct {
	ID     int             `json:"-"`
	Round  string          `json:"-"`
	Dealer poker.Set       `json:"-"`
	No     int             `json:"no"`
	Action map[string]Pile `json:"-"`
	Bet    float64         `json:"bet"`
}

type GameResult struct {
	ID    int     `json:"id"`
	Round string  `json:"round"`
	Win   float64 `json:"win"`
}
