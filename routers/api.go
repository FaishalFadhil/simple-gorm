package routers

import (
	"gin-api/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.GET("/cars/:id", controllers.GetOneCars)
	router.GET("/cars", controllers.GetAllCars)
	router.POST("/cars", controllers.CreateCars)
	router.PATCH("/cars/:id", controllers.UpdateCars)
	router.DELETE("/cars/:id", controllers.DeleteCars)

	router.GET("/users/:id", controllers.GetOneUser)
	router.GET("/users", controllers.GetAllUsers)
	router.POST("/users", controllers.CreateUser)
	router.PATCH("/users/:id", controllers.UpdateUser)
	router.DELETE("/users/:id", controllers.DeleteUser)

	return router
}
