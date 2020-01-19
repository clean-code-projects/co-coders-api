package filter

import (
	"github.com/clean-code-projects/co-coders-api/internal/criteria"
	"github.com/clean-code-projects/co-coders-api/internal/user"
)

// OnCoStyle ..
func OnCoStyle(users []user.User, coStyle criteria.CoStyle) (matchedUsers []user.User) {
	matchedUsers = []user.User{}
	for _, aUser := range users {
		if aUser.HasCoStyle(coStyle) {
			matchedUsers = append(matchedUsers, aUser)
		}
	}
	return matchedUsers
}
