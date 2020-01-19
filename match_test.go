package user_test

import (
	"reflect"
	"testing"
)

type CoStyle int

const (
	Team CoStyle = iota
	Pair
	Mob
)

func (c CoStyle) ToString() string {
	return []string{"Team", "Pair", "Mob"}[c]
}

// User ..
type User struct {
	coStyles []CoStyle
}

func (u User) HasCoStyle(coStyle CoStyle) bool {
	for _, style := range u.coStyles {
		if style == coStyle {
			return true
		}
	}
	return false
}

func FilterOnCoStyle(users []User, coStyle CoStyle) (matchedUser []User) {
	matchedUser = []User{}
	for _, user := range users {
		if user.HasCoStyle(coStyle) {
			matchedUser = append(matchedUser, user)
		}
	}
	return matchedUser
}

func TestMatchOnCoStyleReturnsEmpty(t *testing.T) {
	users := []User{}
	style := Team
	actual := FilterOnCoStyle(users, style)
	AssertEqual(t, []User{}, actual)
}

func TestMatchOnCoStyleReturnsNoMatch(t *testing.T) {
	user := createUser([]CoStyle{Team})
	users := []User{user}
	style := Mob
	actual := FilterOnCoStyle(users, style)
	AssertEqual(t, []User{}, actual)
}

func TestMatchOnCoStyleReturnsAMatch(t *testing.T) {
	user := createUser([]CoStyle{Team})
	users := []User{user}
	style := Team
	actual := FilterOnCoStyle(users, style)
	AssertEqual(t, []User{user}, actual)
}

func TestMatchOnCoStyleSubsetReturnsAMatch(t *testing.T) {
	user := createUser([]CoStyle{Team, Pair})
	users := []User{user}
	criterion := Team
	actual := FilterOnCoStyle(users, criterion)
	AssertEqual(t, []User{user}, actual)
}

func TestCoStyleToString(t *testing.T) {
	AssertEqual(t, Team.ToString(), "Team")
	AssertEqual(t, Pair.ToString(), "Pair")
	AssertEqual(t, Mob.ToString(), "Mob")
}

func createUser(coStyles []CoStyle) User {
	return User{coStyles: coStyles}
}

func AssertEqual(t *testing.T, expected interface{}, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("%+v is not equal to %+v", expected, actual)
	}
}
