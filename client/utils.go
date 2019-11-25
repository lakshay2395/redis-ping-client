package client

import (
	"errors"

	"github.com/go-redis/redis"
)

func generateBatches(requestCount int, batchSize int) ([]int, error) {
	if requestCount <= 0 {
		return nil, errors.New("Request count cannot be less than or equal to 0")
	}
	batches := []int{}
	for requestCount >= batchSize {
		batches = append(batches, batchSize)
		requestCount -= batchSize
	}
	if requestCount > 0 {
		batches = append(batches, requestCount)
	}
	return batches, nil
}

func wait(numberOfBatches int, worker, exit chan string) {
	for i := 0; i < numberOfBatches; i++ {
		<-worker
	}
	exit <- "Done"
}

func makeRequest(redisClient *redis.Client, key string) (string, string, error) {
	val, err := redisClient.Get(key).Result()
	if err != nil {
		return "", "", err
	}
	return "Success", val, nil
}
