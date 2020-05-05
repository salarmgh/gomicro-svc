package gomicrosvc

import (
	"strconv"

	"github.com/streadway/amqp"
)

func bindHandler(key string, handler func(d amqp.Delivery) bool) {
	ch, err := GetChannel(&Connection)
	if err != nil {
		panic(err)
	}

	err = ch.StartConsumer(key, handler, 2)

	if err != nil {
		panic(err)
	}

	forever := make(chan bool)
	<-forever
}

func Binder(handlers []func(d amqp.Delivery) bool) {
	for i, handler := range handlers {
		go bindHandler("handler-"+strconv.Itoa(i), handler)
	}
}
