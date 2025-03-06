package cmd

import (
	"flutterterm/pkg/model"
	"flutterterm/pkg/utils"
	"fmt"

	"github.com/spf13/cobra"
)

const (
	preserveConfig = "preserve-config"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Creates a new .fterm_config.json file",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		force, _ := cmd.Flags().GetBool(force)
		preserveConfig, _ := cmd.Flags().GetBool(preserveConfig)
		err := model.InitConfig(model.DefaultConfigPath, force, preserveConfig)
		if err != nil {
			utils.PrintError(fmt.Sprintf("Could not create new config: %s", err.Error()))
		} else if preserveConfig {
			utils.PrintSuccess(fmt.Sprintf("Successfully created %s. Preserved %d run configs", model.DefaultConfigPath, len(config.Configs)))
		} else {
			utils.PrintSuccess(fmt.Sprintf("Successfully created %s", model.DefaultConfigPath))
		}
	},
}

func init() {
	configCmd.AddCommand(initCmd)
	initCmd.Flags().Bool(force, false, "")
	initCmd.Flags().BoolP(
		preserveConfig,
		string(preserveConfig[0]),
		false,
		"Reset config while preserving run configs",
	)
}
