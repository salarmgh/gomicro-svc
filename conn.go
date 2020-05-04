package gomicrosvc

import (
	"github.com/streadway/amqp"
)

type Conn struct {
	Connection *amqp.Connection
}

func GetConn(rabbitURL string) (Conn, error) {
	conn, err := amqp.Dial(rabbitURL)
	if err != nil {
		return Conn{}, err
	}

	return Conn{
		Connection: conn,
	}, err
}
