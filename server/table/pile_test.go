package table

import (
	"reflect"
	"testing"

	"gitlab.fbk168.com/gamedevjp/blackjack/server/go-modules/poker"
)

var (
	bjs = []poker.Set{
		poker.Set{poker.ClubsAce, poker.DiamondsKing},
		poker.Set{poker.DiamondsAce, poker.HeartsQueen},
		poker.Set{poker.HeartsAce, poker.Spades10},
		poker.Set{poker.SpadesAce, poker.ClubsJack},
	}

	normals = poker.Set{
		poker.Hearts02, poker.Diamonds03, poker.Clubs04, poker.Spades05,
		poker.Hearts06, poker.Diamonds07, poker.Clubs08, poker.Spades09,
	}
)

func TestHasAce(t *testing.T) {
	for _, s := range bjs {
		p := pile{Cards: s}
		if !p.hasAce() {
			t.Errorf("pile.hasAce failed. Set %v got false", p.Cards)
		}
	}

	p := pile{Cards: normals}
	if p.hasAce() {
		t.Errorf("pile.hasAce failed. Set %v got true", p.Cards)
	}
}

func TestHasTens(t *testing.T) {
	for _, s := range bjs {
		p := pile{Cards: s}
		if !p.hasTens() {
			t.Errorf("pile.hasTens failed. Set %v got false", p.Cards)
		}
	}
}

func TestBlackjack(t *testing.T) {
	for _, s := range bjs {
		p := pile{Cards: s}
		if !p.blackjack() {
			t.Errorf("pile.isBlackJack failed. Set %v got false", p.Cards)
		}
	}
}

func TestHit(t *testing.T) {
	p := pile{Cards: poker.Set{poker.Hearts05}}
	w := poker.Set{poker.Hearts05, poker.DiamondsKing}

	p.hit(poker.DiamondsKing)
	if !reflect.DeepEqual(p.Cards, w) {
		t.Errorf("pile.hit failed:\nGot  : %v\n Want: %v\n", p.Cards, w)
	}
}

func TestSoft17(t *testing.T) {
	m := map[bool][]poker.Set{
		true: []poker.Set{
			poker.Set{poker.Clubs06, poker.DiamondsAce},
			poker.Set{poker.HeartsAce, poker.SpadesAce, poker.Clubs05},
		},
		false: []poker.Set{
			poker.Set{poker.DiamondsAce, poker.Hearts06, poker.SpadesJack},
			poker.Set{poker.Clubs04, poker.Diamonds07, poker.Hearts03, poker.Spades02},
		},
	}

	for b, ss := range m {
		for _, s := range ss {
			p := pile{Cards: s}
			if !p.soft17() == b {
				t.Errorf("pile.soft17 failed. Got: %t, cards: %v", p.soft17(), s)
			}
		}
	}
}
