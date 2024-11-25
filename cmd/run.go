/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"flutterterm/ui"
	"flutterterm/utils"
	"fmt"
	"os"
	"os/exec"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

// The file to look for in a flutter project
const pubspec = "pubspec.yaml"

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A guided flutter run command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if !assertRootPath() {
			return
		}

		utils.PrintInfo(fmt.Sprintf("Flutter directory detected. Getting devices\n"))

		devices, err := utils.GetDevices()

		if err != nil {
			e := fmt.Sprintf("There was an error getting devices: %s\n", err)
			utils.PrintError(e)
			return
		}

		configs, err := utils.GetConfigs()

		// Add a default run config if none exist
		if len(configs) == 0 {
			utils.PrintInfo("No configs found, using default\n\n")
			help := fmt.Sprintf("Try creating a \"%s\" file or adding a config to an already created one", utils.ConfigPath)
			utils.PrintHelp(help)
			defaultConfig, err := utils.DefaultConfig()
			if err != nil {
				utils.PrintError(err.Error())
				return
			}
			configs = append(configs, defaultConfig)
		}

		p := tea.NewProgram(ui.InitialRunModel(devices, configs))

		model, err := p.Run()

		if err != nil {
			fmt.Printf("Error %v", err)
			os.Exit(1)
		}

		runModel, ok := model.(ui.RunModel)

		if !ok {
			fmt.Println("Could not cast tea model to run model")
			return
		}

		if !runModel.IsComplete() {
			return
		}

		setupAndRun(runModel)
	},
}

// Runs command based on the model received
func setupAndRun(m ui.RunModel) {
	fmt.Printf("Running %s on %s\n\n", m.Selected_config.Name, m.Selected_device.Name)

	// Device
	device := m.Selected_device.ID
	config := m.Selected_config

	args := []string{"run", "-d", device}
	if config.Target != "" {
		args = append(args, "-t", config.Target)
	}
	if config.Flavor != "" {
		args = append(args, "--flavor", config.Flavor)
	}
	if config.Mode != "" {
		if !assertFlutterMode(config.Mode) {
			utils.PrintError(fmt.Sprintf("Invalid flutter mode: %s", config.Mode))
			return
		}
	}
	if config.DartDefineFromFile != "" {
		args = append(args, "--dart-define-from-file", config.DartDefineFromFile)
	}

	cmd := exec.Command("flutter", args...)

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
func assertRootPath() bool {
	_, err := os.Stat(pubspec)

	if err != nil {
		utils.PrintError("pubspec.yaml not found. Make sure you are in a flutter directory")
		return false
	}

	return true
}

// Verify proper mode being used
func assertFlutterMode(m string) bool {
	m = strings.ToLower(m)
	for mode := range utils.FlutterModes {
		if mode == mode {
			return true
		}
	}
	return false
}

func init() {
	rootCmd.AddCommand(runCmd)
}
