package main

import (
	"flag"
	"github.com/discomco/go-cart/examples/quadratic-roots/containers/command/cartwheel"
	"os"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "", "Quadric Roots config path")
}

func main() {
	flag.Parse()
	err := cartwheel.Run(configPath)
	if err != nil {
		os.Exit(1)
	}
}
