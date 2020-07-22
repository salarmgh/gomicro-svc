package gomicrosvc

import (
	"fmt"
	"log"
	"strings"

	guuid "github.com/google/uuid"
)

func RPC(routingKey string, message *[]byte) (*[]byte, error) {
	c, err := connection.getChannel()
	if err != nil {
		return nil, err
	}
	log.Println("FIRST")
	defer c.Channel.Close()

	correlationId := strings.Replace(guuid.New().String(), "-", "", -1)
	msg, err := c.Channel.Consume(
		fmt.Sprintf("%s-reply", Config.App),
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return nil, err
	}
	log.Println("SECOND")
	err = c.Publish(routingKey, fmt.Sprintf("%s-reply", Config.App), correlationId, message)
	if err != nil {
		return nil, err
	}
	log.Println("THIRD")
	result := <-msg
	log.Println(result.Body)

	return &result.Body, nil
}

func Publish(routingKey string, correlationId string, message *[]byte) error {
	c, err := connection.getChannel()
	if err != nil {
		return err
	}
	defer c.Channel.Close()
	log.Println("==== Publish ====")
	log.Println(routingKey)
	log.Println(correlationId)
	log.Println(message)
	err = c.Publish(routingKey, "", correlationId, message)
	if err != nil {
		return err
	}
	return nil
}
