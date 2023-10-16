package config

import (
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	DoctorServicePort string
	DoctorServiceHost string
	Environment       string
	LogLevel          string
	PostgresHost      string
	PostgresPort      string
	PostgresUser      string
	PostgresPassword  string
	PostgresDatabase  string
	DatabaseUrl       string
}

func Load() Config {
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatalln("Error failed godotenv")
	// 	panic(err)
	// }

	c := Config{}
	c.DoctorServiceHost = cast.ToString(GetOrReturnDefault("DOCTOR_SERVICE_HOST", "localhost"))
	c.DoctorServicePort = cast.ToString(GetOrReturnDefault("DOCTOR_SERVICE_PORT", "5001"))
	c.Environment = cast.ToString(GetOrReturnDefault("ENVIRONMENT", "develop"))
	c.LogLevel = cast.ToString(GetOrReturnDefault("LOG_LEVEL", "debug"))
	c.PostgresHost = cast.ToString(GetOrReturnDefault("POSTGRES_HOST", "localhost"))
	c.PostgresPort = cast.ToString(GetOrReturnDefault("POSTGRES_PORT", "5432"))
	c.PostgresUser = cast.ToString(GetOrReturnDefault("POSTGRES_USER", "azizbek"))
	c.PostgresPassword = cast.ToString(GetOrReturnDefault("POSTGRES_PASSWORD", "CloudA2023*"))
	c.PostgresDatabase = cast.ToString(GetOrReturnDefault("POSTGRES_DATABSE", "doctor_db"))
	c.DatabaseUrl = cast.ToString(GetOrReturnDefault("DATABASE_URL", "postgres://azizbek:CloudA2023*@localhost:5432/doctor_db"))
	return c
}

func GetOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}
	return defaultValue
}
