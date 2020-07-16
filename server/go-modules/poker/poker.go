package poker

import "sync"

var (
	mu     sync.Mutex
	mR     bool
	rankOf = defaultRank
)

func defaultRank(c Card) int {
	if c == BackFace {
		return 0
	}

	if c == JokerBlack || c == JokerWhite {
		return 14
	}

	r := int(c % 13)
	if r == 0 {
		return 13
	}
	return r
}

// RegisterRankFunc sets the rule to count card ranks.
// Multiple registration causes panic.
func RegisterRankFunc(f func(Card) int) {
	if f == nil {
		panic("poker: nil func registered")
	}

	mu.Lock()
	defer mu.Unlock()

	if mR {
		panic("poker: multiple RankFunc registered")
	}

	mR = true
	rankOf = f
}
