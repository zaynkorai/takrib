package takrib

// Sponsor represents sponsor domain model
type Sponsor struct {
	Base
	Name        string `json:"name"`
	Description string `json:"description"`
	URL         string `json:"url"`
	LogoURL     string `json:"logo_url"`
	SponsorType string `json:"type"`
	// EventID     string `json:"event_id"`
}
