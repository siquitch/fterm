package utils

import (
	"encoding/json"
	"io"
	"os"
)

const ConfigPath = ".fterm_config.json"

type Config struct {
	RunConfigs      []FlutterRunConfig `json:"configs"`
	FavoriteConfigs []FlutterRunConfig `json:"favorite_configs"`
	FavoriteDevices []Device           `json:"favorite_devices"`
}

func GetConfig() (Config, error) {
	var config Config

	config_file, err := os.Open(ConfigPath)

	if err != nil {
		return config, err
	}

	defer config_file.Close()

	// Read file

	bytes, err := io.ReadAll(config_file)

	if err != nil {
		return config, err
	}
	err = json.Unmarshal(bytes, &config)

	if err != nil {
		return config, err
	}

	return config, err
}

func (c Config) ToString() string {
	json, _ := json.MarshalIndent(c, "", " ")

	return string(json)
}
