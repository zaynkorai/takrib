package mock

import (
	"github.com/zaynkorai/takrib/pkg/utl/model"
)

// JWT mock
type JWT struct {
	GenerateTokenFn func(*takrib.User) (string, string, error)
}

// GenerateToken mock
func (j *JWT) GenerateToken(u *takrib.User) (string, string, error) {
	return j.GenerateTokenFn(u)
}
