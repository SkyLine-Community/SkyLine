package SkyLine_SYSTEM

import "time"

var SleepCodes = map[string]time.Duration{
	"s":   time.Second,
	"m":   time.Minute,
	"h":   time.Hour,
	"d":   time.Hour * 24,
	"w":   time.Hour * 24 * 7,
	"mo":  time.Hour * 24 * 30,
	"y":   time.Hour * 24 * 365,
	"ms":  time.Millisecond,
	"mis": time.Microsecond,
}
