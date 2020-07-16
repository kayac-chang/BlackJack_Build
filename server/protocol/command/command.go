package command

import "gitlab.fbk168.com/gamedevjp/blackjack/server/go-modules/frame/code"

const (
	c2s code.Code = 8003100

	Login       = c2s
	BackToLobby = c2s + 1
	WatchTable  = c2s + 3
	TakeSeat    = c2s + 5
	Bet         = c2s + 6
	Action      = c2s + 8
	LeaveSeat   = c2s + 10
	Exit        = c2s + 20
)

const (
	s2c code.Code = 8003000

	LoginResult  = s2c
	LobbyResult  = s2c + 1
	LobbyUpdate  = s2c + 2
	TableResult  = s2c + 3
	SeatResult   = s2c + 5
	PleaseBet    = s2c + 6
	BetResult    = s2c + 7
	Ask          = s2c + 8
	ActionResult = s2c + 9
	GameResult   = s2c + 11
)

const (
	NewRound    = s2c + 20 // open for bet
	GameStarted = s2c + 21 // stop betting
	GameSettled = s2c + 22
)

const (
	UpdateSeat   = s2c + 30
	UpdateShoe   = s2c + 31 // new shoe
	UpdateBet    = s2c + 32
	UpdateDeal   = s2c + 33
	UpdateTurn   = s2c + 34
	UpdateAction = s2c + 35
	FirstDeal    = s2c + 36
)
