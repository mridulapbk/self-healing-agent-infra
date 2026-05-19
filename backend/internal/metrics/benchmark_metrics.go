package metrics

import (
	"sync"
	"time"
)

type BenchmarkMetrics struct {
	TotalTasks            int     `json:"total_tasks"`
	CompletedTasks        int     `json:"completed_tasks"`
	RecoveredTasks        int     `json:"recovered_tasks"`
	FailedTasks           int     `json:"failed_tasks"`
	SuccessRate           float64 `json:"success_rate"`
	FailureRate           float64 `json:"failure_rate"`
	AverageRecoveryTimeMs float64 `json:"average_recovery_time_ms"`
	ThroughputPerSecond   float64 `json:"throughput_per_second"`
	TotalExecutionTimeSec float64 `json:"total_execution_time_sec"`
}

var (
	benchmarkMetrics = &BenchmarkMetrics{}
	benchmarkLock    sync.RWMutex
	benchmarkStart   time.Time
	totalRecoveryMs  float64
)

func RecordBenchmarkResult(status string, recoveryDuration time.Duration) {
	benchmarkLock.Lock()
	defer benchmarkLock.Unlock()

	if benchmarkMetrics.TotalTasks == 0 {
		benchmarkStart = time.Now()
	}

	benchmarkMetrics.TotalTasks++

	switch status {
	case "COMPLETED":
		benchmarkMetrics.CompletedTasks++
	case "RECOVERED":
		benchmarkMetrics.RecoveredTasks++
	case "FAILED":
		benchmarkMetrics.FailedTasks++
	}

	totalRecoveryMs += float64(recoveryDuration.Milliseconds())

	successfulTasks := benchmarkMetrics.CompletedTasks + benchmarkMetrics.RecoveredTasks

	benchmarkMetrics.SuccessRate = float64(successfulTasks) / float64(benchmarkMetrics.TotalTasks)
	benchmarkMetrics.FailureRate = float64(benchmarkMetrics.FailedTasks) / float64(benchmarkMetrics.TotalTasks)
	benchmarkMetrics.AverageRecoveryTimeMs = totalRecoveryMs / float64(benchmarkMetrics.TotalTasks)

	elapsed := time.Since(benchmarkStart).Seconds()
	benchmarkMetrics.TotalExecutionTimeSec = elapsed

	if elapsed > 0 {
		benchmarkMetrics.ThroughputPerSecond = float64(benchmarkMetrics.TotalTasks) / elapsed
	}
}

func GetBenchmarkMetrics() BenchmarkMetrics {
	benchmarkLock.RLock()
	defer benchmarkLock.RUnlock()

	return *benchmarkMetrics
}
