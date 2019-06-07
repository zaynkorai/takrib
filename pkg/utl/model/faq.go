package  takrib

// FAQ represents FAQ domain model
type FAQ struct {
	Base
	Question string `json:"question"`
	Answer   string `json:"answer"`
	FAQType  string `json:"type"`
}
