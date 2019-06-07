package  takrib

// FeedBack represents feedback domain model
type FeedBack struct {
	Base
	Rating  string `json:"rating"`
	Comment string `json:"comment"`
	//	UserID     string `json:"user_id"`
	//	EventID     string `json:"user_id"`
}
