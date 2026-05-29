export interface BenchmarkMetrics {
  total_tasks: number;
  completed_tasks: number;
  recovered_tasks: number;
  failed_tasks: number;
  success_rate: number;
  failure_rate: number;
  average_recovery_time_ms: number;
  throughput_per_second: number;
  total_execution_time_sec: number;
}