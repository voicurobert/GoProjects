package entity

// Post struct used for POST requests
type Post struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}
