package orchestrator

import (
	"math/rand"
	"time"

	"self-healing-agent-infra/internal/agents"
	"self-healing-agent-infra/internal/config"
	"self-healing-agent-infra/internal/logger"
	"self-healing-agent-infra/internal/metrics"
	"self-healing-agent-infra/internal/models"
	"self-healing-agent-infra/internal/queue"
	"self-healing-agent-infra/internal/recovery"
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

			logger.Log("INFO", "task_picked", workerID, task.ID, string(task.Status), task.RetryCount, "worker picked task")

			if rand.Float64() < config.Config.CrashRate {
				logger.Log("ERROR", "worker_crash", workerID, task.ID, string(task.Status), task.RetryCount, "simulated worker crash")

				task.Status = models.StatusFailed
				task.UpdatedAt = time.Now().Format(time.RFC3339)
				SaveTask(task)

				metrics.RecordWorkerResult(workerID, string(task.Status))
				metrics.RecordSystemResult(string(task.Status))
				metrics.RecordBenchmarkResult(string(task.Status), 0)

				continue
			}

			processWorkflow(workerID, task)
		}
	}()
}

func processWorkflow(workerID int, task models.Task) {
	workflowStartTime := time.Now()

	strategy := recovery.GetStrategyFromEnv()

	task.Status = models.StatusRunning
	task.StartedAt = workflowStartTime.Format(time.RFC3339)
	task.UpdatedAt = workflowStartTime.Format(time.RFC3339)
	SaveTask(task)

	processedTask := agents.ProcessTask(task)

	for processedTask.Status == models.StatusFailed &&
		recovery.ShouldRetry(strategy, processedTask.RetryCount, processedTask.MaxRetries) {

		processedTask.RetryCount++
		SaveTask(processedTask)

		delay := recovery.GetRetryDelay(strategy, processedTask.RetryCount)

		logger.Log(
			"WARN",
			"task_retry",
			workerID,
			processedTask.ID,
			string(processedTask.Status),
			processedTask.RetryCount,
			"retrying failed task with recovery strategy",
		)

		time.Sleep(delay)

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

	metrics.RecordWorkerResult(workerID, string(processedTask.Status))
	metrics.RecordSystemResult(string(processedTask.Status))
	metrics.RecordBenchmarkResult(
		string(processedTask.Status),
		workflowEndTime.Sub(workflowStartTime),
	)

	logger.Log("INFO", "task_finished", workerID, processedTask.ID, string(processedTask.Status), processedTask.RetryCount, "worker finished task")
}
