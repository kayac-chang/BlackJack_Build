package controller

import "gitlab.fbk168.com/gamedevjp/blackjack/server/go-modules/frame"

type Frame struct {
	frame.Frame
	Token     string `json:"token,omitempty"`
	GameToken string `json:"game_token,omitempty"`
	GameID    string `json:"game_id,omitempty"`
}
