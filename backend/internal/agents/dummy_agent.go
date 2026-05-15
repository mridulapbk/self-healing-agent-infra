package agents

import (
	"fmt"
	"math/rand"
	"time"

	"self-healing-agent-infra/internal/models"
)

func ProcessTask(task models.Task) models.Task {
	task.Status = models.StatusRunning
	task.UpdatedAt = time.Now().Format(time.RFC3339)

	fmt.Println("Processing task:", task.ID)

	time.Sleep(2 * time.Second)

	if rand.Intn(100) < 30 {
		task.Status = models.StatusFailed
		fmt.Println("Task failed:", task.ID)
	} else {
		task.Status = models.StatusCompleted
		fmt.Println("Task completed:", task.ID)
	}

	task.UpdatedAt = time.Now().Format(time.RFC3339)
	return task
}
