package query_test

import (
	"testing"

	"github.com/zaynkorai/takrib/pkg/utl/model"

	"github.com/labstack/echo"

	"github.com/zaynkorai/takrib/pkg/utl/query"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	type args struct {
		user *takrib.AuthUser
	}
	cases := []struct {
		name     string
		args     args
		wantData *takrib.ListQuery
		wantErr  error
	}{
		{
			name: "Super admin user",
			args: args{user: &takrib.AuthUser{
				Role: takrib.SuperAdminRole,
			}},
		},
		{
			name: "Company admin user",
			args: args{user: &takrib.AuthUser{
				Role:      takrib.CompanyAdminRole,
				CompanyID: 1,
			}},
			wantData: &takrib.ListQuery{
				Query: "company_id = ?",
				ID:    1},
		},
		{
			name: "Location admin user",
			args: args{user: &takrib.AuthUser{
				Role:       takrib.LocationAdminRole,
				CompanyID:  1,
				LocationID: 2,
			}},
			wantData: &takrib.ListQuery{
				Query: "location_id = ?",
				ID:    2},
		},
		{
			name: "Normal user",
			args: args{user: &takrib.AuthUser{
				Role: takrib.UserRole,
			}},
			wantErr: echo.ErrForbidden,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			q, err := query.List(tt.args.user)
			assert.Equal(t, tt.wantData, q)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
