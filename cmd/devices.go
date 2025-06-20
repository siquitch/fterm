package cmd

import (
	"fmt"
	"fterm/pkg/model"
	"fterm/pkg/utils"

	"github.com/spf13/cobra"
)

// devicesCmd represents the devices command
var devicesCmd = &cobra.Command{
	Use:   "devices",
	Short: "Select a device to run your flutter app",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		execute(*&config.Fvm)
	},
}

func execute(fvm bool) {
	cmd := model.FlutterDevices(fvm)
	output, err := cmd.Output()
	if err != nil {
		utils.PrintError(err.Error())
		return
	}

	devices, err := model.ParseDevices(output)
	if err != nil {
		utils.PrintError(err.Error())
		return
	}

	for i, d := range devices {
		fmt.Printf("%d: %s - %s\n", i, d.Name, d.ID)
	}
}

func init() {
	rootCmd.AddCommand(devicesCmd)
}
