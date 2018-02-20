package gigasecond

import (
	"math"
	"time"
)

// AddGigasecond calculates when someone born at given time will be 10^9 seconds old.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Duration(math.Pow(10, 9)) * time.Second)
}
