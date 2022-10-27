# Elegant
A collection of useful functions that will make your code more elegant

## Date/time

### TimeDate(year int, month time.Month, day, hour, min, sec, nsec int, loc *time.Location) time.Time
TimeDate works like standard time.Date, but support skip any part if passed -1 (nil for Location)

Example:
```
// Creates a time.Time for the current date, with the current timezone at 12:42:12.134
time := TimeDate(-1, -1, -1, 12, 42, 12, 134, nil)
```
