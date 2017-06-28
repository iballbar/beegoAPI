package utils

import "time"

// InTimeSpan check time between start and end
// return boolean
func InTimeSpan(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end)
}
