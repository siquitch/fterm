package model

import (
	"encoding/json"
	"errors"
	"flutterterm/pkg/utils"
	"fmt"
	"os"
	"os/exec"
)

const version = "0.0.3"

const (
	DefaultConfigPath = ".fterm_config.json"
	pubspec           = "pubspec.yaml"
	mainPath          = "main.dart"
	mainLibPath       = "lib/main.dart"
)

// main.dart paths to look for
var mainPaths = []string{mainPath, mainLibPath}

// Config represents the entire configuration structure
type Config struct {
	Version         string          `json:"version"`
	DefaultConfig   string          `json:"default_config"`
	Configs         []FlutterConfig `json:"configs"`
	FavoriteConfigs []string        `json:"favorite_configs"`
	Devices         DeviceSettings  `json:"devices"`
}

// FlutterConfig represents a single Flutter run configuration
type FlutterConfig struct {
	Name           string   `json:"name"`
	Description    string   `json:"description"`
	Mode           string   `json:"mode"`
	Flavor         string   `json:"flavor"`
	Target         string   `json:"target"`
	DartDefineFile string   `json:"dart_define_from_file"`
	AdditionalArgs []string `json:"additional_args"`
}

// DeviceSettings contains all device-related configurations
type DeviceSettings struct {
	FavoriteDevices []string       `json:"favorite_devices"`
	DeviceConfigs   []DeviceConfig `json:"device_configs"`
}

// DeviceConfig represents a single device configuration
type DeviceConfig struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Platform string `json:"platform"`
}

type RunConfig struct {
	SelectedConfig FlutterConfig
	SelectedDevice Device
}

func InitConfig(path string, force bool) error {
	if !AssertRootPath(force) {
		return errors.New("No pubspec.yaml detected")
	}
	target, err := findDefaultTarget()

	if err != nil {
		target = mainLibPath
	}

	_, err = LoadConfig(path)

	if err == nil && !force {
		return errors.New("Config already detected, use --force to reset it")
	}

	c := Config{
		Version:       version,
		DefaultConfig: "default",
		Configs: []FlutterConfig{
			{
				Name:        "default",
				Description: "The default run configuration",
				Mode:        "debug",
				Target:      target,
			},
		},
		FavoriteConfigs: make([]string, 0),
		Devices: DeviceSettings{
			FavoriteDevices: make([]string, 0),
			DeviceConfigs:   make([]DeviceConfig, 0),
		},
	}
	return c.SaveConfig(path)
}

// Looks for main.dart files in default config
func findDefaultTarget() (string, error) {
	for _, path := range mainPaths {
		if _, err := os.Stat(path); err == nil {
			return path, nil
		}
	}
	err := errors.New("main.dart file not found")
	return "", err
}

// LoadConfig loads the configuration from a file
func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	// Fix unmarshaling issue by parsing the raw JSON again
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config structure: %w", err)
	}

	return &config, nil
}

