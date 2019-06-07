package speaker

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/labstack/echo"
	"github.com/zaynkorai/takrib/pkg/api/speaker/platform/pgsql"
	takrib "github.com/zaynkorai/takrib/pkg/utl/model"
)

// Service represents speaker application interface
type Service interface {
	Create(echo.Context, takrib.Speaker) (*takrib.Speaker, error)
	List(echo.Context, *takrib.Pagination) ([]takrib.Speaker, error)
	View(echo.Context, int) (*takrib.Speaker, error)
	Delete(echo.Context, int) error
	Update(echo.Context, *Update) (*takrib.Speaker, error)
}

// New creates new speaker application service
func New(db *pg.DB, udb UDB, rbac RBAC, sec Securer) *Speaker {
	return &Speaker{db: db, udb: udb, rbac: rbac, sec: sec}
}

// Initialize initalizes speaker application service with defaults
func Initialize(db *pg.DB, rbac RBAC, sec Securer) *Speaker {
	return New(db, pgsql.NewSpeaker(), rbac, sec)
}

// Speaker represents speaker application service
type Speaker struct {
	db   *pg.DB
	udb  UDB
	rbac RBAC
	sec  Securer
}

// Securer represents security interface
type Securer interface {
	Hash(string) string
}

// UDB represents speaker repository interface
type UDB interface {
	Create(orm.DB, takrib.Speaker) (*takrib.Speaker, error)
	View(orm.DB, int) (*takrib.Speaker, error)
	List(orm.DB, *takrib.ListQuery, *takrib.Pagination) ([]takrib.Speaker, error)
	Update(orm.DB, *takrib.Speaker) error
	Delete(orm.DB, *takrib.Speaker) error
}

// RBAC represents role-based-access-control interface
type RBAC interface {
	User(echo.Context) *takrib.AuthUser
	EnforceUser(echo.Context, int) error
	IsLowerRole(echo.Context, takrib.AccessRole) error
}
