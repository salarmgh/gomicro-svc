package main

func initialize(ch *Channel) {
	ch.Channel.ExchangeDeclare(
		"rpc-bus", // name
		"topic",   // type
		true,      // durable
		false,     // auto-deleted
		false,     // internal
		false,     // no-wait
		nil,       // arguments
	)
}
