package config

type Config struct {
	RESTPort      string
	CollectorAddr string
}

func Load() (Config, error) {
	return Config{
		RESTPort:      "8080",
		CollectorAddr: "localhost:50051",
	}, nil
}
