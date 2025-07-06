package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"task_managing/internal/models"
	"time"
)

type CreateTaskResponse struct {
	TaskID string `json:"task_id"`
}

type GetTaskResponse struct {
	TaskID      string        `json:"task_id"`
	Status      models.Status `json:"status"`
	CreatedAt   time.Time     `json:"created_at"`
	EndedAt     time.Time     `json:"ended_at"`
	Result      interface{}   `json:"result"`
	Description string        `json:"description"`
	Error       string        `json:"error,omitempty"`
}

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}

type statusResponse struct {
	Status string `json:"status"`
}
