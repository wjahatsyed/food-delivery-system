package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/wjahatsyed/user-service/routes"
	"github.com/wjahatsyed/user-service/services"
)

func main() {
	services.InitMongo()

	router := gin.Default()
	routes.UserRoutes(router)

	log.Println("🚀 Server started at http://localhost:8080")
	router.Run(":8080")
}
