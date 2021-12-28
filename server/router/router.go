package router

import (
	"github.com/gin-gonic/gin"
	"golang-react-todo/server/middleware"
)

func InitRoutes() *gin.Engine {
	router := gin.New()
	router.GET("/api/task", middleware.GetAllTasks)
	router.POST("/api/tasks", middleware.CreateTask)
	router.PUT("/api/tasks/:id", middleware.TaskComplete)
	router.PUT("/api/undoTask/:id", middleware.UndoTask)
	router.DELETE("/api/deleteTask/:id", middleware.DeleteTask)
	return router
}
