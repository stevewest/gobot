// Package config allows loading of application config from a json file.
package config

import (
	"fmt"
	"io"
	"io/ioutil"
	"encoding/json"
)

type Container struct {
	Config configType
}

type configType struct {
	Server ServerType
	Channels []string
	Plugins []string
}

type ServerType struct {
	Host string
	Port int
}

func (config *Container) Load (reader io.Reader) {
	file, fileError := ioutil.ReadAll(reader)

	if fileError != nil {
		fmt.Errorf("File error: %v\n", fileError)
	}

	var loadedJson configType
	jsonError := json.Unmarshal(file, &loadedJson)
	if jsonError != nil {
		fmt.Errorf("JSON error: %v\n", jsonError)
	}

	config.Config = loadedJson
}
