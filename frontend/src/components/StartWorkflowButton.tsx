"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import { Button } from "@/components/ui/button";

export default function StartWorkflowButton() {
  const router = useRouter();
  const [loading, setLoading] = useState(false);
  const [message, setMessage] = useState("");

  async function startWorkflow() {
    setLoading(true);
    setMessage("");

    try {
      const response = await fetch("http://localhost:8080/workflow/start", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          type: "summarization",
          payload: "UI triggered workflow",
        }),
      });

      if (!response.ok) {
        throw new Error("Failed to start workflow");
      }

      setMessage("Workflow started successfully");

      setTimeout(() => {
        router.refresh();
      }, 3000);
    } catch {
      setMessage("Failed to start workflow");
    } finally {
      setLoading(false);
    }
  }

  return (
    <div className="flex items-center gap-4">
      <Button onClick={startWorkflow} disabled={loading}>
        {loading ? "Starting..." : "Start Test Workflow"}
      </Button>

      {message && (
        <p className="text-sm text-slate-600">
          {message}
        </p>
      )}
    </div>
  );
}