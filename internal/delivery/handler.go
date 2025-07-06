package delivery

import (
	"github.com/gin-gonic/gin"
	"task_managing/internal/service"
)

type TaskManagerService interface {
	CreateTask(description string) string
	GetTask(id string) (service.Task, error)
	DeleteTask(id string) error
}

type Handler struct {
	taskManager TaskManagerService
}

func NewHandler(taskManager TaskManagerService) *Handler {
	return &Handler{
		taskManager: taskManager,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	api := router.Group("api")
	{
		api.POST("", h.createTask)
		api.DELETE(":id", h.deleteTask)
		api.GET(":id", h.getTask)
	}
	return router
}
