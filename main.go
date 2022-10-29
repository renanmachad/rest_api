package main

import (
	"fmt"

	"github.com/renanmachad/rest_api.git/models"
	"github.com/renanmachad/rest_api.git/routes"
)

func main() {

	models.Personalities = []models.Personality{
		{
			Name:    "Renan",
			History: "BLALBLALBLALBLLALA",
		},
		{
			Name:    "Renan",
			History: "BLALBLALBLALBLLALA",
		},
	}
	fmt.Println("Fucking hello world")
	routes.HandleRequest()
}
