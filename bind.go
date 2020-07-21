package gomicrosvc

import (
	"log"
)

// Bind starts consumer
func Start() error {
	log.Println("GoMicroSVC started")
	ch, err := connection.getChannel()
	if err != nil {
		return err
	}
	err = ch.StartConsumer(Config.App)
	return err
}
