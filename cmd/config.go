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
		if config != nil {
            s, err := config.ToString()
            if err != nil {
                fmt.Println(err)
            } else {
                fmt.Println(s)
            }
        }
    },
}

func init() {
	rootCmd.AddCommand(configCmd)
}
