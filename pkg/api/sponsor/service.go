package sponsor

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/labstack/echo"
	"github.com/zaynkorai/takrib/pkg/api/sponsor/platform/pgsql"
	takrib "github.com/zaynkorai/takrib/pkg/utl/model"
)

// Service represents sponsor application interface
type Service interface {
	Create(echo.Context, takrib.Sponsor) (*takrib.Sponsor, error)
	List(echo.Context, *takrib.Pagination) ([]takrib.Sponsor, error)
	View(echo.Context, int) (*takrib.Sponsor, error)
	Delete(echo.Context, int) error
	Update(echo.Context, *Update) (*takrib.Sponsor, error)
}

// New creates new sponsor application service
func New(db *pg.DB, udb UDB, rbac RBAC, sec Securer) *Sponsor {
	return &Sponsor{db: db, udb: udb, rbac: rbac, sec: sec}
}

// Initialize initalizes sponsor application service with defaults
func Initialize(db *pg.DB, rbac RBAC, sec Securer) *Sponsor {
	return New(db, pgsql.NewSponsor(), rbac, sec)
}

// Sponsor represents sponsor application service
type Sponsor struct {
	db   *pg.DB
	udb  UDB
	rbac RBAC
	sec  Securer
}

// Securer represents security interface
type Securer interface {
	Hash(string) string
}

// UDB represents sponsor repository interface
type UDB interface {
	Create(orm.DB, takrib.Sponsor) (*takrib.Sponsor, error)
	View(orm.DB, int) (*takrib.Sponsor, error)
	List(orm.DB, *takrib.ListQuery, *takrib.Pagination) ([]takrib.Sponsor, error)
	Update(orm.DB, *takrib.Sponsor) error
	Delete(orm.DB, *takrib.Sponsor) error
}

// RBAC represents role-based-access-control interface
type RBAC interface {
	User(echo.Context) *takrib.AuthUser
	EnforceUser(echo.Context, int) error
	IsLowerRole(echo.Context, takrib.AccessRole) error
}
