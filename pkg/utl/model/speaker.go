package  takrib

// Speaker represents speaker domain model
type Speaker struct {
	Base
	Name              string `json:"name"`
	ShortBiography    string `json:"short_biography"`
	LongBiography     string `json:"long_biography"`
	Gender            string `json:"gender"`
	Email             string `json:"email"`
	Mobile            string `json:"mobile"`
	Website           string `json:"website"`
	Twitter           string `json:"twitter"`
	Github            string `json:"github"`
	Linkedin          string `json:"linkedin"`
	Organisation      string `json:"organisation"`
	Position          string `json:"position"`
	Country           string `json:"country"`
	City              string `json:"city"`
	PhotoURL          string `json:"photo_url"`
	ThumbnailImageURL string `json:"thumbnail_image_url"`
	//EventID           string `json:"self.event_id"`
}
