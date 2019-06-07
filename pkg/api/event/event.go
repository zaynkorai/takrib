// Package event contains event application services
package event

import (
	"github.com/labstack/echo"
	"github.com/zaynkorai/takrib/pkg/utl/model"
	"github.com/zaynkorai/takrib/pkg/utl/query"
	"github.com/zaynkorai/takrib/pkg/utl/structs"
)

// Create creates a new event account
func (e *Event) Create(c echo.Context, req takrib.Event) (*takrib.Event, error) {
	return e.udb.Create(e.db, req)
}

// List returns list of events
func (e *Event) List(c echo.Context, p *takrib.Pagination) ([]takrib.Event, error) {
	au := e.rbac.User(c)
	q, err := query.List(au)
	if err != nil {
		return nil, err
	}
	return e.udb.List(e.db, q, p)
}

// View returns single event
func (e *Event) View(c echo.Context, id int) (*takrib.Event, error) {
	if err := e.rbac.EnforceUser(c, id); err != nil {
		return nil, err
	}
	return e.udb.View(e.db, id)
}

// Delete deletes a event
func (e *Event) Delete(c echo.Context, id int) error {
	event, err := e.udb.View(e.db, id)
	if err != nil {
		return err
	}
	// if err := e.rbac.IsLowerRole(c, event.Role.AccessLevel); err != nil {
	// 	return err
	// }
	return e.udb.Delete(e.db, event)
}

// Update contains event's information used for updating
type Update struct {
	ID        int
	Eventname *string
	Location  *string
}

// Update updates event's contact information
func (e *Event) Update(c echo.Context, req *Update) (*takrib.Event, error) {
	// if err := e.rbac.EnforceEvent(c, req.ID); err != nil {
	// 	return nil, err
	// }

	event, err := e.udb.View(e.db, req.ID)
	if err != nil {
		return nil, err
	}

	structs.Merge(event, req)
	if err := e.udb.Update(e.db, event); err != nil {
		return nil, err
	}

	return event, nil
}
