//package gigasecond calculates the
package gigasecond

// import path for the time package from the standard library
import "time"

// AddGigasecond function adds 10e9 seconds to any given time
func AddGigasecond(t time.Time) time.Time {
	t = t.Add(time.Second * 1e9) 
	return t
}
