package config

type AppConfig struct {
	WorkerCount   int
	MaxRetries    int
	FailureRate   float64
	WorkerDelayMs int
	CrashRate     float64
}

var Config = AppConfig{
	WorkerCount:   3,
	MaxRetries:    3,
	FailureRate:   0.50,
	WorkerDelayMs: 1000,
	CrashRate:     0.10,
}
