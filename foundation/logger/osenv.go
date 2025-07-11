package logger

import "os"

func getEnvOrKey(key string) string {
	val := os.Getenv(key)
	if val == "" {
		return "${" + key + "}"
	}
	return val
}
