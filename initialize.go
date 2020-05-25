package gomicrosvc

// Handlers map for dispatch
var Handlers map[string]func(data *[]byte) (*[]byte, error)

// Channels for Sync RPC
var Channels map[string]chan string

var connection broker
var rpcChan *channel

// Initialize gomicrosvc
func Initialize(app string, rabbitmqHost string, rabbitmqUser string,
	rabbitmqPass string, rabbitmqExchange string, threadsNumber int,
	handlers []func(data *[]byte) (*[]byte, error)) error {
	initConfig(app, rabbitmqHost, rabbitmqUser, rabbitmqPass, rabbitmqExchange,
		threadsNumber)

	err := connection.getConn()
	if err != nil {
		return err
	}

	rch, err := connection.getChannel()
	if err != nil {
		return err
	}
	rpcChan = rch

	registerHandlers(handlers)

	Channels = make(map[string]chan string)
	return nil
}
