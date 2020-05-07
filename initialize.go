package gomicrosvc

import (
	"fmt"

	"github.com/streadway/amqp"
)

var Connection Conn
var Handlers map[string]func(message amqp.Delivery) bool
var Channels map[string]chan string

func Initialize(handlers []func(message amqp.Delivery) bool) {
	initConfig()

	rabbitmqURI := fmt.Sprintf("amqp://%s:%s@%s", config.Rabbitmq.User,
		config.Rabbitmq.Password, config.Rabbitmq.Host)
	conn, err := GetConn(rabbitmqURI)
	if err != nil {
		panic(err)
	}
	Connection = conn

	registerHandlers(handlers)

	Channels = make(map[string]chan string)

	err = declareExchange(config.Rabbitmq.Exchange)
	if err != nil {
		panic(err)
	}
}

func declareExchange(exchangeName string) error {
	ch, err := GetChannel(&Connection)
	if err != nil {
		panic(err)
	}
	return ch.Channel.ExchangeDeclare(
		exchangeName,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
}

func registerHandlers(handlers []func(message amqp.Delivery) bool) {
	h := map[string]func(message amqp.Delivery) bool{}
	for _, function := range handlers {
		h[getFunctionName(function)] = function
	}
	Handlers = h
}
