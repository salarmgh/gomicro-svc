package gomicrosvc

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Configuration struct {
	App      string `yaml:"app"`
	Rabbitmq struct {
		Host     string `yaml:"host"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Exchange string `yaml:"exchange"`
	} `yaml:"rabbitmq"`
}

var config Configuration

func initConfig() {
	file, err := ioutil.ReadFile("config.yml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(file, &config)
	if err != nil {
		panic(err)
	}
}
