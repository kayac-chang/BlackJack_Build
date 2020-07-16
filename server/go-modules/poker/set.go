package poker

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// A Set represents a pile of cards.
type Set []Card

// Clone returns a copy of itself.
func (s *Set) Clone() Set {
	return append(Set(nil), *s...)
}

// Add appends the given cards to s.
func (s *Set) Add(cards ...Card) {
	*s = append(*s, cards...)
}

// Codes returns a list of card codes.
func (s Set) Codes() []string {
	var codes []string
	for _, c := range s {
		codes = append(codes, c.Code())
	}
	return codes
}

// Deal pops n cards from s, returning the popped cards and the number
// of the rest.
func (s *Set) Deal(n int) (Set, int) {
	if n < 0 {
		panic(fmt.Sprintf("poker: dealing %d cards", n))
	}

	if n >= len(*s) {
		defer func() { *s = (*s)[:0] }()
		return *s, 0
	}

	d := (*s)[:n]
	*s = (*s)[n:]
	return d, len(*s)
}

// Draw pops a card from s, returning the popped card.
// If s is empty, it returns an error.
func (s *Set) Draw() (Card, error) {
	if len(*s) == 0 {
		return 0, errors.New("poker: not enough cards")
	}

	defer func() { *s = (*s)[1:] }()
	return (*s)[0], nil
}

// Has checks if any of the given cards appears in the set.
func (s Set) Has(cards ...Card) bool {
	for i := range s {
		for j := range cards {
			if s[i] == cards[j] {
				return true
			}
		}
	}
	return false
}

// HasAll checks if all the given cards appear in the set.
func (s Set) HasAll(cards ...Card) bool {
	for _, c := range cards {
		if !s.Has(c) {
			return false
		}
	}
	return true
}

// Remove deletes the cards in the set from the s, returning the
// removed elements.
func (s *Set) Remove(cards ...Card) (Set, error) {
	if len(*s) == 0 {
		return *s, errors.New("poker: not enough cards")
	}

	removed := Set{}
	for _, c := range cards {
		for i := 0; i < len(*s); i++ {
			if (*s)[i] == c {
				removed.Add(c)
				(*s)[i] = (*s)[len(*s)-1]
				*s = (*s)[:len(*s)-1]
				break
			}
		}
	}

	return removed, nil
}

// IndexOf returns the indexes of the given cards respectively.
// If the tile doesn't exist, it returns -1.
func (s Set) IndexOf(cards ...Card) []int {
	idx := make([]int, len(cards))

Outer:
	for i := range cards {
		for j := range s {
			if s[j] == cards[i] {
				idx[i] = j
				continue Outer
			}
		}
		idx[i] = -1
	}

	return idx
}

// PointsBlackjack returns total points of the set under the rule
// of Blackjack. The J, Q and K are counted as 10 points, while
// the Ace may be 1 or 11, depending on the rest of the cards.
func (s Set) PointsBlackjack() int {
	var total, aces int
	for _, c := range s {
		r := c.Rank()
		switch {
		case r == 1:
			aces++
		case r > 10:
			r = 10
			fallthrough
		default:
			total += r
		}
	}

	for ; aces > 1; aces-- {
		total++
	}
	if aces == 1 {
		if total < 11 {
			total += 11
		} else {
			total++
		}
	}
	return total
}

// Shuffle randomises the sequence of the Set and returns itself.
func (s Set) Shuffle() Set {
	rand.Shuffle(len(s), func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})
	return s
}

// Decks returns a Set which consisted of n decks of cards. Which each
// deck contains j jokers. If j > 0, white Joker assigned before the
// black one.
func Decks(n, j int) Set {
	if n < 0 || j < 0 {
		msg := fmt.Sprintf("poker: invalid number of decks %d or jokers %d", n, j)
		panic(msg)
	}

	current := 0
	s := make(Set, n*(52+j))

	for i := 0; i < n; i++ {

		// assign 52 cards
		for c := 1; c < 53; c++ {
			s[current] = Card(c)
			current++
		}

		// assign jokers
		for c := 0; c < j; c++ {
			s[current] = Card(54 - c%2) // assign white Joker first
			current++
		}
	}

	return s
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
