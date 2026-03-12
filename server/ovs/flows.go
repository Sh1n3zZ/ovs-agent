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

// AddBridgeFlow adds a single OpenFlow flow to the specified bridge and returns its textual representation.
func AddBridgeFlow(client *ovs.Client, bridge string, flow *ovs.Flow) (string, error) {
	if err := client.OpenFlow.AddFlow(bridge, flow); err != nil {
		return "", err
	}

	if flow == nil {
		return "", nil
	}

	b, err := flow.MarshalText()
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// AddBridgeFlowBundle applies a bundle of flow additions and deletions atomically on the specified bridge
// and returns textual representations for the added flows.
func AddBridgeFlowBundle(client *ovs.Client, bridge string, addFlows []*ovs.Flow, deleteFlows []*ovs.MatchFlow) ([]string, error) {
	err := client.OpenFlow.AddFlowBundle(bridge, func(tx *ovs.FlowTransaction) error {
		if len(addFlows) > 0 {
			tx.Add(addFlows...)
		}
		if len(deleteFlows) > 0 {
			tx.Delete(deleteFlows...)
		}
		return tx.Commit()
	})
	if err != nil {
		return nil, err
	}

	out := make([]string, 0, len(addFlows))
	for _, f := range addFlows {
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
