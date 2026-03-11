package ovs

import ovs "github.com/digitalocean/go-openvswitch/ovs"

// ListBridgeFlows retrieves statistics about all flows for the specified bridge
// and returns them as human-readable textual representations.
func ListBridgeFlows(client *ovs.Client, bridge string) ([]string, error) {
	flows, err := client.OpenFlow.DumpFlows(bridge)
	if err != nil {
		return nil, err
	}

	out := make([]string, 0, len(flows))
	for _, f := range flows {
		if f == nil {
			out = append(out, "")
			continue
		}

		if b, err := f.MarshalText(); err == nil {
			out = append(out, string(b))
			continue
		}

		out = append(out, "")
	}

	return out, nil
}

// ListBridgeFlowsWithMatchArgs retrieves statistics about all flows for the specified bridge,
// allowing the caller to provide a MatchFlow to filter results (for example by table or cookie).
// It returns the flows as human-readable textual representations.
func ListBridgeFlowsWithMatchArgs(client *ovs.Client, bridge string, flow *ovs.MatchFlow) ([]string, error) {
	flows, err := client.OpenFlow.DumpFlowsWithFlowArgs(bridge, flow)
	if err != nil {
		return nil, err
	}

	out := make([]string, 0, len(flows))
	for _, f := range flows {
		if f == nil {
			out = append(out, "")
			continue
		}

		if b, err := f.MarshalText(); err == nil {
			out = append(out, string(b))
			continue
		}

		out = append(out, "")
	}

	return out, nil
}

