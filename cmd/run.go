package cmd

import (
	"flutterterm/pkg/flows"
	"flutterterm/pkg/model"
	"flutterterm/pkg/utils"

	"github.com/spf13/cobra"
)

const (
	force     = "force"
	favorites = "favorites"
	def       = "default"
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

		if !model.AssertRootPath(force) {
			return
		}

		var runConfig model.RunConfig

		argLen := len(args)

		if argLen == 0 && !def {
			runConfig, _ = flows.RunFlow(*config)
		} else if (argLen == 1 && !def) || (argLen == 0 && def) {
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
			d, _ := flows.DeviceFlow()
			if !d.Verified() {
				return
			}
			runConfig = model.RunConfig{
				SelectedConfig: *c,
				SelectedDevice: d,
			}
		}

		if !runConfig.IsComplete() {
			return
		}

		config := runConfig.SelectedConfig
		config.Run(runConfig.SelectedDevice)
	},
}

func init() {
	runCmd.Flags().BoolP(favorites, string(favorites[0]), false, "Show favorites")
	runCmd.Flags().Bool(force, false, "")
	runCmd.Flags().BoolP(def, string(def[0]), false, "Run default config")
	rootCmd.AddCommand(runCmd)
}
