package datetime

import "time"

// InTimeSpan check time between start and end
// return boolean
func InTimeSpan(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end)
}

// InTimeSpanEqual check time between start and end including check time.
// return boolean
func InTimeSpanEqual(start, end, check time.Time) bool {
	return (check.After(start) || check.Equal(start)) && (check.Before(end) || check.Equal(end))
}
