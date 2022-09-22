package main

import (
	"flag"
	"github.com/discomco/go-cart/examples/robby/execute-game/containers/CMD/cartwheel"
	"os"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "", "EBS-ConfigureEvent-CMD config path")
}

func main() {
	flag.Parse()
	err := cartwheel.RunCarthweel(configPath)
	if err != nil {
		os.Exit(1)
	}
}
