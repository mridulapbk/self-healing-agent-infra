# Self-Healing Agent Infrastructure

## Overview

Self-Healing Agent Infrastructure is a distributed workflow orchestration platform that simulates resilient AI agents capable of automatically recovering from failures in a distributed environment.

The platform leverages concurrent workers, durable Redis queues, PostgreSQL persistence, fault injection, retry mechanisms, benchmarking, and real-time monitoring to demonstrate self-healing behavior commonly found in modern cloud-native systems.

### Technologies Used

* Golang
* PostgreSQL
* Redis
* Next.js
* React
* Recharts
* Docker
* Tailwind CSS

---

## System Architecture Flow

```text
Client Request
      |
      v
API Layer (Go HTTP Server)
      |
      v
Redis Queue
      |
      v
Worker Pool (3 Concurrent Workers)
      |
      +---- Success ----> PostgreSQL
      |
      +---- Failure ----> Retry Engine
                               |
                               v
                        Self-Healing Recovery
                               |
                               v
                          PostgreSQL

Metrics API
      |
      v
Next.js Dashboard
```

---

## Architecture

### Components

### 1. API Layer

* Accepts workflow requests
* Exposes metrics endpoints
* Manages workflow orchestration APIs
* Provides communication between frontend and backend

### 2. Redis Queue

* Durable task storage
* Decouples producers and consumers
* Supports asynchronous task execution
* Enables distributed processing

### 3. Worker Pool

* Multiple concurrent workers
* Pull tasks from Redis
* Execute workflows independently
* Simulate distributed AI agents

### 4. Self-Healing Engine

* Detects failures
* Retries failed tasks automatically
* Recovers worker crashes
* Tracks recovery metrics

### 5. PostgreSQL

* Stores workflow execution history
* Persists task states
* Stores benchmark metrics
* Maintains worker statistics

### 6. Monitoring Dashboard

* Real-time system monitoring
* Worker utilization analytics
* Benchmark reporting
* Mobile-responsive visualization

---

## Features

### Workflow Orchestration

* Create and execute distributed workflows
* Concurrent task processing
* Queue-based workload distribution
* Asynchronous execution model

### Self-Healing Recovery

* Automatic retry mechanism
* Recovery from transient failures
* Worker crash simulation
* Recovery rate tracking
* Fault tolerance demonstration

### Distributed Queue

* Redis-backed durable queue
* Reliable task delivery
* Concurrent consumer support
* Scalable processing architecture

### Fault Injection

* Simulated worker crashes
* Random task failures
* Recovery benchmarking
* Reliability testing

### Benchmarking

* Success rate measurement
* Throughput calculation
* Recovery time tracking
* Worker utilization analytics
* Failure analysis

### Monitoring Dashboard

* System metrics
* Benchmark metrics
* Worker metrics
* Worker utilization charts
* Task distribution visualization
* Mobile responsive design
* Workflow execution controls

---

## Screenshots

### Dashboard Overview

![Dashboard Overview](docs/screenshots/dashboard-overview.jpg)

Complete dashboard displaying system metrics, benchmark metrics, worker metrics, workflow controls, and monitoring capabilities.

### Mobile Responsive View

![Mobile View](docs/screenshots/mobile-view.png)

Responsive dashboard layout optimized for mobile devices.

### Worker Metrics

![Worker Metrics](docs/screenshots/worker-metrics.jpg)

Worker performance statistics showing processed, completed, recovered, and failed tasks.

### Charts & Visualizations

![Charts](docs/screenshots/charts.jpg)

Worker utilization and task distribution visualizations generated from live backend data.

### Workflow Execution Logs

![Workflow Execution](docs/screenshots/workflow-execution.png)

Terminal logs demonstrating workflow processing, retries, worker crashes, recovery events, and successful task completion.

---

## Running Locally

### Clone Repository

```bash
git clone https://github.com/mridulapbk/self-healing-agent-infra.git
cd self-healing-agent-infra
```

### Start Infrastructure

```bash
docker compose up -d
```

### Run Backend

```bash
cd backend
go run cmd/server/main.go
```

Backend starts at:

```text
http://localhost:8080
```

### Run Frontend

```bash
cd frontend
npm install
npm run dev
```

Frontend starts at:

```text
http://localhost:3000
```

---

## API Endpoints

### Start Workflow

```http
POST /workflow/start
```

### Workflow Status

```http
GET /workflow/status
```

### Worker Metrics

```http
GET /metrics/workers
```

### System Metrics

```http
GET /metrics/system
```

### Benchmark Metrics

```http
GET /metrics/benchmark
```

---

## Example Metrics

```json
{
  "total_tasks": 20,
  "completed_tasks": 9,
  "recovered_tasks": 8,
  "failed_tasks": 3,
  "recovery_rate": 0.40,
  "failure_rate": 0.15
}
```

---

## Key Capabilities Demonstrated

* Distributed task orchestration
* Fault-tolerant workflow execution
* Worker crash simulation
* Automatic retry and recovery
* Queue-based architecture
* Benchmarking and performance monitoring
* Real-time dashboard visualization
* Mobile responsive frontend
* Persistent task storage
* Concurrent worker processing

---

## Future Enhancements

* WebSocket-based real-time updates
* Kubernetes deployment
* Dynamic worker auto-scaling
* AI-driven failure prediction
* Distributed tracing
* Prometheus integration
* Grafana dashboards
* GitHub Actions CI/CD pipeline

---

## Author

**Mridula Prabhakar**

Master of Science in Software Engineering Systems
Northeastern University

* Software Engineer
* Distributed Systems Enthusiast
* Cloud & Backend Development
* AI Infrastructure and Automation
