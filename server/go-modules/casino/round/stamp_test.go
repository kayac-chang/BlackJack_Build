package round

import (
	"encoding/json"
	"testing"
	"time"
)

func TestStampMarshalText(t *testing.T) {
	s := Stamp{time.Date(2019, time.February, 1, 0, 0, 0, 0, time.UTC), 54321, 6789}

	b, err := json.Marshal(s)
	if err != nil {
		t.Fatal("Stamp.MarshalText failed:", err)
	}

	w := `"` + sample + `"`
	if string(b) != w {
		t.Errorf("Stamp.MarshalText failed.\nGot : %s\nWant: %s", b, w)
	}
}

func TestStampUnmarshalText(t *testing.T) {
	s := Stamp{}
	b := []byte(`"` + sample + `"`)

	if err := json.Unmarshal(b, &s); err != nil {
		t.Fatal("Stamp.UnmarshalText failed:", err)
	}

	if g := s.String(); g != sample {
		t.Errorf("Stamp.UnmarshalText failed.\nGot : %s\nWant: %s", g, sample)
	}
}
