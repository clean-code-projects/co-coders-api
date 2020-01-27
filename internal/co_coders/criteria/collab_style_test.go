package criteria

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMatchWithNoCollabStylesReturnsEmpty(t *testing.T) {
	collabStyles := CollabStyles{Team}
	noCollabStyles := CollabStyles{}

	result := collabStyles.Match(noCollabStyles)

	assert.Equal(t, noCollabStyles, result.Matches)
	assert.Equal(t, 0.0, result.Score)
}

func TestMatchWithNoMatchingCollabStylesReturnsNoMatches(t *testing.T) {
	collabStyles := CollabStyles{Team}
	stylesToMatch := CollabStyles{Mob}

	result := collabStyles.Match(stylesToMatch)

	assert.Equal(t, CollabStyles{}, result.Matches)
	assert.Equal(t, 0.0, result.Score)
}

func TestMatchWithOneMatchingCollabStyleReturnsMatch(t *testing.T) {
	matchingCollabStyle := CollabStyles{Team}

	result := matchingCollabStyle.Match(matchingCollabStyle)

	assert.Equal(t, matchingCollabStyle, result.Matches)
	assert.Equal(t, 1.0, result.Score)
}

func TestMatchWithSubsetOfMatchingCollabStylesReturnsOnlyMatchingStyles(t *testing.T) {
	collabStyles := CollabStyles{Team, Pair}
	stylesToMatch := CollabStyles{Pair, Mob}

	result := collabStyles.Match(stylesToMatch)

	assert.Equal(t, CollabStyles{Pair}, result.Matches)
	assert.Equal(t, 0.5, result.Score)
}
