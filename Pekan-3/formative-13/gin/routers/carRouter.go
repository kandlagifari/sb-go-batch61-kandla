package routers

import (
	"github.com/kandlagifari/sb-go-batch61-kandla/Pekan-3/formative-13/gin/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/cars", controllers.CreateCar)

	router.PUT("/cars/:carID", controllers.UpdateCar)

	router.GET("/cars/:carID", controllers.GetCar)

	router.DELETE("/cars/:carID", controllers.DeleteCar)

	return router
}
