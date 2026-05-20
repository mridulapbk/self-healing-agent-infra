# Self-Healing Agent Infrastructure

## Overview

Distributed workflow orchestration platform that simulates
self-healing AI agents using:

- Redis durable queue
- Concurrent workers
- Retry recovery
- PostgreSQL persistence
- Fault injection
- Benchmarking
- Structured logging

## Architecture

[diagram]

## Features

### Workflow Orchestration
...

### Self-Healing Recovery
...

### Distributed Queue
...

### Fault Injection
...

### Benchmarking
...

## Running Locally

docker compose up -d

go run cmd/server/main.go

## Metrics

GET /metrics/workers
GET /metrics/system
GET /metrics/benchmark