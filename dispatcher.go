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
	if h, ok := Handlers[handler]; ok {
		h(d)
	}
	return true
}
