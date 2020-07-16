package protocol

import "gitlab.fbk168.com/gamedevjp/blackjack/server/go-modules/frame/code"

const (
	BadRequest        code.Code = 400
	PleaseLogin       code.Code = 401
	MoneyNotEnough    code.Code = 402
	SeatOccupied      code.Code = 403
	NotFound          code.Code = 404
	CommandNotAllowed code.Code = 405
	Timeout           code.Code = 408
	AlreadyLoggedIn   code.Code = 409
	BetOutOfRange     code.Code = 413
	Closed            code.Code = 503
)

const (
	Stayed code.Code = 20 + iota
	Won
	Paid
	Tied
	Lost
	GaveUp
)
