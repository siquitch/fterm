package cmd

import (
	"fterm/pkg/flows"
	"fterm/pkg/model"
	"fterm/pkg/utils"

	"github.com/spf13/cobra"
)

const (
	force     = "force"
	favorites = "favorites"
	def       = "default"
	last      = "last"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A guided flutter run command",
	Long:  ``,
	Args:  cobra.MaximumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		force, err := cmd.Flags().GetBool(force)

		if err != nil {
			utils.PrintError(err.Error())
		}

		def, err := cmd.Flags().GetBool(def)

		if err != nil {
			utils.PrintError(err.Error())
		}

		last, err := cmd.Flags().GetBool(last)

		if err != nil {
			utils.PrintError(err.Error())
		}

		if !model.AssertRootPath(force) {
			return
		}

		var runConfigs model.RunConfig

		argLen := len(args)

		if argLen == 0 && !def && !last {
			runConfigs, _ = flows.RunFlow(*config)
		} else if (argLen == 1 && !def && !last) || (argLen == 0 && def) || (argLen == 0 && last) {
			var c *model.FlutterConfig
			var err error

			if def {
				c, err = config.GetConfigByName(config.DefaultConfig)
			} else {
				c, err = config.GetConfigByName(args[0])
			}

			if err != nil {
				utils.PrintError(err.Error())
				return
			}
			d, _ := flows.DeviceFlow(*config)
			if !d.Verified() {
				return
			}
			runConfigs = model.RunConfig{
				SelectedConfig: *c,
				SelectedDevice: d,
			}
		}

		if !runConfigs.IsComplete() {
			return
		}

		rc := runConfigs.SelectedConfig

		go config.SaveConfig(model.DefaultConfigPath)

		rc.Run(runConfigs.SelectedDevice, config.Fvm)
	},
}

func init() {
	runCmd.Flags().BoolP(favorites, string(favorites[0]), false, "Show favorites")
	runCmd.Flags().Bool(force, false, "")
	runCmd.Flags().BoolP(def, string(def[0]), false, "Run default config")
	runCmd.Flags().BoolP(last, string(last[0]), false, "Run last config")
	rootCmd.AddCommand(runCmd)
}
