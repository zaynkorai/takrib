package  takrib

// Event represents event domain model
type Event struct {
	Base
	Eventname string `json:"name"`
	Location  string `json:"location"`
}
