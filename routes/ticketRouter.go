package routes

import (
	"BiletAlSatArkaUc/controllers"

	"github.com/gin-gonic/gin"
)

func Ticket_Routes(incoming_routes *gin.Engine) {

	incoming_routes.GET("/ticket", controllers.Query_Tickets())
	incoming_routes.GET("/ticket/:id", controllers.Query())
	incoming_routes.GET("/ticket/cancel", controllers.Cancel())

}
