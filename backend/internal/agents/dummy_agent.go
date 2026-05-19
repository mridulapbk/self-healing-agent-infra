package agents

import (
	"fmt"
	"math/rand"
	"time"

	"self-healing-agent-infra/internal/config"
	"self-healing-agent-infra/internal/models"
)

func ProcessTask(task models.Task) models.Task {
	task.Status = models.StatusRunning
	task.UpdatedAt = time.Now().Format(time.RFC3339)

	fmt.Println("Processing task:", task.ID, "Attempt:", task.RetryCount+1)

	time.Sleep(time.Duration(config.Config.WorkerDelayMs) * time.Millisecond)

	if rand.Float64() < config.Config.FailureRate {
		task.Status = models.StatusFailed
		fmt.Println("Task failed:", task.ID)
	} else {
		task.Status = models.StatusCompleted
		fmt.Println("Task completed:", task.ID)
	}

	task.UpdatedAt = time.Now().Format(time.RFC3339)
	return task
}
