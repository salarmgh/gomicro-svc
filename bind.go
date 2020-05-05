package gomicrosvc

import (
	"strconv"

	"github.com/streadway/amqp"
)

func bindHandler(ch *Channel, key string, handler func(d amqp.Delivery) bool) {
	err := ch.StartConsumer(key, handler, 2)

	if err != nil {
		panic(err)
	}

	forever := make(chan bool)
	<-forever
}

func Binder(ch *Channel, handlers []func(d amqp.Delivery) bool) {
	for i, handler := range handlers {
		go bindHandler(ch, "handler-"+strconv.Itoa(i), handler)
	}
}
