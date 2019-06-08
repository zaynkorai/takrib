package transport

import (
	"net/http"
	"strconv"

	"github.com/zaynkorai/takrib/pkg/api/sponsor"

	takrib "github.com/zaynkorai/takrib/pkg/utl/model"

	"github.com/labstack/echo"
)

// HTTP represents sponsor http service
type HTTP struct {
	svc sponsor.Service
}

// NewHTTP creates new sponsor http service
func NewHTTP(svc sponsor.Service, er *echo.Group) {
	h := HTTP{svc}
	ur := er.Group("/sponsors")
	// swagger:route POST /v1/sponsors sponsors sponsorCreate
	// Creates new sponsor account.
	// responses:
	//  200: sponsorResp
	//  400: errMsg
	//  401: err
	//  403: errMsg
	//  500: err
	ur.POST("", h.create)

	// swagger:operation GET /v1/sponsors sponsors listsponsors
	// ---
	// summary: Returns list of sponsors.
	// description: Returns list of sponsors. Depending on the sponsor role requesting it, it may return all sponsors for SuperAdmin/Admin sponsors, all company/location sponsors for Company/Location admins, and an error for non-admin sponsors.
	// parameters:
	// - name: limit
	//   in: query
	//   description: number of results
	//   type: int
	//   required: false
	// - name: page
	//   in: query
	//   description: page number
	//   type: int
	//   required: false
	// responses:
	//   "200":
	//     "$ref": "#/responses/sponsorListResp"
	//   "400":
	//     "$ref": "#/responses/errMsg"
	//   "401":
	//     "$ref": "#/responses/err"
	//   "403":
	//     "$ref": "#/responses/err"
	//   "500":
	//     "$ref": "#/responses/err"
	ur.GET("", h.list)

	// swagger:operation GET /v1/sponsors/{id} sponsors getsponsor
	// ---
	// summary: Returns a single sponsor.
	// description: Returns a single sponsor by its ID.
	// parameters:
	// - name: id
	//   in: path
	//   description: id of sponsor
	//   type: int
	//   required: true
	// responses:
	//   "200":
	//     "$ref": "#/responses/sponsorResp"
	//   "400":
	//     "$ref": "#/responses/err"
	//   "401":
	//     "$ref": "#/responses/err"
	//   "403":
	//     "$ref": "#/responses/err"
	//   "404":
	//     "$ref": "#/responses/err"
	//   "500":
	//     "$ref": "#/responses/err"
	ur.GET("/:id", h.view)

	// swagger:operation PATCH /v1/sponsors/{id} sponsors sponsorUpdate
	// ---
	// summary: Updates sponsor's contact information
	// description: Updates sponsor's contact information -> first name, last name, mobile, phone, address.
	// parameters:
	// - name: id
	//   in: path
	//   description: id of sponsor
	//   type: int
	//   required: true
	// - name: request
	//   in: body
	//   description: Request body
	//   required: true
	//   schema:
	//     "$ref": "#/definitions/sponsorUpdate"
	// responses:
	//   "200":
	//     "$ref": "#/responses/sponsorResp"
	//   "400":
	//     "$ref": "#/responses/errMsg"
	//   "401":
	//     "$ref": "#/responses/err"
	//   "403":
	//     "$ref": "#/responses/err"
	//   "500":
	//     "$ref": "#/responses/err"
	ur.PATCH("/:id", h.update)

	// swagger:operation DELETE /v1/sponsors/{id} sponsors sponsorDelete
	// ---
	// summary: Deletes a sponsor
	// description: Deletes a sponsor with requested ID.
	// parameters:
	// - name: id
	//   in: path
	//   description: id of sponsor
	//   type: int
	//   required: true
	// responses:
	//   "200":
	//     "$ref": "#/responses/ok"
	//   "400":
	//     "$ref": "#/responses/err"
	//   "401":
	//     "$ref": "#/responses/err"
	//   "403":
	//     "$ref": "#/responses/err"
	//   "500":
	//     "$ref": "#/responses/err"
	ur.DELETE("/:id", h.delete)
}

// Sponsor create request
// swagger:model sponsorCreate
type createReq struct {
	Name        string `json:"name" validate:"required,min=5"`
	Description string `json:"description"  validate:"required"`
	URL         string `json:"url"`
	LogoURL     string `json:"logo_url"`
	SponsorType string `json:"type"  validate:"required"`
}

func (h *HTTP) create(c echo.Context) error {
	req := new(createReq)

	if err := c.Bind(req); err != nil {

		return err
	}
	usr, err := h.svc.Create(c, takrib.Sponsor{
		Name:        req.Name,
		Description: req.Description,
		URL:         req.URL,
		LogoURL:     req.LogoURL,
		SponsorType: req.SponsorType,
	})

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, usr)
}

type listResponse struct {
	Sponsors []takrib.Sponsor `json:"sponsors"`
	Page     int              `json:"page"`
}

func (h *HTTP) list(c echo.Context) error {
	p := new(takrib.PaginationReq)
	if err := c.Bind(p); err != nil {
		return err
	}

	result, err := h.svc.List(c, p.Transform())

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, listResponse{result, p.Page})
}

func (h *HTTP) view(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return takrib.ErrBadRequest
	}

	result, err := h.svc.View(c, id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}

// Sponsor update request
// swagger:model sponsorUpdate
type updateReq struct {
	ID          int     `json:"-"`
	Name        *string `json:"name,omitempty" validate:"omitempty,min=3"`
	Description *string `json:"description"  validate:"omitempty,min=20"`
	URL         *string `json:"url"`
	LogoURL     *string `json:"logo_url"`
	SponsorType *string `json:"type"  validate:"omitempty,min=3"`
}

func (h *HTTP) update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return takrib.ErrBadRequest
	}

	req := new(updateReq)
	if err := c.Bind(req); err != nil {
		return err
	}

	usr, err := h.svc.Update(c, &sponsor.Update{
		ID:          id,
		Name:        req.Name,
		Description: req.Description,
		URL:         req.URL,
		LogoURL:     req.LogoURL,
		SponsorType: req.SponsorType,
	})

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, usr)
}

func (h *HTTP) delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return takrib.ErrBadRequest
	}

	if err := h.svc.Delete(c, id); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}
