package cmd

import (
	"github.com/0ya-sh0/keyforge/internal/api"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the server at port 8080. If the port is in use, it will try the next port",
	Run: func(cmd *cobra.Command, args []string) {
		err := api.StartGRPCServer()
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
