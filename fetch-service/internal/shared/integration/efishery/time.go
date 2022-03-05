package efishery

import (
	"encoding/json"
	"strconv"
	"time"
)

const (
	dateOnlyFormat = "2006-01-02"
	dateTimeFormat = "2006-01-02 15:04:05"
)

type Time time.Time

func (t *Time) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}

	if s == "" {
		return nil
	}

	ti, err := time.Parse(time.RFC3339Nano, s)
	if err == nil {
		*t = (Time)(ti)
		return nil
	}

	ti, err = time.Parse(dateOnlyFormat, s)
	if err == nil {
		*t = (Time)(ti)
		return nil
	}

	ti, err = time.Parse(dateTimeFormat, s)
	if err == nil {
		*t = (Time)(ti)
		return nil
	}

	i, err := strconv.ParseInt(s[:10], 10, 64)
	if err == nil {
		*t = (Time)(time.Unix(i, 0))
		return nil
	}

	return err
}
