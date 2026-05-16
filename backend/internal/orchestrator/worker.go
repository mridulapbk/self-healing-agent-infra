package orchestrator

import (
	"fmt"
	"time"

	"self-healing-agent-infra/internal/agents"
	"self-healing-agent-infra/internal/models"
	"self-healing-agent-infra/internal/queue"
)

func StartWorker() {
	go func() {
		fmt.Println("Background worker started")

		for task := range queue.GetTaskQueue() {
			processWorkflow(task)
		}
	}()
}

func processWorkflow(task models.Task) {
	workflowStartTime := time.Now()

	task.Status = models.StatusRunning
	task.StartedAt = workflowStartTime.Format(time.RFC3339)
	task.UpdatedAt = workflowStartTime.Format(time.RFC3339)
	SaveTask(task)

	processedTask := agents.ProcessTask(task)

	for processedTask.Status == models.StatusFailed && processedTask.RetryCount < processedTask.MaxRetries {
		processedTask.RetryCount++
		SaveTask(processedTask)

		processedTask = agents.ProcessTask(processedTask)

		if processedTask.Status == models.StatusCompleted && processedTask.RetryCount > 0 {
			processedTask.Status = models.StatusRecovered
			break
		}
	}

	workflowEndTime := time.Now()
	processedTask.CompletedAt = workflowEndTime.Format(time.RFC3339)
	processedTask.RecoveryDuration = workflowEndTime.Sub(workflowStartTime).String()
	processedTask.UpdatedAt = workflowEndTime.Format(time.RFC3339)

	SaveTask(processedTask)

	fmt.Println("Workflow finished:", processedTask.ID, processedTask.Status)
}
