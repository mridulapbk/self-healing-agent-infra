"use client";

import {
  Bar,
  BarChart,
  CartesianGrid,
  ResponsiveContainer,
  Tooltip,
  XAxis,
  YAxis,
} from "recharts";

interface WorkerChartData {
  worker: string;
  processed: number;
}

interface WorkerChartProps {
  data: WorkerChartData[];
}

export default function WorkerChart({ data }: WorkerChartProps) {
  return (
    <ResponsiveContainer width="100%" height={350}>
      <BarChart data={data}>
        <CartesianGrid strokeDasharray="3 3" />
        <XAxis dataKey="worker" />
        <YAxis />
        <Tooltip />
        <Bar
  dataKey="processed"
  fill="#3b82f6"
  radius={[8, 8, 0, 0]}
/>
      </BarChart>
    </ResponsiveContainer>
  );
}