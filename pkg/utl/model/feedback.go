package takrib

// FeedBack represents feedback domain model
type FeedBack struct {
	Base
	Rating  string `json:"rating"`
	Comment string `json:"comment"`
	// UserID  int    `json:"user_id"`
	// EventID int    `json:"event_id"`
}
