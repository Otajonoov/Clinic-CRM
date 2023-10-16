package config

import (
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	PostgresHost     string
	PostgresPort     string
	PostgresUser     string
	PostgresDatabase string
	PostgresPassword string
	Environment      string
	LogLevel         string
	DatabaseUrl      string
	LabServiceHost   string
	LabServicePort   string
}

func Load() Config {
	c := Config{}

	c.PostgresHost = cast.ToString(GetOrReturnDefault("POSTGRES_HOST", "localhost"))
	c.PostgresPort = cast.ToString(GetOrReturnDefault("POSTGRES_PORT", "5432"))
	c.PostgresUser = cast.ToString(GetOrReturnDefault("POSTGRES_USER", "azizbek"))
	c.PostgresPassword = cast.ToString(GetOrReturnDefault("POSTGRES_PASSWORD", "CloudA2023*"))
	c.PostgresDatabase = cast.ToString(GetOrReturnDefault("POSTGRES_DATABSE", "medical_crm"))
	c.DatabaseUrl = cast.ToString(GetOrReturnDefault("DATABASE_URL", "postgres://azizbek:CloudA2023*@localhost:5432/medical_crm"))
	c.Environment = cast.ToString(GetOrReturnDefault("ENVIRONMENT", "developer"))
	c.LogLevel = cast.ToString(GetOrReturnDefault("LOG_LEVEL", "debug"))
	c.LabServiceHost = cast.ToString(GetOrReturnDefault("LAB_SERVICE_HOST", "localhost"))
	c.LabServicePort = cast.ToString(GetOrReturnDefault("LAB_SERVICE_PORT", "5002"))

	return c
}

func GetOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}
	return defaultValue
}

