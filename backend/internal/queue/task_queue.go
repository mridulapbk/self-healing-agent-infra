package queue

import "self-healing-agent-infra/internal/models"

var TaskQueue = make(chan models.Task, 100)

func EnqueueTask(task models.Task) {
	TaskQueue <- task
}

func GetTaskQueue() <-chan models.Task {
	return TaskQueue
}
