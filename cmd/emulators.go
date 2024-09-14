/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"flutterterm/utils"
	"fmt"

	"github.com/spf13/cobra"
)

// emulatorsCmd represents the emulators command
var emulatorsCmd = &cobra.Command{
	Use:   "emulators",
	Short: "Start a new emulator",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		_, err := utils.GetEmulators()

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
}
