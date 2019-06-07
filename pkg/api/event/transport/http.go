package transport

import (
	"net/http"
	"strconv"

	"github.com/zaynkorai/takrib/pkg/api/event"

	"github.com/zaynkorai/takrib/pkg/utl/model"

	"github.com/labstack/echo"
)

// HTTP represents event http service
type HTTP struct {
	svc event.Service
}

// NewHTTP creates new event http service
func NewHTTP(svc event.Service, er *echo.Group) {
	h := HTTP{svc}
	ur := er.Group("/events")
	// swagger:route POST /v1/events events eventCreate
	// Creates new event account.
	// responses:
	//  200: eventResp
	//  400: errMsg
	//  401: err
	//  403: errMsg
	//  500: err
	ur.POST("", h.create)

	// swagger:operation GET /v1/events events listEvents
	// ---
	// summary: Returns list of events.
	// description: Returns list of events. Depending on the event role requesting it, it may return all events for SuperAdmin/Admin events, all company/location events for Company/Location admins, and an error for non-admin events.
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
	//     "$ref": "#/responses/eventListResp"
	//   "400":
	//     "$ref": "#/responses/errMsg"
	//   "401":
	//     "$ref": "#/responses/err"
	//   "403":
	//     "$ref": "#/responses/err"
	//   "500":
	//     "$ref": "#/responses/err"
	ur.GET("", h.list)

	// swagger:operation GET /v1/events/{id} events getevent
	// ---
	// summary: Returns a single event.
	// description: Returns a single event by its ID.
	// parameters:
	// - name: id
	//   in: path
	//   description: id of event
	//   type: int
	//   required: true
	// responses:
	//   "200":
	//     "$ref": "#/responses/eventResp"
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

	// swagger:operation PATCH /v1/events/{id} events eventUpdate
	// ---
	// summary: Updates event's contact information
	// description: Updates event's contact information -> first name, last name, mobile, phone, address.
	// parameters:
	// - name: id
	//   in: path
	//   description: id of event
	//   type: int
	//   required: true
	// - name: request
	//   in: body
	//   description: Request body
	//   required: true
	//   schema:
	//     "$ref": "#/definitions/eventUpdate"
	// responses:
	//   "200":
	//     "$ref": "#/responses/eventResp"
	//   "400":
	//     "$ref": "#/responses/errMsg"
	//   "401":
	//     "$ref": "#/responses/err"
	//   "403":
	//     "$ref": "#/responses/err"
	//   "500":
	//     "$ref": "#/responses/err"
	ur.PATCH("/:id", h.update)

	// swagger:operation DELETE /v1/events/{id} events eventDelete
	// ---
	// summary: Deletes a event
	// description: Deletes a event with requested ID.
	// parameters:
	// - name: id
	//   in: path
	//   description: id of event
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


// Event create request
// swagger:model eventCreate
type createReq struct {
	Eventname string `json:"eventname" validate:"required,min=5"`
	Location  string `json:"location" validate:"required"`
}

func (h *HTTP) create(c echo.Context) error {
	r := new(createReq)

	if err := c.Bind(r); err != nil {

		return err
	}
	usr, err := h.svc.Create(c, takrib.Event{
		Eventname: r.Eventname,
		Location:  r.Location,
	})

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, usr)
}

type listResponse struct {
	Events []takrib.Event `json:"events"`
	Page   int           `json:"page"`
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

// Event update request
// swagger:model eventUpdate
type updateReq struct {
	ID        int     `json:"-"`
	Eventname *string `json:"eventname,omitempty" validate:"omitempty,min=2"`
	Location  *string `json:"location,omitempty" validate:"omitempty,min=2"`

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

	usr, err := h.svc.Update(c, &event.Update{
		ID:        id,
		Eventname: req.Eventname,
		Location:  req.Location,
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
