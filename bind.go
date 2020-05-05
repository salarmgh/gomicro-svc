package gomicrosvc

import (
	"fmt"

	"github.com/streadway/amqp"
)

func bindHandler(key string, handler func(d amqp.Delivery) bool) {
	routingKey := fmt.Sprintf("%s.%s", config.App, key)
	ch, err := GetChannel(&Connection)
	if err != nil {
		panic(err)
	}

	err = ch.StartConsumer(routingKey, handler, 2)

	if err != nil {
		panic(err)
	}

	forever := make(chan bool)
	<-forever
}

func Binder(handlers []func(d amqp.Delivery) bool) {
	for _, handler := range handlers {
		go bindHandler(getFunctionName(handler), handler)
	}
}
