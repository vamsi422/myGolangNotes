package routes

import (
	controllers "golang-jwt-project/controllers/"

	"github.com/gin-gonic/gin"
)

func AuthRouters(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/signup", controllers.Signup())
	incomingRoutes.POST("users/login", controllers.Login())
}
