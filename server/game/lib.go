package game

import (
	"database/sql"

	"gitlab.fbk168.com/gamedevjp/blackjack/server/conf"
	"gitlab.fbk168.com/gamedevjp/blackjack/server/go-modules/casino/lobby"
	"gitlab.fbk168.com/gamedevjp/blackjack/server/go-modules/casino/round"
	"gitlab.fbk168.com/gamedevjp/blackjack/server/table"
	"go.uber.org/zap"
)

func newLogger() (*zap.Logger, error) {
	if conf.Env == "dev" {
		return zap.NewDevelopment()
	}
	return zap.NewProduction()
}

func applyTables(db *sql.DB, lby *lobby.Lobby, z *zap.Logger) error {
	q := "SELECT id, max_bet, min_bet, seats_num, last_round FROM rooms;"
	rows, err := db.Query(q)
	if err != nil {
		return err
	}
	defer rows.Close()

	opts := table.Options{
		Database: db,
		Lobby:    lby,
		Logger:   z,
	}
	var last sql.NullString
	for rows.Next() {
		if err = rows.Scan(&opts.ID, &opts.MaxBet, &opts.MinBet, &opts.SeatsNum, &last); err != nil {
			return err
		}

		if last.Valid {
			r, err := round.Parse(last.String)
			if err != nil {
				return err
			}
			opts.LastShoe = r.Stamp().Shoe
		}

		t := table.New(opts)
		lby.Add(t)
	}
	return rows.Err()
}
