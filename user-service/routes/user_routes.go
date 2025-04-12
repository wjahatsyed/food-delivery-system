package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wjahatsyed/user-service/controllers"
)

func UserRoutes(router *gin.Engine) {
	router.POST("/users", controllers.CreateUser)
}
