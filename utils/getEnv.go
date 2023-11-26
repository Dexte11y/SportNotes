package utils

import "os"

func DefaultGetEnv(key, defaultVal string) string {
	val := os.Getenv(key)
	if val == "" {
		val = defaultVal
	}
	return val
}
