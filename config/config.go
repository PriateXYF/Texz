package config

import (
	_ "embed"
)

//go:embed embed/config.json
var baseConfig string

func GetConfig() string {
	configPath := ExecutablePathJoin("config.json")
	config, err := ReadFile(configPath)
	if err != nil {
		GenerateConfig()
		return baseConfig
	}
	return config
}

func GenerateConfig() error {
	configPath := ExecutablePathJoin("config.json")
	err := OutputFile(configPath, baseConfig)
	return err
}
