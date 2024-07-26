package routes

import (
	controllers "golang-jwt-project/controllers"

	"github.com/gin-gonic/gin"

	"golang-jwt-project/middleware"
)

func UserRouters(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controllers.GetUser())
	incomingRoutes.GET("users/:user_id", controllers.GetUser())
}
