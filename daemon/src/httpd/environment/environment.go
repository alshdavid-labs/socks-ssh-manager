package environment

import (
	"flag"
)

type Environment struct {
	ConfigPath string
}

func NewEnvironment() *Environment {
	env := &Environment{}

	configPath := flag.String("config-path", "", "Text to parse.")
	flag.Parse()

	if configPath != nil {
		env.ConfigPath = *configPath
	}

	return env
}
