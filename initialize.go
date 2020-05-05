package gomicrosvc

var Connection = GetConn("amqp://guest:guest@localhost")

func Initialize() {
	ch, err := GetChannel(&Connection)
	if err != nil {
		panic(err)
	}

	err = ch.Channel.ExchangeDeclare(
		"rpc-bus", // name
		"topic",   // type
		true,      // durable
		false,     // auto-deleted
		false,     // internal
		false,     // no-wait
		nil,       // arguments
	)

	if err != nil {
		panic(err)
	}
}
