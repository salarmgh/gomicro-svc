package gomicrosvc

import (
	"reflect"
	"runtime"
	"strings"
)

func registerHandlers(handlers []func(data *[]byte) (*[]byte, error)) {
	h := map[string]func(data *[]byte) (*[]byte, error){}
	for _, function := range handlers {
		h[getFunctionName(function)] = function
	}
	Handlers = h
}

func getFunctionName(f interface{}) string {
	return strings.Split(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), ".")[1]
}
