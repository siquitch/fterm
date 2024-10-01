/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"flutterterm/ui"
	"flutterterm/utils"
	"fmt"
	"os/exec"

	// tea "github.com/charmbracelet/bubbletea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

// emulatorsCmd represents the emulators command
var emulatorsCmd = &cobra.Command{
	Use:   "emulators",
	Short: "Start an emulator as detected by flutter emulators",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		emulators, err := utils.GetEmulators()

		// Whether to cold start the emulator
		isCold, err := cmd.Flags().GetBool("cold")

		if err != nil {
			utils.PrintError(err.Error())
			return
		}

		p := tea.NewProgram(ui.InitialEmulatorModel(emulators))

		model, err := p.Run()

		if err != nil {
			fmt.Println("Emulators exited with error")
		}

		emulatorModel, ok := model.(ui.EmulatorModel)

		if !ok {
			fmt.Println("Could not cast model to emulatorModel")
		}

		if !emulatorModel.IsComplete() {
			return
		}

		device := emulatorModel.SelectedEmulator

		cold := ""

		if isCold {
			cold = "--cold"
		}

		s := fmt.Sprintf("Opening %s", device.Name)
		utils.PrintInfo(s)

		// Run the final command
		flutterCmd := exec.Command("flutter", "emulators", "--launch", device.ID, cold)

		err = flutterCmd.Run()

		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(emulatorsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// emulatorsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// emulatorsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	emulatorsCmd.Flags().BoolP("cold", "c", false, "Cold start the emulator")
}
