package gomicrosvc

import (
	"fmt"
	"strings"

	guuid "github.com/google/uuid"
)

func RPC(routingKey string, message *[]byte) (*[]byte, error) {
	c, err := connection.getChannel()
	if err != nil {
		return nil, err
	}
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
	err = c.Publish(routingKey, fmt.Sprintf("%s-reply", Config.App), correlationId, message)
	if err != nil {
		return nil, err
	}

	result := <-msg

	return &result.Body, nil
}

func Publish(routingKey string, correlationId string, message *[]byte) error {
	c, err := connection.getChannel()
	if err != nil {
		return err
	}
	defer c.Channel.Close()

	err = c.Publish(routingKey, "", correlationId, message)
	if err != nil {
		return err
	}
	return nil
}
