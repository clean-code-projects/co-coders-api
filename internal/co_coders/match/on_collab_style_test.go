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

