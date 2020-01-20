package user

import "github.com/clean-code-projects/co-coders-api/internal/co_coders/criteria"

// User ..
type User struct {
	CollabStyles []criteria.CollabStyle
}

// HasCollabStyle ..
func (u User) HasCollabStyle(collabStyle criteria.CollabStyle) bool {
	for _, style := range u.CollabStyles {
		if style == collabStyle {
			return true
		}
	}
	return false
}

// New ..
func New(collabStyles []criteria.CollabStyle) User {
	return User{CollabStyles: collabStyles}
}
