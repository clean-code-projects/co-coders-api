package filter

import (
	"github.com/clean-code-projects/co-coders-api/internal/criteria"
	"github.com/clean-code-projects/co-coders-api/internal/user"
)

// OnCollabStyle ..
func OnCollabStyle(users []user.User, coStyle criteria.CollabStyle) (matchedUsers []user.User) {
	matchedUsers = []user.User{}
	for _, aUser := range users {
		if aUser.HasCollabStyle(coStyle) {
			matchedUsers = append(matchedUsers, aUser)
		}
	}
	return matchedUsers
}
