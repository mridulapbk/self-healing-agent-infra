package orchestrator

import (
	"sync"

	"self-healing-agent-infra/internal/database"
	"self-healing-agent-infra/internal/models"
)

var (
	taskStore = make(map[string]models.Task)
	storeLock sync.RWMutex
)

func SaveTask(task models.Task) {
	storeLock.Lock()
	taskStore[task.ID] = task
	storeLock.Unlock()

	err := database.SaveTaskToDB(task)
	if err != nil {
		// Keep app running even if DB write fails
		println("Failed to save task to DB:", err.Error())
	}
}

func GetTask(taskID string) (models.Task, bool) {
	task, err := database.GetTaskFromDB(taskID)
	if err == nil {
		return task, true
	}

	storeLock.RLock()
	defer storeLock.RUnlock()

	task, exists := taskStore[taskID]
	return task, exists
}
