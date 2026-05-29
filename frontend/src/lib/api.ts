import { BenchmarkMetrics } from "@/types/benchmark-metrics";
import { SystemMetrics } from "@/types/system-metrics";
import { WorkerMetricsResponse } from "@/types/worker-metrics";

export async function getWorkerMetrics(): Promise<WorkerMetricsResponse> {
  const response = await fetch(
    "http://localhost:8080/metrics/workers",
    {
      cache: "no-store",
    }
  );

  if (!response.ok) {
    throw new Error("Failed to fetch worker metrics");
  }

  return response.json();
}
export async function getSystemMetrics(): Promise<SystemMetrics> {
  const response = await fetch("http://localhost:8080/metrics/system", {
    cache: "no-store",
  });

  if (!response.ok) {
    throw new Error("Failed to fetch system metrics");
  }

  return response.json();
}

export async function getBenchmarkMetrics(): Promise<BenchmarkMetrics> {
  const response = await fetch("http://localhost:8080/metrics/benchmark", {
    cache: "no-store",
  });

  if (!response.ok) {
    throw new Error("Failed to fetch benchmark metrics");
  }

  return response.json();
}