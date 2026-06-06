package recovery

import "time"

type Strategy string

const (
	NoRecovery         Strategy = "NO_RECOVERY"
	FixedRetry         Strategy = "FIXED_RETRY"
	ExponentialBackoff Strategy = "EXPONENTIAL_BACKOFF"
)

func GetRetryDelay(strategy Strategy, retryCount int) time.Duration {
	switch strategy {
	case NoRecovery:
		return 0
	case FixedRetry:
		return 500 * time.Millisecond
	case ExponentialBackoff:
		return time.Duration(500*(1<<retryCount)) * time.Millisecond
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
