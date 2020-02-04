package criteria

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	Team = NewCollabStyle("Team")
	Pair = NewCollabStyle("Pair")
	Mob = NewCollabStyle("Mob")
)

func TestMatchWithNoCollabStylesReturnsEmpty(t *testing.T) {
	collabStyles := NewCollabStyles(Team)
	noCollabStyles := NewCollabStyles()

	result := collabStyles.Match(noCollabStyles)

	assert.Equal(t, []Criterion{}, result)
}

func TestMatchWithNoMatchingCollabStylesReturnsNoMatches(t *testing.T) {
	collabStyles := NewCollabStyles(Team)
	stylesToMatch := NewCollabStyles(Mob)

	result := collabStyles.Match(stylesToMatch)

	assert.Equal(t, []Criterion{}, result)
}

func TestMatchWithOneMatchingCollabStyleReturnsMatch(t *testing.T) {
	matchingCollabStyle := NewCollabStyles(Team)

	result := matchingCollabStyle.Match(matchingCollabStyle)

	assert.Equal(t, []Criterion{Team}, result)
}

func TestMatchWithSubsetOfMatchingCollabStylesReturnsOnlyMatchingStyles(t *testing.T) {
	collabStyles := NewCollabStyles(Team, Pair)
	stylesToMatch := NewCollabStyles(Pair, Mob)

	result := collabStyles.Match(stylesToMatch)

	assert.Equal(t, []Criterion{Pair}, result)
}
