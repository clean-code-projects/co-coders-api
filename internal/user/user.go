package user

import "github.com/clean-code-projects/co-coders-api/internal/criteria"

// User ..
type User struct {
	CoStyles []criteria.CoStyle
}

// HasCoStyle ..
func (u User) HasCoStyle(coStyle criteria.CoStyle) bool {
	for _, style := range u.CoStyles {
		if style == coStyle {
			return true
		}
	}
	return false
}

// New ..
func New(coStyles []criteria.CoStyle) User {
	return User{CoStyles:coStyles}
}
