package types

import "time"

func (d *Duration) Duration() time.Duration {
	return time.Duration(d.Nanos) + time.Duration(d.Seconds*int64(time.Second))
}

func NewDurationFromTm(d time.Duration) *Duration {
	sec := int64(d.Seconds())
	nanos := int32(d.Nanoseconds() - sec*int64(time.Second))
	return &Duration{
		Seconds: sec,
		Nanos:   nanos,
	}
}
