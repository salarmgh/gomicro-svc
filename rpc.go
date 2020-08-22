package gomicrosvc

import (
	"context"
	"errors"
	"log"
	"strings"
	"time"

	guuid "github.com/google/uuid"
	"github.com/streadway/amqp"
	"google.golang.org/protobuf/proto"
)

func RPC(routingKey string, message *Data) (*Data, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	go func() {

		select {
		case <-ctx.Done():
			err := ctx.Err()
			log.Println(err)
			break
		}

	}()

	c, err := connection.getChannel()
	if err != nil {
		return nil, err
	}

	defer c.Channel.Close()

	q, err := c.callBackQueue()
	if err != nil {
		return nil, err
	}

	correlationId := strings.Replace(guuid.New().String(), "-", "", -1)
	msgs, err := c.Channel.Consume(
		q.Name,
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

	log.Printf("RPC publish routingKey: %s, replyQueue: %s, corrID: %s", routingKey, q.Name, correlationId)
	err = c.Publish(routingKey, q.Name, correlationId, &marshalledData)
	if err != nil {
		return nil, err
	}

	var result amqp.Delivery
	for d := range msgs {
		if correlationId == d.CorrelationId {
			log.Printf("my %s == dst %s", correlationId, d.CorrelationId)
			result = d
			break
		}
		log.Printf("my %s != dst %s", correlationId, d.CorrelationId)
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
	err = c.CallBackPublish(routingKey, "", correlationId, message)
	if err != nil {
		return err
	}
	return nil
}
