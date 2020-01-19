package user_test

import (
	"reflect"
	"testing"
)

type CoStyle int

const (
	Team CoStyle = iota + 1
	Mob
)

// User ..
type User struct {
	coStyle CoStyle
}

func FilterOnCoStyle(users []User, coStyle CoStyle) (matchedUser []User) {
	matchedUser = []User{}
	for _, user := range users {
		if user.coStyle == coStyle {
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
	users := []User{User{coStyle: Team}}
	style := Mob
	actual := FilterOnCoStyle(users, style)
	AssertEqual(t, []User{}, actual)
}

func TestMatchOnCoStyleReturnsAMatch(t *testing.T) {
	users := []User{User{coStyle: Team}}
	style := Team
	actual := FilterOnCoStyle(users, style)
	AssertEqual(t, []User{User{coStyle: Team}}, actual)
}

func AssertEqual(t *testing.T, expected interface{}, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("%+v is not equal to %+v", expected, actual)
	}
}
