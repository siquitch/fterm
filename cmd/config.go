package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// configCmd represents the configs command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "View your current config",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(config.ToString())
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
