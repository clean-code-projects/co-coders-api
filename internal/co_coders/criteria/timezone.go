package criteria

import (
	"fmt"
	"math"
)

type TimeZone struct {
	offset float64
}

func (t TimeZone) String() string {
	return fmt.Sprintf("TimeZone{offset: %.2f}", t.offset)
}

func NewTimeZone(offset float64) TimeZone {
	return TimeZone{offset: offset}
}

func NewTimeZoneRange(min, max float64) TimeZoneRange {
	return TimeZoneRange{min, max}
}

func makeTimeZoneRange(min, max float64) []float64 {
	var result []float64
	result = append(result, min)
	tz := min
	for tz < math.Floor(max) {
		zone, ok := nextIncrementalTimeZone(tz)
		if ok {
			result = append(result, zone)
		}
		nextTz := nextFullTimeZone(tz)
		result = append(result, nextTz)
		tz = nextTz
	}
	return result
}

func nextFullTimeZone(tz float64) float64 {
	nftz := tz + 1.0
	if nftz >= 12.0 {
		nftz = -11.0
	}
	return nftz
}

func nextIncrementalTimeZone(tz float64) (float64, bool) {
	switch tz {
	case -10.00:
		return -9.5, true
	case -4.00:
		return -3.5, true
	case 3.00:
		return 3.5, true
	case 4.00:
		return 4.5, true
	case 5.00:
		return 5.5, true
	case 5.5:
		return 5.75, true
	case 6.0:
		return 6.5, true
	case 8.0:
		return 8.75, true
	case 9.0:
		return 9.5, true
	case 10.0:
		return 10.5, true
	case 12.0:
		return 12.75, true
	}
	return 0.0, false
}

type TimeZoneRange struct {
	min float64
	max float64
}

func (t TimeZoneRange) Match(targetTimeZoneRange Matchable) (matches []Criterion) {
	booleanMap := make(map[float64]bool)

	timeZoneRange, ok := targetTimeZoneRange.(TimeZoneRange)
	if !ok {
		return []Criterion{}
	}

	for _, offset := range makeTimeZoneRange(timeZoneRange.min, timeZoneRange.max) {
		booleanMap[offset] = true
	}

	for _, offset := range makeTimeZoneRange(t.min, t.max) {
		if _, ok := booleanMap[offset]; ok {
			matches = append(matches, NewTimeZone(offset))
		}
	}
	return matches
}

