package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

const ConfigPath = ".fterm_config.json"

type FlutterConfig struct {
	Name               string `json:"name"`
	Mode               string `json:"mode"`
	Flavor             string `json:"flavor"`
	Target             string `json:"target"`
	DartDefineFromFile string `json:"dart_define_from_file"`
}

func (config FlutterConfig) ToString() string {
	var s string
	s = fmt.Sprintf("Config: %s\n", config.Name)
	s += fmt.Sprintf("Mode: %s\n", config.Mode)
	s += fmt.Sprintf("Flavor: %s\n", config.Flavor)
	s += fmt.Sprintf("Target: %s\n", config.Target)
    s += fmt.Sprintf("Dart define file: %s\n", config.DartDefineFromFile)
	return s
}

func DefaultConfig() FlutterConfig {
	return FlutterConfig{
		Name:   "Default",
		Mode:   "debug",
		Target: "main.dart",
	}
}

func GetConfigs() ([]FlutterConfig, error) {
	PrintInfo("Fetching configs\n\n")
	var configs []FlutterConfig

	config_file, err := os.Open(ConfigPath)

	if err != nil {
		return configs, err
	}

	defer config_file.Close()

	// Read file

	bytes, err := io.ReadAll(config_file)

	if err != nil {
		return configs, err
	}

	err = json.Unmarshal(bytes, &configs)

	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < len(configs); i++ {
		configs[i].ToString()
	}

	return configs, nil
}
