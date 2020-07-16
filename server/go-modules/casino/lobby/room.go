package lobby

import (
	"gitlab.fbk168.com/gamedevjp/blackjack/server/go-modules/casino/player"
	"gitlab.fbk168.com/gamedevjp/blackjack/server/go-modules/frame"
	"gitlab.fbk168.com/gamedevjp/blackjack/server/go-modules/frame/code"
)

// A Room represents a the place where a game plays.
type Room interface {
	Wellcome(*player.Player) error
	Player(id string) *player.Player
	Kick(id string, c code.Code, reason interface{}) *player.Player

	ID() string
	Stat() interface{}

	Broadcast(f frame.Frame)
	Open()
	Close(soft bool, reason interface{})
}
