"use client";

import {
  PieChart,
  Pie,
  Cell,
  Tooltip,
  ResponsiveContainer,
} from "recharts";

interface Props {
  completed: number;
  recovered: number;
  failed: number;
}

const COLORS = [
  "#22c55e",
  "#3b82f6",
  "#ef4444",
];

export default function TaskDistributionChart({
  completed,
  recovered,
  failed,
}: Props) {
  const data = [
    {
      name: "Completed",
      value: completed,
    },
    {
      name: "Recovered",
      value: recovered,
    },
    {
      name: "Failed",
      value: failed,
    },
  ];

  return (
    <ResponsiveContainer width="100%" height={350}>
      <PieChart>
        <Pie
          data={data}
          dataKey="value"
          outerRadius={120}
          label
        >
          {data.map((_, index) => (
            <Cell
              key={index}
              fill={COLORS[index]}
            />
          ))}
        </Pie>

        <Tooltip />
      </PieChart>
    </ResponsiveContainer>
  );
}