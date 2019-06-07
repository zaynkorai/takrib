package pgsql

import (
	"net/http"
	"strings"

	"github.com/go-pg/pg"

	"github.com/go-pg/pg/orm"
	"github.com/labstack/echo"
	takrib "github.com/zaynkorai/takrib/pkg/utl/model"
)

// NewSpeaker returns a new speaker database instance
func NewSpeaker() *Speaker {
	return &Speaker{}
}

// Speaker represents the client for speaker table
type Speaker struct{}

// Custom errors
var (
	ErrAlreadyExists = echo.NewHTTPError(http.StatusInternalServerError, "Speaker name already exists.")
)

// Create creates a new speaker on database
func (e *Speaker) Create(db orm.DB, usr takrib.Speaker) (*takrib.Speaker, error) {
	var speaker = new(takrib.Speaker)
	err := db.Model(speaker).Where("lower(name) = ? and deleted_at is null",
		strings.ToLower(usr.Name)).Select()

	if err != nil && err != pg.ErrNoRows {
		return nil, ErrAlreadyExists

	}

	if err := db.Insert(&usr); err != nil {
		return nil, err
	}
	return &usr, nil
}

// View returns single speaker by ID
func (e *Speaker) View(db orm.DB, id int) (*takrib.Speaker, error) {
	var speaker = new(takrib.Speaker)
	sql := `SELECT "speaker".* 
	FROM "speakers" AS "speaker"
	WHERE ("speaker"."id" = ? and deleted_at is null)`
	_, err := db.QueryOne(speaker, sql, id)
	if err != nil {
		return nil, err
	}
	return speaker, nil
}

// Update updates speaker's contact info
func (e *Speaker) Update(db orm.DB, speaker *takrib.Speaker) error {
	return db.Update(speaker)
}

// List returns list of all speakers retrievable for the current speaker, depending on role
func (e *Speaker) List(db orm.DB, qp *takrib.ListQuery, p *takrib.Pagination) ([]takrib.Speaker, error) {
	var speakers []takrib.Speaker
	q := db.Model(&speakers).Column("speaker.*").Limit(p.Limit).Offset(p.Offset).Where("deleted_at is null").Order("speaker.id desc")
	if qp != nil {
		q.Where(qp.Query, qp.ID)
	}
	if err := q.Select(); err != nil {
		return nil, err
	}
	return speakers, nil
}

// Delete sets deleted_at for a speaker
func (e *Speaker) Delete(db orm.DB, speaker *takrib.Speaker) error {
	return db.Delete(speaker)
}
