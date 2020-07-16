package table

import (
	"database/sql"

	"gitlab.fbk168.com/gamedevjp/blackjack/server/go-modules/casino/lobby"
	"gitlab.fbk168.com/gamedevjp/blackjack/server/protocol"
	"go.uber.org/zap"
)

type Options struct {
	protocol.TableStat
	Database *sql.DB
	Logger   *zap.Logger
	Lobby    *lobby.Lobby
	LastShoe int
}
