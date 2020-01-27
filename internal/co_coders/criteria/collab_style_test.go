package criteria

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMatchWithNoCollabStylesReturnsEmpty(t *testing.T) {
	collabStyles := CollabStyles{Team}
	noCollabStyles := CollabStyles{}

	result := collabStyles.Match(noCollabStyles)

	assert.Equal(t, noCollabStyles, result)
}

func TestMatchWithNoMatchingCollabStylesReturnsNoMatches(t *testing.T) {
	collabStyles := CollabStyles{Team}
	stylesToMatch := CollabStyles{Mob}

	result := collabStyles.Match(stylesToMatch)

	assert.Equal(t, CollabStyles{}, result)
}

func TestMatchWithOneMatchingCollabStyleReturnsMatch(t *testing.T) {
	matchingCollabStyle := CollabStyles{Team}

	result := matchingCollabStyle.Match(matchingCollabStyle)

	assert.Equal(t, matchingCollabStyle, result)
}

func TestMatchWithSubsetOfMatchingCollabStylesReturnsOnlyMatchingStyles(t *testing.T) {
	collabStyles := CollabStyles{Team, Pair}
	stylesToMatch := CollabStyles{Pair, Mob}

	result := collabStyles.Match(stylesToMatch)

	assert.Equal(t, CollabStyles{Pair}, result)
}
