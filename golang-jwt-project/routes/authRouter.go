package routes

import (
	controllers "github.com/vamsi422/golang-jwt-project/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRouters(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/signup", controllers.Signup())
	incomingRoutes.POST("users/login", controllers.Login())
}
