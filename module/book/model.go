package book

type BookEntity struct {
	ISBN   string `json:"isbn,omitempty"`
	Title  string `json:"title,omitempty"`
	Author string `json:"author,omitempty"`
}
