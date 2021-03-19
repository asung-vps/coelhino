package cli

import (
	"github.com/spf13/cobra"

	"fmt"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "coelhino",
	Short: "Coelhino CLI",
}

func init() {
	// Set passed function(s) to be run when each command's Execute method is called.
	cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(initCmd)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initConfig() {

}
