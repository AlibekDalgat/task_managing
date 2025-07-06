package models

import "time"

type Status string

const (
	StatusPending   Status = "pending"
	StatusRunning   Status = "running"
	StatusCompleted Status = "completed"
	StatusFailed    Status = "failed"
)

type BaseTask struct {
	id          string
	status      Status
	createdAt   time.Time
	endedAt     time.Time
	result      interface{}
	description string
	err         error
}

func NewBaseTask(id, description string) *BaseTask {
	return &BaseTask{
		id:          id,
		status:      StatusPending,
		createdAt:   time.Now(),
		description: description,
	}
}

func (t *BaseTask) ID() string {
	return t.id
}

func (t *BaseTask) Status() Status {
	return t.status
}

func (t *BaseTask) CreatedAt() time.Time {
	return t.createdAt
}

func (t *BaseTask) Result() interface{} {
	return t.result
}

func (t *BaseTask) Error() error {
	return t.err
}

func (t *BaseTask) SetStatus(status Status) {
	t.status = status
	if status == StatusCompleted || status == StatusFailed {
		t.endedAt = time.Now()
	}
}

func (t *BaseTask) SetResult(result interface{}) {
	t.result = result
}

func (t *BaseTask) SetError(err error) {
	t.err = err
}

func (t *BaseTask) Description() string {
	return t.description
}
