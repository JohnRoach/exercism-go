// Package gigasecond was added here for reasons...
package gigasecond

import "time"

// Constant declaration.
const testVersion = 4

// AddGigasecond tell you when your time is up
func AddGigasecond(birthTime time.Time) time.Time {
	return time.Unix(birthTime.Unix()+1000000000, 0)
}
