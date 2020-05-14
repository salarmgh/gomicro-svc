package gomicrosvc

type Rabbitmq struct {
	Host     string
	User     string
	Password string
	Exchange string
}

type Configuration struct {
	App      string
	Rabbitmq Rabbitmq
	Threads  int
}

var Config Configuration

func initConfig(app string, rabbitmqHost string, rabbitmqUser string,
	rabbitmqPass string, rabbitmqExchange string, threadsNumber int) {

	Config = Configuration{
		App: app,
		Rabbitmq: Rabbitmq{
			Host:     rabbitmqHost,
			User:     rabbitmqUser,
			Password: rabbitmqPass,
			Exchange: rabbitmqExchange,
		},
		Threads: threadsNumber,
	}
}
