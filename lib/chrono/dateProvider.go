package chrono

import "time"

type ITimeProvider interface {
	GetTime() time.Time
}

type TimeProvider struct {
}



func (s *TimeProvider) GetTime() time.Time {
	return time.Now()
}
