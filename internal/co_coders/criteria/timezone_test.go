package criteria

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)


type TimeZone struct {
	min float64
	max float64
}


func makeTimeZoneRange(min, max float64) []float64 {
	var timeZoneSlice []float64
	timeZoneSlice = append(timeZoneSlice, min)
	tz := min
	for tz < math.Floor(max) {
		zone, ok := nextIncrementalTimeZone(tz)
		if ok {
			timeZoneSlice = append(timeZoneSlice, zone)
		}
		nextTz := nextFullTimeZone(tz)
		timeZoneSlice = append(timeZoneSlice, nextTz)
		tz = nextTz
	}
	return timeZoneSlice
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
		case -10.00: return -9.5, true
		case -4.00: return -3.5, true
		case 3.00: return 3.5, true
		case 4.00: return 4.5, true
		case 5.00: return 5.5, true
		case 5.5: return 5.75, true
		case 6.0: return 6.5, true
		case 8.0: return 8.75, true
		case 9.0: return 9.5, true
		case 10.0: return 10.5, true
		case 12.0: return 12.75, true
	}
	return 0.0, false
}


func (t TimeZone) Match(targetTimeZone TimeZone) (matches []float64) {
	 booleanMap := make(map[float64]bool)
	 for _, num := range makeTimeZoneRange(targetTimeZone.min, targetTimeZone.max) {
	 	booleanMap[num] = true
	 }

	 ourTimeZone := makeTimeZoneRange(t.min, t.max)

	 for _, num := range ourTimeZone {
	 	if _, ok := booleanMap[num]; ok {
	 		matches = append(matches, num)
		}
	 }
	 return matches
}

func NewTimeZone(min, max float64) TimeZone {
	return TimeZone{min, max}
}

func TestGMTMatchesGMT(t *testing.T) {
	gmt := 0.0
	timezone := NewTimeZone(gmt, gmt)
	targetTimeZone := NewTimeZone(gmt, gmt)
	result := timezone.Match(targetTimeZone)
	expected := []float64{0.0}
	assert.Equal(t, expected, result)
}

func TestNZDMatchesAST(t *testing.T) {
	nzd := 11.0
	ast := -9.0
	timezone := NewTimeZone(nzd, ast)
	targetTimeZone := NewTimeZone(nzd, ast)
	result := timezone.Match(targetTimeZone)
	expected := []float64{11.0, 12.0, -11.0, -10.0, -9.0}
	assert.Subset(t, expected, result)
}

func TestPSTMatchesEST(t *testing.T) {
	pst := -8.0
	est := -5.0
	timezone := NewTimeZone(pst, est)
	targetTimeZone := NewTimeZone(pst, est)
	result := timezone.Match(targetTimeZone)
	expected := []float64{-8.0, -7.0, -6.0, -5.0}
	assert.Subset(t, expected, result)
}

func TestILSTMatchesINST(t *testing.T) {
	ilst := 3.0
	inst := 5.5
	timezone := NewTimeZone(ilst, inst)
	targetTimeZone := NewTimeZone(ilst, inst)
	result := timezone.Match(targetTimeZone)
	expected := []float64{3.5, 3.0, 4.0, 4.5, 5.0, 5.5}
	assert.Subset(t, expected, result)
}
