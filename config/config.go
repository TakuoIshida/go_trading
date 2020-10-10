package config

import (
	"gopkg.in/go-ini/ini.v1"
)

type ConfigList struct {
	Apikey    string
	ApiSecret string
}

var Config ConfigList

func main() {
	cfg, err := ini.Load(config.ini)

}
