package criteria

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestMatchHasAScore(t *testing.T) {
	criteria := Criteria{CollabStyles:  NewCollabStyles(Team)}
	allMatchCriteria := Criteria{CollabStyles:  NewCollabStyles(Team)}

	result := criteria.Match(allMatchCriteria)

	assert.Equal(t, 1.0, result.Score)
}

func TestMatchHasAScoreOfZeroWithNoMatch(t *testing.T) {
	criteria := Criteria{CollabStyles: NewCollabStyles(Pair)}
	noMatchCriteria := Criteria{CollabStyles:  NewCollabStyles(Team)}

	result := criteria.Match(noMatchCriteria)

	assert.Equal(t, 0.0, result.Score)
}

func TestMatchScore50Percent(t *testing.T) {
	criteria := Criteria{CollabStyles:  NewCollabStyles(Pair, Mob)}
	halfMatchCriteria := Criteria{CollabStyles:  NewCollabStyles(Mob)}

	result := criteria.Match(halfMatchCriteria)

	assert.Equal(t, 0.5, result.Score)
}

func TestMatchForTimeZoneCollabStyleAndCodingLanguage(t *testing.T){
	tz := NewTimeZoneRange(0.0, 0.0)
	cs := NewCollabStyles(Pair, Team)
	cl := NewCodingLanguages(JavaScript, Java)
	criteria := NewCriteria(tz, cl , cs)
	assert.Equal(t, criteria, criteria)
}

func NewCriteria(matchables ...Matchable) Criteria {
	return New()
}
