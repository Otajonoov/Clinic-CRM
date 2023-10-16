package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	PatientServicePort string
	PatientServiceHost string
	LabServicePort string
	LabServiceHost string
	DoctorServicePort string
	DoctorServiceHost string
	Environment        string
	LogLevel           string
	PostgresHost       string
	PostgresPort       string
	PostgresUser       string
	PostgresPassword   string
	PostgresDatabase   string
	DatabaseUrl        string
}

func Load() Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("Error failed godotenv")
		panic(err)
	}

	c := Config{}
	c.PatientServiceHost = cast.ToString(GetOrReturnDefault("PATIENT_SERVICE_HOST", "localhost"))
	c.PatientServicePort = cast.ToString(GetOrReturnDefault("PATIENT_SERVICE_PORT", "5000"))
	c.DoctorServiceHost = cast.ToString(GetOrReturnDefault("DOCTOR_SERVICE_HOST", "localhost"))
	c.DoctorServicePort = cast.ToString(GetOrReturnDefault("DOCTOR_SERVICE_PORT", "5001"))
	c.LabServiceHost = cast.ToString(GetOrReturnDefault("LAB_SERVICE_HOST", "localhost"))
	c.LabServicePort = cast.ToString(GetOrReturnDefault("LAB_SERVICE_PORT", "5002"))
	c.Environment = cast.ToString(GetOrReturnDefault("ENVIRONMENT", "develop"))
	c.LogLevel = cast.ToString(GetOrReturnDefault("LOG_LEVEL", "debug"))
	c.PostgresHost = cast.ToString(GetOrReturnDefault("POSTGRES_HOST", "localhost"))
	c.PostgresPort = cast.ToString(GetOrReturnDefault("POSTGRES_PORT", "5432"))
	c.PostgresUser = cast.ToString(GetOrReturnDefault("POSTGRES_USER", "postgres"))
	c.PostgresPassword = cast.ToString(GetOrReturnDefault("POSTGRES_PASSWORD", "12345"))
	c.PostgresDatabase = cast.ToString(GetOrReturnDefault("POSTGRES_DATABSE", "medical_crm"))
	c.DatabaseUrl = cast.ToString(GetOrReturnDefault("DATABASE_URL", "postgres://azizbek:CloudA2023*@localhost:5432/medical_crm"))
	return c
}

func GetOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}
	return defaultValue
}
