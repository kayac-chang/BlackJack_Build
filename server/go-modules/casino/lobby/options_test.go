package lobby

import "testing"

type myStringer struct {
	s string
}

func (m myStringer) String() string {
	return m.s
}

func TestDefaultRoomIDer(t *testing.T) {
	cases := map[string]interface{}{
		"":       false,
		"100":    100,
		"room-a": "room-a",
		"room-b": myStringer{"room-b"},
	}

	for want, v := range cases {
		if got := defaultRoomIDer(v); got != want {
			t.Errorf("defaultRoomIDer failed.\nGot: %s, want: %s, type: %T", got, want, v)
		}
	}
}

func TestDefaultOptions(t *testing.T) {
	if err := DefaultOptions().Valid(); err != nil {
		t.Error("DefaultOptions failed:", err)
	}
}
