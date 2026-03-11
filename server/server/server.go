package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"strings"

	libovs "github.com/digitalocean/go-openvswitch/ovs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/Sh1n3zZ/ovs-agent/api/ovsagentpb"
	"github.com/Sh1n3zZ/ovs-agent/bootstrap"
)

// RunGRPCServer starts the OVSAgent gRPC server on the given address, using the provided OVS client.
// addr should be in the form ":50051" or "0.0.0.0:50051".
func RunGRPCServer(addr string, ovsClient *libovs.Client) error {
	env := bootstrap.NewConfig()

	authInterceptor := grpc.UnaryInterceptor(func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(codes.Unauthenticated, "missing metadata")
		}

		values := md.Get("authorization")
		if len(values) == 0 {
			return nil, status.Error(codes.Unauthenticated, "missing authorization header")
		}

		raw := values[0]
		const prefix = "Bearer "

		if !strings.HasPrefix(raw, prefix) {
			return nil, status.Error(codes.Unauthenticated, "invalid authorization header format")
		}

		token := strings.TrimPrefix(raw, prefix)
		if token == "" || token != env.APISecret {
			return nil, status.Error(codes.Unauthenticated, "invalid token")
		}

		return handler(ctx, req)
	})

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", addr, err)
	}

	grpcServer := grpc.NewServer(authInterceptor)
	ovsagentpb.RegisterOVSAgentServer(grpcServer, NewOVSAgentServer(ovsClient))

	log.Printf("OVSAgent gRPC server listening on %s", addr)

	if err := grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("gRPC server stopped: %w", err)
	}

	return nil
}