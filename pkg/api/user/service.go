package user

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/labstack/echo"
	"github.com/zaynkorai/takrib/pkg/api/user/platform/pgsql"
	"github.com/zaynkorai/takrib/pkg/utl/model"
)

// Service represents user application interface
type Service interface {
	Create(echo.Context, takrib.User) (*takrib.User, error)
	List(echo.Context, *takrib.Pagination) ([]takrib.User, error)
	View(echo.Context, int) (*takrib.User, error)
	Delete(echo.Context, int) error
	Update(echo.Context, *Update) (*takrib.User, error)
}

// New creates new user application service
func New(db *pg.DB, udb UDB, rbac RBAC, sec Securer) *User {
	return &User{db: db, udb: udb, rbac: rbac, sec: sec}
}

// Initialize initalizes User application service with defaults
func Initialize(db *pg.DB, rbac RBAC, sec Securer) *User {
	return New(db, pgsql.NewUser(), rbac, sec)
}

// User represents user application service
type User struct {
	db   *pg.DB
	udb  UDB
	rbac RBAC
	sec  Securer
}

// Securer represents security interface
type Securer interface {
	Hash(string) string
}

// UDB represents user repository interface
type UDB interface {
	Create(orm.DB, takrib.User) (*takrib.User, error)
	View(orm.DB, int) (*takrib.User, error)
	List(orm.DB, *takrib.ListQuery, *takrib.Pagination) ([]takrib.User, error)
	Update(orm.DB, *takrib.User) error
	Delete(orm.DB, *takrib.User) error
}

// RBAC represents role-based-access-control interface
type RBAC interface {
	User(echo.Context) *takrib.AuthUser
	EnforceUser(echo.Context, int) error
	AccountCreate(echo.Context, takrib.AccessRole, int, int) error
	IsLowerRole(echo.Context, takrib.AccessRole) error
}
