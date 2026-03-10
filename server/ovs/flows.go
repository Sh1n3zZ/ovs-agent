package ovs

import ovs "github.com/digitalocean/go-openvswitch/ovs"

// ListBridgeFlows retrieves statistics about all flows for the specified bridge.
// It delegates to the underlying go-openvswitch OpenFlowService.DumpFlows.
func ListBridgeFlows(client *ovs.Client, bridge string) ([]*ovs.Flow, error) {
	return client.OpenFlow.DumpFlows(bridge)
}

// ListBridgeFlowsWithMatchArgs retrieves statistics about all flows for the specified bridge,
// allowing the caller to provide a MatchFlow to filter results (for example by table or cookie).
// It delegates to the underlying go-openvswitch OpenFlowService.DumpFlowsWithFlowArgs.
func ListBridgeFlowsWithMatchArgs(client *ovs.Client, bridge string, flow *ovs.MatchFlow) ([]*ovs.Flow, error) {
	return client.OpenFlow.DumpFlowsWithFlowArgs(bridge, flow)
}

