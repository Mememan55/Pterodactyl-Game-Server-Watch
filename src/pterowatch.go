package main

import (
	"./config"
	"fmt"
	"time"
)

func main() {
	configFile := "/etc/pterowatch/pterowatch.conf"

	cfg := config.Config{}

	config.ReadConfig(&cfg, configFile)

	for i := 0; i < len(cfg.Servers); i++ {
		fmt.Println(cfg.Servers[i].IP)
	}
}
