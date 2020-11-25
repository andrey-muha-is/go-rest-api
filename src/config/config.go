package config

import (
	"go.uber.org/zap"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	AppPort         string
	DbHost          string
	DbPort          string
	DbUser          string
	DbPassword      string
	DbName          string
	DbChannelsTable string
	DbProgramsTable string
}

// GetAppConfig - init app config with env variables
func GetAppConfig() AppConfig {
	err := godotenv.Load()
	if err != nil {
		zap.S().Error("Error loading .env file")
	}

	config := AppConfig{
		AppPort: 		 getEnvWithFallback("APP_PORT", "3000"),
		DbHost: 		 getEnvWithFallback("DB_HOST", "localhost"),
		DbPort: 		 getEnvWithFallback("DB_PORT", "3306"),
		DbUser: 		 getEnvWithFallback("DB_USER", "root"),
		DbPassword: 	 getEnvWithFallback("DB_PASSWORD", ""),
		DbName: 		 getEnvWithFallback("DB_NAME", "default"),
		DbChannelsTable: getEnvWithFallback("DB_CHANNELS_TABLE", "channel"),
		DbProgramsTable: getEnvWithFallback("DB_PROGRAMS_TABLE", "program"),
	}

	return config
}

func getEnvWithFallback(envVariableKey string, fallbackValue string) string {
	value, ok := os.LookupEnv(envVariableKey)
	if !ok {
		zap.S().Errorf("Error loading env variable %s. Using fallback value = %s", envVariableKey, fallbackValue)
		return fallbackValue
	}
	return value
}
