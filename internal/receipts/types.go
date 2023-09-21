package receipts

import (
	"fmt"
	"strings"
	"time"
)

type Date struct {
	time.Time
}

func (c *Date) UnmarshalJSON(data []byte) (err error) {
	s := strings.Trim(string(data), "\"")
	if s == "null" {
		c.Time = time.Time{}
		return
	}
	c.Time, err = time.Parse(time.DateOnly, s)
	return
}

func (c *Date) MarshalJSON() ([]byte, error) {
	if c.Time.IsZero() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", c.Time.Format(time.DateOnly))), nil
}
