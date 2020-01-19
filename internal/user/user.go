package user

import "github.com/clean-code-projects/co-coders-api/internal/criterion"


// User ..
type User struct {
	CoStyles []criterion.CoStyle
}

// HasCoStyle ..
func (u User) HasCoStyle(coStyle criterion.CoStyle) bool {
	for _, style := range u.CoStyles {
		if style == coStyle {
			return true
		}
	}
	return false
}

func New(coStyles []criterion.CoStyle) User {
	return User{CoStyles:coStyles}
}
