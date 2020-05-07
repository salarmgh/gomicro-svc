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
	Threads int `yaml:"threads"`
}

var config Configuration

func initConfig() {
	configFile, err := ioutil.ReadFile("config.yml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		panic(err)
	}
}
