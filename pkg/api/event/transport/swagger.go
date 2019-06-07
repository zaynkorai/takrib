package transport

import (
	"github.com/zaynkorai/takrib/pkg/utl/model"
)

// Event model response
// swagger:response eventResp
type swaggEventResponse struct {
	// in:body
	Body struct {
		*takrib.Event
	}
}

// Event model response
// swagger:response eventListResp
type swaggEventListResponse struct {
	// in:body
	Body struct {
		Events []takrib.Event `json:"event"`
		Page   int           `json:"page"`
	}
}
