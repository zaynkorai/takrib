// Package speaker contains speaker application services
package speaker

import (
	"github.com/labstack/echo"
	takrib "github.com/zaynkorai/takrib/pkg/utl/model"
	"github.com/zaynkorai/takrib/pkg/utl/query"
	"github.com/zaynkorai/takrib/pkg/utl/structs"
)

// Create creates a new speaker account
func (e *Speaker) Create(c echo.Context, req takrib.Speaker) (*takrib.Speaker, error) {
	return e.udb.Create(e.db, req)
}

// List returns list of speakers
func (e *Speaker) List(c echo.Context, p *takrib.Pagination) ([]takrib.Speaker, error) {
	au := e.rbac.User(c)
	q, err := query.List(au)
	if err != nil {
		return nil, err
	}
	return e.udb.List(e.db, q, p)
}

// View returns single speaker
func (e *Speaker) View(c echo.Context, id int) (*takrib.Speaker, error) {
	if err := e.rbac.EnforceUser(c, id); err != nil {
		return nil, err
	}
	return e.udb.View(e.db, id)
}

// Delete deletes a speaker
func (e *Speaker) Delete(c echo.Context, id int) error {
	speaker, err := e.udb.View(e.db, id)
	if err != nil {
		return err
	}

	return e.udb.Delete(e.db, speaker)
}

// Update contains speaker's information used for updating
type Update struct {
	ID                int
	Name              *string
	ShortBiography    *string
	LongBiography     *string
	Gender            *string
	Email             *string
	Mobile            *string
	Website           *string
	Twitter           *string
	Github            *string
	Linkedin          *string
	Organisation      *string
	Position          *string
	Country           *string
	City              *string
	PhotoURL          *string
	ThumbnailImageURL *string
}

// Update updates speaker's contact information
func (e *Speaker) Update(c echo.Context, req *Update) (*takrib.Speaker, error) {

	speaker, err := e.udb.View(e.db, req.ID)
	if err != nil {
		return nil, err
	}

	structs.Merge(speaker, req)
	if err := e.udb.Update(e.db, speaker); err != nil {
		return nil, err
	}

	return speaker, nil
}
