package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/renanmachad/rest_api.git/controllers"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/person", controllers.AllPersons).Methods("Get")
	r.HandleFunc("/api/person/{id}", controllers.Person).Methods("Get")

	log.Fatal(http.ListenAndServe(":8000", r))
}
