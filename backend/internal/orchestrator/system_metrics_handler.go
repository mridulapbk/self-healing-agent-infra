package orchestrator

import (
	"encoding/json"
	"net/http"

	"self-healing-agent-infra/internal/metrics"
)

func GetSystemMetricsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(metrics.GetSystemMetrics())
}
