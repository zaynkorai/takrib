package event

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/labstack/echo"
	"github.com/zaynkorai/takrib/pkg/api/event/platform/pgsql"
	"github.com/zaynkorai/takrib/pkg/utl/model"
)

// Service represents event application interface
type Service interface {
	Create(echo.Context, takrib.Event) (*takrib.Event, error)
	List(echo.Context, *takrib.Pagination) ([]takrib.Event, error)
	View(echo.Context, int) (*takrib.Event, error)
	Delete(echo.Context, int) error
	Update(echo.Context, *Update) (*takrib.Event, error)
}

// New creates new event application service
func New(db *pg.DB, udb UDB, rbac RBAC, sec Securer) *Event {
	return &Event{db: db, udb: udb, rbac: rbac, sec: sec}
}

// Initialize initalizes Event application service with defaults
func Initialize(db *pg.DB, rbac RBAC, sec Securer) *Event {
	return New(db, pgsql.NewEvent(), rbac, sec)
}

// Event represents event application service
type Event struct {
	db   *pg.DB
	udb  UDB
	rbac RBAC
	sec  Securer
}

// Securer represents security interface
type Securer interface {
	Hash(string) string
}

// UDB represents event repository interface
type UDB interface {
	Create(orm.DB, takrib.Event) (*takrib.Event, error)
	View(orm.DB, int) (*takrib.Event, error)
	List(orm.DB, *takrib.ListQuery, *takrib.Pagination) ([]takrib.Event, error)
	Update(orm.DB, *takrib.Event) error
	Delete(orm.DB, *takrib.Event) error
}


// RBAC represents role-based-access-control interface
type RBAC interface {
	User(echo.Context) *takrib.AuthUser
	EnforceUser(echo.Context, int) error
	IsLowerRole(echo.Context, takrib.AccessRole) error
}
