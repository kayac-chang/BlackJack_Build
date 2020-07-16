package round

import (
	"encoding/json"
	"testing"
	"time"
)

const sample = "2019-02-01-54321-6789"

func TestRoundMarshalText(t *testing.T) {
	r, err := Parse(sample)
	if err != nil {
		t.Fatal("Parse failed:", err)
	}

	b, err := json.Marshal(r)
	if err != nil {
		t.Fatal("Round.MarshalText failed:", err)
	}

	w := `"` + sample + `"`
	if string(b) != w {
		t.Errorf("Round.MarshalText failed.\nGot : %s\nWant: %s", b, w)
	}
}

func TestRoundUnmarshalText(t *testing.T) {
	r := New(0, 0)
	b := []byte(`"` + sample + `"`)

	if err := json.Unmarshal(b, r); err != nil {
		t.Fatal("Round.UnmarshalText failed:", err)
	}

	if g := r.String(); g != sample {
		t.Errorf("Round.UnmarshalText failed.\nGot : %s\nWant: %s", g, sample)
	}
}

func TestParse(t *testing.T) {
	var cases = map[string][]int{
		"2018-12-10-001-002":    []int{2018, 12, 10, 1, 2},
		"2019-01-02-12345-6789": []int{2019, 1, 2, 12345, 6789},
	}

	for k, v := range cases {
		r, err := Parse(k)
		if err != nil {
			t.Error("Parse failed:", err)
		}

		s := r.Stamp()
		d := time.Date(v[0], time.Month(v[1]), v[2], 0, 0, 0, 0, time.UTC)
		if !s.Date.Equal(d) || s.Shoe != v[3] || s.Round != v[4] {
			t.Errorf("Parse failed.\nGot : %s\nWant: %s", k, r)
		}
	}
}
