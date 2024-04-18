package main

import (
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	BotToken string `yaml:"token"`
}

func InitCfg() Config {
	var cfg Config = Config{}
	val, err := os.ReadFile("config.yml")
	if err != nil {
		log.Fatal().Msg("Can't read file - " + err.Error())
	}
	err = yaml.Unmarshal(val, &cfg)
	if err != nil {
		log.Fatal().Msg("Can't get value from config - " + err.Error())
	}
	return cfg
}
