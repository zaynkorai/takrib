package pgsql_test

// import (
// 	"testing"

// 	"github.com/zaynkorai/takrib/pkg/utl/model"

// 	"github.com/zaynkorai/takrib/pkg/api/event/platform/pgsql"
// 	"github.com/zaynkorai/takrib/pkg/utl/mock"
// 	"github.com/stretchr/testify/assert"
// )

// func TestCreate(t *testing.T) {
// 	cases := []struct {
// 		name     string
// 		wantErr  bool
// 		req      takrib.Event
// 		wantData *takrib.Event
// 	}{
// 		{
// 			name:    "Event already exists",
// 			wantErr: true,
// 			req: takrib.Event{
// 				Email:    "johndoe@mail.com",
// 				Eventname: "johndoe",
// 			},
// 		},
// 		{
// 			name:    "Fail on insert duplicate ID",
// 			wantErr: true,
// 			req: takrib.Event{
// 				Eventname:   "tomjones",
// 				Base: takrib.Base{
// 					ID: 1,
// 				},
// 			},
// 		},
// 		{
// 			name: "Success",
// 			req: takrib.Event{
// 				Eventname:   "newtomjones",
// 				Base: takrib.Base{
// 					ID: 2,
// 				},
// 			},
// 			wantData: &takrib.Event{
// 				Eventname:   "newtomjones",
// 				Base: takrib.Base{
// 					ID: 2,
// 				},
// 			},
// 		},
// 	}

// 	dbCon := mock.NewPGContainer(t)
// 	defer dbCon.Shutdown()

// 	db := mock.NewDB(t, dbCon, &takrib.Role{}, &takrib.Event{})

// 	if err := mock.InsertMultiple(db, &takrib.Role{
// 		ID:          1,
// 		AccessLevel: 1,
// 		Name:        "SUPER_ADMIN"}, &cases[1].req); err != nil {
// 		t.Error(err)
// 	}

// 	udb := pgsql.NewEvent()

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
// 		wantData *takrib.Event
// 	}{
// 		{
// 			name:    "Event does not exist",
// 			wantErr: true,
// 			id:      1000,
// 		},
// 		{
// 			name: "Success",
// 			id:   2,
// 			wantData: &takrib.Event{
// 				Eventname:   "tomjones",
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

// 	db := mock.NewDB(t, dbCon, &takrib.Role{}, &takrib.Event{})

// 	if err := mock.InsertMultiple(db, &takrib.Role{
// 		ID:          1,
// 		AccessLevel: 1,
// 		Name:        "SUPER_ADMIN"}, cases[1].wantData); err != nil {
// 		t.Error(err)
// 	}

// 	udb := pgsql.NewEvent()

// 	for _, tt := range cases {
// 		t.Run(tt.name, func(t *testing.T) {
// 			event, err := udb.View(db, tt.id)
// 			assert.Equal(t, tt.wantErr, err != nil)
// 			if tt.wantData != nil {
// 				if event == nil {
// 					t.Errorf("response was nil due to: %v", err)
// 				} else {
// 					assert.Equal(t, tt.wantData, event)
// 				}
// 			}
// 		})
// 	}
// }

// func TestUpdate(t *testing.T) {
// 	cases := []struct {
// 		name     string
// 		wantErr  bool
// 		usr      *takrib.Event
// 		wantData *takrib.Event
// 	}{
// 		{
// 			name: "Success",
// 			usr: &takrib.Event{
// 				Base: takrib.Base{
// 					ID: 2,
// 				},
// 				Eventname:  "newEventname",
// 			},
// 			wantData: &takrib.Event{
// 				Eventname:   "tomjones",
// 				Base: takrib.Base{
// 					ID: 2,
// 				},
// 			},
// 		},
// 	}

// 	dbCon := mock.NewPGContainer(t)
// 	defer dbCon.Shutdown()

// 	db := mock.NewDB(t, dbCon, &takrib.Role{}, &takrib.Event{})

// 	if err := mock.InsertMultiple(db, &takrib.Role{
// 		ID:          1,
// 		AccessLevel: 1,
// 		Name:        "SUPER_ADMIN"}, cases[0].usr); err != nil {
// 		t.Error(err)
// 	}

// 	udb := pgsql.NewEvent()

// 	for _, tt := range cases {
// 		t.Run(tt.name, func(t *testing.T) {
// 			err := udb.Update(db, tt.wantData)
// 			assert.Equal(t, tt.wantErr, err != nil)
// 			if tt.wantData != nil {
// 				event := &takrib.Event{
// 					Base: takrib.Base{
// 						ID: tt.usr.ID,
// 					},
// 				}
// 				if err := db.Select(event); err != nil {
// 					t.Error(err)
// 				}
// 				tt.wantData.UpdatedAt = event.UpdatedAt
// 				tt.wantData.CreatedAt = event.CreatedAt
// 				tt.wantData.LastLogin = event.LastLogin
// 				tt.wantData.DeletedAt = event.DeletedAt
// 				assert.Equal(t, tt.wantData, event)
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
// 		wantData []takrib.Event
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
// 			wantData: []takrib.Event{
// 				{
// 					Email:      "tomjones@mail.com",
// 					FirstName:  "Tom",
// 					LastName:   "Jones",
// 					Eventname:   "tomjones",
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
// 					Eventname:   "johndoe",
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

// 	db := mock.NewDB(t, dbCon, &takrib.Role{}, &takrib.Event{})

// 	if err := mock.InsertMultiple(db, &takrib.Role{
// 		ID:          1,
// 		AccessLevel: 1,
// 		Name:        "SUPER_ADMIN"}, &cases[1].wantData); err != nil {
// 		t.Error(err)
// 	}

// 	udb := pgsql.NewEvent()

// 	for _, tt := range cases {
// 		t.Run(tt.name, func(t *testing.T) {
// 			events, err := udb.List(db, tt.qp, tt.pg)
// 			assert.Equal(t, tt.wantErr, err != nil)
// 			if tt.wantData != nil {
// 				for i, v := range events {
// 					tt.wantData[i].CreatedAt = v.CreatedAt
// 					tt.wantData[i].UpdatedAt = v.UpdatedAt
// 				}
// 				assert.Equal(t, tt.wantData, events)
// 			}
// 		})
// 	}
// }

// func TestDelete(t *testing.T) {
// 	cases := []struct {
// 		name     string
// 		wantErr  bool
// 		usr      *takrib.Event
// 		wantData *takrib.Event
// 	}{
// 		{
// 			name: "Success",
// 			usr: &takrib.Event{
// 				Base: takrib.Base{
// 					ID:        2,
// 					DeletedAt: mock.TestTime(2018),
// 				},
// 			},
// 			wantData: &takrib.Event{
// 				Email:      "tomjones@mail.com",
// 				FirstName:  "Tom",
// 				LastName:   "Jones",
// 				Eventname:   "tomjones",
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

// 	db := mock.NewDB(t, dbCon, &takrib.Role{}, &takrib.Event{})

// 	if err := mock.InsertMultiple(db, &takrib.Role{
// 		ID:          1,
// 		AccessLevel: 1,
// 		Name:        "SUPER_ADMIN"}, cases[0].wantData); err != nil {
// 		t.Error(err)
// 	}

// 	udb := pgsql.NewEvent()

// 	for _, tt := range cases {
// 		t.Run(tt.name, func(t *testing.T) {

// 			err := udb.Delete(db, tt.usr)
// 			assert.Equal(t, tt.wantErr, err != nil)

// 			// Check if the deleted_at was set
// 		})
// 	}
// }
