package delivery

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createTask(c *gin.Context) {
	var input CreateTaskRequest
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "неверное содержание json")
		return
	}
	id := h.taskManager.CreateTask(input.Description)
	c.JSON(http.StatusOK, CreateTaskResponse{
		TaskID: id},
	)
}

func (h *Handler) deleteTask(c *gin.Context) {
	id := c.Param("id")
	err := h.taskManager.DeleteTask(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{Status: "Задача удалена"})
}

func (h *Handler) getTask(c *gin.Context) {
	id := c.Param("id")
	task, err := h.taskManager.GetTask(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	var errorMessage string
	if task.Error() != nil {
		errorMessage = task.Error().Error()
	}
	c.JSON(http.StatusOK, GetTaskResponse{
		TaskID:      task.ID(),
		Status:      task.Status(),
		CreatedAt:   task.CreatedAt(),
		EndedAt:     task.EndedAt(),
		Result:      task.Result(),
		Description: task.Description(),
		Error:       errorMessage,
	})
}
