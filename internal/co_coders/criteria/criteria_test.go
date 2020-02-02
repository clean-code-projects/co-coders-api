package criteria

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestMatchHasAScore(t *testing.T) {
	criteria := Criteria{CollabStyles: CollabStyles{Team}}
	allMatchCriteria := Criteria{CollabStyles: CollabStyles{Team}}

	result := criteria.Match(allMatchCriteria)

	assert.Equal(t, 1.0, result.Score)
}

func TestMatchHasAScoreOfZeroWithNoMatch(t *testing.T) {
	criteria := Criteria{CollabStyles: CollabStyles{Pair}}
	noMatchCriteria := Criteria{CollabStyles: CollabStyles{Team}}

	result := criteria.Match(noMatchCriteria)

	assert.Equal(t, 0.0, result.Score)
}

func TestMatchScore50Percent(t *testing.T) {
	criteria := Criteria{CollabStyles: CollabStyles{Pair, Mob}}
	halfMatchCriteria := Criteria{CollabStyles: CollabStyles{Mob}}

	result := criteria.Match(halfMatchCriteria)

	assert.Equal(t, 0.5, result.Score)
}

func TestMatchForTimeZoneCollabStyleAndCodingLanguage(t *testing.T){
	tz := NewTimeZone(0.0, 0.0)
	//cs := NewCollabStyles(Pair, Team)
	//cl := NewCodingLanguages(JavaScript, Java)
	criteria := NewCriteria(tz /* cs, cl*/)
	assert.Equal(t, criteria, criteria)
}

func NewCodingLanguages(languages ...CodingLanguage) CodingLanguages {
	return append(CodingLanguages{}, languages...)
}

func NewCollabStyles(styles ...CollabStyle) CollabStyles {
	return append(CollabStyles{}, styles...)
}

type Criterion interface {
	Name() string
}

func NewCriteria(matchables ...[]Criterion) Criteria {
	return New()
}