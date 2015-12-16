package core

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

var (
	Cfg        Config
	configFile = "config.yaml"
)

type Config struct {
	Global struct {
		Port string
	}
	Database struct {
		DriverName string `yaml:"driver_name"`
		Database   string
		UserName   string `yaml:"user_name"`
		Password   string
		Host       string
		Charset    string
		Engine     string
		Encoding   string
	}
	Kafka struct {
		Brokers      string
		DefaultTopic string `yaml:"default_topic"`
	}
}

func InitCfg() {

	data, fileReadErr := ioutil.ReadFile(configFile)
	if fileReadErr != nil {
		log.Fatalf("Failed to read", fileReadErr)
	}

	err := yaml.Unmarshal(data, &Cfg)

	if err != nil {
		log.Fatalf("Failed to parse yaml data: %s", err)
	}

}
