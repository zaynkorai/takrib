package transport

import (
	"github.com/zaynkorai/takrib/pkg/utl/model"
)

// User model response
// swagger:response userResp
type swaggUserResponse struct {
	// in:body
	Body struct {
		*takrib.User
	}
}

// Users model response
// swagger:response userListResp
type swaggUserListResponse struct {
	// in:body
	Body struct {
		Users []takrib.User `json:"users"`
		Page  int          `json:"page"`
	}
}
