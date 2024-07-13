package routes

import (
	"github.com/gin-gonic/gin"
	"goAPI/controllers"
)

type TasksRoute struct {
	tasksController controllers.TasksController
}

func NewTasksRoute(tasksController controllers.TasksController) TasksRoute {
	return TasksRoute{tasksController}
}

func (tr TasksRoute) Register(router *gin.Engine) {
	router.POST("/people/:id/task/start", tr.tasksController.StartTask)
	router.POST("/people/:id/task/end", tr.tasksController.EndTask)
	router.GET("/people/:id/tasks", tr.tasksController.GetTasksByPeople)
}
