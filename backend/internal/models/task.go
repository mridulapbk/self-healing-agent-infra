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
	ID               string     `json:"id"`
	Type             string     `json:"type"`
	Payload          string     `json:"payload"`
	Status           TaskStatus `json:"status"`
	RetryCount       int        `json:"retry_count"`
	MaxRetries       int        `json:"max_retries"`
	CreatedAt        string     `json:"created_at"`
	UpdatedAt        string     `json:"updated_at"`
	StartedAt        string     `json:"started_at"`
	CompletedAt      string     `json:"completed_at"`
	RecoveryDuration string     `json:"recovery_duration"`
}
