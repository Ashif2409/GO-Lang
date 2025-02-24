package models

type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Name   string  `json:"name"`
	Author *Author `json:"author"`
}

type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}
