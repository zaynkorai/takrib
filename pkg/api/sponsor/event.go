// Package sponsor contains Sponsor application services
package sponsor

import (
	"github.com/labstack/echo"
	takrib "github.com/zaynkorai/takrib/pkg/utl/model"
	"github.com/zaynkorai/takrib/pkg/utl/query"
	"github.com/zaynkorai/takrib/pkg/utl/structs"
)

// Create creates a new Sponsor account
func (e *Sponsor) Create(c echo.Context, req takrib.Sponsor) (*takrib.Sponsor, error) {
	return e.udb.Create(e.db, req)
}

// List returns list of Sponsors
func (e *Sponsor) List(c echo.Context, p *takrib.Pagination) ([]takrib.Sponsor, error) {
	au := e.rbac.User(c)
	q, err := query.List(au)
	if err != nil {
		return nil, err
	}
	return e.udb.List(e.db, q, p)
}

// View returns single Sponsor
func (e *Sponsor) View(c echo.Context, id int) (*takrib.Sponsor, error) {
	if err := e.rbac.EnforceUser(c, id); err != nil {
		return nil, err
	}
	return e.udb.View(e.db, id)
}

// Delete deletes a Sponsor
func (e *Sponsor) Delete(c echo.Context, id int) error {
	sponsor, err := e.udb.View(e.db, id)
	if err != nil {
		return err
	}
	// if err := e.rbac.IsLowerRole(c, sponsor.Role.AccessLevel); err != nil {
	// 	return err
	// }
	return e.udb.Delete(e.db, sponsor)
}

// Update contains sponsor's information used for updating
type Update struct {
	ID       int
	Name     *string
	Location *string
}

// Update updates sponsor's contact information
func (e *Sponsor) Update(c echo.Context, req *Update) (*takrib.Sponsor, error) {
	// if err := e.rbac.EnforceSponsor(c, req.ID); err != nil {
	// 	return nil, err
	// }

	sponsor, err := e.udb.View(e.db, req.ID)
	if err != nil {
		return nil, err
	}

	structs.Merge(sponsor, req)
	if err := e.udb.Update(e.db, sponsor); err != nil {
		return nil, err
	}

	return sponsor, nil
}
