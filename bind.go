package gomicrosvc

import "fmt"

func Start() {
	ch, err := GetChannel(&Connection)
	if err != nil {
		panic(err)
	}

	err = ch.StartConsumer(Config.Threads)
	if err != nil {
		panic(err)
	}

	fmt.Println("GoMicroSVC started ...")
}
