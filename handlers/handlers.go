package handlers

import (
	"encoding/json"
	"github.com/sougata/book_keeping/models"
	"github.com/sougata/book_keeping/parser"
	"net/http"
)

var NewBook models.Books

func FetchBooks(writer http.ResponseWriter, _ *http.Request) {
	newBooks := models.FetchBooks()
	response, _ := json.Marshal(newBooks)
	writer.Header().Set("Content-Type", "pkglication/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(response)
}

func InsertBooks(writer http.ResponseWriter, req *http.Request) {
	createBooks := &models.Books{}
	parser.ParseJsonBody(req, createBooks)
	books := createBooks.InsertBooks()
	response, _ := json.Marshal(books)
	writer.WriteHeader(http.StatusOK)
	writer.Write(response)
}
