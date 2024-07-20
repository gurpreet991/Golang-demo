package request_stream

import (
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

var (
	M                     int
	N                     int
	K                     int
	U                     int
	RequestsPerUser       int
	RequestProcessingTime int
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Getting error in loading .env file: %v", err)
	}

	M = convertEnvToInt("REQUESTS_PER_SECOND_PER_USER", 10)
	N = convertEnvToInt("REQUESTS_PER_MINTUE_PER_USER", 600)
	K = convertEnvToInt("REQUESTS_PER_MINTUE", 6000)
	U = convertEnvToInt("NUMBER_OF_USER", 6)
	RequestsPerUser = convertEnvToInt("REQUESTS_PER_USER", 4000)
	RequestProcessingTime = 1
}

func convertEnvToInt(name string, defaultValue int) int {
	valueStr := os.Getenv(name)
	if valueStr == "" {
		return defaultValue
	}
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		log.Fatalf("Error while converting env variavle %s to int: %v", name, err)
	}
	return value
}

var (
	Mu                sync.Mutex
	QueuedRequests    int
	ProcessedRequests int
	InputChannel      = make(chan Request)
)
