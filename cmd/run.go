/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"flutterterm/ui"
	"flutterterm/utils"
	"fmt"
	"os"

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
		// if !assertRootPath() {
		// 	return
		// }
		// fmt.Println("Flutter directory detected. Getting devices")

		devices, err := utils.GetDevices()

		if err != nil {
			fmt.Printf("There was an error getting devices: %s", err)
			return
		}

		configs, err := utils.GetConfigs()

		if err != nil {
			fmt.Printf("There was an error getting configs: %s", err)
			return
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
		}

		if !runModel.IsComplete() {
			return
		}

		setupAndRun(runModel)
	},
}

// Runs command based on the model received
func setupAndRun(m ui.RunModel) {
	fmt.Printf("Running %s on %s", m.Selected_config.Name, m.Selected_device.Name)

	// var args string
	//
	// // Device
	// args = fmt.Sprint("-d ")
}

// Check if in a flutter project
func assertRootPath() bool {
	_, err := os.Stat(pubspec)

	if err != nil {
		fmt.Println("pubspec.yaml not found. Make sure you are in a flutter directory")
		return false
	}

	return true
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
