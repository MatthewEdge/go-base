package env

import "os"

// GetString returns the ENV var at the given key, if found, or defaultValue if not found
func GetString(key, defaultValue string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	return val
}
