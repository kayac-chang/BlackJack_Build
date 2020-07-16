package table

import (
	"gitlab.fbk168.com/gamedevjp/blackjack/server/go-modules/frame/code"
	"gitlab.fbk168.com/gamedevjp/blackjack/server/go-modules/poker"
	"gitlab.fbk168.com/gamedevjp/blackjack/server/protocol"
)

type pile protocol.Pile

func (p pile) codes() []string {
	return p.Cards.Codes()
}

func (p pile) hasAce() bool {
	for _, c := range p.Cards {
		if c.Rank() == 1 {
			return true
		}
	}
	return false
}

func (p pile) hasTens() bool {
	for _, c := range p.Cards {
		if c.Rank() > 9 {
			return true
		}
	}
	return false
}

func (p *pile) hit(c poker.Card) {
	p.Cards.Add(c)
}

func (p pile) blackjack() bool {
	return len(p.Cards) == 2 && p.hasAce() && p.hasTens()
}

func (p pile) bust() bool {
	return p.points() > 21
}

func (p pile) pair() bool {
	return len(p.Cards) == 2 && p.Cards[0].Rank() == p.Cards[1].Rank()
}

func (p pile) soft17() bool {
	if p.points() == 17 && p.hasAce() {
		s := append(p.Cards, poker.Clubs08)
		return s.PointsBlackjack() < 21
	}
	return false
}

func (p *pile) set(result code.Code, odds float64) {
	p.Result = result
	p.Pay = p.Bet * odds
}

func (p *pile) stay() {
	p.Result = protocol.Stayed
}

func (p pile) points() int {
	return p.Cards.PointsBlackjack()
}

func newPile() *pile {
	return &pile{Cards: poker.Set{}}
}
