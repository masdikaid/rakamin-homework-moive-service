package contract

type MovieContract struct {
	Id          uint   `json:"id"`
	Title       string `json:"title"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	Duration    int    `json:"duration"`
	Image       string `json:"image"`
}
