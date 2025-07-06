package service

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"sync"
	"task_managing/internal/models"
	"time"
)

type Task interface {
	ID() string
	Status() models.Status
	CreatedAt() time.Time
	Run(ctx context.Context) error
	Result() interface{}
	Description() string
	Error() error
	SetStatus(status models.Status)
	SetResult(result interface{})
	SetError(err error)
}

type TaskInfo struct {
	Task
	Cancel context.CancelFunc
}

type TaskManager struct {
	tasks map[string]TaskInfo
	mu    sync.RWMutex
}

func NewTaskManager() *TaskManager {
	return &TaskManager{
		tasks: make(map[string]TaskInfo),
	}
}

func (tm *TaskManager) CreateTask(description string) string {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	task := NewIOBoundTask(description)
	ctx, cancel := context.WithCancel(context.Background())
	taskInfo := TaskInfo{
		task,
		cancel,
	}
	tm.tasks[task.ID()] = taskInfo

	go func() {
		if err := task.Run(ctx); err != nil {
			logrus.Errorf("Задача %s провалилась: %v", task.ID(), err)
		}
	}()

	logrus.Infof("Задача %s создалась", task.ID())
	return task.ID()
}

func (tm *TaskManager) DeleteTask(id string) error {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	taskInfo, ok := tm.tasks[id]
	if !ok {
		return fmt.Errorf("Задача %s не найдена", id)
	}

	taskInfo.Cancel()
	delete(tm.tasks, id)

	logrus.Infof("Задача %s удалена", id)
	return nil
}

func (tm *TaskManager) GetTask(id string) (Task, error) {
	tm.mu.RLock()
	defer tm.mu.RUnlock()

	taskInfo, ok := tm.tasks[id]
	if !ok {
		return nil, fmt.Errorf("Задача %s не найдена", id)
	}
	return taskInfo.Task, nil
}
