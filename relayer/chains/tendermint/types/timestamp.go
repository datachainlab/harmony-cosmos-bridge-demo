package types

import (
	"time"
)

func (t *Timestamp) Time() time.Time {
	return time.Unix(t.Seconds, int64(t.Nanos))
}

func NewTimestampFromTime(t time.Time) *Timestamp {
	return &Timestamp{
		Seconds: t.Unix(),
		Nanos:   int32(t.Nanosecond()),
	}
}
