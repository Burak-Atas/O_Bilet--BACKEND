package routes

import (
	"BiletAlSatArkaUc/controllers"

	"github.com/gin-gonic/gin"
)

func Sefer_routes(incoming_routes *gin.Engine) {
	incoming_routes.GET("/seferekle", controllers.Sefer_Ekle())
	incoming_routes.GET("/sefersil", controllers.Sefer_Sil())
}
