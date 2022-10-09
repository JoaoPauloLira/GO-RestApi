package main

import (
	"fmt"
	"go-rest-api/models"
	"go-rest-api/routes"
)

func main() {
	testes()
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}

func testes() {
	models.Personalidades = []models.Personalidade{
		{"Nome 1", "Historia 1"},
		{"Nome 2", "Historia 2"},
	}
}
