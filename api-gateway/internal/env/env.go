package env

import (
	"log"
	"os"
)

func Get(name string, defaultValue string) string {
	v, ok := os.LookupEnv(name)
	if !ok {
		if defaultValue == "" {
			log.Fatalf("missing required environment variable " + name)
		}

		return defaultValue
	}
	return v
}
