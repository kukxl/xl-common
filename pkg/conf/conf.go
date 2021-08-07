package conf

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"path/filepath"
)

type Conf struct {
	Service ServiceConf `yaml:"services"`
	Client  ServiceConf `yaml:"client"`
}

type ServiceConf struct {
	App    AppConf    `yaml:"app"`
	Consul ConsulConf `yaml:"consul"`
	Logger LoggerConf `yaml:"logger"`
}

type AppConf struct {
	Name 		string	`yaml:"name"`
	Port 		int 	`yaml:"port"`
}

type ConsulConf struct {
	Addr 		string 	`yaml:"addr"`
	Timeout 	string `yaml:"timeOut"`
	Interval  	string `yaml:"interval"`
	OverTime 	string `yaml:"overTime"`
}

type LoggerConf struct {
	Model 		string	`yaml:"model"`
	Level 		string 	`yaml:"level"`
	FileName 	string	`yaml:"fileName"`
	MaxAge 		int		`yaml:"maxAge"`
	MaxBackups  int		`yaml:"maxBackups"`
	MaxSize		int		`yaml:"maxSize"`
}

var Config Conf

func ReadConf(configPath string)  {
	dir, err := filepath.Abs(".")
	if err != nil {
		panic(err)
	}
	yamlByte, err := ioutil.ReadFile(fmt.Sprintf("%s/conf/%s", dir, configPath))
	if err != nil {
		log.Fatalf("read yaml fail %v",err)
	}
	err = yaml.Unmarshal(yamlByte, &Config)
	if err != nil {
		log.Println(err.Error())
		panic(err)
	}
	log.Printf("config is init")
}