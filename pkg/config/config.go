package config

import (
	"auth/internal/storages"
	"fmt"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	IsDebug *bool `yaml:"is_debug" env-default:"false"`

	DBConfig storages.PostgresConfig `yaml:"database"`

	JWT struct {
		Secret string `yaml:"secret_key" env-required:"true"`
	} `yaml:"jwt"`

	SMTP struct {
		Host     string `yaml:"smtp_host" env-required:"true"`
		Port     int    `yaml:"smtp_port" env-required:"true"`
		Email    string `yaml:"smtp_email" env-required:"true"`
		Password string `yaml:"smtp_password" env-required:"true"`
	} `yaml:"smtp"`

	Services struct {
		UserService string `yaml:"user_service" env-required:"true"`
	} `yaml:"services"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{}
		err := cleanenv.ReadConfig("./config.yml", instance)
		fmt.Println(instance.JWT)
		if err != nil {
			errorHandler, _ := cleanenv.GetDescription(instance, nil)
			fmt.Println(errorHandler)
		}
	})

	return instance
}
