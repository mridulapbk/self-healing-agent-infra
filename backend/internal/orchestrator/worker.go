package orchestrator

import (
	"time"

	"self-healing-agent-infra/internal/agents"
	"self-healing-agent-infra/internal/logger"
	"self-healing-agent-infra/internal/metrics"
	"self-healing-agent-infra/internal/models"
	"self-healing-agent-infra/internal/queue"
)

func StartWorker(workerID int) {
	go func() {
		logger.Log("INFO", "worker_started", workerID, "", "", 0, "background worker started")

		for {
			task, err := queue.DequeueTask()
			if err != nil {
				logger.Log("ERROR", "dequeue_failed", workerID, "", "", 0, err.Error())
				continue
			}

			logger.Log(
				"INFO",
				"task_picked",
				workerID,
				task.ID,
				string(task.Status),
				task.RetryCount,
				"worker picked task",
			)

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

	for processedTask.Status == models.StatusFailed &&
		processedTask.RetryCount < processedTask.MaxRetries {

		processedTask.RetryCount++

		SaveTask(processedTask)

		logger.Log(
			"WARN",
			"task_retry",
			workerID,
			processedTask.ID,
			string(processedTask.Status),
			processedTask.RetryCount,
			"retrying failed task",
		)

		processedTask = agents.ProcessTask(processedTask)

		if processedTask.Status == models.StatusCompleted &&
			processedTask.RetryCount > 0 {

			processedTask.Status = models.StatusRecovered
			break
		}
	}

	workflowEndTime := time.Now()

	processedTask.CompletedAt = workflowEndTime.Format(time.RFC3339)
	processedTask.RecoveryDuration = workflowEndTime.Sub(workflowStartTime).String()
	processedTask.UpdatedAt = workflowEndTime.Format(time.RFC3339)

	SaveTask(processedTask)

	metrics.RecordWorkerResult(
		workerID,
		string(processedTask.Status),
	)

	metrics.RecordSystemResult(
		string(processedTask.Status),
	)

	logger.Log(
		"INFO",
		"task_finished",
		workerID,
		processedTask.ID,
		string(processedTask.Status),
		processedTask.RetryCount,
		"worker finished task",
	)
}
