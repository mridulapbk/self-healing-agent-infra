package models

type TaskStatus string

const (
	StatusPending   TaskStatus = "PENDING"
	StatusRunning   TaskStatus = "RUNNING"
	StatusCompleted TaskStatus = "COMPLETED"
	StatusFailed    TaskStatus = "FAILED"
	StatusRecovered TaskStatus = "RECOVERED"
)

type Task struct {
	ID         string
	Type       string
	Payload    string
	Status     TaskStatus
	RetryCount int
	MaxRetries int
	CreatedAt  string
	UpdatedAt  string
}
