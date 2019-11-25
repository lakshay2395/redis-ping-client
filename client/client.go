package client

import (
	"log"

	"github.com/go-redis/redis"
	"github.com/lakshay2395/redis-ping-client/counter"
	"github.com/lakshay2395/redis-ping-client/param"
)

//Client - client to execute operation
type Client interface {
	Execute()
}

type client struct {
	parameters  param.Parameters
	redisClient *redis.Client
	counter     counter.Counter
}

func New(parameters param.Parameters) *client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     parameters.Host() + ":" + parameters.Port(),
		Password: "",
		DB:       0,
	})
	return &client{
		parameters:  parameters,
		redisClient: redisClient,
		counter:     counter.New(),
	}
}

func (c *client) Execute() error {
	batches, err := generateBatches(c.parameters.RequestCount(), c.parameters.BatchSize())
	if err != nil {
		return err
	}
	c.executeBatches(batches)
	return nil
}

func (c *client) executeBatches(batches []int) {
	exit := make(chan string)
	worker := make(chan string, len(batches))
	go wait(len(batches), worker, exit)
	for _, batchSize := range batches {
		go c.executeBatch(batchSize, worker)
	}
	<-exit
}

func (c *client) executeBatch(batchSize int, worker chan string) {
	for i := 0; i < batchSize; i++ {
		status, body, err := makeRequest(c.redisClient, c.parameters.Key())
		c.counter.Increment()
		if err != nil {
			log.Println("Response[", c.counter.Get(), "]: Error occurred while making request = ", err)
		} else {
			log.Println("Response[", c.counter.Get(), "]: status = ", status, " , body = ", body)
		}
	}
	worker <- "Done"
}
