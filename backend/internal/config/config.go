package config

type AppConfig struct {
	WorkerCount int
	MaxRetries  int
}

var Config = AppConfig{
	WorkerCount: 3,
	MaxRetries:  3,
}
