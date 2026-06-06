#!/bin/bash

TOTAL_TASKS=${1:-100}

echo "Starting benchmark with $TOTAL_TASKS workflows..."

for i in $(seq 1 $TOTAL_TASKS)
do
  curl -s -X POST http://localhost:8080/workflow/start \
  -H "Content-Type: application/json" \
  -d '{"type":"benchmark","payload":"benchmark test"}' > /dev/null
done

echo ""
echo "Submitted $TOTAL_TASKS workflows."
echo "Wait 30-60 seconds for workers to finish."