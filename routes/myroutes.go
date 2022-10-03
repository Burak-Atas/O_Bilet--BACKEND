package routes

import (
	"BiletAlSatArkaUc/controllers"

	"github.com/gin-gonic/gin"
)

func My_routes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/loadmoney", controllers.Load_money())
}
