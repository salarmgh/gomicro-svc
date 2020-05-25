package gomicrosvc

import (
	"fmt"

	"github.com/streadway/amqp"
)

type broker struct {
	Connection *amqp.Connection
}

func (b *broker) getConn() error {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s",
		Config.Rabbitmq.User, Config.Rabbitmq.Password, Config.Rabbitmq.Host))
	if err != nil {
		return err
	}

	b.Connection = conn
	return nil
}

func (b *broker) getChannel() (channel, error) {
	ch, err := b.Connection.Channel()
	return channel{
		channel: ch,
	}, err
}
