package main

import (
	"fmt"
	"time"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := GetConn("amqp://guest:guest@localhost")
	if err != nil {
		panic(err)
	}

	ch, err := GetChannel(&conn)

	initialize(&ch)

	go func() {
		for {
			time.Sleep(time.Second)
			ch.Publish("test-key", []byte(`{"message":"test"}`))
		}
	}()

	err = ch.StartConsumer("test-queue", "test-key", handler, 2)

	if err != nil {
		panic(err)
	}

	forever := make(chan bool)
	<-forever
}

func handler(d amqp.Delivery) bool {
	if d.Body == nil {
		fmt.Println("Error, no message body!")
		return false
	}
	fmt.Println(string(d.Body))
	return true
}
