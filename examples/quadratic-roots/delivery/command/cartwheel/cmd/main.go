package main

import (
	"flag"
	"github.com/discomco/go-cart/examples/quadratic-roots/delivery/command/cartwheel"
	"os"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "", "Quadric Roots config path")
}

func main() {
	flag.Parse()
	err := cartwheel.RunCartwheel(configPath)
	if err != nil {
		os.Exit(1)
	}
}
