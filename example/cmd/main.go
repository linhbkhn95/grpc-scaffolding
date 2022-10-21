package main

import (
	"github.com/spf13/cobra"

	"github.com/linhbkhn95/golang-british/logger"

	"github.com/linhbkhn95/example/server"
)

func main() {
	cmd := &cobra.Command{
		Use: "rpc-application",
	}

	cmd.AddCommand(&cobra.Command{
		Use: "server",
		Run: func(cmd *cobra.Command, args []string) {
			server.Serve()
		},
	})
	if err := cmd.Execute(); err != nil {
		logger.Fatalf("Error when executed %v", err)
	}
}
