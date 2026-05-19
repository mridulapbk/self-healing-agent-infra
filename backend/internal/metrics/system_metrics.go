package metrics

import "sync"

type SystemMetrics struct {
	TotalTasks     int     `json:"total_tasks"`
	CompletedTasks int     `json:"completed_tasks"`
	RecoveredTasks int     `json:"recovered_tasks"`
	FailedTasks    int     `json:"failed_tasks"`
	RecoveryRate   float64 `json:"recovery_rate"`
	FailureRate    float64 `json:"failure_rate"`
}

var (
	systemMetrics = &SystemMetrics{}
	systemLock    sync.RWMutex
)

func RecordSystemResult(status string) {
	systemLock.Lock()
	defer systemLock.Unlock()

	systemMetrics.TotalTasks++

	switch status {
	case "COMPLETED":
		systemMetrics.CompletedTasks++
	case "RECOVERED":
		systemMetrics.RecoveredTasks++
	case "FAILED":
		systemMetrics.FailedTasks++
	}

	if systemMetrics.TotalTasks > 0 {
		systemMetrics.RecoveryRate = float64(systemMetrics.RecoveredTasks) / float64(systemMetrics.TotalTasks)
		systemMetrics.FailureRate = float64(systemMetrics.FailedTasks) / float64(systemMetrics.TotalTasks)
	}
}

func GetSystemMetrics() SystemMetrics {
	systemLock.RLock()
	defer systemLock.RUnlock()

	return *systemMetrics
}
