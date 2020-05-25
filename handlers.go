package gomicrosvc

import (
	"reflect"
	"runtime"
	"strings"

	"github.com/streadway/amqp"
)

func registerHandlers(handlers []func(message amqp.Delivery) bool) {
	h := map[string]func(message amqp.Delivery) bool{}
	for _, function := range handlers {
		h[getFunctionName(function)] = function
	}
	Handlers = h
}

func getFunctionName(f interface{}) string {
	return strings.Split(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), ".")[1]
}
