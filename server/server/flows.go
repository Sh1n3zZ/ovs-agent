package server

import (
	"context"
	"fmt"

	ovsclient "github.com/digitalocean/go-openvswitch/ovs"

	"github.com/Sh1n3zZ/ovs-agent/api/ovsagentpb"
	serverovs "github.com/Sh1n3zZ/ovs-agent/server/ovs"
)

// ovsAgentServer implements the OVSAgent gRPC service defined in api/ovsagentpb/agent.proto.
type ovsAgentServer struct {
	ovsagentpb.UnimplementedOVSAgentServer

	ovsClient *ovsclient.Client
}

// NewOVSAgentServer creates a new ovsAgentServer with the provided OVS client.
func NewOVSAgentServer(client *ovsclient.Client) *ovsAgentServer {
	return &ovsAgentServer{
		ovsClient: client,
	}
}

// ListFlows handles the ListFlows RPC and maps ListFlowsRequest to underlying OVS flow queries.
func (s *ovsAgentServer) ListFlows(ctx context.Context, req *ovsagentpb.ListFlowsRequest) (*ovsagentpb.ListFlowsResponse, error) {
	if req.GetBridge() == "" {
		return nil, fmt.Errorf("bridge must be specified")
	}

	// When no filters are provided, list all flows on the bridge.
	if req.GetTable() == 0 && req.GetCookie() == 0 && req.GetMatchExpr() == "" {
		rawFlows, err := serverovs.ListBridgeFlows(s.ovsClient, req.GetBridge())
		if err != nil {
			return nil, fmt.Errorf("list bridge flows: %w", err)
		}

		resp := &ovsagentpb.ListFlowsResponse{}
		for _, raw := range rawFlows {
			resp.Flows = append(resp.Flows, &ovsagentpb.Flow{Raw: raw})
		}
		return resp, nil
	}

	// Build a MatchFlow based on proto filters.
	matchFlow := &ovsclient.MatchFlow{}

	// Table: 0 means "not specified" here, so only set when non-zero.
	if req.GetTable() != 0 {
		matchFlow.Table = int(req.GetTable())
	}

	// Cookie: only apply when non-zero.
	if req.GetCookie() != 0 {
		matchFlow.Cookie = req.GetCookie()
		// CookieMask = 0 means exact match according to MatchFlow semantics.
		matchFlow.CookieMask = 0
	}

	// NOTE: Advanced raw match_expr from the request is not parsed here.
	// If you need full support, extend this to parse req.MatchExpr into MatchFlow.Matches.

	rawFlows, err := serverovs.ListBridgeFlowsWithMatchArgs(s.ovsClient, req.GetBridge(), matchFlow)
	if err != nil {
		return nil, fmt.Errorf("list bridge flows with match args: %w", err)
	}

	resp := &ovsagentpb.ListFlowsResponse{}
	for _, raw := range rawFlows {
		resp.Flows = append(resp.Flows, &ovsagentpb.Flow{Raw: raw})
	}

	return resp, nil
}
