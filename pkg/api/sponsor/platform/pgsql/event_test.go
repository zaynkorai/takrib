package pgsql_test

// import (
// 	"testing"

// 	"github.com/zaynkorai/takrib/pkg/utl/model"

// 	"github.com/zaynkorai/takrib/pkg/api/sponsor/platform/pgsql"
// 	"github.com/zaynkorai/takrib/pkg/utl/mock"
// 	"github.com/stretchr/testify/assert"
// )

// func TestCreate(t *testing.T) {
// 	cases := []struct {
// 		name     string
// 		wantErr  bool
// 		req      takrib.Sponsor
// 		wantData *takrib.Sponsor
// 	}{
// 		{
// 			name:    "Sponsor already exists",
// 			wantErr: true,
// 			req: takrib.Sponsor{
// 				Email:    "johndoe@mail.com",
// 				Sponsorname: "johndoe",
// 			},
// 		},
// 		{
// 			name:    "Fail on insert duplicate ID",
// 			wantErr: true,
// 			req: takrib.Sponsor{
// 				Sponsorname:   "tomjones",
// 				Base: takrib.Base{
// 					ID: 1,
// 				},
// 			},
// 		},
// 		{
// 			name: "Success",
// 			req: takrib.Sponsor{
// 				Sponsorname:   "newtomjones",
// 				Base: takrib.Base{
// 					ID: 2,
// 				},
// 			},
// 			wantData: &takrib.Sponsor{
// 				Sponsorname:   "newtomjones",
// 				Base: takrib.Base{
// 					ID: 2,
// 				},
// 			},
// 		},
// 	}

// 	dbCon := mock.NewPGContainer(t)
// 	defer dbCon.Shutdown()

// 	db := mock.NewDB(t, dbCon, &takrib.Role{}, &takrib.Sponsor{})

// 	if err := mock.InsertMultiple(db, &takrib.Role{
// 		ID:          1,
// 		AccessLevel: 1,
// 		Name:        "SUPER_ADMIN"}, &cases[1].req); err != nil {
// 		t.Error(err)
// 	}

// 	udb := pgsql.NewSponsor()

// 	for _, tt := range cases {
// 		t.Run(tt.name, func(t *testing.T) {
// 			resp, err := udb.Create(db, tt.req)
// 			assert.Equal(t, tt.wantErr, err != nil)
// 			if tt.wantData != nil {
// 				if resp == nil {
// 					t.Error("Expected data, but received nil.")
// 					return
// 				}
// 				assert.Equal(t, tt.wantData, resp)
// 			}
// 		})
// 	}
// }

// func TestView(t *testing.T) {
// 	cases := []struct {
// 		name     string
// 		wantErr  bool
// 		id       int
// 		wantData *takrib.Sponsor
// 	}{
// 		{
// 			name:    "Sponsor does not exist",
// 			wantErr: true,
// 			id:      1000,
// 		},
// 		{
// 			name: "Success",
// 			id:   2,
// 			wantData: &takrib.Sponsor{
// 				Sponsorname:   "tomjones",
// 				Base: takrib.Base{
// 					ID: 2,
// 				},
// 				Role: &takrib.Role{
// 					ID:          1,
// 					AccessLevel: 1,
// 					Name:        "SUPER_ADMIN",
// 				},
// 			},
// 		},
// 	}

// 	dbCon := mock.NewPGContainer(t)
// 	defer dbCon.Shutdown()

// 	db := mock.NewDB(t, dbCon, &takrib.Role{}, &takrib.Sponsor{})

// 	if err := mock.InsertMultiple(db, &takrib.Role{
// 		ID:          1,
// 		AccessLevel: 1,
// 		Name:        "SUPER_ADMIN"}, cases[1].wantData); err != nil {
// 		t.Error(err)
// 	}

// 	udb := pgsql.NewSponsor()

// 	for _, tt := range cases {
// 		t.Run(tt.name, func(t *testing.T) {
// 			sponsor, err := udb.View(db, tt.id)
// 			assert.Equal(t, tt.wantErr, err != nil)
// 			if tt.wantData != nil {
// 				if sponsor == nil {
// 					t.Errorf("response was nil due to: %v", err)
// 				} else {
// 					assert.Equal(t, tt.wantData, sponsor)
// 				}
// 			}
// 		})
// 	}
// }

// func TestUpdate(t *testing.T) {
// 	cases := []struct {
// 		name     string
// 		wantErr  bool
// 		usr      *takrib.Sponsor
// 		wantData *takrib.Sponsor
// 	}{
// 		{
// 			name: "Success",
// 			usr: &takrib.Sponsor{
// 				Base: takrib.Base{
// 					ID: 2,
// 				},
// 				Sponsorname:  "newSponsorname",
// 			},
// 			wantData: &takrib.Sponsor{
// 				Sponsorname:   "tomjones",
// 				Base: takrib.Base{
// 					ID: 2,
// 				},
// 			},
// 		},
// 	}

// 	dbCon := mock.NewPGContainer(t)
// 	defer dbCon.Shutdown()

// 	db := mock.NewDB(t, dbCon, &takrib.Role{}, &takrib.Sponsor{})

// 	if err := mock.InsertMultiple(db, &takrib.Role{
// 		ID:          1,
// 		AccessLevel: 1,
// 		Name:        "SUPER_ADMIN"}, cases[0].usr); err != nil {
// 		t.Error(err)
// 	}

// 	udb := pgsql.NewSponsor()

