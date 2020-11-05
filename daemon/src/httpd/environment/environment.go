package environment

import (
	"flag"
	"os"
	"path"
	"strings"
)

const DEFAULT_CONFIG = "state.json"

type Environment struct {
	ConfigPath string
}

func NewEnvironment() *Environment {
	env := &Environment{}

	configPathRef := flag.String("config-path", "", "Text to parse.")
	flag.Parse()
	configPath := *configPathRef

	if configPath == "" {
		cwd, _ := os.Getwd()
		env.ConfigPath = path.Join(cwd, DEFAULT_CONFIG)
	} else if strings.HasPrefix(configPath, "/") {
		env.ConfigPath = configPath
	} else {
		cwd, _ := os.Getwd()
		env.ConfigPath = path.Join(cwd, configPath)
	}

	return env
}
