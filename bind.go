package gomicrosvc

import (
	"log"
)

// Bind starts consumer
func Bind(foreground bool) error {
	ch, err := connection.getChannel()
	if err != nil {
		return err
	}

	err = ch.StartConsumer(Config.Concurrency)
	if err != nil {
		return err
	}

	log.Println("GoMicroSVC started")
	if foreground {
		forever := make(chan bool)
		<-forever
	}

	return nil
}
