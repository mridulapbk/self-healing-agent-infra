export interface WorkerMetric {
  processed: number;
  completed: number;
  recovered: number;
  failed: number;
}

export interface WorkerMetricsResponse {
  [workerId: string]: WorkerMetric;
}