package poker

import "testing"

func TestSuitMarshalText(t *testing.T) {
	b, _ := Diamonds.MarshalText()
	if string(b) != "66" {
		t.Errorf("Suit.MarshalText failed. Got: %s, Want: d", b)
	}
}

func TestSuitUnmarshalText(t *testing.T) {
	var s Suit
	if err := s.UnmarshalText([]byte("68")); err == nil {
		t.Error(`Suit.UnmarshalText failed: Unmarshal "68" didn't return error`)
	}

	for k, str := range suitCode {
		if err := s.UnmarshalText([]byte(str)); err != nil {
			t.Error(err)
		}
		if s != k {
			t.Errorf("Suit.UnmarshalText failed. Want: %s, got: %s", k, s)
		}
	}
}

func TestSuit(t *testing.T) {
	unknowns := []Card{BackFace, JokerBlack, JokerWhite}
	for _, c := range unknowns {
		if s := c.Suit(); s != UnknownSuit {
			t.Errorf("Card.Suit failed. Card: %s, got: %s", c, s)
		}
	}

	for i := 0; i < 4; i++ {
		for j := 1; j < 14; j++ {
			c := Card(i*13 + j)
			if s := c.Suit(); s != Suit(i) {
				t.Errorf("Card.Suit failed. Card: %s, got %s, want: %s", c, s, Suit(i))
			}
		}
	}
}
