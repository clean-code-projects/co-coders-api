package filter_test

import (
	"github.com/clean-code-projects/co-coders-api/internal/criteria"
	"github.com/clean-code-projects/co-coders-api/internal/filter"
	"github.com/clean-code-projects/co-coders-api/internal/user"
	"github.com/clean-code-projects/co-coders-api/pkg/assert"

	"testing"
)

func TestFilterOnCoStyleReturnsEmpty(t *testing.T) {
	style := criteria.Team
	actual := filter.OnCoStyle([]user.User{}, style)
	assert.Equals(t, []user.User{}, actual)
}

func TestFilterOnCoStyleReturnsNoMatch(t *testing.T) {
	aUser := user.New([]criteria.CoStyle{criteria.Team})
	users := []user.User{aUser}
	style := criteria.Mob
	actual := filter.OnCoStyle(users, style)
	assert.Equals(t, []user.User{}, actual)
}

func TestFilterOnCoStyleReturnsAMatch(t *testing.T) {
	aUser := user.New([]criteria.CoStyle{criteria.Team})
	users := []user.User{aUser}
	style := criteria.Team
	actual := filter.OnCoStyle(users, style)
	assert.Equals(t, []user.User{aUser}, actual)
}

func TestFilterOnCoStyleSubsetReturnsAMatch(t *testing.T) {
	aUser := user.New([]criteria.CoStyle{criteria.Team, criteria.Pair})
	users := []user.User{aUser}
	criterion := criteria.Team
	actual := filter.OnCoStyle(users, criterion)
	assert.Equals(t, []user.User{aUser}, actual)
}

