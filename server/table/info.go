package table

import (
	"strconv"

	"gitlab.fbk168.com/gamedevjp/blackjack/server/go-modules/casino/player"
	"gitlab.fbk168.com/gamedevjp/blackjack/server/protocol"
)

func (t *Table) ID() string {
	return strconv.Itoa(t.id)
}

func (t *Table) Player(id string) *player.Player {
	t.mu.RLock()
	defer t.mu.RUnlock()

	return t.players[id]
}

func (t *Table) stat() protocol.TableStat {
	t.mu.RLock()
	defer t.mu.RUnlock()

	var current int
	for _, s := range t.seats {
		if s.player != nil {
			current++
		}
	}

	return protocol.TableStat{
		ID:       t.id,
		SeatsNum: len(t.seats),
		Occupied: current,
		MaxBet:   t.maxBet,
		MinBet:   t.minBet,
		History:  t.history,
	}
}

func (t *Table) Stat() interface{} {
	return t.stat()
}

func (t *Table) Seats() []protocol.Seat {
	t.mu.RLock()
	defer t.mu.RUnlock()

	var seats []protocol.Seat
	for i, s := range t.seats {
		if s.player == nil {
			seats = append(seats, protocol.Seat{No: i})
			continue
		}

		var piles []protocol.Pile
		for _, p := range s.piles {
			piles = append(piles, protocol.Pile(*p))
		}

		seats = append(seats, protocol.Seat{
			No:        i,
			Account:   s.player.ID(),
			Player:    s.player.Name(),
			TotalBet:  s.piles[0].Bet + s.piles[1].Bet,
			Insurance: s.insurance,
			Pay:       s.pay,
			Piles:     piles,
		})
	}
	return seats
}

func (t *Table) Detail() protocol.Table {
	t.mu.RLock()
	defer t.mu.RUnlock()

	return protocol.Table{
		t.Stat().(protocol.TableStat),
		protocol.TableUpdate{
			Round:       t.round.String(),
			State:       t.state,
			ShoeNum:     len(t.shoe),
			Dealer:      "",
			DealerCards: t.dealer.Cards,
			Seats:       t.Seats(),
		},
	}
}
