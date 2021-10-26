package models

import (
	"github.com/jinzhu/gorm"
	"github.com/sougata/book_keeping/config"
	"os"
	"strings"
)

type Books struct {
	Books []Book `json:"books"`
}

type Book struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
	Year   int    `json:"year"`
	Price  int    `json:"price"`
}

var db *gorm.DB

func init() {
	config_file_path := os.Args[1:]
	config.GetConn(strings.Join(config_file_path, " "))
	db = config.GetDB()
	db.AutoMigrate(&Books{})
}

func (book_list *Books) InsertBooks() *Books {
	for _, book := range book_list.Books {
		db.NewRecord(book)
		db.Create(&book)
	}
	return book_list
}

func FetchBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}
