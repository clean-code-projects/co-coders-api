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
	matchCriteria := criteria.Criteria{
		CollabStyles: []criteria.CollabStyle{criteria.Team}}
	allMatchUser := user.New(matchCriteria.CollabStyles)
	usersToMatch := []user.User{allMatchUser}
	aMatch := OnCriteria(matchCriteria, usersToMatch)
	assert.Equal(t,1.0, aMatch.Score)
}

func TestMatchHasAScoreOfZeroWithNoMatch(t *testing.T) {
	matchCriteria := criteria.Criteria{
		CollabStyles: []criteria.CollabStyle{criteria.Pair}}
	noMatchUser := user.New([]criteria.CollabStyle{criteria.Team})
	usersToMatch := []user.User{noMatchUser}
	aMatch := OnCriteria(matchCriteria, usersToMatch)
	assert.Equal(t,0.0, aMatch.Score)
}

func TestMatchScore50Percent(t *testing.T) {
	matchCriteria := criteria.Criteria{
		CollabStyles: []criteria.CollabStyle{criteria.Pair, criteria.Mob}}
	halfMatchUser := user.New([]criteria.CollabStyle{criteria.Mob})
	usersToMatch := []user.User{halfMatchUser}
	aMatch := OnCriteria(matchCriteria, usersToMatch)
	assert.Equal(t,0.5, aMatch.Score)
}

func OnCriteria(aCriteria criteria.Criteria, users []user.User) criteria.Match {
	total := 0.0
	matches := 0.0
	for _, collabStyle := range aCriteria.CollabStyles {
		result := match.OnCollabStyle(users, collabStyle)
		if len(result) > 0 {
			matches ++
		}
		total ++
	}
	return criteria.Match{Score: matches / total}
}
