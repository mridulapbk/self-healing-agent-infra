package main

import (
	"fmt"
	"log"
	"net/http"

	"self-healing-agent-infra/internal/database"
	"self-healing-agent-infra/internal/orchestrator"
)

func main() {
	database.InitDB()

	orchestrator.StartWorker(1)
	orchestrator.StartWorker(2)
	orchestrator.StartWorker(3)

	http.HandleFunc("/workflow/start", orchestrator.StartWorkflowHandler)
	http.HandleFunc("/workflow/status", orchestrator.GetWorkflowStatusHandler)
	http.HandleFunc("/metrics/workers", orchestrator.GetWorkerMetricsHandler)

	fmt.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
