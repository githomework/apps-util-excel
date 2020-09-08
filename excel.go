package excel

import (
	"time"
)

const layout = "2006-01-02 15:04:05"

// Force time to UTC to work with another Excel package.
func TimeToUTC(t time.Time) time.Time {
	t, _ = time.ParseInLocation(layout, t.Format(layout), time.UTC)
	return t
}
