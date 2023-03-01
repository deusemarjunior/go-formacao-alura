package main

import (
	"github.com/deusemarjunior/gin-api-rest/database"
	"github.com/deusemarjunior/gin-api-rest/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
