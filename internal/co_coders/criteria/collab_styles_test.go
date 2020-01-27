package criteria

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMatchOnCollabStyleReturnsEmpty(t *testing.T) {
	collabStyles := CollabStyles{Team}
	noCollabStyles := CollabStyles{}

	result := collabStyles.Match(noCollabStyles)

	assert.Equal(t, noCollabStyles, result)
}

func TestMatchOnCollabStyleReturnsNoMatch(t *testing.T) {
	collabStyles := CollabStyles{Team}
	stylesToMatch := CollabStyles{Mob}

	result := collabStyles.Match(stylesToMatch)

	assert.Equal(t, CollabStyles{}, result)
}

func TestMatchOnCollabStyleReturnsAMatch(t *testing.T) {
	matchingCollabStyle := CollabStyles{Team}

	result := matchingCollabStyle.Match(matchingCollabStyle)

	assert.Equal(t, matchingCollabStyle, result)
}

func TestMatchOnCollabStyleSubsetReturnsAMatch(t *testing.T) {
	collabStyles := CollabStyles{Team, Pair}
	stylesToMatch := CollabStyles{Pair}

	result := collabStyles.Match(stylesToMatch)

	assert.Equal(t, stylesToMatch, result)
}
