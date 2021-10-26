package main

import (
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sougata/book_keeping/end_points"
	"log"
	"net/http"
)

func main() {
	fmt.Println("starting server ...")
	router := mux.NewRouter()
	end_points.RegBooksKeepRoutes(router)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe("localhost:9000", router))
}
