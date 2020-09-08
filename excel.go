package excel

import (
	"time"
)

// Force time to UTC to work with another Excel package.
func TimeToUTC(t time.Time) time.Time {
	t, _ = time.ParseInLocation("2006-01-02 15:04", t.Format("2006-01-02 15:04"), time.UTC)
	return t
}
