package orchestrator

import (
	"encoding/json"
	"net/http"
	"time"

	"self-healing-agent-infra/internal/agents"
	"self-healing-agent-infra/internal/models"

	"github.com/google/uuid"
)

type StartWorkflowRequest struct {
	Type    string `json:"type"`
	Payload string `json:"payload"`
}

type StartWorkflowResponse struct {
	Message          string            `json:"message"`
	TaskID           string            `json:"task_id"`
	Status           models.TaskStatus `json:"status"`
	RetryCount       int               `json:"retry_count"`
	RecoveryDuration string            `json:"recovery_duration"`
}

func StartWorkflowHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method allowed", http.StatusMethodNotAllowed)
		return
	}

	var req StartWorkflowRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	workflowStartTime := time.Now()

	task := models.Task{
		ID:         uuid.New().String(),
		Type:       req.Type,
		Payload:    req.Payload,
		Status:     models.StatusPending,
		RetryCount: 0,
		MaxRetries: 3,
		CreatedAt:  workflowStartTime.Format(time.RFC3339),
		UpdatedAt:  workflowStartTime.Format(time.RFC3339),
		StartedAt:  workflowStartTime.Format(time.RFC3339),
	}

	processedTask := agents.ProcessTask(task)

	for processedTask.Status == models.StatusFailed && processedTask.RetryCount < processedTask.MaxRetries {
		processedTask.RetryCount++
		processedTask = agents.ProcessTask(processedTask)

		if processedTask.Status == models.StatusCompleted && processedTask.RetryCount > 0 {
			processedTask.Status = models.StatusRecovered
			break
		}
	}

	workflowEndTime := time.Now()
	processedTask.CompletedAt = workflowEndTime.Format(time.RFC3339)
	processedTask.RecoveryDuration = workflowEndTime.Sub(workflowStartTime).String()

	message := "workflow processed"
	if processedTask.Status == models.StatusRecovered {
		message = "workflow recovered after retry"
	} else if processedTask.Status == models.StatusFailed {
		message = "workflow failed after maximum retries"
	}

	response := StartWorkflowResponse{
		Message:          message,
		TaskID:           processedTask.ID,
		Status:           processedTask.Status,
		RetryCount:       processedTask.RetryCount,
		RecoveryDuration: processedTask.RecoveryDuration,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
