package transport

import (
	"net/http"
	"strconv"

	"github.com/zaynkorai/takrib/pkg/api/speaker"

	takrib "github.com/zaynkorai/takrib/pkg/utl/model"

	"github.com/labstack/echo"
)

// HTTP represents speaker http service
type HTTP struct {
	svc speaker.Service
}

// NewHTTP creates new speaker http service
func NewHTTP(svc speaker.Service, er *echo.Group) {
	h := HTTP{svc}
	ur := er.Group("/speakers")
	// swagger:route POST /v1/speakers speakers speakerCreate
	// Creates new speaker account.
	// responses:
	//  200: speakerResp
	//  400: errMsg
	//  401: err
	//  403: errMsg
	//  500: err
	ur.POST("", h.create)

	// swagger:operation GET /v1/speakers speakers listSpeakers
	// ---
	// summary: Returns list of speakers.
	// description: Returns list of speakers. Depending on the speaker role requesting it, it may return all speakers for SuperAdmin/Admin speakers, all company/location speakers for Company/Location admins, and an error for non-admin speakers.
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
	//     "$ref": "#/responses/speakerListResp"
	//   "400":
	//     "$ref": "#/responses/errMsg"
	//   "401":
	//     "$ref": "#/responses/err"
	//   "403":
	//     "$ref": "#/responses/err"
	//   "500":
	//     "$ref": "#/responses/err"
	ur.GET("", h.list)

	// swagger:operation GET /v1/speakers/{id} speakers getSpeaker
	// ---
	// summary: Returns a single speaker.
	// description: Returns a single speaker by its ID.
	// parameters:
	// - name: id
	//   in: path
	//   description: id of speaker
	//   type: int
	//   required: true
	// responses:
	//   "200":
	//     "$ref": "#/responses/speakerResp"
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

	// swagger:operation PATCH /v1/speakers/{id} speakers speakerUpdate
	// ---
	// summary: Updates speaker's contact information
	// description: Updates speaker's contact information -> first name, last name, mobile, phone, address.
	// parameters:
	// - name: id
	//   in: path
	//   description: id of speaker
	//   type: int
	//   required: true
	// - name: request
	//   in: body
	//   description: Request body
	//   required: true
	//   schema:
	//     "$ref": "#/definitions/speakerUpdate"
	// responses:
	//   "200":
	//     "$ref": "#/responses/speakerResp"
	//   "400":
	//     "$ref": "#/responses/errMsg"
	//   "401":
	//     "$ref": "#/responses/err"
	//   "403":
	//     "$ref": "#/responses/err"
	//   "500":
	//     "$ref": "#/responses/err"
	ur.PATCH("/:id", h.update)

	// swagger:operation DELETE /v1/speakers/{id} speakers speakerDelete
	// ---
	// summary: Deletes a speaker
	// description: Deletes a speaker with requested ID.
	// parameters:
	// - name: id
	//   in: path
	//   description: id of speaker
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

// Speaker create request
// swagger:model speakerCreate
type createReq struct {
	Name              string `json:"name" validate:"required,min=5"`
	ShortBiography    string `json:"short_biography" validate:"required,min=20"`
	LongBiography     string `json:"long_biography" validate:"min=50"`
	Gender            string `json:"gender" validate:"required"`
	Email             string `json:"email" validate:"required,min=6"`
	Mobile            string `json:"mobile" validate:"required,min=9"`
	Website           string `json:"website"`
	Twitter           string `json:"twitter"`
	Github            string `json:"github" validate:"min=5"`
	Linkedin          string `json:"linkedin" validate:"required,min=4"`
	Organisation      string `json:"organisation" validate:"required,min=2"`
	Position          string `json:"position" validate:"required,min=2"`
	Country           string `json:"country" validate:"required,min=3"`
	City              string `json:"city" validate:"required,min=3"`
	PhotoURL          string `json:"photo_url" validate:"required,min=5"`
	ThumbnailImageURL string `json:"thumbnail_image_url"`
}

func (h *HTTP) create(c echo.Context) error {
	req := new(createReq)

	if err := c.Bind(req); err != nil {

		return err
	}
	usr, err := h.svc.Create(c, takrib.Speaker{
		Name:              req.Name,
		ShortBiography:    req.ShortBiography,
		LongBiography:     req.LongBiography,
		Gender:            req.Gender,
		Email:             req.Email,
		Mobile:            req.Mobile,
		Website:           req.Website,
		Twitter:           req.Twitter,
		Github:            req.Github,
		Linkedin:          req.Linkedin,
		Organisation:      req.Organisation,
		Position:          req.Position,
		Country:           req.Country,
		City:              req.City,
		PhotoURL:          req.PhotoURL,
		ThumbnailImageURL: req.ThumbnailImageURL,
	})

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, usr)
}

type listResponse struct {
	Speakers []takrib.Speaker `json:"speakers"`
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

// Speaker update request
// swagger:model speakerUpdate
type updateReq struct {
	ID                int     `json:"-"`
	Name              *string `json:"name,omitempty" validate:"omitempty,min=2"`
	ShortBiography    *string `json:"short_biography" validate:"omitempty,min=20"`
	LongBiography     *string `json:"long_biography" validate:"min=50"`
	Gender            *string `json:"gender" validate:"omitempty"`
	Email             *string `json:"email" validate:"omitempty,min=6"`
	Mobile            *string `json:"mobile" validate:"omitempty,min=9"`
	Website           *string `json:"website"`
	Twitter           *string `json:"twitter"`
	Github            *string `json:"github" validate:"min=5"`
	Linkedin          *string `json:"linkedin" validate:"omitempty,min=4"`
	Organisation      *string `json:"organisation" validate:"omitempty,min=2"`
	Position          *string `json:"position" validate:"omitempty,min=2"`
	Country           *string `json:"country" validate:"omitempty,min=3"`
	City              *string `json:"city" validate:"omitempty,min=3"`
	PhotoURL          *string `json:"photo_url" validate:"omitempty,min=5"`
	ThumbnailImageURL *string `json:"thumbnail_image_url"`
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

	usr, err := h.svc.Update(c, &speaker.Update{
		ID:                id,
		Name:              req.Name,
		ShortBiography:    req.ShortBiography,
		LongBiography:     req.LongBiography,
		Gender:            req.Gender,
		Email:             req.Email,
		Mobile:            req.Mobile,
		Website:           req.Website,
		Twitter:           req.Twitter,
		Github:            req.Github,
		Linkedin:          req.Linkedin,
		Organisation:      req.Organisation,
		Position:          req.Position,
		Country:           req.Country,
		City:              req.City,
		PhotoURL:          req.PhotoURL,
		ThumbnailImageURL: req.ThumbnailImageURL,
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
