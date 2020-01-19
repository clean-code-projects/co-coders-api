package filter_test

import (
	"testing"

	"github.com/clean-code-projects/co-coders-api/assert"
	"github.com/clean-code-projects/co-coders-api/criterion"
	"github.com/clean-code-projects/co-coders-api/filter"
	"github.com/clean-code-projects/co-coders-api/user"
)


func TestFilterOnCoStyleReturnsEmptyWhenUseHasNoCoStyles(t *testing.T) {
	style := criterion.Team

	actual := filter.OnCoStyle([]user.User{}, style)

	assert.Equals(t, []user.User{}, actual)
}

func TestFilterOnCoStyleReturnsNoMatchWhenNoMatchingCoStyle(t *testing.T) {
	aUser := user.New([]criterion.CoStyle{criterion.Team})
	users := []user.User{aUser}
	style := criterion.Mob

	actual := filter.OnCoStyle(users, style)

	assert.Equals(t, []user.User{}, actual)
}

func TestFilterOnCoStyleReturnsAMatch(t *testing.T) {
	aUser := user.New([]criterion.CoStyle{criterion.Team})
	users := []user.User{aUser}
	style := criterion.Team

	actual := filter.OnCoStyle(users, style)

	assert.Equals(t, []user.User{aUser}, actual)
}

func TestFilterOnCoStyleReturnsAMatchWhenUserHasMultipleCoStyles(t *testing.T) {
	aUser := user.New([]criterion.CoStyle{criterion.Team, criterion.Pair})
	users := []user.User{aUser}
	style := criterion.Team

	actual := filter.OnCoStyle(users, style)

	assert.Equals(t, []user.User{aUser}, actual)
}


