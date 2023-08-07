package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	RabbitmqHost     string `envconfig:"RABBITMQ_HOST"`
	RabbitmqPort     int    `envconfig:"RABBITMQ_PORT"`
	RabbitmqVHost    string `envconfig:"RABBITMQ_VHOST"`
	RabbitmqUser     string `envconfig:"RABBITMQ_USERNAME"`
	RabbitmqPassword string `envconfig:"RABBITMQ_PASSWORD"`
	RabbitmqQueue    string `envconfig:"RABBITMQ_QUEUE"`
}

var CFG Config

func Setup() {
	var config Config
	if err := envconfig.Process("", &config); err != nil {
		panic(err)
	}
	CFG = config
}
