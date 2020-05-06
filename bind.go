package gomicrosvc

func Bind() {
	ch, err := GetChannel(&Connection)
	if err != nil {
		panic(err)
	}

	err = ch.StartConsumer(2)
	if err != nil {
		panic(err)
	}

	forever := make(chan bool)
	<-forever
}
