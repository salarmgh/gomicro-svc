package gomicrosvc

import (
	"fmt"

	guuid "github.com/google/uuid"
)

func RPCCall(routingKey string, message string) string {
	ch, err := GetChannel(&Connection)
	if err != nil {
		panic(err)
	}

	uid := guuid.New().String()
	callerID := fmt.Sprintf("%s%s%s", config.App,
		".reply_", uid)

	c := make(chan string)
	Channels[uid] = c

	ch.Publish(routingKey, callerID, []byte(message))

	return <-c
}

func AsyncRPCCall(routingKey string, message string) {
	ch, err := GetChannel(&Connection)
	if err != nil {
		panic(err)
	}

	ch.Publish(routingKey, "", []byte(message))
}
