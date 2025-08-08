package config

import "os"

func IsProduction() bool {
	appEnv := os.Getenv("APP_ENV")

	if appEnv == "PRODUCTION" || appEnv == "STAGING" {
		return true
	}

	return false
}
