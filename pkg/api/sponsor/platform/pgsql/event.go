package pgsql

import (
	"net/http"
	"strings"

	"github.com/go-pg/pg"

	"github.com/go-pg/pg/orm"
	"github.com/labstack/echo"
	takrib "github.com/zaynkorai/takrib/pkg/utl/model"
)

// NewSponsor returns a new sponsor database instance
func NewSponsor() *Sponsor {
	return &Sponsor{}
}

// Sponsor represents the client for sponsor table
type Sponsor struct{}

// Custom errors
var (
	ErrAlreadyExists = echo.NewHTTPError(http.StatusInternalServerError, "Sponsor Name already exists.")
)

// Create creates a new sponsor on database
func (e *Sponsor) Create(db orm.DB, usr takrib.Sponsor) (*takrib.Sponsor, error) {
	var sponsor = new(takrib.Sponsor)
	err := db.Model(sponsor).Where("lower(name) = ? and deleted_at is null",
		strings.ToLower(usr.Name)).Select()

	if err != nil && err != pg.ErrNoRows {
		return nil, ErrAlreadyExists

	}

	if err := db.Insert(&usr); err != nil {
		return nil, err
	}
	return &usr, nil
}

// View returns single sponsor by ID
func (e *Sponsor) View(db orm.DB, id int) (*takrib.Sponsor, error) {
	var sponsor = new(takrib.Sponsor)
	sql := `SELECT "sponsor".* 
	FROM "sponsors" AS "sponsor"
	WHERE ("sponsor"."id" = ? and deleted_at is null)`
	_, err := db.QueryOne(sponsor, sql, id)
	if err != nil {
		return nil, err
	}
	return sponsor, nil
}

// Update updates sponsor's contact info
func (e *Sponsor) Update(db orm.DB, sponsor *takrib.Sponsor) error {
	return db.Update(sponsor)
}

// List returns list of all sponsors retrievable for the current sponsor, depending on role
func (e *Sponsor) List(db orm.DB, qp *takrib.ListQuery, p *takrib.Pagination) ([]takrib.Sponsor, error) {
	var sponsors []takrib.Sponsor
	q := db.Model(&sponsors).Column("sponsor.*").Limit(p.Limit).Offset(p.Offset).Where("deleted_at is null").Order("Sponsor.id desc")
	if qp != nil {
		q.Where(qp.Query, qp.ID)
	}
	if err := q.Select(); err != nil {
		return nil, err
	}
	return sponsors, nil
}

// Delete sets deleted_at for a sponsor
func (e *Sponsor) Delete(db orm.DB, sponsor *takrib.Sponsor) error {
	return db.Delete(sponsor)
}
