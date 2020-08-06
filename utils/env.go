package utils

import (
	"os"
)

// EnvGet returns env variable or fallback.
func EnvGet(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}
