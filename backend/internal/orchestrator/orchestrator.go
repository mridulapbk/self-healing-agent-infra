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
	Message string            `json:"message"`
	TaskID  string            `json:"task_id"`
	Status  models.TaskStatus `json:"status"`
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

	task := models.Task{
		ID:         uuid.New().String(),
		Type:       req.Type,
		Payload:    req.Payload,
		Status:     models.StatusPending,
		RetryCount: 0,
		MaxRetries: 3,
		CreatedAt:  time.Now().Format(time.RFC3339),
		UpdatedAt:  time.Now().Format(time.RFC3339),
	}

	processedTask := agents.ProcessTask(task)

	response := StartWorkflowResponse{
		Message: "workflow processed",
		TaskID:  processedTask.ID,
		Status:  processedTask.Status,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
