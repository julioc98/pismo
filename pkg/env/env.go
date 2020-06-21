package env

import "os"

// Get get a env var or set to default
func Get(env string, defaultEnv string) string {
	if e := os.Getenv(env); e != "" {
		return os.Getenv(env)
	}
	return defaultEnv
}
