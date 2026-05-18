package orchestrator

import (
	"fmt"
	"time"

	"self-healing-agent-infra/internal/agents"
	"self-healing-agent-infra/internal/metrics"
	"self-healing-agent-infra/internal/models"
	"self-healing-agent-infra/internal/queue"
)

func StartWorker(workerID int) {
	go func() {
		fmt.Println("Background worker started:", workerID)

		for task := range queue.GetTaskQueue() {
			fmt.Println("Worker", workerID, "picked task:", task.ID)
			processWorkflow(workerID, task)
		}
	}()
}

func processWorkflow(workerID int, task models.Task) {
	workflowStartTime := time.Now()

	task.Status = models.StatusRunning
	task.StartedAt = workflowStartTime.Format(time.RFC3339)
	task.UpdatedAt = workflowStartTime.Format(time.RFC3339)
	SaveTask(task)

	processedTask := agents.ProcessTask(task)

	for processedTask.Status == models.StatusFailed && processedTask.RetryCount < processedTask.MaxRetries {
		processedTask.RetryCount++
		SaveTask(processedTask)

		fmt.Println("Worker", workerID, "retrying task:", processedTask.ID, "Retry:", processedTask.RetryCount)

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
	metrics.RecordWorkerResult(workerID, string(processedTask.Status))

	fmt.Println("Worker", workerID, "finished task:", processedTask.ID, "Status:", processedTask.Status)
}
