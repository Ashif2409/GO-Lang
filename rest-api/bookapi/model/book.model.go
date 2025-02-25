package models

// import "gorm.io/gorm"

type Book struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Isbn     string `json:"isbn"`
	Name     string `json:"name"`
	AuthorID uint   `json:"author_id"`
	Author   Author `json:"author" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Author struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}
