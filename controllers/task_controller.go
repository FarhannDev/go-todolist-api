package controllers

import (
	"net/http"
	"strconv"
	"todolist-api/models"
	"todolist-api/services"

	"github.com/gin-gonic/gin"
)

func AddTask(c *gin.Context) {
	var requestTask models.Task
	if err := c.ShouldBindJSON(&requestTask); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse("Data tidak valid"))
		return
	}

	tasks := services.AddTask(requestTask.Task, requestTask.Description)
	c.JSON(http.StatusCreated, models.SuccessResponse("Task berhasil ditambahkan", tasks))
}
func GetTasks(c *gin.Context) {
	tasks := services.Gettask()
	c.JSON(http.StatusOK, models.SuccessResponse("Berhasil menampilkan data task", tasks))
}
func CompleteTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	task := services.CompleteTask(id)
	message := "Task berhasil diperbarui"
	c.JSON(http.StatusOK, models.SuccessResponse(message, task))
}
func DeleteTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	task := services.DeleteTask(id)
	message := "Task berhasil dihapus"
	c.JSON(http.StatusOK, models.SuccessResponse(message, task))
}
