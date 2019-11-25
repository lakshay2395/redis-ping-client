package param

type Parameters interface {
	Host() string
	Key() string
	Port() string
	RequestCount() int
	BatchSize() int
}

//Parameter - default parameter structure
type params struct {
	host         string
	key          string
	port         string
	requestCount int
	batchSize    int
}

//New - get new parameter object
func New(host string, key string, port string, requestCount int, batchSize int) params {
	return params{
		host:         host,
		key:          key,
		port:         port,
		requestCount: requestCount,
		batchSize:    batchSize,
	}
}

//Host - return host name
func (p params) Host() string {
	return p.host
}

//key - returns key to access
func (p params) Key() string {
	return p.key
}

//port - return port number
func (p params) Port() string {
	return p.port
}

//requestCount - returns count of requests to execute
func (p params) RequestCount() int {
	return p.requestCount
}

//batchSize - returns size of batches for workers to execute
func (p params) BatchSize() int {
	return p.batchSize
}
