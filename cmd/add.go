package cmd

import (
	"fterm/pkg/flows"
	"fterm/pkg/utils"
	"fmt"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a run configuration to the config file",
	Run: func(cmd *cobra.Command, args []string) {
		nc, _ := flows.AddFlow()

		err := config.AddRunConfig(nc)

		if err != nil {
			utils.PrintError(err.Error())
		} else {
			utils.PrintSuccess(fmt.Sprintf("%s created successfully", nc.Name))
		}
	},
}

func init() {
	configCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
