package cmd

import (
	"flutterterm/pkg/utils"
	"fmt"

	"github.com/spf13/cobra"
)

// devicesCmd represents the devices command
var devicesCmd = &cobra.Command{
	Use:   "devices",
	Short: "Select a device to run your flutter app",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		execute()
	},
}

func execute() {
	cmd := utils.FlutterDevices()
	output, err := cmd.Output()
	if err != nil {
		utils.PrintError(err.Error())
		return
	}

	devices, err := utils.ParseDevices(output)
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
