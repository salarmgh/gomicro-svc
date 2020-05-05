package gomicrosvc

import (
	"github.com/streadway/amqp"
)

func bindHandler(key string, handler func(d amqp.Delivery) bool) {
	const BASE_NAME = "test" + "."
	ch, err := GetChannel(&Connection)
	if err != nil {
		panic(err)
	}

	err = ch.StartConsumer(BASE_NAME+key, handler, 2)

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
