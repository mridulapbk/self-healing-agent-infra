package logger

import (
	"encoding/json"
	"fmt"
	"time"
)

type LogEntry struct {
	Timestamp string `json:"timestamp"`
	Level     string `json:"level"`
	Event     string `json:"event"`
	WorkerID  int    `json:"worker_id,omitempty"`
	TaskID    string `json:"task_id,omitempty"`
	Status    string `json:"status,omitempty"`
	Retry     int    `json:"retry_count,omitempty"`
	Message   string `json:"message,omitempty"`
}

func Log(level, event string, workerID int, taskID, status string, retry int, message string) {
	entry := LogEntry{
		Timestamp: time.Now().Format(time.RFC3339),
		Level:     level,
		Event:     event,
		WorkerID:  workerID,
		TaskID:    taskID,
		Status:    status,
		Retry:     retry,
		Message:   message,
	}

	logBytes, err := json.Marshal(entry)
	if err != nil {
		fmt.Println("failed to marshal log:", err)
		return
	}

	fmt.Println(string(logBytes))
}
