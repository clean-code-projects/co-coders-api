package filter

import (
	"github.com/clean-code-projects/co-coders-api/criterion"
	"github.com/clean-code-projects/co-coders-api/user"
)


// OnCoStyle ..
func OnCoStyle(users []user.User, coStyle criterion.CoStyle) (matchedUser []user.User) {
	matchedUser = []user.User{}
	for _, u := range users {
		if u.HasCoStyle(coStyle) {
			matchedUser = append(matchedUser, u)
		}
	}
	return matchedUser
}
