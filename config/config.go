package config

type Config struct {
	PostgresHost     string
	PostgresPort     string
	PostgresUser     string
	PostgresPassword string
	PostgresDB       string
}

func Load() Config {
	cfg := Config{}

	cfg.PostgresHost = "localhost"
	cfg.PostgresPort = "5432"
	cfg.PostgresPassword = "1999"
	cfg.PostgresUser = "admin"
	cfg.PostgresDB = "shopingconnect"

	return cfg
}
