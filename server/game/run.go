package game

import (
	"database/sql"
	"strconv"

	"gitlab.fbk168.com/gamedevjp/blackjack/server/conf"
	"gitlab.fbk168.com/gamedevjp/blackjack/server/controller"
	"gitlab.fbk168.com/gamedevjp/blackjack/server/go-modules/casino/lobby"
	"gitlab.fbk168.com/gamedevjp/blackjack/server/protocol"
	"gitlab.fbk168.com/gamedevjp/blackjack/server/protocol/command"
	"go.uber.org/zap"
)

var opts = lobby.Options{
	Commands: lobby.CommandSet{
		Wellcome: command.LobbyResult,
		Move:     command.WatchTable,
		Exit:     command.Exit,
	},
	RoomID: func(v interface{}) string {
		if d, ok := v.(protocol.Data); ok {
			return strconv.Itoa(d.ID)
		}
		return ""
	},
	HandleDuplicated: lobby.RemoveFormer,
}

func Run() error {

	// create logger with config
	z, err := newLogger()
	if err != nil {
		return err
	}
	opts.Logger = z

	// create lobby
	lby, err := lobby.New(&opts)
	if err != nil {
		z.Fatal("failed to create lobby", zap.Error(err))
		return err
	}
	z.Info("rtp control", zap.Int("level", conf.RTPctrl))

	// read database with config
	db, err := sql.Open("mysql", conf.MysqlConf.ConverToPath())
	if err != nil {
		z.Fatal("failed to open database", zap.Error(err))
		return err
	}

	if err = applyTables(db, lby, z); err != nil {
		z.Fatal("failed to create tables", zap.Error(err))
		return err
	}

	// apply lobby to controller
	controller.ApplyLobby(lby)
	controller.ApplyLogger(z)

	// lobby open
	go lby.Open()

	controller.Start()
	return nil
}
