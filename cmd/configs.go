package cmd

import (
	"flutterterm/utils"
	"fmt"

	"github.com/spf13/cobra"
)

// configsCmd represents the configs command
var configsCmd = &cobra.Command{
	Use:   "configs",
	Short: "View your current configs",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		configs, err := utils.GetConfigs()

		if err != nil {
			utils.PrintError(err.Error())
		}

		for c := range configs {
			s := fmt.Sprintf("%s\n", configs[c].ToString())
			utils.PrintInfo(s)
		}
	},
}

func init() {
	rootCmd.AddCommand(configsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
