package cmd

import (
	"flutterterm/pkg/flows"
	"flutterterm/pkg/utils"

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

		err = flows.EmulatorFlow(isCold)

		if err != nil {
			utils.PrintError(err.Error())
		}
	},
}

func init() {
	emulatorsCmd.Flags().BoolP(cold, string(cold[0]), false, "Cold start the emulator")
	rootCmd.AddCommand(emulatorsCmd)
}
