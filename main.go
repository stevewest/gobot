package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"os"
	"github.com/stevewest/gobot/config"
)

var opts struct {
	Config string `short:"c" long:"config" description:"Path to config file, ./config.json by default" default:"config.json"`
}

func main() {

	_, err := flags.Parse(&opts)

	if err != nil {
		panic(err)
		os.Exit(1)
	}

	// Try and load our config
	configFile, err := os.Open(opts.Config)

	if err != nil {
		fmt.Errorf("Cannot load config file!")
		panic(err)
		os.Exit(2)
	}

	container := config.Container{}
	container.Load(configFile)

	fmt.Println(container.Config.Server.Host)
}
