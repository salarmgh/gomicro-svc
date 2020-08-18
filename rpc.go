package gomicrosvc

import (
	"errors"
	"log"
	"strings"

	guuid "github.com/google/uuid"
	"google.golang.org/protobuf/proto"
)

func RPC(routingKey string, message *Data) (*Data, error) {
	c, err := connection.getChannel()
	if err != nil {
		return nil, err
	}

	defer c.Channel.Close()

	err = c.declareQueue(replyQueue, true)
	if err != nil {
		return nil, err
	}

	err = c.queueBind(replyQueue, "")
	if err != nil {
		return nil, err
	}

	correlationId := strings.Replace(guuid.New().String(), "-", "", -1)
	msg, err := c.Channel.Consume(
		replyQueue,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return nil, err
	}

	data := Message{Result: message, Error: ""}
	marshalledData, err := proto.Marshal(&data)
	if err != nil {
		return nil, errors.New("Couldn't marshal")
	}

	err = c.Publish(routingKey, replyQueue, correlationId, &marshalledData)
	if err != nil {
		return nil, err
	}
	result := <-msg
	if result.CorrelationId == correlationId {
		log.Println("CorrelationId Matched")
	} else {
		log.Println("CorrelationId Not Matched")
	}
	resp := Message{}
	err = proto.Unmarshal(result.Body, &resp)

	if resp.Error != "" {
		return nil, errors.New(resp.Error)
	}

	return resp.Result, nil
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
