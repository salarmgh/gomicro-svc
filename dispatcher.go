package gomicrosvc

import (
	"log"
	"strings"

	"github.com/streadway/amqp"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func dispatcher(message amqp.Delivery) bool {
	go func() {
		method := strings.Split(message.RoutingKey, ".")[1]
		if handler, ok := Handlers[method]; ok {
			msg := Message{}
			err := proto.Unmarshal(message.Body, &msg)
			result, err := handler(msg.Result)
			if err != nil {
				log.Println(err)
				res := anypb.Any{}
				data := Message{Result: &res, Error: err.Error()}
				marshalledData, err := proto.Marshal(&data)
				if err != nil {
					log.Println(err)
					res := anypb.Any{}
					data := Message{Result: &res, Error: "Couldn't Marshal"}
					marshalledData, _ := proto.Marshal(&data)
					err = Publish(message.ReplyTo, message.CorrelationId, &marshalledData)
					//return true
				}
				err = Publish(message.ReplyTo, message.CorrelationId, &marshalledData)
				//return true
			}
			data := Message{Result: result, Error: ""}
			marshalledData, err := proto.Marshal(&data)
			if err != nil {
				log.Println(err)
				res := anypb.Any{}
				data := Message{Result: &res, Error: "Couldn't Marshal"}
				marshalledData, _ := proto.Marshal(&data)
				err = Publish(message.ReplyTo, message.CorrelationId, &marshalledData)
				//return true
			}
			log.Printf("Sent reply to %s, corrID: %s", message.ReplyTo, message.CorrelationId)
			err = Publish(message.ReplyTo, message.CorrelationId, &marshalledData)
			//return true
		}
	}()

	return true
}
