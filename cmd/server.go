package cmd

import (
	"github.com/GeekTree0101/iSpygo/server"
	"github.com/spf13/cobra"
)

var (
	runCmd = &cobra.Command{
		Use:   "run",
		Short: "run iGospy",
		Run: func(cmd *cobra.Command, args []string) {
			port := "7777"

			// FIXME: How to get the port number in a more efficient way
			if len(args) > 0 {
				port = args[0]
			}

			serv := server.NewServer(port)
			serv.Mount()
			serv.Route()
			serv.Run()
		},
	}
)

func init() {
	rootCmd.AddCommand(runCmd)
}
