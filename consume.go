package gomicrosvc

import (
	"fmt"
	"os"
)

func (conn Channel) StartConsumer(concurrency int) error {
	queueName := Config.App
	_, err := conn.Channel.QueueDeclare(queueName, true, false, false, false,
		nil)
	if err != nil {
		return err
	}

	err = conn.Channel.QueueBind(queueName, queueName+".*",
		Config.Rabbitmq.Exchange, false, nil)
	if err != nil {
		return err
	}

	msgs, err := conn.Channel.Consume(
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
