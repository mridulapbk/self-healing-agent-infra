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

func enableCORS(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		handler(w, r)
	}
}

func main() {
	database.InitDB()
	queue.InitRedisQueue()

	for i := 1; i <= config.Config.WorkerCount; i++ {
		orchestrator.StartWorker(i)
	}

	http.HandleFunc(
		"/workflow/start",
		enableCORS(orchestrator.StartWorkflowHandler),
	)

	http.HandleFunc(
		"/workflow/status",
		enableCORS(orchestrator.GetWorkflowStatusHandler),
	)

	http.HandleFunc(
		"/metrics/workers",
		enableCORS(orchestrator.GetWorkerMetricsHandler),
	)

	http.HandleFunc(
		"/metrics/system",
		enableCORS(orchestrator.GetSystemMetricsHandler),
	)

	http.HandleFunc(
		"/metrics/benchmark",
		enableCORS(orchestrator.GetBenchmarkMetricsHandler),
	)

	fmt.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
