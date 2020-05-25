package gomicrosvc

import (
	"github.com/streadway/amqp"
)

// Handlers map for dispatch
var Handlers map[string]func(message amqp.Delivery) bool

// Channels for Sync RPC
var Channels map[string]chan string

var connection broker
var rpcChan channel

// Initialize gomicrosvc
func Initialize(app string, rabbitmqHost string, rabbitmqUser string,
	rabbitmqPass string, rabbitmqExchange string, threadsNumber int,
	handlers []func(message amqp.Delivery) bool) error {
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
