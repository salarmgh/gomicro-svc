package gomicrosvc

import (
	"github.com/streadway/amqp"
)

type Conn struct {
	Connection *amqp.Connection
}

func GetConn(rabbitURL string) (Conn, error) {
	conn, err := amqp.Dial(rabbitURL)

	return Conn{
		Connection: conn,
	}, err
}
