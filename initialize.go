package gomicrosvc

import "fmt"

var Connection Conn

func Initialize() {
	initConfig()

	Connection = GetConn(fmt.Sprintf("amqp://%s:%s@%s", config.Rabbitmq.User,
		config.Rabbitmq.Password, config.Rabbitmq.Host))

	ch, err := GetChannel(&Connection)
	if err != nil {
		panic(err)
	}

	err = ch.Channel.ExchangeDeclare(
		config.Rabbitmq.Exchange, // name
		"topic",                  // type
		true,                     // durable
		false,                    // auto-deleted
		false,                    // internal
		false,                    // no-wait
		nil,                      // arguments
	)

	if err != nil {
		panic(err)
	}
}
