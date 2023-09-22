package receipts

import (
	"fmt"
	"strings"
	"time"
)

type Date time.Time

func (c *Date) UnmarshalJSON(data []byte) (err error) {
	s := strings.Trim(string(data), "\"")
	if s == "null" {
		return
	}
	parsedTime, err := time.Parse(time.DateOnly, s)
	if err != nil {
		return
	}
	*c = Date(parsedTime)
	return
}

func (c Date) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", time.Time(c).Format(time.DateOnly))), nil
}

func (c Date) Day() int {
	return time.Time(c).Day()
}
