package recovery

import (
	"os"
	"time"

	"self-healing-agent-infra/internal/metrics"
	"self-healing-agent-infra/internal/queue"
)

type Strategy string

const (
	NoRecovery         Strategy = "NO_RECOVERY"
	FixedRetry         Strategy = "FIXED_RETRY"
	ExponentialBackoff Strategy = "EXPONENTIAL_BACKOFF"
	AdaptiveRecovery   Strategy = "ADAPTIVE_RECOVERY"
)

const (
	highFailureRateThreshold = 0.30
	highQueueDepthThreshold  = int64(50)
	minimumObservedTasks     = 10
)

// Decision records the runtime conditions used by Adaptive Recovery.
type Decision struct {
	ConfiguredStrategy Strategy
	SelectedStrategy   Strategy
	FailureRate        float64
	QueueDepth         int64
	Reason             string
}

func GetStrategyFromEnv() Strategy {
	strategy := Strategy(os.Getenv("RECOVERY_STRATEGY"))

	switch strategy {
	case NoRecovery:
		return NoRecovery
	case FixedRetry:
		return FixedRetry
	case ExponentialBackoff:
		return ExponentialBackoff
	case AdaptiveRecovery:
		return AdaptiveRecovery
	default:
		return FixedRetry
	}
}

// ResolveStrategy chooses the effective strategy using current system state.
//
// Adaptive rules:
//  1. High cumulative failure rate -> Exponential Backoff.
//  2. Large queue backlog -> Fixed Retry.
//  3. Normal conditions -> Hybrid Adaptive Recovery.
func ResolveStrategy(configured Strategy) Decision {
	decision := Decision{
		ConfiguredStrategy: configured,
		SelectedStrategy:   configured,
		Reason:             "configured strategy selected",
	}

	if configured != AdaptiveRecovery {
		return decision
	}

	systemState := metrics.GetSystemMetrics()
	decision.FailureRate = systemState.FailureRate

	queueDepth, err := queue.GetQueueLength()
	if err == nil {
		decision.QueueDepth = queueDepth
	}

	if systemState.TotalTasks >= minimumObservedTasks &&
		systemState.FailureRate >= highFailureRateThreshold {
		decision.SelectedStrategy = ExponentialBackoff
		decision.Reason = "high system failure rate detected"
		return decision
	}

	if decision.QueueDepth >= highQueueDepthThreshold {
		decision.SelectedStrategy = FixedRetry
		decision.Reason = "large queue backlog detected"
		return decision
	}

	decision.SelectedStrategy = AdaptiveRecovery
	decision.Reason = "normal system conditions"
	return decision
}

func GetRetryDelay(strategy Strategy, retryCount int) time.Duration {
	if retryCount < 1 {
		retryCount = 1
	}

	switch strategy {
	case NoRecovery:
		return 0

	case FixedRetry:
		return 500 * time.Millisecond

	case ExponentialBackoff:
		return time.Duration(500*(1<<(retryCount-1))) * time.Millisecond

	case AdaptiveRecovery:
		// Hybrid behavior under normal system conditions:
		// Retry 1: 500 ms
		// Retry 2: 500 ms
		// Retry 3: 1 second
		// Retry 4: 2 seconds
		// Retry 5: 4 seconds
		switch retryCount {
		case 1, 2:
			return 500 * time.Millisecond
		case 3:
			return time.Second
		default:
			return time.Duration(1<<(retryCount-3)) * time.Second
		}

	default:
		return 500 * time.Millisecond
	}
}

func ShouldRetry(strategy Strategy, retryCount int, maxRetries int) bool {
	if strategy == NoRecovery {
		return false
	}

	return retryCount < maxRetries
}
