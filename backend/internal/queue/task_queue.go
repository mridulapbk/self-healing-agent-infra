package queue

import (
	"context"
	"encoding/json"
	"fmt"

	"self-healing-agent-infra/internal/models"

	"github.com/redis/go-redis/v9"
)

var (
	ctx        = context.Background()
	RedisQueue *redis.Client
	queueName  = "workflow_tasks"
)

func InitRedisQueue() {
	RedisQueue = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	_, err := RedisQueue.Ping(ctx).Result()
	if err != nil {
		panic("Failed to connect to Redis: " + err.Error())
	}

	fmt.Println("Connected to Redis")
}

func EnqueueTask(task models.Task) {
	taskBytes, err := json.Marshal(task)
	if err != nil {
		fmt.Println("Failed to marshal task:", err)
		return
	}

	err = RedisQueue.LPush(ctx, queueName, taskBytes).Err()
	if err != nil {
		fmt.Println("Failed to enqueue task:", err)
	}
}

func DequeueTask() (models.Task, error) {
	result, err := RedisQueue.BRPop(ctx, 0, queueName).Result()
	if err != nil {
		return models.Task{}, err
	}

	var task models.Task
	err = json.Unmarshal([]byte(result[1]), &task)
	if err != nil {
		return models.Task{}, err
	}

	return task, nil
}
