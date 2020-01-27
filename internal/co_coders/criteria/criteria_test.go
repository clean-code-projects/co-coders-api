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
