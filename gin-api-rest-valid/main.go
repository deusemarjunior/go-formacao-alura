package main

import (
	"github.com/deusemarjunior/gin-api-rest-valid/database"
	"github.com/deusemarjunior/gin-api-rest-valid/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequest()
}

