package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
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
			logrus.Errorf("Задача %s провалилась: %v", t.ID(), t.Error())
		} else {
			t.SetStatus(models.StatusCompleted)
			logrus.Infof("Задача %s выполнилась", t.ID())
		}
	}()

	duration := 4 * time.Minute

	select {
	case <-time.After(duration):
		if rand.Intn(10) < 1 { // 10% шанс ошибки
			err := fmt.Errorf("симитированная ошибка во время выполнения задачи")
			t.SetError(err)
			t.SetEndedAt(time.Now())
			return err
		}

		result := "Задача выполнена успешно"
		t.SetResult(result)
		t.SetEndedAt(time.Now())
		return nil

	case <-ctx.Done():
		t.SetError(fmt.Errorf("задача отменена"))
		t.SetStatus(models.StatusFailed)
		t.SetEndedAt(time.Now())
		return ctx.Err()
	}
}
