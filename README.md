# Redis ping client
A minimalistic client to make get requests to a redis server for a specific key

# Generate Executable
`go build -o redis-ping-client`

# Usage
```
redis-ping-client
    -HOST_NAME string
        hostname (default "__HOST_NAME__")
    -KEY string
        redis key to access (default "__KEY__")
    -PORT string
        port number for redis server (default "6379")
    -REQUEST_COUNT int
        total number of requests to make (default 50000)
    -BATCH_SIZE int
        Size of batch for concurrent requests (default 1)
    -h
        help option
```