package routes

import (
	"todolist-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Endpoint task
	r.POST("/task", controllers.AddTask)
	r.GET("/task", controllers.GetTasks)
	r.PUT("/task/:id/complete", controllers.CompleteTask)
	r.DELETE("/task/:id", controllers.DeleteTask)

	return r
}
