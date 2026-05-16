package orchestrator

import (
	"sync"

	"self-healing-agent-infra/internal/models"
)

var (
	taskStore = make(map[string]models.Task)
	storeLock sync.RWMutex
)

func SaveTask(task models.Task) {
	storeLock.Lock()
	defer storeLock.Unlock()

	taskStore[task.ID] = task
}

func GetTask(taskID string) (models.Task, bool) {
	storeLock.RLock()
	defer storeLock.RUnlock()

	task, exists := taskStore[taskID]
	return task, exists
}