// SaveConfig saves the configuration to a file
func (c *Config) SaveConfig(path string) error {
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	err = os.WriteFile(path, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}

// Add new run config to config file
func (c *Config) AddRunConfig(fc FlutterConfig) error {
	if fc.Name == "" {
		return errors.New("Name cannot be empty")
	}
	_, err := c.GetConfigByName(fc.Name)

	// Config already exists
	if err == nil {
		return errors.New(fmt.Sprintf("Config %s already exists", fc.Name))
	}

	c.Configs = append(c.Configs, fc)

	err = c.SaveConfig(DefaultConfigPath)

	if err != nil {
		return err
	}

	return nil
}

func (c *Config) RemoveRunConfig(name string) error {
	_, err := c.GetConfigByName(name)

	if err != nil {
		return err
	}

    configs := make([]FlutterConfig, 0)

    for _, config := range c.Configs {
        if config.Name != name {
            configs = append(configs, config)
        }
    }

    c.Configs = configs

    err = c.SaveConfig(DefaultConfigPath)

	return err
}

func (c *Config) GetConfigByName(name string) (*FlutterConfig, error) {
	for _, config := range c.Configs {
		if config.Name == name {
			return &config, nil
		}
	}
	return nil, fmt.Errorf("config with name '%s' not found", name)
}

func (c *Config) GetDefaultConfig() (*FlutterConfig, error) {
	return c.GetConfigByName(c.DefaultConfig)
}

func (c *Config) IsFavoriteConfig(name string) bool {
	for _, favorite := range c.FavoriteConfigs {
		if favorite == name {
			return true
		}
	}
	return false
}

func (c *Config) IsFavoriteDevice(deviceID string) bool {
	for _, favorite := range c.Devices.FavoriteDevices {
		if favorite == deviceID {
			return true
		}
	}
	return false
}

func (c *Config) AddFavoriteConfig(name string) error {
	// Check if config exists
	_, err := c.GetConfigByName(name)
	if err != nil {
		return err
	}

	// Don't add if already favorite
	if c.IsFavoriteConfig(name) {
		return nil
	}

	c.FavoriteConfigs = append(c.FavoriteConfigs, name)
	return nil
}

// RemoveFavoriteConfig removes a config from favorites
func (c *Config) RemoveFavoriteConfig(name string) {
	newFavorites := []string{}
	for _, favorite := range c.FavoriteConfigs {
		if favorite != name {
			newFavorites = append(newFavorites, favorite)
		}
	}
	c.FavoriteConfigs = newFavorites
}

// ToString returns the configuration as a formatted JSON string
func (c *Config) ToString() (string, error) {
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal config to JSON: %w", err)
	}
	return string(data), nil
}

// BuildFlutterCommand builds the flutter run command for a given config
func (fc *FlutterConfig) BuildFlutterCommand(deviceID string) *exec.Cmd {
	args := []string{"run"}

	// Set mode
	if fc.Mode != "" {
		args = append(args, fmt.Sprintf("--%s", fc.Mode))
	}

	// Set flavor
	if fc.Flavor != "" {
		args = append(args, "--flavor", fc.Flavor)
	}

	// Set target
	if fc.Target != "" {
		args = append(args, "--target", fc.Target)
	}

	// Set dart-define-from-file
	if fc.DartDefineFile != "" {
		args = append(args, "--dart-define-from-file", fc.DartDefineFile)
	}

	// Set device
	if deviceID != "" {
		args = append(args, "-d", deviceID)
	}

	// Add additional arguments
	args = append(args, fc.AdditionalArgs...)

	// Create the command
	cmd := exec.Command("flutter", args...)
	return cmd
}

// Whether the model has enough information to run
func (rc *RunConfig) IsComplete() bool {
	return rc.SelectedConfig.Name != "" && rc.SelectedDevice.ID != ""
}

// Whether the flutter config has enough info to be valid
func (fc *FlutterConfig) Validate() error {
	if fc.Name == "" {
		return errors.New("Name must not be empty")
	}

	return nil
}

func (fc *FlutterConfig) Run(device Device) {
	utils.PrintInfo(fmt.Sprintf("Running %s on %s\n\n", fc.Name, device.Name))
	cmd := fc.BuildFlutterCommand(device.ID)

	// For color and input handling
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Start()

	if err != nil {
		utils.PrintError(err.Error())
		return
	}

	if err := cmd.Wait(); err != nil {
		s := fmt.Sprintf("Flutterterm finished with error: %s", err)
		utils.PrintError(s)
	} else {
		utils.PrintSuccess("Flutterterm finished successfully")
	}
}

// Check if in a flutter project
func AssertRootPath(force bool) bool {
	if force {
		return true
	}
	_, err := os.Stat(pubspec)

	if err != nil {
		utils.PrintError("pubspec.yaml not found. Make sure you are in a flutter directory")
		return false
	}

	return true
}
