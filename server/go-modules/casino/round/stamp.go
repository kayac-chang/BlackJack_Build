package round

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// A Stamp records the counters time, the shoe and the round number.
type Stamp struct {
	Date        time.Time
	Shoe, Round int
}

// String implements the stringer interface. The output format
// is "2006-01-02-003-004".
func (s Stamp) String() string {
	return fmt.Sprintf("%s-%03d-%03d", s.Date.Format("2006-01-02"), s.Shoe, s.Round)
}

// MarshalText implements the encoding.TextMarshaler interface.
func (s Stamp) MarshalText() ([]byte, error) {
	return []byte(s.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (s *Stamp) UnmarshalText(b []byte) error {
	var err error
	s.Date, s.Shoe, s.Round, err = parse(string(b))
	return err
}

// Scan implements the sql.Scanner interface.
func (s *Stamp) Scan(v interface{}) error {
	b, ok := v.([]byte)
	if !ok {
		return fmt.Errorf("round: unsupport scanning type %T", v)
	}

	return s.UnmarshalText(b)
}

// Value implements the driver.Valuer interface.
func (s Stamp) Value() (driver.Value, error) {
	return s.String(), nil
}
