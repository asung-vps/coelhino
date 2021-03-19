package cli

import (
	"github.com/spf13/cobra"

	"log"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize Starfire project",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Calling init command")
	},
}
