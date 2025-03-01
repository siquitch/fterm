package cmd

import (
	"flutterterm/pkg/command"
	"flutterterm/pkg/flows"
	"flutterterm/pkg/model"
	"flutterterm/pkg/utils"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

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

		if !model.AssertRootPath(force) {
			return
		}

		if len(args) < 1 {

			p := tea.NewProgram(flows.InitialRunModel(*config))

			model, err := p.Run()

			if err != nil {
				utils.PrintError(fmt.Sprintf("Error %s", err.Error()))
				return
			}

			runModel, _ := model.(flows.RunModel)

			if !runModel.IsComplete() {
				return
			}

			setupAndRun(runModel)
			return
		}

		if len(args) == 1 {
			fmt.Println(args[0])
		}

	},
}

// Runs command based on the model received
func setupAndRun(m flows.RunModel) {
	fmt.Printf("Running %s on %s\n\n", m.SelectedConfig().Name, m.SelectedDevice().Name)

	// Device
	device := m.SelectedDevice().ID
	config := m.SelectedConfig()

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
	if config.DartDefineFile != "" {
		args = append(args, "--dart-define-from-file", config.DartDefineFile)
	}

	cmd := command.FlutterRun(args)

	// For color and input handling
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Start()

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

func init() {
	runCmd.Flags().BoolP(favorites, string(favorites[0]), false, "Show favorites")
    runCmd.Flags().Bool(force, false, "")
	rootCmd.AddCommand(runCmd)
}
