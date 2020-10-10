package main

import (
	"fmt"
	"std/ishida/go_trading/config"
)

func main() {
	fmt.Println(config.Config.Apikey)
	fmt.Println(config.Config.ApiSecret)
}
