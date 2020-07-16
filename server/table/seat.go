package table

import (
	"gitlab.fbk168.com/gamedevjp/blackjack/server/go-modules/casino/player"
	"gitlab.fbk168.com/gamedevjp/blackjack/server/go-modules/poker"
	"gitlab.fbk168.com/gamedevjp/blackjack/server/protocol"
	"gitlab.fbk168.com/gamedevjp/blackjack/server/protocol/action"
)

type seat struct {
	player    *player.Player
	insurance float64
	pay       float64
	piles     []*pile
	quited    bool
}

func (s *seat) clean() {
	if s.quited && s.player != nil {
		s.player.Quit()
		s.player = nil
		s.quited = false
	}

	s.piles = []*pile{newPile(), newPile()}
	s.insurance = 0
	s.pay = 0
}

func (s *seat) hasBet() bool {
	return s.player != nil && s.piles[0].Bet != 0
}

func (s *seat) quit() bool {
	if !s.hasBet() {
		s.player = nil
		return true
	}

	s.piles[0].Result = protocol.Stayed
	s.piles[1].Result = protocol.Stayed
	s.quited = true
	return false
}

func (s *seat) reset() {
	s.insurance = 0
	s.piles = []*pile{newPile(), newPile()}

	if s.quited && s.player != nil {
		s.player = nil
	}
}

func (s *seat) settle(dpt int, dbj bool) {
	if !s.hasBet() {
		return
	}

	for i, p := range s.piles {

		// didn't bet
		if p.Bet == 0 {
			continue
		}

		// lost if bust
		if p.bust() {
			p.set(protocol.Lost, 0)
			continue
		}

		// player chose paid out mmediately
		if p.Result == protocol.Paid {
			p.Pay = 2 * p.Bet
			continue
		}

		// player gave up
		if p.Result == protocol.GaveUp {
			p.Pay = p.Bet / 2
			continue
		}

		// player doubled, check the first result
		if i == 1 && p.Action == action.Double {
			p.Pay = s.piles[0].Pay
			p.Result = s.piles[0].Result
			continue
		}

		pbj := p.blackjack() && s.piles[1].Bet == 0
		switch {

		// both Blackjack, ties
		case dbj && pbj:
			p.set(protocol.Tied, 1)

		// dealer Blackjack, lost
		case dbj:
			p.set(protocol.Lost, 0)

		// player Blackjack, more pay
		case pbj:
			p.set(protocol.Won, 2.5)

		// dealer bust, won if not bust
		case dpt > 21:
			p.set(protocol.Won, 2)

		// if not bust, higher point wins
		default:
			ppt := p.points()
			switch {
			case ppt > dpt:
				p.set(protocol.Won, 2)
			case ppt == dpt:
				p.set(protocol.Tied, 1)
			default:
				p.set(protocol.Lost, 0)
			}
		}
	}

	if dbj && s.insurance != 0 {
		s.pay = s.piles[0].Bet
	} else {
		s.pay = 0
	}
}

type dealerSeat struct {
	pile
	player *player.Player
	quited bool
}

func (s *dealerSeat) clean() {
	if s.quited == true {
		s.player = nil
		s.quited = false
	}

	s.Cards = poker.Set{}
}

func newSeat() *seat {
	return &seat{piles: []*pile{newPile(), newPile()}}
}
