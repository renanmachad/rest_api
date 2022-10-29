package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/renanmachad/rest_api.git/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "home page")
}

func AllPersons(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalities)
}
