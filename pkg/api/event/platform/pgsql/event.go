package pgsql

import (
	"net/http"
	"strings"

	"github.com/go-pg/pg"

	"github.com/go-pg/pg/orm"
	"github.com/labstack/echo"
	"github.com/zaynkorai/takrib/pkg/utl/model"
)

// NewEvent returns a new event database instance
func NewEvent() *Event {
	return &Event{}
}

// Event represents the client for event table
type Event struct{}

// Custom errors
var (
	ErrAlreadyExists = echo.NewHTTPError(http.StatusInternalServerError, "Eventname already exists.")
)

// Create creates a new event on database
func (e *Event) Create(db orm.DB, usr takrib.Event) (*takrib.Event, error) {
	var event = new(takrib.Event)
	err := db.Model(event).Where("lower(eventname) = ? and deleted_at is null",
		strings.ToLower(usr.Eventname)).Select()

	if err != nil && err != pg.ErrNoRows {
		return nil, ErrAlreadyExists

	}

	if err := db.Insert(&usr); err != nil {
		return nil, err
	}
	return &usr, nil
}

// View returns single event by ID
func (e *Event) View(db orm.DB, id int) (*takrib.Event, error) {
	var event = new(takrib.Event)
	sql := `SELECT "event".* 
	FROM "events" AS "event"
	WHERE ("event"."id" = ? and deleted_at is null)`
	_, err := db.QueryOne(event, sql, id)
	if err != nil { 
		return nil, err
	}
	return event, nil
}

// Update updates event's contact info
func (e *Event) Update(db orm.DB, event *takrib.Event) error {
	return db.Update(event)
}

// List returns list of all events retrievable for the current event, depending on role
func (e *Event) List(db orm.DB, qp *takrib.ListQuery, p *takrib.Pagination) ([]takrib.Event, error) {
	var events []takrib.Event
	q := db.Model(&events).Column("event.*").Limit(p.Limit).Offset(p.Offset).Where("deleted_at is null").Order("event.id desc")
	if qp != nil {
		q.Where(qp.Query, qp.ID)
	}
	if err := q.Select(); err != nil {
		return nil, err
	}
	return events, nil
}

// Delete sets deleted_at for a event
func (e *Event) Delete(db orm.DB, event *takrib.Event) error {
	return db.Delete(event)
}
