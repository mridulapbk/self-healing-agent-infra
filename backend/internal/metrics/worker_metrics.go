package metrics

import "sync"

type WorkerMetrics struct {
	Processed int `json:"processed"`
	Completed int `json:"completed"`
	Recovered int `json:"recovered"`
	Failed    int `json:"failed"`
}

var (
	workerMetrics = make(map[int]*WorkerMetrics)
	metricsLock   sync.RWMutex
)

func RecordWorkerResult(workerID int, status string) {
	metricsLock.Lock()
	defer metricsLock.Unlock()

	if workerMetrics[workerID] == nil {
		workerMetrics[workerID] = &WorkerMetrics{}
	}

	workerMetrics[workerID].Processed++

	switch status {
	case "COMPLETED":
		workerMetrics[workerID].Completed++
	case "RECOVERED":
		workerMetrics[workerID].Recovered++
	case "FAILED":
		workerMetrics[workerID].Failed++
	}
}

func GetWorkerMetrics() map[int]*WorkerMetrics {
	metricsLock.RLock()
	defer metricsLock.RUnlock()

	result := make(map[int]*WorkerMetrics)
	for id, metric := range workerMetrics {
		result[id] = metric
	}

	return result
}