// 	for _, tt := range cases {
// 		t.Run(tt.name, func(t *testing.T) {
// 			err := udb.Update(db, tt.wantData)
// 			assert.Equal(t, tt.wantErr, err != nil)
// 			if tt.wantData != nil {
// 				sponsor := &takrib.Sponsor{
// 					Base: takrib.Base{
// 						ID: tt.usr.ID,
// 					},
// 				}
// 				if err := db.Select(sponsor); err != nil {
// 					t.Error(err)
// 				}
// 				tt.wantData.UpdatedAt = sponsor.UpdatedAt
// 				tt.wantData.CreatedAt = sponsor.CreatedAt
// 				tt.wantData.LastLogin = sponsor.LastLogin
// 				tt.wantData.DeletedAt = sponsor.DeletedAt
// 				assert.Equal(t, tt.wantData, sponsor)
// 			}
// 		})
// 	}
// }

// func TestList(t *testing.T) {
// 	cases := []struct {
// 		name     string
// 		wantErr  bool
// 		qp       *takrib.ListQuery
// 		pg       *takrib.Pagination
// 		wantData []takrib.Sponsor
// 	}{
// 		{
// 			name:    "Invalid pagination values",
// 			wantErr: true,
// 			pg: &takrib.Pagination{
// 				Limit: -100,
// 			},
// 		},
// 		{
// 			name: "Success",
// 			pg: &takrib.Pagination{
// 				Limit:  100,
// 				Offset: 0,
// 			},
// 			qp: &takrib.ListQuery{
// 				ID:    1,
// 				Query: "company_id = ?",
// 			},
// 			wantData: []takrib.Sponsor{
// 				{
// 					Email:      "tomjones@mail.com",
// 					FirstName:  "Tom",
// 					LastName:   "Jones",
// 					Sponsorname:   "tomjones",
// 					RoleID:     1,
// 					CompanyID:  1,
// 					LocationID: 1,
// 					Password:   "newPass",
// 					Base: takrib.Base{
// 						ID: 2,
// 					},
// 					Role: &takrib.Role{
// 						ID:          1,
// 						AccessLevel: 1,
// 						Name:        "SUPER_ADMIN",
// 					},
// 				},
// 				{
// 					Email:      "johndoe@mail.com",
// 					FirstName:  "John",
// 					LastName:   "Doe",
// 					Sponsorname:   "johndoe",
// 					RoleID:     1,
// 					CompanyID:  1,
// 					LocationID: 1,
// 					Password:   "hunter2",
// 					Base: takrib.Base{
// 						ID: 1,
// 					},
// 					Role: &takrib.Role{
// 						ID:          1,
// 						AccessLevel: 1,
// 						Name:        "SUPER_ADMIN",
// 					},
// 					Token: "loginrefresh",
// 				},
// 			},
// 		},
// 	}

// 	dbCon := mock.NewPGContainer(t)
// 	defer dbCon.Shutdown()

// 	db := mock.NewDB(t, dbCon, &takrib.Role{}, &takrib.Sponsor{})

// 	if err := mock.InsertMultiple(db, &takrib.Role{
// 		ID:          1,
// 		AccessLevel: 1,
// 		Name:        "SUPER_ADMIN"}, &cases[1].wantData); err != nil {
// 		t.Error(err)
// 	}

// 	udb := pgsql.NewSponsor()

// 	for _, tt := range cases {
// 		t.Run(tt.name, func(t *testing.T) {
// 			sponsors, err := udb.List(db, tt.qp, tt.pg)
// 			assert.Equal(t, tt.wantErr, err != nil)
// 			if tt.wantData != nil {
// 				for i, v := range sponsors {
// 					tt.wantData[i].CreatedAt = v.CreatedAt
// 					tt.wantData[i].UpdatedAt = v.UpdatedAt
// 				}
// 				assert.Equal(t, tt.wantData, sponsors)
// 			}
// 		})
// 	}
// }

// func TestDelete(t *testing.T) {
// 	cases := []struct {
// 		name     string
// 		wantErr  bool
// 		usr      *takrib.Sponsor
// 		wantData *takrib.Sponsor
// 	}{
// 		{
// 			name: "Success",
// 			usr: &takrib.Sponsor{
// 				Base: takrib.Base{
// 					ID:        2,
// 					DeletedAt: mock.TestTime(2018),
// 				},
// 			},
// 			wantData: &takrib.Sponsor{
// 				Email:      "tomjones@mail.com",
// 				FirstName:  "Tom",
// 				LastName:   "Jones",
// 				Sponsorname:   "tomjones",
// 				RoleID:     1,
// 				CompanyID:  1,
// 				LocationID: 1,
// 				Password:   "newPass",
// 				Base: takrib.Base{
// 					ID: 2,
// 				},
// 			},
// 		},
// 	}

// 	dbCon := mock.NewPGContainer(t)
// 	defer dbCon.Shutdown()

// 	db := mock.NewDB(t, dbCon, &takrib.Role{}, &takrib.Sponsor{})

// 	if err := mock.InsertMultiple(db, &takrib.Role{
// 		ID:          1,
// 		AccessLevel: 1,
// 		Name:        "SUPER_ADMIN"}, cases[0].wantData); err != nil {
// 		t.Error(err)
// 	}

// 	udb := pgsql.NewSponsor()

// 	for _, tt := range cases {
// 		t.Run(tt.name, func(t *testing.T) {

// 			err := udb.Delete(db, tt.usr)
// 			assert.Equal(t, tt.wantErr, err != nil)

// 			// Check if the deleted_at was set
// 		})
// 	}
// }
