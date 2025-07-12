package logger

import "os"

// function 'getEnvOrKey' returns the value of the environment variable with the given key,
// if the environment variable is not set, it returns the key name itself
func getEnvOrKey(key string) string {
	val := os.Getenv(key)
	if val == "" {
		return "${" + key + "}"
	}
	return val
}
