package query

import (
	"github.com/labstack/echo"
	"github.com/zaynkorai/takrib/pkg/utl/model"
)

// List prepares data for list queries
func List(u *takrib.AuthUser) (*takrib.ListQuery, error) {
	switch true {
	case u.Role <= takrib.AdminRole: // user is SuperAdmin or Admin
		return nil, nil
	case u.Role == takrib.CompanyAdminRole:
		return &takrib.ListQuery{Query: "company_id = ?", ID: u.CompanyID}, nil
	case u.Role == takrib.LocationAdminRole:
		return &takrib.ListQuery{Query: "location_id = ?", ID: u.LocationID}, nil
	default:
		return nil, echo.ErrForbidden
	}
}
