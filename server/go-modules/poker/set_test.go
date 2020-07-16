package poker

import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	s := Set{SpadesAce, Diamonds02}
	s.Add(Clubs03)

	n := Set{Hearts04, JokerBlack}
	s.Add(n...)

	want := Set{SpadesAce, Diamonds02, Clubs03, Hearts04, JokerBlack}
	if !reflect.DeepEqual(s, want) {
		t.Errorf("Set.Add failed.\nGot:  %v\nWant: %v", s, want)
	}
}

func TestDeal(t *testing.T) {
	s := Set{Spades03, Clubs05, HeartsJack, DiamondsAce}
	n := 2

	want, rest := s[:n], s[n:]
	got, left := s.Deal(n)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Set.Deal failed.\nGot:  %v\nWant: %v", got, want)
	}
	if !reflect.DeepEqual(s, rest) || left != len(rest) {
		t.Errorf("Set.Deal failed.\nRest: %v\nWant: %v", s, rest)
	}

	// test dealing all cards
	got, left = s.Deal(len(s))
	if !reflect.DeepEqual(got, rest) {
		t.Errorf("Set.Deal failed.\nGot:  %v\nWant: %v", got, rest)
	}
	if left != 0 || len(s) != 0 {
		t.Errorf("Set.Deal failed.\nRest: %v\nWant: %v", s, Set{})
	}
}

func TestDraw(t *testing.T) {
	s := Set{Spades03, Clubs05, HeartsJack, DiamondsAce}

	c, err := s.Draw()
	if err != nil {
		t.Error("Set.Draw failed: ", err)
	}

	if c != Spades03 {
		t.Errorf("Set.Draw failed. Drawn card: %v, want: %v", c, Spades03)
	}

	rest := Set{Clubs05, HeartsJack, DiamondsAce}
	if !reflect.DeepEqual(s, rest) {
		t.Errorf("Set.Draw failed.\nRest: %v\nWant: %v", s, rest)
	}
}

func TestHas(t *testing.T) {
	s := Decks(1, 2).Shuffle()
	w := Set{Spades08, Hearts09, ClubsAce, DiamondsKing, JokerWhite}

	for _, c := range w {
		if !s.Has(c) {
			t.Error("Set.Has failed")
		}
	}

	if !s.Has(w...) {
		t.Error("Set.Has failed")
	}

	if w.Has(Spades07, Hearts03) || w.Has(Diamonds08) {
		t.Error("Set.Has failed")
	}
}

func TestHasAll(t *testing.T) {
	s := Decks(1, 2).Shuffle()
	w := Set{Clubs05, Diamonds03, HeartsAce, SpadesQueen}

	if !s.HasAll(w...) {
		t.Error("Set.HasAll failed")
	}

	if w.HasAll(s...) {
		t.Error("Set.HasAll failed")
	}
}

func TestPointsBlackjack(t *testing.T) {
	m := map[int][]Set{
		8: []Set{Set{Spades03, Hearts05}},
		12: []Set{
			Set{ClubsAce, DiamondsAce},
			Set{HeartsAce, Spades04, Diamonds07},
			Set{DiamondsAce, Clubs07, Hearts04},
			Set{Spades02, Diamonds08, ClubsAce, HeartsAce},
		},
		13: []Set{
			Set{SpadesQueen, ClubsAce, DiamondsAce, HeartsAce},
			Set{Hearts08, Diamonds05},
			Set{Spades03, ClubsKing},
		},
		21: []Set{
			Set{Hearts10, ClubsAce},
			Set{ClubsQueen, DiamondsJack, SpadesAce},
		},
		23: []Set{Set{DiamondsKing, HeartsJack, ClubsAce, SpadesAce, HeartsAce}},
	}

	for i := range m {
		for _, s := range m[i] {
			if p := s.PointsBlackjack(); p != i {
				t.Errorf("Set.PointsBlackjack failed. Got: %d, cards: %v", p, s)
			}
		}
	}
}

func TestShuffle(t *testing.T) {
	s := Set{SpadesAce, Spades02, Spades08, SpadesJack}
	o := make(Set, len(s))
	copy(o, s)

	o.Shuffle()
	if reflect.DeepEqual(o, s) {
		t.Error("Set.Shuffle failed. Cards not shuffled")
	}
}

func TestIndexOf(t *testing.T) {
	s := Set{Diamonds03, ClubsAce, SpadesJack, Hearts08, JokerBlack, Clubs08, HeartsKing, Spades10, Diamonds07}
	cases := []Set{
		Set{JokerBlack, HeartsKing, Diamonds03},
		Set{JokerWhite, Clubs08, SpadesJack, Hearts10},
		Set{SpadesAce, DiamondsKing, Hearts02, Clubs07},
		nil,
	}
	wants := [][]int{
		[]int{4, 6, 0},
		[]int{-1, 5, 2, -1},
		[]int{-1, -1, -1, -1},
		[]int{},
	}

	t.Log("Set :", s)
	for i, want := range wants {
		t.Log("Case:", cases[i])
		if got := s.IndexOf(cases[i]...); !reflect.DeepEqual(got, wants[i]) {
			t.Errorf("Set.IndexOf failed.\nGot : %v\nWant: %v", got, want)
		}
	}
}

func TestDecks(t *testing.T) {
	cases := [][]int{
		[]int{3, 0}, // decks, jokers per deck
		[]int{2, 3},
		[]int{5, 1},
	}

	for _, c := range cases {
		t.Logf("Decks: assigning %d decks, %d jokers per deck.", c[0], c[1])

		s := Decks(c[0], c[1])
		if len(s) != c[0]*(52+c[1]) {
			t.Errorf("Decks failed: Got %d cards", len(s))
		}

		// test if each suit and rank are assigned
		for i, got := range s {
			want := i%(52+c[1]) + 1
			if want > 52 {
				want = 53 + want%2
			}

			if got != Card(want) {
				t.Errorf("Decks failed: Got %v at index %d, want: %v", got, i, want)
			}
		}
	}
}
