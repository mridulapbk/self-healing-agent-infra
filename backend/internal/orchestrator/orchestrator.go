package orchestrator

import (
	"encoding/json"
	"net/http"
	"time"

	"self-healing-agent-infra/internal/config"
	"self-healing-agent-infra/internal/models"
	"self-healing-agent-infra/internal/queue"

	"github.com/google/uuid"
)

type StartWorkflowRequest struct {
	Type    string `json:"type"`
	Payload string `json:"payload"`
}

type StartWorkflowResponse struct {
	Message    string            `json:"message"`
	TaskID     string            `json:"task_id"`
	Status     models.TaskStatus `json:"status"`
	RetryCount int               `json:"retry_count"`
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

	now := time.Now()

	task := models.Task{
		ID:         uuid.New().String(),
		Type:       req.Type,
		Payload:    req.Payload,
		Status:     models.StatusPending,
		RetryCount: 0,
		MaxRetries: config.Config.MaxRetries,
		CreatedAt:  now.Format(time.RFC3339),
		UpdatedAt:  now.Format(time.RFC3339),
	}

	SaveTask(task)
	queue.EnqueueTask(task)

	response := StartWorkflowResponse{
		Message:    "workflow accepted and queued",
		TaskID:     task.ID,
		Status:     task.Status,
		RetryCount: task.RetryCount,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetWorkflowStatusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method allowed", http.StatusMethodNotAllowed)
		return
	}

	taskID := r.URL.Query().Get("id")
	if taskID == "" {
		http.Error(w, "Task ID is required", http.StatusBadRequest)
		return
	}

	task, exists := GetTask(taskID)
	if !exists {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}
