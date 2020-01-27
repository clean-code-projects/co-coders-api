package user

import "github.com/clean-code-projects/co-coders-api/internal/co_coders/criteria"

// User ..
type User struct {
	Criteria criteria.Criteria
}

// New ..
func New() User {
	return User{criteria.New()}
}
