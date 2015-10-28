package config_test

import (
	"config"
	"testing"
)

func TestConfigGet(t *testing.T) {

	c := config.Container{Config: map[string]string{"foo": "bar"}}

	actual := c.Get("foo", "")

	if actual != "bar" {
		t.Fail()
	}

}

func TestConfigGetDefault(t *testing.T) {

	c := config.Container{Config: map[string]string{}}

	actual := c.Get("fake", "blue")

	if actual != "blue" {
		t.Fail()
	}

}
