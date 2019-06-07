package transport

import (
	takrib "github.com/zaynkorai/takrib/pkg/utl/model"
)

// Speaker model response
// swagger:response speakerResp
type swaggSpeakerResponse struct {
	// in:body
	Body struct {
		*takrib.Speaker
	}
}

// Speaker model response
// swagger:response speakerListResp
type swaggSpeakerListResponse struct {
	// in:body
	Body struct {
		Speakers []takrib.Speaker `json:"speaker"`
		Page     int             `json:"page"`
	}
}
