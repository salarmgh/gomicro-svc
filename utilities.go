package gomicrosvc

import (
	"reflect"
	"runtime"
	"strings"
)

func getFunctionName(f interface{}) string {
	return strings.Split(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), ".")[1]
}
