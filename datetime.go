package elegant

import "time"

// TimeDate works like standard time.Date, but support skip any part if passed -1 (nil for Location)
func TimeDate(year int, month time.Month, day, hour, min, sec, nsec int, loc *time.Location) time.Time {
	now := time.Now()
	if loc != nil {
		now = now.In(loc)
	}

	if year == -1 {
		year = now.Year()
	}

	if month == -1 {
		month = now.Month()
	}

	if day == -1 {
		day = now.Day()
	}

	if hour == -1 {
		hour = now.Hour()
	}

	if min == -1 {
		min = now.Minute()
	}

	if sec == -1 {
		sec = now.Second()
	}

	if nsec == -1 {
		nsec = now.Nanosecond()
	}

	if loc == nil {
		loc = now.Location()
	}

	return time.Date(year, month, day, hour, min, sec, nsec, loc)
}
