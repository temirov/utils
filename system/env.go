package system

import (
	"log"
	"os"
)

func GetEnvOrFail(name string) string {
	value := os.Getenv(name)
	if value == "" {
		log.Fatalf("%s environment variable not set", name)
	}
	return value
}
