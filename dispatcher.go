package gomicrosvc

import (
	"fmt"
	"strings"

	"github.com/streadway/amqp"
)

func Dispatcher(d amqp.Delivery) bool {
	if d.Body == nil {
		fmt.Println("Error, no message body!")
		return false
	}

	handler := strings.Split(d.RoutingKey, ".")[1]
	isReply := strings.Split(handler, "_")
	if isReply[0] == "reply" {
		go func() {
			if h, ok := Channels[isReply[1]]; ok {
				h <- string(d.Body)
				delete(Channels, isReply[1])
			}
		}()
	} else {
		if h, ok := Handlers[handler]; ok {
			go h(d)
		}
	}
	return true
}
