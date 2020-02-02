package criteria

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)


type TimeZone struct {
	min float64
	max float64
	name string
}

func (t TimeZone) Name() string {
	return ""
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


func (t TimeZone) Match(targetTimeZone TimeZone) (matches []Criterion) {
	 booleanMap := make(map[float64]bool)
	 for _, num := range makeTimeZoneRange(targetTimeZone.min, targetTimeZone.max) {
	 	booleanMap[num] = true
	 }

	 ourTimeZone := makeTimeZoneRange(t.min, t.max)

	 for _, num := range ourTimeZone {
	 	if _, ok := booleanMap[num]; ok {
	 		matches = append(matches, NewTimeZone(num, num)...)
		}
	 }
	 return matches
}

type TimeZoneMatcher struct {

}

func (t TimeZoneMatcher) Match(criteria, targetCriteria []Criterion) []Criterion {
	targetTimeZone := targetCriteria[0].(TimeZone)
	timeZone  := criteria[0].(TimeZone)
	return timeZone.Match(targetTimeZone)
}


func NewTimeZone(min, max float64) []Criterion {
	return []Criterion{TimeZone{min, max, ""}}
}

func TestGMTMatchesGMT(t *testing.T) {
	gmt := 0.0
	timezone := NewTimeZone(gmt, gmt)
	targetTimeZone := NewTimeZone(gmt, gmt)
	matcher := TimeZoneMatcher{}
	result := matcher.Match(timezone, targetTimeZone)
	expected := generateTimeZones(0.0)
	assert.Equal(t, expected, result)
}

func TestNZDMatchesAST(t *testing.T) {
	nzd := 11.0
	ast := -9.0
	timezone := NewTimeZone(nzd, ast)
	targetTimeZone := NewTimeZone(nzd, ast)
	matcher := TimeZoneMatcher{}
	result := matcher.Match(timezone, targetTimeZone)
	expected := generateTimeZones(11.0, 12.0, -11.0, -10.0, -9.0)
	assert.Subset(t, expected, result)
}

func generateTimeZones(timeZones ...float64) (results []Criterion) {
	for _, tz := range timeZones {
		results = append(results, TimeZone{tz, tz, ""})
	}
	return results
}

func TestPSTMatchesEST(t *testing.T) {
	pst := -8.0
	est := -5.0
	timezone := NewTimeZone(pst, est)
	targetTimeZone := NewTimeZone(pst, est)
	matcher := TimeZoneMatcher{}
	result := matcher.Match(timezone, targetTimeZone)
	expected := generateTimeZones(-8.0, -7.0, -6.0, -5.0)
	assert.Subset(t, expected, result)
}

func TestILSTMatchesINST(t *testing.T) {
	ilst := 3.0
	inst := 5.5
	timezone := NewTimeZone(ilst, inst)
	targetTimeZone := NewTimeZone(ilst, inst)
	matcher := TimeZoneMatcher{}
	result := matcher.Match(timezone, targetTimeZone)
	expected := generateTimeZones(3.5, 3.0, 4.0, 4.5, 5.0, 5.5)
	assert.Subset(t, expected, result)
}
