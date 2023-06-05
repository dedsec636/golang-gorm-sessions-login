package model

import "ginbasic/helper"

type Author struct {
	ID    uint `gorm:"primary_key;auto_increment"`
	Name  string
	Books []Book
}

type Book struct {
	BookID   uint `gorm:"primary_key;auto_increment"`
	Title    string
	AuthorID uint   `gorm:"index;not null"`      // Foreign key
	Author   Author `gorm:"foreignkey:AuthorID"` // Association
}

func (author *Author) Save() (*Author, error) {
	err := helper.Database.Preload("Books").Create(&author).Error
	if err != nil {
		return &Author{}, err
	}
	return author, err
}

func (book *Book) Save() (*Book, error) {
	err := helper.Database.Preload("Author").Create(&book).Error
	if err != nil {
		return &Book{}, err
	}
	return book, err
}
