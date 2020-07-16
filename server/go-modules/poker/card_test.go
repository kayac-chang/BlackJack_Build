package poker

import "testing"

func TestCardMarshalText(t *testing.T) {
	b, _ := DiamondsJack.MarshalText()
	if string(b) != "CB" {
		t.Errorf("Card.MarshalText failed. Got: %s, Want: CB", b)
	}
}

func TestCardUnmarshalText(t *testing.T) {
	var c Card
	if err := c.UnmarshalText([]byte("KK")); err == nil {
		t.Error(`Card.UnmarshalText failed: Unmarshal "KK" didn't return error`)
	}

	for k, str := range code {
		if err := c.UnmarshalText([]byte(str)); err != nil {
			t.Error(err)
		}
		if c != k {
			t.Errorf("Card.UnmarshalText failed. Want: %s, got: %s", k, c)
		}
	}
}

func TestRank(t *testing.T) {
	for i := 0; i < 4; i++ {
		for j := 1; j < 14; j++ {
			c := Card(i*13 + j)
			if c.Rank() != j {
				t.Errorf("Card.Rank failed. Card: %v, got: %d", c, j)
			}
		}
	}
}
