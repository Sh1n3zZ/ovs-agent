package server

import (
	"fmt"
	"log"
	"net"

	libovs "github.com/digitalocean/go-openvswitch/ovs"
	"google.golang.org/grpc"

	"github.com/Sh1n3zZ/ovs-agent/api/ovsagentpb"
)

// RunGRPCServer starts the OVSAgent gRPC server on the given address, using the provided OVS client.
// addr should be in the form ":50051" or "0.0.0.0:50051".
func RunGRPCServer(addr string, ovsClient *libovs.Client) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", addr, err)
	}

	grpcServer := grpc.NewServer()
	ovsagentpb.RegisterOVSAgentServer(grpcServer, NewOVSAgentServer(ovsClient))

	log.Printf("OVSAgent gRPC server listening on %s", addr)

	if err := grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("gRPC server stopped: %w", err)
	}

	return nil
}