package round

import (
	"database/sql/driver"
	"fmt"
	"sync"
	"time"
)

const parseFmt = "%4d-%2d-%2d-%d-%d"

// A Round is a counter to identify a game.
type Round struct {
	mu          sync.RWMutex
	date        time.Time
	shoe, round int
}

// NextShoe adds 1 to shoe number. If another day arrives,
// it resets the value to 1. It returns a Stamp with new round.
func (r *Round) NextShoe() Stamp {
	r.mu.Lock()
	defer r.mu.Unlock()

	// if new date
	now := time.Now()
	if now.Day() != r.date.Day() {
		r.date, r.shoe = now, 0
	}

	// same date
	r.shoe++
	r.round = 1
	return Stamp{r.date, r.shoe, r.round}
}

// NextRound adds 1 to round number. If another day arrives,
// it resets the value to 1. It returns a Stamp with new round.
func (r *Round) NextRound() Stamp {
	r.mu.Lock()
	defer r.mu.Unlock()

	// if new date
	now := time.Now()
	if now.Day() != r.date.Day() {
		r.date, r.shoe, r.round = now, 1, 0
	}

	// same date
	r.round++
	return Stamp{r.date, r.shoe, r.round}
}

// MarshalText implements the encoding.TextMarshaler interface.
func (r *Round) MarshalText() ([]byte, error) {
	return []byte(r.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (r *Round) UnmarshalText(b []byte) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	var err error
	r.date, r.shoe, r.round, err = parse(string(b))
	return err
}

// String implements the stringer interface. The output format
// is "2006-01-02-003-004".
func (r *Round) String() string {
	return r.Stamp().String()
}

// Stamp returns a Stamp consisted with the counters time, shoe number and round number.
func (r *Round) Stamp() Stamp {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return Stamp{r.date, r.shoe, r.round}
}

// Scan implements the sql.Scanner interface.
func (r *Round) Scan(v interface{}) error {
	b, ok := v.([]byte)
	if !ok {
		return fmt.Errorf("round: unsupport scanning type %T", v)
	}

	return r.UnmarshalText(b)
}

// Value implements the driver.Valuer interface.
func (r *Round) Value() (driver.Value, error) {
	return r.String(), nil
}

// New returns a new round counter by given start shoe and round.
func New(shoe, round int) *Round {
	return &Round{
		date:  time.Now(),
		shoe:  shoe,
		round: round,
	}
}

func parse(s string) (time.Time, int, int, error) {
	var y, d, shoe, rnd int
	var m time.Month

	_, err := fmt.Sscanf(s, parseFmt, &y, &m, &d, &shoe, &rnd)
	return time.Date(y, m, d, 0, 0, 0, 0, time.UTC), shoe, rnd, err
}

// Parse returns the Round parsed from the given string.
func Parse(s string) (*Round, error) {
	date, shoe, rnd, err := parse(s)

	return &Round{
		date:  date,
		shoe:  shoe,
		round: rnd,
	}, err
}
