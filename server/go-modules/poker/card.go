package poker

import "fmt"

// Count is the number of cards in a deck, jokers excluded.
const Count = 52

// A deck of French playing cards.
const (
	BackFace Card = iota

	ClubsAce
	Clubs02
	Clubs03
	Clubs04
	Clubs05
	Clubs06
	Clubs07
	Clubs08
	Clubs09
	Clubs10
	ClubsJack
	ClubsQueen
	ClubsKing

	DiamondsAce
	Diamonds02
	Diamonds03
	Diamonds04
	Diamonds05
	Diamonds06
	Diamonds07
	Diamonds08
	Diamonds09
	Diamonds10
	DiamondsJack
	DiamondsQueen
	DiamondsKing

	HeartsAce
	Hearts02
	Hearts03
	Hearts04
	Hearts05
	Hearts06
	Hearts07
	Hearts08
	Hearts09
	Hearts10
	HeartsJack
	HeartsQueen
	HeartsKing

	SpadesAce
	Spades02
	Spades03
	Spades04
	Spades05
	Spades06
	Spades07
	Spades08
	Spades09
	Spades10
	SpadesJack
	SpadesQueen
	SpadesKing

	JokerBlack
	JokerWhite
)

var (
	faces = map[Card]string{
		BackFace: "🂠",

		ClubsAce:   "🃑",
		Clubs02:    "🃒",
		Clubs03:    "🃓",
		Clubs04:    "🃔",
		Clubs05:    "🃕",
		Clubs06:    "🃖",
		Clubs07:    "🃗",
		Clubs08:    "🃘",
		Clubs09:    "🃙",
		Clubs10:    "🃚",
		ClubsJack:  "🃛",
		ClubsQueen: "🃝",
		ClubsKing:  "🃞",

		DiamondsAce:   "🃁",
		Diamonds02:    "🃂",
		Diamonds03:    "🃃",
		Diamonds04:    "🃄",
		Diamonds05:    "🃅",
		Diamonds06:    "🃆",
		Diamonds07:    "🃇",
		Diamonds08:    "🃈",
		Diamonds09:    "🃉",
		Diamonds10:    "🃊",
		DiamondsJack:  "🃋",
		DiamondsQueen: "🃍",
		DiamondsKing:  "🃎",

		HeartsAce:   "🂱",
		Hearts02:    "🂲",
		Hearts03:    "🂳",
		Hearts04:    "🂴",
		Hearts05:    "🂵",
		Hearts06:    "🂶",
		Hearts07:    "🂷",
		Hearts08:    "🂸",
		Hearts09:    "🂹",
		Hearts10:    "🂺",
		HeartsJack:  "🂻",
		HeartsQueen: "🂽",
		HeartsKing:  "🂾",

		SpadesAce:   "🂡",
		Spades02:    "🂢",
		Spades03:    "🂣",
		Spades04:    "🂤",
		Spades05:    "🂥",
		Spades06:    "🂦",
		Spades07:    "🂧",
		Spades08:    "🂨",
		Spades09:    "🂩",
		Spades10:    "🂪",
		SpadesJack:  "🂫",
		SpadesQueen: "🂭",
		SpadesKing:  "🂮",

		JokerBlack: "🃏",
		JokerWhite: "🃟",
	}

	code = map[Card]string{
		BackFace: "A0",

		ClubsAce:   "D1",
		Clubs02:    "D2",
		Clubs03:    "D3",
		Clubs04:    "D4",
		Clubs05:    "D5",
		Clubs06:    "D6",
		Clubs07:    "D7",
		Clubs08:    "D8",
		Clubs09:    "D9",
		Clubs10:    "DA",
		ClubsJack:  "DB",
		ClubsQueen: "DD",
		ClubsKing:  "DE",

		DiamondsAce:   "C1",
		Diamonds02:    "C2",
		Diamonds03:    "C3",
		Diamonds04:    "C4",
		Diamonds05:    "C5",
		Diamonds06:    "C6",
		Diamonds07:    "C7",
		Diamonds08:    "C8",
		Diamonds09:    "C9",
		Diamonds10:    "CA",
		DiamondsJack:  "CB",
		DiamondsQueen: "CD",
		DiamondsKing:  "CE",

		HeartsAce:   "B1",
		Hearts02:    "B2",
		Hearts03:    "B3",
		Hearts04:    "B4",
		Hearts05:    "B5",
		Hearts06:    "B6",
		Hearts07:    "B7",
		Hearts08:    "B8",
		Hearts09:    "B9",
		Hearts10:    "BA",
		HeartsJack:  "BB",
		HeartsQueen: "BD",
		HeartsKing:  "BE",

		SpadesAce:   "A1",
		Spades02:    "A2",
		Spades03:    "A3",
		Spades04:    "A4",
		Spades05:    "A5",
		Spades06:    "A6",
		Spades07:    "A7",
		Spades08:    "A8",
		Spades09:    "A9",
		Spades10:    "AA",
		SpadesJack:  "AB",
		SpadesQueen: "AD",
		SpadesKing:  "AE",

		JokerBlack: "CF",
		JokerWhite: "DF",
	}
)

// A Card represents a card in a standard 52-card deck.
type Card int

// Code returns the card code.
func (c Card) Code() string {
	return code[c]
}

// MarshalText implements the encoding.TextMarshaler interface.
func (c Card) MarshalText() ([]byte, error) {
	return []byte(code[c]), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (c *Card) UnmarshalText(b []byte) error {
	for *c = 0; int(*c) < len(code); *c++ {
		if code[*c] == string(b) {
			break
		}
	}

	if int(*c) == len(code) {
		return fmt.Errorf("poker: unknown code %s", b)
	}
	return nil
}

// Rank returns the card's rank.
// By default, the rank of BackFace is 0 and the jokers are 14.
// One can modify the rule by calling RegisterRankFunc().
func (c Card) Rank() int {
	return rankOf(c)
}

func (c Card) String() string {
	return faces[c]
}
