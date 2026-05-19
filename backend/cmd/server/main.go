package main

import (
	"fmt"
	"log"
	"net/http"

	"self-healing-agent-infra/internal/config"
	"self-healing-agent-infra/internal/database"
	"self-healing-agent-infra/internal/orchestrator"
	"self-healing-agent-infra/internal/queue"
)

func main() {
	database.InitDB()
	queue.InitRedisQueue()

	for i := 1; i <= config.Config.WorkerCount; i++ {
		orchestrator.StartWorker(i)
	}

	http.HandleFunc("/workflow/start", orchestrator.StartWorkflowHandler)
	http.HandleFunc("/workflow/status", orchestrator.GetWorkflowStatusHandler)
	http.HandleFunc("/metrics/workers", orchestrator.GetWorkerMetricsHandler)
	http.HandleFunc("/metrics/system", orchestrator.GetSystemMetricsHandler)
	http.HandleFunc("/metrics/benchmark", orchestrator.GetBenchmarkMetricsHandler)

	fmt.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
