package gomicrosvc

func Start() {
	ch, err := GetChannel(&Connection)
	if err != nil {
		panic(err)
	}

	err = ch.StartConsumer(config.Threads)
	if err != nil {
		panic(err)
	}

}
