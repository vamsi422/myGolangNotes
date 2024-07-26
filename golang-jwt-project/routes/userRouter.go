package routes

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/vamsi422/golang-jwt-project/controllers"
	"github.com/vamsi422/golang-jwt-project/middleware"
)

func UserRouters(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controllers.GetUsers())
	incomingRoutes.GET("users/:user_id", controllers.GetUser())
}
