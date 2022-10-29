package main

import (
	"fmt"

	"github.com/renanmachad/rest_api.git/models"
	"github.com/renanmachad/rest_api.git/routes"
)

func main() {

	models.Personalities = []models.Personality{
		{
			Id:      1,
			Name:    "Renan",
			History: "BLALBLALBLALBLLALA",
		},
		{
			Id:      2,
			Name:    "Renan",
			History: "BLALBLALBLALBLLALA",
		},
	}
	fmt.Println("Starting server ...")
	routes.HandleRequest()
}
