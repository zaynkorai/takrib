package mockdb

import (
	"github.com/go-pg/pg/orm"
	"github.com/zaynkorai/takrib/pkg/utl/model"
)

// User database mock
type User struct {
	CreateFn         func(orm.DB, takrib.User) (*takrib.User, error)
	ViewFn           func(orm.DB, int) (*takrib.User, error)
	FindByUsernameFn func(orm.DB, string) (*takrib.User, error)
	FindByTokenFn    func(orm.DB, string) (*takrib.User, error)
	ListFn           func(orm.DB, *takrib.ListQuery, *takrib.Pagination) ([]takrib.User, error)
	DeleteFn         func(orm.DB, *takrib.User) error
	UpdateFn         func(orm.DB, *takrib.User) error
}

// Create mock
func (u *User) Create(db orm.DB, usr takrib.User) (*takrib.User, error) {
	return u.CreateFn(db, usr)
}

// View mock
func (u *User) View(db orm.DB, id int) (*takrib.User, error) {
	return u.ViewFn(db, id)
}

// FindByUsername mock
func (u *User) FindByUsername(db orm.DB, uname string) (*takrib.User, error) {
	return u.FindByUsernameFn(db, uname)
}

// FindByToken mock
func (u *User) FindByToken(db orm.DB, token string) (*takrib.User, error) {
	return u.FindByTokenFn(db, token)
}

// List mock
func (u *User) List(db orm.DB, lq *takrib.ListQuery, p *takrib.Pagination) ([]takrib.User, error) {
	return u.ListFn(db, lq, p)
}

// Delete mock
func (u *User) Delete(db orm.DB, usr *takrib.User) error {
	return u.DeleteFn(db, usr)
}

// Update mock
func (u *User) Update(db orm.DB, usr *takrib.User) error {
	return u.UpdateFn(db, usr)
}
