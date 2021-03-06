package gomicrosvc

import (
	"fmt"
	"strings"

	guuid "github.com/google/uuid"
	"google.golang.org/protobuf/types/known/anypb"
)

type Data = anypb.Any

// Handlers map for dispatch
var Handlers map[string]func(data *Data) (*Data, error)

var connection broker

var replyQueue string

// Initialize gomicrosvc
func Initialize(app string, rabbitmqHost string, rabbitmqUser string,
	rabbitmqPass string, rabbitmqExchange string, threadsNumber int,
	handlers []func(data *Data) (*Data, error)) error {
	initConfig(app, rabbitmqHost, rabbitmqUser, rabbitmqPass, rabbitmqExchange,
		threadsNumber)

	replyID := strings.Replace(guuid.New().String(), "-", "", -1)

	replyQueue = fmt.Sprintf("%s-reply-%s", Config.App, replyID)

	err := connection.getConn()
	if err != nil {
		return err
	}

	registerHandlers(handlers)

	err = initRabbit()
	if err != nil {
		return err
	}

	return nil
}

func initRabbit() error {
	c, err := connection.getChannel()
	if err != nil {
		return err
	}
	defer c.Channel.Close()

	err = c.declareExchange("rpc-bus")
	if err != nil {
		return err
	}

	err = c.declareQueue(Config.App)
	if err != nil {
		return err
	}

	err = c.queueBind(Config.App, ".*")
	if err != nil {
		return err
	}

	return nil
}
