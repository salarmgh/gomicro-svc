package gomicrosvc

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var Config map[interface{}]interface{}

func InitConfig() {
	Config = make(map[interface{}]interface{})

	configFile, err := ioutil.ReadFile("config.yml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(configFile, &Config)
	if err != nil {
		panic(err)
	}
}
