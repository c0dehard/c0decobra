package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// twoCmd represents the two command
var encryptCmd = &cobra.Command{
	Use:   "two",
	Short: "two",
	Long:  `two two two two two`,
	Run: func(cmd *cobra.Command, args []string) {

		// define execution after call just below this line
		fmt.Println("called two")
	},
}

func init() {
	rootCmd.AddCommand(twoCmd)
}
