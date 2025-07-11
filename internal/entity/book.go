package entity

type Book struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	ISBN     string `json:"isbn"`
	Quantity int    `json:"quantity"`
	Category string `json:"category"`
}
