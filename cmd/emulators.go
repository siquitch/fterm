package cmd

import (
	"flutterterm/pkg/ui"
	"flutterterm/pkg/utils"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

const (
	cold = "cold"
)

// runs flutter emulators
var emulatorsCmd = &cobra.Command{
	Use:   "emulators",
	Short: "Runs 'flutter emulators'",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		isCold, err := cmd.Flags().GetBool(cold)

		p := tea.NewProgram(ui.InitialEmulatorModel(isCold))

		_, err = p.Run()

		if err != nil {
			e := fmt.Sprintf("Emulators exited with error: %s", err)
			utils.PrintError(e)
		}
	},
}

func init() {
	emulatorsCmd.Flags().BoolP(cold, string(cold[0]), false, "Cold start the emulator")
	rootCmd.AddCommand(emulatorsCmd)
}
