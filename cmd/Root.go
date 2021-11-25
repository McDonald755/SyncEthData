package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.AddCommand(
		ScanCmd(),
		)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err.Error())
	}
}
