package transport

import (
	takrib "github.com/zaynkorai/takrib/pkg/utl/model"
)

// Sponsor model response
// swagger:response sponsorResp
type swaggSponsorResponse struct {
	// in:body
	Body struct {
		*takrib.Sponsor
	}
}

// Sponsor model response
// swagger:response sponsorListResp
type swaggSponsorListResponse struct {
	// in:body
	Body struct {
		Sponsors []takrib.Sponsor `json:"sponsor"`
		Page     int             `json:"page"`
	}
}
