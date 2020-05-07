package gomicrosvc

import "fmt"

func RPCCall(routingKey string) string {
	ch, err := GetChannel(&Connection)
	if err != nil {
		panic(err)
	}

	uid := "5a4009ec-da68-42e0-8912-e53345e37444"
	queueName := fmt.Sprintf("%s%s%s", config.App,
		".reply_", uid)

	//result, err := ch.SyncConsumer("server.testHandler1", queueName,
	//	[]byte(`{"message":"handler 1"}`))
	//if err != nil {
	//	panic(err)
	//}

	c := make(chan string)
	Channels[uid] = c

	ch.Publish(routingKey, queueName, "2", []byte("testestedte"))

	return <-c
}

// StartConsumer -
func (conn Channel) SyncConsumer(
	dstRoutingKey string,
	queueName string,
	message []byte) (string, error) {

	// create the queue if it doesn't already exist
	_, err := conn.Channel.QueueDeclare(queueName, true, false, false, false,
		nil)
	if err != nil {
		return "", err
	}

	// bind the queue to the routing key
	err = conn.Channel.QueueBind(queueName, queueName, config.Rabbitmq.Exchange, false, nil)
	if err != nil {
		return "", err
	}

	msgs, err := conn.Channel.Consume(
		queueName, // queue
		"",        // consumer
		false,     // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	if err != nil {
		return "", err
	}

	ch, err := GetChannel(&Connection)
	if err != nil {
		panic(err)
	}
	ch.Publish("test.testHandler1", "1", "2", message)

	msg := <-msgs
	return string(msg.Body), nil
}
