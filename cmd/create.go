package cmd

import (
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new emulator",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// p := tea.NewProgram(pkg.InitialCreateEmulatorModel())
		//
		// model, err := p.Run()
		// if err != nil {
		// 	utils.PrintError(err.Error())
		// 	return
		// }
		//
		// createModel, ok := model.(pkg.CreateEmulatorModel)
		//
		// if !ok {
		// 	utils.PrintError("Could not case tea model to create emulator model")
		// 	return
		// }
		//
		// if !createModel.IsComplete() {
		// 	return
		// }
		//
		// runArgs := []string{"emulators", "--create", "--name", createModel.Text()}
		//
		// runCmd := exec.Command("flutter", runArgs...)
		//
		// // Using cmd.Run() doesn't work for some reason
		// _, err = runCmd.CombinedOutput()
		//
		// if err != nil {
		// 	utils.PrintError(fmt.Sprintf("Error: %s\n", err.Error()))
		// 	return
		// }
		//
		// utils.PrintInfo(fmt.Sprintf("Successfully created %s", createModel.Text()))
	},
}

func init() {
	emulatorsCmd.AddCommand(createCmd)
}
