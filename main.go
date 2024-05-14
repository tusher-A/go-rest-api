package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tusher-A/go-rest-api/configs"
	"github.com/tusher-A/go-rest-api/routes"
)


func main() {
	router := gin.Default()

	// connect to db
	configs.ConnectDB()

	// rest api initiated
	routes.InitiateRoutes(router)

	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	routerInitiateError := router.Run()
	if routerInitiateError != nil {
		log.Fatal(routerInitiateError)
	}
}
