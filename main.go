package main

import (
	"dts-task/app"
	"dts-task/handler"
	"dts-task/task"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db := app.NewDB()

	taskRepository := task.NewRepository(db)
	taskService := task.NewService(taskRepository)
	taskHandler := handler.NewTaskHandler(taskService)

	router := gin.Default()
	router.Use(cors.Default())

	api := router.Group("/api/v1")

	api.GET("/tasks", taskHandler.GetTasks)
	api.GET("/tasks/:id", taskHandler.GetTask)
	api.POST("/tasks", taskHandler.CreateTask)
	api.PUT("/tasks/:id", taskHandler.UpdateTask)
	api.DELETE("/tasks/:id", taskHandler.DeleteTask)

	router.Run()
}
