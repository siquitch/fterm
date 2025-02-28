package cmd

import (
	"flutterterm/pkg/ui"
	"flutterterm/pkg/utils"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

// The file to look for in a flutter project
const pubspec = "pubspec.yaml"

const (
	force     = "force"
	favorites = "favorites"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A guided flutter run command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		force, err := cmd.Flags().GetBool(force)

		if err != nil {
			utils.PrintError(err.Error())
		}

		if !assertRootPath() && !force {
			return
		}

		utils.PrintInfo(fmt.Sprintf("Flutter directory detected. Getting devices\n"))

		configs := config.RunConfigs

		// Add a default run config if none exist
		if len(configs) == 0 {
			utils.PrintInfo("No configs found, using default\n\n")
			help := fmt.Sprintf("Try creating a \"%s\" file or adding a config to an already created one", utils.ConfigPath)
			utils.PrintHelp(help)
			defaultConfig, err := utils.DefaultRunConfig()
			if err != nil {
				utils.PrintError(err.Error())
				return
			}
			configs = append(configs, defaultConfig)
		} else {
			utils.PrintSuccess(fmt.Sprintf("%d configs found\n\n", len(configs)))
		}

		p := tea.NewProgram(ui.InitialRunModel(*config))

		model, err := p.Run()

		if err != nil {
			utils.PrintError(fmt.Sprintf("Error %s", err.Error()))
			return
		}

		runModel, _ := model.(ui.RunModel)

		if !runModel.IsComplete() {
			return
		}

		setupAndRun(runModel)
	},
}

// Runs command based on the model received
func setupAndRun(m ui.RunModel) {
	fmt.Printf("Running %s on %s\n\n", m.SelectedConfig().Name, m.SelectedDevice().Name)

	// Device
	device := m.SelectedDevice().ID
	config := m.SelectedConfig()

	err := config.AssertConfig()

	if err != nil {
		e := fmt.Sprintf("Invalid configuration: %s", err)
		utils.PrintError(e)
		return
	}

	args := []string{"run", "-d", device}
	if config.Target != "" {
		args = append(args, "-t", config.Target)
	}
	if config.Mode != "" {
		arg := fmt.Sprintf("--%s", config.Mode)
		args = append(args, arg)
	}
	if config.Flavor != "" {
		args = append(args, "--flavor", config.Flavor)
	}
	if config.DartDefineFromFile != "" {
		args = append(args, "--dart-define-from-file", config.DartDefineFromFile)
	}

	cmd := utils.FlutterRun(args)

	// For color and input handling
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err = cmd.Start()

	if err != nil {
		utils.PrintError(err.Error())
		return
	}

	if err := cmd.Wait(); err != nil {
		s := fmt.Sprintf("Flutterterm finished with error: %s", err)
		utils.PrintError(s)
	} else {
		utils.PrintSuccess("Flutterterm finished successfully")
	}
}

// Check if in a flutter project
func assertRootPath() bool {
	_, err := os.Stat(pubspec)

	if err != nil {
		utils.PrintError("pubspec.yaml not found. Make sure you are in a flutter directory")
		return false
	}

	return true
}

func init() {
	runCmd.Flags().Bool(force, false, "")
	runCmd.Flags().BoolP(favorites, string(favorites[0]), false, "Show favorites")
	rootCmd.AddCommand(runCmd)
}
