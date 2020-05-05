package gomicrosvc

import (
	"github.com/streadway/amqp"
)

func (conn Channel) Publish(routingKey string, data []byte) error {
	return conn.Channel.Publish(
		// exchange - yours may be different
		"rpc-bus",
		routingKey,
		// mandatory - we don't care if there I no queue
		false,
		// immediate - we don't care if there is no consumer on the queue
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			Body:         data,
			DeliveryMode: amqp.Persistent,
		})
}
