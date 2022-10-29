package routes

import (
	"log"
	"net/http"

	"github.com/renanmachad/rest_api.git/controllers"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/api/person", controllers.AllPersons)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
