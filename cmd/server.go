package main

import (
	"log"

	"github.com/Sh1n3zZ/ovs-agent/bootstrap"
	grpcserver "github.com/Sh1n3zZ/ovs-agent/server/server"

	"github.com/spf13/cobra"
)

var (
	withConsumer bool
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run the HTTP server (optionally with RocketMQ consumer in the same process)",
	Long:  "Run the HTTP server. Use --with-consumer to also start the RocketMQ SimpleConsumer in a goroutine.",
	RunE:  runServerCmd,
}

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.Flags().BoolVar(&withConsumer, "with-consumer", false, "also start the RocketMQ consumer in the same process")
}

func runServerCmd(cmd *cobra.Command, args []string) error {
	runServer(withConsumer)
	return nil
}

func runServer(withConsumer bool) {
	app := bootstrap.App()

	env := app.Env
	ovsClient := app.OVSClient

	// env.ServerAddress is expected to be a host:port or :port style address
	addr := env.ServerAddress
	if addr == "" {
		// Fallback to a sane default if configuration is missing.
		addr = ":50051"
	}

	// Start the gRPC server using the pre-initialized OVS client from bootstrap.
	// This blocks until the server stops or an error occurs.
	if err := grpcserver.RunGRPCServer(addr, ovsClient); err != nil {
		log.Fatalf("failed to start gRPC server: %v", err)
	}
}
