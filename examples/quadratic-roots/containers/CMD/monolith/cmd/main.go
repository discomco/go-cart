package main

import (
	"flag"
	"github.com/discomco/go-cart/examples/quadratic-roots/containers/CMD/monolith/cartwheel"
	"os"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "", "EBS-ConfigureEvent-CMD config path")
}

func main() {
	flag.Parse()
	err := cartwheel.RunCartwheel(configPath)
	if err != nil {
		os.Exit(1)
	}
}
