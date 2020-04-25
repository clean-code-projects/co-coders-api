package criteria

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatchWithNoCollabStylesReturnsEmpty(t *testing.T) {
	collabStyles := NewCollabStyles("Team")
	noCollabStyles := NewCollabStyles()

	result := collabStyles.Match(noCollabStyles)

	assert.Equal(t, generateNamedMatches(), result)
}

func TestMatchWithNoMatchingCollabStylesReturnsNoMatches(t *testing.T) {
	collabStyles := NewCollabStyles("Team")
	stylesToMatch := NewCollabStyles("Mob")

	result := collabStyles.Match(stylesToMatch)

	assert.Equal(t, generateNamedMatches(), result)
}

func TestMatchWithOneMatchingCollabStyleReturnsMatch(t *testing.T) {
	matchingCollabStyle := NewCollabStyles("Team")

	result := matchingCollabStyle.Match(matchingCollabStyle)

	assert.Equal(t, generateNamedMatches("Team"), result)
}

func TestMatchWithSubsetOfMatchingCollabStylesReturnsOnlyMatchingStyles(t *testing.T) {
	collabStyles := NewCollabStyles("Team", "Pair")
	stylesToMatch := NewCollabStyles("Pair", "Mob")

	result := collabStyles.Match(stylesToMatch)

	assert.Equal(t, generateNamedMatches("Pair"), result)
}
