package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"gitlab.fbk168.com/gamedevjp/blackjack/server/game"
	"gitlab.fbk168.com/gamedevjp/blackjack/server/go-modules/ulgsdk/ulg168utils"
)

func main() {
	ulg168utils.InitConf(&ulg168utils.ULG168Conf{
		// APIHost: conf.ULG168APIHost,
		// MaintainAPI:   conf.MaintainAPI,
		// MaintainToken: conf.MaintainToken,
		// GameID:        conf.GameID,
		// GameType:      "blackjack",
		// ENV:           conf.Env,
		CMDPrefix: 8003000,
	})
	if err := game.Run(); err != nil {
		log.Println(err)
	}
}
