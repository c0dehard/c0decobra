package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// decryptCmd represents the decrypt command
var decryptCmd = &cobra.Command{
	Use:   "one",
	Short: "one",
	Long:  `one one one one one`,
	Run: func(cmd *cobra.Command, args []string) {

		// define execution after call just below this line
		fmt.Println("called one")
	},
}

func init() {
	rootCmd.AddCommand(decryptCmd)
}
