package takrib

import "time"

// Event represents event domain model
type Event struct {
	Base
	Name                 string     `json:"name"`
	Email                string     `json:"email"`
	Description          string     `json:"description"`
	StartTime            time.Time  `json:"start_time"`
	EndTime              time.Time  `json:"end_time"`
	OrganizerName        string     `json:"organizer_name"`
	OrganizerDescription string     `json:"organizer_description"`
	Latitude             string     `json:"latitude"`
	Longitude            string     `json:"longitude"`
	State                string     `json:"state"`
	LogoURL              string     `json:"logo_url"`
	LargeImageURL        string     `json:"large_imageUrl"`
	CopyRight            string     `json:"copyright"`
	CodeOfConduct        string     `json:"code_of_conduct"`
	Speakers             []*Speaker `json:"-"`
	Sponsors             []*Sponsor `json:"-"`
}
