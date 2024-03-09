package main

import (
	"github.com/MauricioUhlig/go-rest-api/database"
	"github.com/MauricioUhlig/go-rest-api/models"
	"github.com/MauricioUhlig/go-rest-api/routers"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Teste", Historia: "Historia"},
		{Id: 2, Nome: "Teste 2", Historia: "Historia 2"},
	}
	database.ConnectToDB()
	routers.HandleRequest()
}
