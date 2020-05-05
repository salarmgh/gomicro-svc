package gomicrosvc

import (
	"github.com/streadway/amqp"
)

type Conn struct {
	Connection *amqp.Connection
}

func GetConn(rabbitURL string) Conn {
	conn, err := amqp.Dial(rabbitURL)
	if err != nil {
		panic(err)
	}

	return Conn{
		Connection: conn,
	}
}
