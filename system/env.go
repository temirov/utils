package system

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func GetEnvOrFail(name string) string {
	value := os.Getenv(name)
	if value == "" {
		log.Fatalf("%s environment variable not set", name)
	}
	return value
}

func ExpandEnvVar(envVar string) (string, error) {
	trimmedEnvVar := strings.TrimSpace(envVar)
	if envValue := os.ExpandEnv(trimmedEnvVar); envValue != "" {
		return strings.TrimSpace(envValue), nil
	}
	return "", fmt.Errorf("environment variable %s is not setup", trimmedEnvVar)
}
