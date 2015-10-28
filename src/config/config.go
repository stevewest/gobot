// Package config allows loading of application config from a json file.
package config

type Container struct {
	File string
	Config map[string]string
}

// Get returns the value of the given key, or the default if it does not exist.
// Acts as a minor shortcut.
func (config Container) Get (key, def string) string {

	if val, ok := config.Config[key]; ok {
		return val
	}

	return def
}
