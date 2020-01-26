package match_test

import (
	"github.com/clean-code-projects/co-coders-api/internal/co_coders/criteria"
	"github.com/clean-code-projects/co-coders-api/internal/co_coders/match"
	"github.com/clean-code-projects/co-coders-api/internal/co_coders/user"
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestMatchOnCollabStyleReturnsEmpty(t *testing.T) {
	style := criteria.Team
	actual := match.OnCollabStyle([]user.User{}, style)
	assert.Equal(t, []user.User{}, actual)
}

func TestMatchOnCollabStyleReturnsNoMatch(t *testing.T) {
	aUser := user.New([]criteria.CollabStyle{criteria.Team})
	users := []user.User{aUser}
	style := criteria.Mob
	actual := match.OnCollabStyle(users, style)
	assert.Equal(t, []user.User{}, actual)
}

func TestMatchOnCollabStyleReturnsAMatch(t *testing.T) {
	aUser := user.New([]criteria.CollabStyle{criteria.Team})
	users := []user.User{aUser}
	style := criteria.Team
	actual := match.OnCollabStyle(users, style)
	assert.Equal(t, []user.User{aUser}, actual)
}

func TestMatchOnCollabStyleSubsetReturnsAMatch(t *testing.T) {
	aUser := user.New([]criteria.CollabStyle{criteria.Team, criteria.Pair})
	users := []user.User{aUser}
	criterion := criteria.Team
	actual := match.OnCollabStyle(users, criterion)
	assert.Equal(t, []user.User{aUser}, actual)
}

func TestMatchHasAScore(t *testing.T) {
	aCriteria := user.New([]criteria.CollabStyle{criteria.Team, criteria.Pair})
	users := []user.User{aCriteria}
	aMatch := OnCriteria(aCriteria, users)
	assert.Equal(t,1.0, aMatch.Score)
}

func TestMatchHasAScoreOfZeroWithNoMatch(t *testing.T) {
	matchCriteria := user.New([]criteria.CollabStyle{criteria.Pair})
	badMatch := user.New([]criteria.CollabStyle{criteria.Team})
	users := []user.User{badMatch}
	aMatch := OnCriteria(matchCriteria, users)
	assert.Equal(t,0.0, aMatch.Score)
}

func TestMatchScore50Percent(t *testing.T) {
	// TODO name does not match type user to criteria
	matchCriteria := user.New([]criteria.CollabStyle{criteria.Pair, criteria.Mob})
	halfMatch := user.New([]criteria.CollabStyle{criteria.Mob})
	users := []user.User{halfMatch}
	aMatch := OnCriteria(matchCriteria, users)
	assert.Equal(t,0.5, aMatch.Score)
}

type Match struct{
	Score float64
}

func OnCriteria(aCriteria user.User, users []user.User) Match {
	total := 0.0
	matches := 0.0
	for _, collabStyle := range aCriteria.CollabStyles {
		result := match.OnCollabStyle(users, collabStyle)
		if len(result) > 0 {
			matches ++
		}
		total ++
	}
	return Match{matches / total}
}