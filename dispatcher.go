package gomicrosvc

import (
	"log"
	"strings"

	"github.com/streadway/amqp"
)

func dispatcher(message amqp.Delivery) bool {
	if message.Body == nil {
		log.Println("Error, no message body!")
		return false
	}

	method := strings.Split(message.RoutingKey, ".")[1]
	reply := strings.Split(method, "_")

	if reply[0] == "reply" {
		go func() {
			callID := reply[1]
			if handler, ok := Channels[callID]; ok {
				handler <- string(message.Body)
				delete(Channels, callID)
			}
		}()
	} else {
		if handler, ok := Handlers[method]; ok {
			handler(message)
		}
	}
	return true
}
