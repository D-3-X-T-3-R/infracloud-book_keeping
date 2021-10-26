package end_points

import (
	"github.com/gorilla/mux"
	"github.com/sougata/book_keeping/handlers"
)

var RegBooksKeepRoutes = func(route *mux.Router) {
	route.HandleFunc("/book/", handlers.InsertBooks).Methods("POST")
	route.HandleFunc("/book/", handlers.FetchBooks).Methods("GET")
}
