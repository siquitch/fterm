/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"flutterterm/pkg/model"
	"flutterterm/pkg/utils"
	"fmt"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Creates a new .fterm_config.json file",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		force, _ := cmd.Flags().GetBool(force)
		err := model.InitConfig(model.DefaultConfigPath, force)
		if err != nil {
			utils.PrintError(fmt.Sprintf("Could not create new config: %s", err.Error()))
		} else {
			utils.PrintSuccess(fmt.Sprintf("Successfully created %s", model.DefaultConfigPath))
		}
	},
}

func init() {
	configCmd.AddCommand(initCmd)
	initCmd.Flags().Bool(force, false, "")
}
