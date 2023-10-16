package config

import (
	"os"

	"github.com/spf13/cast"
)

// Config...
type Config struct {
	Environment string

	PatientServiceHost string
	PatientServicePort int
	DoctorServiceHost  string
	DoctorServicePort  int
	LabServiceHost     string
	LabServicePort     int
	MaxImageSize       int
	BaseUrl            string

	CtxTimeout int // context timeout in second

	LogLevel string
	HTTPPort string
}

func Load() Config {
	c := Config{}

	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))
	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debud"))
	c.HTTPPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ":8080"))

	c.PatientServiceHost = cast.ToString(getOrReturnDefault("PATIENT_SERVICE_HOST", "reception_service"))
	c.PatientServicePort = cast.ToInt(getOrReturnDefault("PATIENT_SERVICE_PORT", 5000))
	c.DoctorServiceHost = cast.ToString(getOrReturnDefault("DOCTOR_SERVICE_HOST", "medical_doctor_service"))
	c.DoctorServicePort = cast.ToInt(getOrReturnDefault("DOCTOR_SERVICE_PORT", 5001))
	c.LabServiceHost = cast.ToString(getOrReturnDefault("LAB_SERVICE_HOST", "medical_lab_service"))
	c.LabServicePort = cast.ToInt(getOrReturnDefault("LAB_SERVICE_PORT", 5002))
	c.MaxImageSize = cast.ToInt(getOrReturnDefault("MAX_IMAGE_SIZE", 5))
	c.CtxTimeout = cast.ToInt(getOrReturnDefault("CTX_TIMEOUT", 7))
	c.BaseUrl = cast.ToString(getOrReturnDefault("BASE_URL", "https://medical.samandardev.uz/v1/"))

	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
