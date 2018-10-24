// Package gigasecond impliments the methods to calculate moment in time
package gigasecond

// import path for the time package from the standard library
import "time"

// AddGigasecond accepts time, adds 1B seconds to it and returns time 
func AddGigasecond(t time.Time) time.Time {
	return t.Local().Add(time.Second * time.Duration(1000000000))
}
