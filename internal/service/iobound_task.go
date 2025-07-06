package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"task_managing/internal/models"
	"time"
)

type IOBoundTask struct {
	*models.BaseTask
}

func NewIOBoundTask(description string) *IOBoundTask {
	return &IOBoundTask{
		BaseTask: models.NewBaseTask(uuid.New().String(), description),
	}
}

func (t *IOBoundTask) Run(ctx context.Context) error {
	t.SetStatus(models.StatusRunning)
	defer func() {
		if t.Error() != nil {
			t.SetStatus(models.StatusFailed)
		} else {
			t.SetStatus(models.StatusCompleted)
		}
	}()

	duration := time.Duration(rand.Intn(3) + 3)
	done := make(chan bool)

	go func() {
		defer close(done)

		time.Sleep(duration)

		if rand.Intn(10) < 1 { // 10% шанс ошибки
			err := fmt.Errorf("ошибка во время выполнения задачи")
			t.SetError(err)
			return
		}

		result := fmt.Sprintf("Задача выполненв успешно. Описание: %s", t.Description())
		t.SetResult(result)
	}()

	select {
	case <-done:
		return t.Error()
	case <-ctx.Done():
		t.SetError(fmt.Errorf("задача удалена"))
		return ctx.Err()
	}
}
