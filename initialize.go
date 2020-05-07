package gomicrosvc

import (
	"fmt"

	"github.com/streadway/amqp"
)

var Connection Conn

var Handlers map[string]func(d amqp.Delivery) bool
var Channels map[string]chan string

func Initialize(handlers []func(d amqp.Delivery) bool) {
	initConfig()
	h := map[string]func(d amqp.Delivery) bool{}
	for _, function := range handlers {
		h[getFunctionName(function)] = function
	}
	Handlers = h
	Channels = make(map[string]chan string)

	Connection = GetConn(fmt.Sprintf("amqp://%s:%s@%s", config.Rabbitmq.User,
		config.Rabbitmq.Password, config.Rabbitmq.Host))

	ch, err := GetChannel(&Connection)
	if err != nil {
		panic(err)
	}

	err = ch.Channel.ExchangeDeclare(
		config.Rabbitmq.Exchange, // name
		"topic",                  // type
		true,                     // durable
		false,                    // auto-deleted
		false,                    // internal
		false,                    // no-wait
		nil,                      // arguments
	)
	if err != nil {
		panic(err)
	}
}
