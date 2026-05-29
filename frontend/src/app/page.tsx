import WorkerChart from "@/components/WorkerChart";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import TaskDistributionChart from "@/components/TaskDistributionChart";
import {
  getBenchmarkMetrics,
  getSystemMetrics,
  getWorkerMetrics,
} from "@/lib/api";

export const revalidate = 5;

export default async function Home() {
  const systemMetrics = await getSystemMetrics();
  const benchmarkMetrics = await getBenchmarkMetrics();
  const workerMetrics = await getWorkerMetrics();

  const workerChartData = Object.entries(workerMetrics).map(
    ([workerId, metric]) => ({
      worker: `Worker ${workerId}`,
      processed: metric.processed,
    })
  );

  return (
    <main className="min-h-screen bg-slate-50 p-8">
      <div className="mb-8">
        <h1 className="text-4xl font-bold tracking-tight">
          Self-Healing Agent Infrastructure Dashboard
        </h1>
        <p className="mt-2 text-slate-600">
          Live monitoring of distributed workflows
        </p>
      </div>

      <section className="mb-10">
        <h2 className="mb-4 text-2xl font-semibold">System Metrics</h2>

        <div className="grid gap-6 md:grid-cols-3">
          <Card>
            <CardHeader>
              <CardTitle>Total Tasks</CardTitle>
            </CardHeader>
            <CardContent>
              <p className="text-4xl font-bold">
                {systemMetrics.total_tasks}
              </p>
            </CardContent>
          </Card>

          <Card>
            <CardHeader>
              <CardTitle>Completed</CardTitle>
            </CardHeader>
            <CardContent>
              <p className="text-4xl font-bold text-green-600">
                {systemMetrics.completed_tasks}
              </p>
            </CardContent>
          </Card>

          <Card>
            <CardHeader>
              <CardTitle>Recovered</CardTitle>
            </CardHeader>
            <CardContent>
              <p className="text-4xl font-bold text-blue-600">
                {systemMetrics.recovered_tasks}
              </p>
            </CardContent>
          </Card>

          <Card>
            <CardHeader>
              <CardTitle>Failed</CardTitle>
            </CardHeader>
            <CardContent>
              <p className="text-4xl font-bold text-red-600">
                {systemMetrics.failed_tasks}
              </p>
            </CardContent>
          </Card>

          <Card>
            <CardHeader>
              <CardTitle>Recovery Rate</CardTitle>
            </CardHeader>
            <CardContent>
              <p className="text-4xl font-bold">
                {(systemMetrics.recovery_rate * 100).toFixed(1)}%
              </p>
            </CardContent>
          </Card>

          <Card>
            <CardHeader>
              <CardTitle>Failure Rate</CardTitle>
            </CardHeader>
            <CardContent>
              <p className="text-4xl font-bold">
                {(systemMetrics.failure_rate * 100).toFixed(1)}%
              </p>
            </CardContent>
          </Card>
        </div>
      </section>

      <section className="mb-10">
        <h2 className="mb-4 text-2xl font-semibold">Benchmark Metrics</h2>

        <div className="grid gap-6 md:grid-cols-3">
          <Card>
            <CardHeader>
              <CardTitle>Success Rate</CardTitle>
            </CardHeader>
            <CardContent>
              <p className="text-4xl font-bold text-green-600">
                {(benchmarkMetrics.success_rate * 100).toFixed(1)}%
              </p>
            </CardContent>
          </Card>

          <Card>
            <CardHeader>
              <CardTitle>Average Recovery Time</CardTitle>
            </CardHeader>
            <CardContent>
              <p className="text-4xl font-bold">
                {benchmarkMetrics.average_recovery_time_ms.toFixed(0)} ms
              </p>
            </CardContent>
          </Card>

          <Card>
            <CardHeader>
              <CardTitle>Throughput / Second</CardTitle>
            </CardHeader>
            <CardContent>
              <p className="text-4xl font-bold">
                {benchmarkMetrics.throughput_per_second.toFixed(2)}
              </p>
            </CardContent>
          </Card>

          <Card>
            <CardHeader>
              <CardTitle>Total Execution Time</CardTitle>
            </CardHeader>
            <CardContent>
              <p className="text-4xl font-bold">
                {benchmarkMetrics.total_execution_time_sec.toFixed(2)}s
              </p>
            </CardContent>
          </Card>

          <Card>
            <CardHeader>
              <CardTitle>Benchmark Completed</CardTitle>
            </CardHeader>
            <CardContent>
              <p className="text-4xl font-bold">
                {benchmarkMetrics.completed_tasks}
              </p>
            </CardContent>
          </Card>

          <Card>
            <CardHeader>
              <CardTitle>Benchmark Failed</CardTitle>
            </CardHeader>
            <CardContent>
              <p className="text-4xl font-bold text-red-600">
                {benchmarkMetrics.failed_tasks}
              </p>
            </CardContent>
          </Card>
        </div>
      </section>

      <section className="mb-10">
        <h2 className="mb-4 text-2xl font-semibold">Worker Metrics</h2>

        <div className="overflow-hidden rounded-lg border bg-white">
          <table className="w-full">
            <thead className="bg-slate-100">
              <tr>
                <th className="p-4 text-left">Worker</th>
                <th className="p-4 text-left">Processed</th>
                <th className="p-4 text-left">Completed</th>
                <th className="p-4 text-left">Recovered</th>
                <th className="p-4 text-left">Failed</th>
              </tr>
            </thead>

            <tbody>
              {Object.entries(workerMetrics).map(([workerId, metric]) => (
                <tr key={workerId} className="border-t">
                  <td className="p-4 font-medium">
                    Worker {workerId}
                  </td>
                  <td className="p-4">{metric.processed}</td>
                  <td className="p-4 text-green-600">
                    {metric.completed}
                  </td>
                  <td className="p-4 text-blue-600">
                    {metric.recovered}
                  </td>
                  <td className="p-4 text-red-600">
                    {metric.failed}
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      </section>

      <section className="mb-10">
        <h2 className="mb-4 text-2xl font-semibold">
          Worker Utilization
        </h2>

        <Card>
          <CardContent className="pt-6">
            <WorkerChart data={workerChartData} />
          </CardContent>
        </Card>
      </section>

      <section className="mt-10">
        <h2 className="mb-4 text-2xl font-semibold">
          Task Distribution
        </h2>

        <Card>
          <CardContent className="pt-6">
            <TaskDistributionChart
              completed={systemMetrics.completed_tasks}
              recovered={systemMetrics.recovered_tasks}
              failed={systemMetrics.failed_tasks}
            />
          </CardContent>
        </Card>
      </section>
    </main>
  );
}