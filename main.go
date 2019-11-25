package main

import (
	"flag"
	"log"
	"time"

	"github.com/lakshay2395/redis-ping-client/client"
	"github.com/lakshay2395/redis-ping-client/param"
)

func main() {
	host := flag.String("HOST_NAME", "__HOST_NAME__", "hostname")
	port := flag.String("PORT", "6379", "port number for redis server")
	key := flag.String("KEY", "__KEY__", "redis key to access")
	requestCount := flag.Int("REQUEST_COUNT", 50000, "total number of requests to make")
	batchSize := flag.Int("BATCH_SIZE", 1, "Size of batch for concurrent requests")
	flag.Parse()
	startTime := time.Now()
	c := client.New(param.New(*host, *key, *port, *requestCount, *batchSize))
	err := c.Execute()
	if err != nil {
		log.Fatal("Error : ", err)
	}
	endTime := time.Now()
	log.Println("Time Taken = ", endTime.Unix()-startTime.Unix())
}
