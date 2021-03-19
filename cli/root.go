package cli

import (
	"github.com/spf13/cobra"

	"fmt"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "coel",
	Short: "Coelhino CLI",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Calling Coelhino root command")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
