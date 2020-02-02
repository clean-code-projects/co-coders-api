package criteria

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func TestTimeZoneString(t *testing.T) {
	timeZone := NewTimeZone(1.0)
	assert.Equal(t, "TimeZone{offset: 1.00}", timeZone.String())
}

func generateTimeZones(timeZones ...float64) (results []Criterion) {
	for _, tz := range timeZones {
		results = append(results, NewTimeZone(tz))
	}
	return results
}

type invalidMatchable struct{}

func (i invalidMatchable) Match(m Matchable) []Criterion {
	return []Criterion{}
}
