package main

import (
	"github.com/streadway/amqp"
)

type Channel struct {
	Channel *amqp.Channel
}

func GetChannel(conn *Conn) (Channel, error) {
	ch, err := conn.Connection.Channel()
	return Channel{
		Channel: ch,
	}, err
}
