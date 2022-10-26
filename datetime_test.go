package elegant

import (
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	now := time.Now()
	timeNow = func() time.Time { return now }

	m.Run()
}

func TestMakeCurrentDateUTC(t *testing.T) {
	now := timeNow().In(time.UTC)
	resultTime := TimeDate(-1, -1, -1, 12, 42, 12, 134, time.UTC)

	checkDate(t, resultTime, now.Year(), now.Month(), now.Day())
	checkTime(t, resultTime, 12, 42, 12, 134)
}

func TestMakeCurrentDate(t *testing.T) {
	now := timeNow()
	resultTime := TimeDate(-1, -1, -1, 12, 42, 12, 134, nil)

	checkDate(t, resultTime, now.Year(), now.Month(), now.Day())
	checkTime(t, resultTime, 12, 42, 12, 134)
}

func TestMakeAnotherDate(t *testing.T) {
	now := timeNow()
	resultTime := TimeDate(1872, 3, 23, -1, -1, -1, -1, nil)

	checkDate(t, resultTime, 1872, 3, 23)
	checkTime(t, resultTime, now.Hour(), now.Minute(), now.Second(), now.Nanosecond())
}

func checkDate(t *testing.T, date time.Time, year int, month time.Month, day int) {
	t.Helper()

	if date.Year() != year {
		t.Errorf("Year does not match, Want %d Have %d", date.Year(), year)
	}

	if date.Month() != month {
		t.Errorf("Month does not match, Want %d Have %d", date.Month(), month)
	}

	if date.Day() != day {
		t.Errorf("Day does not match. Want %d Have %d", date.Day(), day)
	}
}

func checkTime(t *testing.T, date time.Time, hour, minute, second, nanosecond int) {
	t.Helper()

	if date.Hour() != hour {
		t.Errorf("Hour does not match. Want %d Have %d", date.Hour(), hour)
	}

	if date.Minute() != minute {
		t.Errorf("Minutes does not match. Want %d Have %d", date.Minute(), minute)
	}

	if date.Second() != second {
		t.Errorf("Seconds does not match. Want %d Have %d", date.Second(), second)
	}

	if date.Nanosecond() != nanosecond {
		t.Errorf("Nanosecond does not match. Want %d Have %d", date.Nanosecond(), nanosecond)
	}
}
