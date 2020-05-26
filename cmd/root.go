package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile     string
	userLicense string

	rootCmd = &cobra.Command{
		Use:   "c0decobra",
		Short: "c0decobra info short",
		Long:  `Long info for c0decobra just add more text`,
	}
	// Simply add "global" variables below this line
)

const (
// You can also add "global" const just here
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
	rootCmd.PersistentFlags().StringP("author", "a", "c0deN", "author name for copyright attribution")
	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "name of license for the project")
	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
	viper.SetDefault("license", "apache")

}

func er(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}

// Add "global" functions that should be available for all modules in cmd below this line
