package envutils

import (
	"log"
	"os"
	"strconv"
)

func MustGetString(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Environment variable %s not set", key)
	}
	return value
}

func MustGetInt64(key string) int64 {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Environment variable %s not set", key)
	}

	intValue, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		log.Fatalf("Error parsing %s: %v", key, err)
	}

	return intValue
}

func MustGetInt32(key string) int32 {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Environment variable %s not set", key)
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		log.Fatalf("Error parsing %s: %v", key, err)
	}

	return int32(intValue)
}

func MustGetInt(key string) int {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Environment variable %s not set", key)
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		log.Fatalf("Error parsing %s: %v", key, err)
	}

	return intValue
}

func MustGetBool(key string) bool {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Environment variable %s not set", key)
	}

	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		log.Fatalf("Error parsing %s: %v", key, err)
	}

	return boolValue
}

func MustGetf64(key string) float64 {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Environment variable %s not set", key)
	}

	fValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		log.Fatalf("Error parsing %s: %v", key, err)
	}

	return fValue
}

func MustGetf32(key string) float32 {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Environment variable %s not set", key)
	}

	fValue, err := strconv.ParseFloat(value, 32)
	if err != nil {
		log.Fatalf("Error parsing %s: %v", key, err)
	}

	return float32(fValue)
}
