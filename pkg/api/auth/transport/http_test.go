package transport_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo"

	"github.com/zaynkorai/takrib/pkg/api/auth"
	"github.com/zaynkorai/takrib/pkg/api/auth/transport"
	"github.com/zaynkorai/takrib/pkg/utl/middleware/jwt"
	"github.com/zaynkorai/takrib/pkg/utl/mock"
	"github.com/zaynkorai/takrib/pkg/utl/mock/mockdb"
	"github.com/zaynkorai/takrib/pkg/utl/model"
	"github.com/zaynkorai/takrib/pkg/utl/server"

	"github.com/go-pg/pg/orm"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	cases := []struct {
		name       string
		req        string
		wantStatus int
		wantResp   *takrib.AuthToken
		udb        *mockdb.User
		jwt        *mock.JWT
		sec        *mock.Secure
	}{
		{
			name:       "Invalid request",
			req:        `{"username":"juzernejm"}`,
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "Fail on FindByUsername",
			req:        `{"username":"juzernejm","password":"hunter123"}`,
			wantStatus: http.StatusInternalServerError,
			udb: &mockdb.User{
				FindByUsernameFn: func(orm.DB, string) (*takrib.User, error) {
					return nil, takrib.ErrGeneric
				},
			},
		},
		{
			name:       "Success",
			req:        `{"username":"juzernejm","password":"hunter123"}`,
			wantStatus: http.StatusOK,
			udb: &mockdb.User{
				FindByUsernameFn: func(orm.DB, string) (*takrib.User, error) {
					return &takrib.User{
						Password: "hunter123",
						Active:   true,
					}, nil
				},
				UpdateFn: func(db orm.DB, u *takrib.User) error {
					return nil
				},
			},
			jwt: &mock.JWT{
				GenerateTokenFn: func(*takrib.User) (string, string, error) {
					return "jwttokenstring", mock.TestTime(2018).Format(time.RFC3339), nil
				},
			},
			sec: &mock.Secure{
				HashMatchesPasswordFn: func(string, string) bool {
					return true
				},
				TokenFn: func(string) string {
					return "refreshtoken"
				},
			},
			wantResp: &takrib.AuthToken{Token: "jwttokenstring", Expires: mock.TestTime(2018).Format(time.RFC3339), RefreshToken: "refreshtoken"},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			r := server.New()
			transport.NewHTTP(auth.New(nil, tt.udb, tt.jwt, tt.sec, nil), r, nil)
			ts := httptest.NewServer(r)
			defer ts.Close()
			path := ts.URL + "/login"
			res, err := http.Post(path, "application/json", bytes.NewBufferString(tt.req))
			if err != nil {
				t.Fatal(err)
			}
			defer res.Body.Close()
			if tt.wantResp != nil {
				response := new(takrib.AuthToken)
				if err := json.NewDecoder(res.Body).Decode(response); err != nil {
					t.Fatal(err)
				}
				tt.wantResp.RefreshToken = response.RefreshToken
				assert.Equal(t, tt.wantResp, response)
			}
			assert.Equal(t, tt.wantStatus, res.StatusCode)
		})
	}
}

func TestRefresh(t *testing.T) {
	cases := []struct {
		name       string
		req        string
		wantStatus int
		wantResp   *takrib.RefreshToken
		udb        *mockdb.User
		jwt        *mock.JWT
	}{
		{
			name:       "Fail on FindByToken",
			req:        "refreshtoken",
			wantStatus: http.StatusInternalServerError,
			udb: &mockdb.User{
				FindByTokenFn: func(orm.DB, string) (*takrib.User, error) {
					return nil, takrib.ErrGeneric
				},
			},
		},
		{
			name:       "Success",
			req:        "refreshtoken",
			wantStatus: http.StatusOK,
			udb: &mockdb.User{
				FindByTokenFn: func(orm.DB, string) (*takrib.User, error) {
					return &takrib.User{
						Username: "johndoe",
						Active:   true,
					}, nil
				},
			},
			jwt: &mock.JWT{
				GenerateTokenFn: func(*takrib.User) (string, string, error) {
					return "jwttokenstring", mock.TestTime(2018).Format(time.RFC3339), nil
				},
			},
			wantResp: &takrib.RefreshToken{Token: "jwttokenstring", Expires: mock.TestTime(2018).Format(time.RFC3339)},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			r := server.New()
			transport.NewHTTP(auth.New(nil, tt.udb, tt.jwt, nil, nil), r, nil)
			ts := httptest.NewServer(r)
			defer ts.Close()
			path := ts.URL + "/refresh/" + tt.req
			res, err := http.Get(path)
			if err != nil {
				t.Fatal(err)
			}
			defer res.Body.Close()
			if tt.wantResp != nil {
				response := new(takrib.RefreshToken)
				if err := json.NewDecoder(res.Body).Decode(response); err != nil {
					t.Fatal(err)
				}
				assert.Equal(t, tt.wantResp, response)
			}
			assert.Equal(t, tt.wantStatus, res.StatusCode)
		})
	}
}

func TestMe(t *testing.T) {
	cases := []struct {
		name       string
		wantStatus int
		wantResp   *takrib.User
		header     string
		udb        *mockdb.User
		rbac       *mock.RBAC
	}{
		{
			name:       "Fail on user view",
			wantStatus: http.StatusInternalServerError,
			udb: &mockdb.User{
				ViewFn: func(orm.DB, int) (*takrib.User, error) {
					return nil, takrib.ErrGeneric
				},
			},
			rbac: &mock.RBAC{
				UserFn: func(echo.Context) *takrib.AuthUser {
					return &takrib.AuthUser{ID: 1}
				},
			},
			header: mock.HeaderValid(),
		},
		{
			name:       "Success",
			wantStatus: http.StatusOK,
			udb: &mockdb.User{
				ViewFn: func(db orm.DB, i int) (*takrib.User, error) {
					return &takrib.User{
						Base: takrib.Base{
							ID: i,
						},
						CompanyID:  2,
						LocationID: 3,
						Email:      "john@mail.com",
						FirstName:  "John",
						LastName:   "Doe",
					}, nil
				},
			},
			rbac: &mock.RBAC{
				UserFn: func(echo.Context) *takrib.AuthUser {
					return &takrib.AuthUser{ID: 1}
				},
			},
			header: mock.HeaderValid(),
			wantResp: &takrib.User{
				Base: takrib.Base{
					ID: 1,
				},
				CompanyID:  2,
				LocationID: 3,
				Email:      "john@mail.com",
				FirstName:  "John",
				LastName:   "Doe",
			},
		},
	}

	client := &http.Client{}
	jwtMW := jwt.New("jwtsecret", "HS256", 60)

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			r := server.New()
			transport.NewHTTP(auth.New(nil, tt.udb, nil, nil, tt.rbac), r, jwtMW.MWFunc())
			ts := httptest.NewServer(r)
			defer ts.Close()
			path := ts.URL + "/me"
			req, err := http.NewRequest("GET", path, nil)
			req.Header.Set("Authorization", tt.header)
			if err != nil {
				t.Fatal(err)
			}
			res, err := client.Do(req)
			if err != nil {
				t.Fatal(err)
			}
			defer res.Body.Close()
			if tt.wantResp != nil {
				response := new(takrib.User)
				if err := json.NewDecoder(res.Body).Decode(response); err != nil {
					t.Fatal(err)
				}
				assert.Equal(t, tt.wantResp, response)
			}
			assert.Equal(t, tt.wantStatus, res.StatusCode)
		})
	}
}
