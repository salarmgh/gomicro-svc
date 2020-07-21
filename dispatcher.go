package gomicrosvc

import (
	"log"
	"strings"

	"github.com/streadway/amqp"
)

func dispatcher(message amqp.Delivery) bool {
	method := strings.Split(message.RoutingKey, ".")[1]
	if handler, ok := Handlers[method]; ok {
		result, err := handler(&message.Body)
		if err != nil {
			log.Println(err)
		}
		err = Publish(message.ReplyTo, message.CorrelationId, result)
		if err != nil {
			return false
		}
	}

	return true
}
