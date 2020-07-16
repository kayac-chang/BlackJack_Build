package poker

import "fmt"

// list of suits
const (
	Clubs Suit = iota
	Diamonds
	Hearts
	Spades
	UnknownSuit
)

var (
	suitStr = map[Suit]string{
		Clubs:       "♣",
		Diamonds:    "♦",
		Hearts:      "♥",
		Spades:      "♠",
		UnknownSuit: "Unknown",
	}

	suitCode = map[Suit]string{
		Clubs:       "63",
		Diamonds:    "66",
		Hearts:      "65",
		Spades:      "60",
		UnknownSuit: "u",
	}
)

// A Suit represents a suit of a card.
type Suit int

// Code returns the suit code.
func (s Suit) Code() string {
	return suitCode[s]
}

// MarshalText implements the encoding.TextMarshaler interface.
func (s Suit) MarshalText() ([]byte, error) {
	return []byte(suitCode[s]), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (s *Suit) UnmarshalText(b []byte) error {
	for *s = range suitCode {
		if suitCode[*s] == string(b) {
			return nil
		}
	}

	return fmt.Errorf("poker: unknown suit code %s", b)
}

func (s Suit) String() string {
	return suitStr[s]
}

// Suit returns the card's suit.
func (c Card) Suit() Suit {
	if c == BackFace {
		return UnknownSuit
	}
	return Suit((c - 1) / 13)
}
