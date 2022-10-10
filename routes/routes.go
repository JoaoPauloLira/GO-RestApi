package routes

import (
	"github.com/gorilla/mux"
	"go-rest-api/controllers"
	"log"
	"net/http"
)

const (
	GET    string = "Get"
	POST          = "Post"
	PUT           = "Put"
	DELETE        = "Delete"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods(GET)
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaUmaPersonalidade).Methods(GET)

	log.Fatal(http.ListenAndServe(":8000", r))
}
