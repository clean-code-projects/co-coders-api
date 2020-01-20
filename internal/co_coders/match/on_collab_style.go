package match

import (
	"github.com/clean-code-projects/co-coders-api/internal/co_coders/criteria"
	"github.com/clean-code-projects/co-coders-api/internal/co_coders/user"
)

// OnCollabStyle ..
func OnCollabStyle(users []user.User, collabStyle criteria.CollabStyle) (matchedUsers []user.User) {
	matchedUsers = []user.User{}
	for _, aUser := range users {
		if aUser.HasCollabStyle(collabStyle) {
			matchedUsers = append(matchedUsers, aUser)
		}
	}
	return matchedUsers
}
