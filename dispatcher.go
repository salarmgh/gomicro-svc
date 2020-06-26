package gomicrosvc

import (
	"log"
	"strings"

	"github.com/streadway/amqp"
)

func dispatcher(message amqp.Delivery) bool {
	method := strings.Split(message.RoutingKey, ".")[1]
	reply := strings.Split(method, "_")

	if reply[0] == "reply" {
		go func() {
			callID := reply[1]
			if handler, ok := Channels[callID]; ok {
				handler <- &message.Body
				delete(Channels, callID)
			}
		}()
	} else {
		if handler, ok := Handlers[method]; ok {
			result, err := handler(&message.Body)
			if err != nil {
				log.Println(err)
				result.Error = err.Error()
			}
			AsyncRPCCall(message.ReplyTo, result)
		}
	}
	return true
}
