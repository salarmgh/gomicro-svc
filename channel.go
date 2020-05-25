package gomicrosvc

import (
	"fmt"
	"os"

	"github.com/streadway/amqp"
)

type channel struct {
	channel *amqp.Channel
}

func (c channel) Publish(routingKey string, replyTo string,
	data []byte) error {
	return c.channel.Publish(
		Config.Rabbitmq.Exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType:  "application/octet-stream",
			ReplyTo:      replyTo,
			Priority:     0,
			Body:         data,
			DeliveryMode: amqp.Persistent,
		})
}

func (c channel) StartConsumer(concurrency int) error {
	err := c.declareExchange("rpc-bus")
	if err != nil {
		return err
	}

	err = c.declareQueue(Config.App)
	if err != nil {
		return err
	}

	err = c.consume(Config.App, Config.Concurrency)
	if err != nil {
		return err
	}

	return nil
}

func (c channel) declareExchange(exchangeName string) error {
	return c.channel.ExchangeDeclare(
		exchangeName,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
}

func (c channel) declareQueue(queueName string) error {
	_, err := c.channel.QueueDeclare(queueName, true, false, false, false,
		nil)
	if err != nil {
		return err
	}
	return nil
}

func (c channel) queueBind(queueName string, pattern string) error {
	err := c.channel.QueueBind(queueName, queueName+".*",
		Config.Rabbitmq.Exchange, false, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c channel) consume(queueName string, concurrency int) error {
	msgs, err := c.channel.Consume(
		queueName,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	for i := 0; i < concurrency; i++ {
		go func() {
			for msg := range msgs {
				if dispatcher(msg) {
					msg.Ack(false)
				} else {
					msg.Nack(false, true)
				}
			}
			fmt.Println("Rabbit consumer closed - critical Error")
			os.Exit(1)
		}()
	}
	return nil
}
