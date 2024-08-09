package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Fasi7",
	Long:  `All software has versions. This is Fasi7's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Fasi7 v0.1")
	},
}
