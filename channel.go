package gomicrosvc

import (
	"errors"
	"log"

	"github.com/streadway/amqp"
)

type channel struct {
	Channel *amqp.Channel
}

func (c channel) Publish(routingKey string, replyTo string, correlationId string,
	data *[]byte) error {
	return c.Channel.Publish(
		Config.Rabbitmq.Exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType:   "application/octet-stream",
			ReplyTo:       replyTo,
			CorrelationId: correlationId,
			Priority:      0,
			Body:          *data,
			DeliveryMode:  amqp.Persistent,
		})
}

//func (c channel) RPCCall(routingKey string, data *[]byte) error {
//
//}

func (c channel) StartConsumer(queueName string) error {
	err := c.consume(queueName)
	if err != nil {
		return err
	}

	return nil
}

func (c channel) ClientConsumer(queueName string) error {
	err := c.consume(queueName)
	if err != nil {
		return err
	}

	return nil
}

func (c channel) declareExchange(exchangeName string) error {
	return c.Channel.ExchangeDeclare(
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
	_, err := c.Channel.QueueDeclare(queueName, true, false, false, false,
		nil)
	if err != nil {
		return err
	}
	return nil
}

func (c channel) queueBind(queueName string, path string) error {
	err := c.Channel.QueueBind(queueName, queueName+path,
		Config.Rabbitmq.Exchange, false, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c channel) consume(queueName string, exclusive bool) error {
	msgs, err := c.Channel.Consume(
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

	for msg := range msgs {
		err = msg.Ack(false)
		if err != nil {
			log.Println(err)
		}
		dispatcher(msg)
	}

	return errors.New("Consumer exited")
}
