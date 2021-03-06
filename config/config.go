package config

import (
	"log"
	"os"

	"gopkg.in/go-ini/ini.v1"
)

type ConfigList struct {
	Apikey    string
	ApiSecret string
	Logfile   string
}

var Config ConfigList

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Printf("Failed to read file: %v", err)
		os.Exit(1)
	}
	Config = ConfigList{
		Apikey:    cfg.Section("bitflyer").Key("api_key").String(),
		ApiSecret: cfg.Section("bitflyer").Key("api_secret").String(),
		Logfile:   cfg.Section("bitflyer").Key("log_file").String(),
	}
}
