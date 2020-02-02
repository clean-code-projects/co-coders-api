package criteria

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)


type TimeZone struct {
	offset float64
}

func (t TimeZone) String() string {
	return fmt.Sprintf("TimeZone{offset: %.2f}", t.offset)
}

func NewTimeZone(offset float64) TimeZone {
	return TimeZone{offset:offset}
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

type TimeZoneRange struct {
	min float64
	max float64
}

func (t TimeZoneRange) Match(target Matchable) (matches []Criterion) {
	booleanMap := make(map[float64]bool)

	targetTimeZoneRange, ok := target.(TimeZoneRange)
	if !ok {
		return []Criterion{}
	}

	for _, offset := range makeTimeZoneRange(targetTimeZoneRange.min, targetTimeZoneRange.max) {
		booleanMap[offset] = true
	}

	for _, offset := range makeTimeZoneRange(t.min, t.max) {
		if _, ok := booleanMap[offset]; ok {
			matches = append(matches, NewTimeZone(offset))
		}
	}
	return matches
}

func TestGMTMatchesGMT(t *testing.T) {
	gmt := 0.0
	timezone := NewTimeZoneRange(gmt, gmt)
	targetTimeZone := NewTimeZoneRange(gmt, gmt)
	result := timezone.Match(targetTimeZone)
	expected := generateTimeZones(0.0)
	assert.Equal(t, expected, result)
}

func TestNZDMatchesAST(t *testing.T) {
	nzd := 11.0
	ast := -9.0
	timezone := NewTimeZoneRange(nzd, ast)
	targetTimeZone := NewTimeZoneRange(nzd, ast)
	result := timezone.Match(targetTimeZone)
	expected := generateTimeZones(11.0, 12.0, -11.0, -10.0, -9.0)
	assert.Subset(t, expected, result)
}


func TestPSTMatchesEST(t *testing.T) {
	pst := -8.0
	est := -5.0
	timezone := NewTimeZoneRange(pst, est)
	targetTimeZone := NewTimeZoneRange(pst, est)
	result := timezone.Match(targetTimeZone)
	expected := generateTimeZones(-8.0, -7.0, -6.0, -5.0)
	assert.Subset(t, expected, result)
}

func TestILSTMatchesINST(t *testing.T) {
	ilst := 3.0
	inst := 5.5
	timezone := NewTimeZoneRange(ilst, inst)
	targetTimeZone := NewTimeZoneRange(ilst, inst)
	result := timezone.Match(targetTimeZone)
	expected := generateTimeZones(3.5, 3.0, 4.0, 4.5, 5.0, 5.5)
	assert.Subset(t, expected, result)
}

func TestMatchWithInvalidTypeReturnsNoMatches(t *testing.T) {
	timezone := NewTimeZoneRange(0.0, 0.0)
	result := timezone.Match(invalidMatchable{})
	assert.Equal(t, []Criterion{}, result)
}

func TestTimeZoneString(t *testing.T){
	timeZone := NewTimeZone(1.0)
	assert.Equal(t, "TimeZone{offset: 1.00}",timeZone.String())
}

func generateTimeZones(timeZones ...float64) (results []Criterion) {
	for _, tz := range timeZones {
		results = append(results, NewTimeZone(tz))
	}
	return results
}

type invalidMatchable struct {}

func (i invalidMatchable) Match(m Matchable) []Criterion {
	return []Criterion{}
}