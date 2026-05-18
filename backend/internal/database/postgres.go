package database

import (
	"database/sql"
	"fmt"
	"log"

	"self-healing-agent-infra/internal/models"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	connStr := "host=localhost port=5432 user=postgres password=postgres dbname=self_healing_db sslmode=disable"

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	fmt.Println("Connected to PostgreSQL")
}

func SaveTaskToDB(task models.Task) error {
	query := `
		INSERT INTO tasks (
			id, type, payload, status, retry_count, max_retries,
			created_at, updated_at, started_at, completed_at, recovery_duration
		)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)
		ON CONFLICT (id) DO UPDATE SET
			status = EXCLUDED.status,
			retry_count = EXCLUDED.retry_count,
			updated_at = EXCLUDED.updated_at,
			started_at = EXCLUDED.started_at,
			completed_at = EXCLUDED.completed_at,
			recovery_duration = EXCLUDED.recovery_duration;
	`

	_, err := DB.Exec(
		query,
		task.ID,
		task.Type,
		task.Payload,
		string(task.Status),
		task.RetryCount,
		task.MaxRetries,
		task.CreatedAt,
		task.UpdatedAt,
		task.StartedAt,
		task.CompletedAt,
		task.RecoveryDuration,
	)

	return err
}

func GetTaskFromDB(taskID string) (models.Task, error) {
	var task models.Task
	var status string

	query := `
		SELECT id, type, payload, status, retry_count, max_retries,
		       created_at, updated_at, started_at, completed_at, recovery_duration
		FROM tasks
		WHERE id = $1;
	`

	err := DB.QueryRow(query, taskID).Scan(
		&task.ID,
		&task.Type,
		&task.Payload,
		&status,
		&task.RetryCount,
		&task.MaxRetries,
		&task.CreatedAt,
		&task.UpdatedAt,
		&task.StartedAt,
		&task.CompletedAt,
		&task.RecoveryDuration,
	)

	task.Status = models.TaskStatus(status)
	return task, err
}
