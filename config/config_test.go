package config_test

import (
	"github.com/stevewest/gobot/config"
	"testing"
	"strings"
	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {

	reader := strings.NewReader(`{
	"Server": {
		"host": "test.dev",
		"port": 1234
	},
	"Channels": [
		"#foo",
		"#bar"
	],
	"Plugins": [
		"one",
		"two"
	]
}`)

	c := config.Container{}
	c.Load(reader)

	assert.Equal(t, "test.dev", c.Config.Server.Host)
	assert.Equal(t, 1234, c.Config.Server.Port)
	assert.Equal(t, []string{"#foo", "#bar"}, c.Config.Channels)
	assert.Equal(t, []string{"one", "two"}, c.Config.Plugins)

}
