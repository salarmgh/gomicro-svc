package gomicrosvc

import (
	"fmt"

	guuid "github.com/google/uuid"
)

func RPCCall(routingKey string, message *[]byte) (*[]byte, error) {
	uid := guuid.New().String()
	callerID := fmt.Sprintf("%s%s%s", Config.App,
		".reply_", uid)
	c := make(chan *[]byte)
	Channels[uid] = c

	rpcChan.Publish(routingKey, callerID, message)

	return <-c, nil
}

func AsyncRPCCall(routingKey string, message *[]byte) {
	rpcChan.Publish(routingKey, "", message)
}
