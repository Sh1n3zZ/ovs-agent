package flows

import (
	"fmt"
	"net"

	ovshelpers "github.com/Sh1n3zZ/ovs-agent/server/ovs"
	libovs "github.com/digitalocean/go-openvswitch/ovs"
)

// InstallStaticARPBinding installs a pair of OpenFlow rules that implement a static ARP binding.
// It allows ARP packets from the given (IP, MAC, inPort) and drops other ARP packets claiming the same IP.
//
// The resulting rules are equivalent to:
//
//	priority=100,in_port=<inPort>,arp,arp_spa=<ip>,arp_sha=<mac>,actions=normal
//	priority=90,in_port=<inPort>,arp,arp_spa=<ip>,actions=drop
func InstallStaticARPBinding(client *libovs.Client, bridge string, inPort int, ip string, macStr string) ([]string, error) {
	mac, err := net.ParseMAC(macStr)
	if err != nil {
		return nil, fmt.Errorf("parse mac %q: %w", macStr, err)
	}

	allowFlow := &libovs.Flow{
		Priority: 100,
		Protocol: libovs.ProtocolARP,
		InPort:   inPort,
		Matches: []libovs.Match{
			libovs.ARPSourceProtocolAddress(ip),
			libovs.ARPSourceHardwareAddress(mac),
		},
		Actions: []libovs.Action{
			libovs.Normal(),
		},
	}

	denyFlow := &libovs.Flow{
		Priority: 90,
		Protocol: libovs.ProtocolARP,
		InPort:   inPort,
		Matches: []libovs.Match{
			libovs.ARPSourceProtocolAddress(ip),
		},
		Actions: []libovs.Action{
			libovs.Drop(),
		},
	}

	return ovshelpers.AddBridgeFlowBundle(client, bridge, []*libovs.Flow{allowFlow, denyFlow}, nil)
}

// RemoveStaticARPBinding removes the pair of OpenFlow rules installed by InstallStaticARPBinding
// for the given (bridge, inPort, ip, mac).
func RemoveStaticARPBinding(client *libovs.Client, bridge string, inPort int, ip string, macStr string) error {
	mac, err := net.ParseMAC(macStr)
	if err != nil {
		return fmt.Errorf("parse mac %q: %w", macStr, err)
	}

	allowMatch := &libovs.MatchFlow{
		Protocol: libovs.ProtocolARP,
		InPort:   inPort,
		Matches: []libovs.Match{
			libovs.ARPSourceProtocolAddress(ip),
			libovs.ARPSourceHardwareAddress(mac),
		},
	}

	denyMatch := &libovs.MatchFlow{
		Protocol: libovs.ProtocolARP,
		InPort:   inPort,
		Matches: []libovs.Match{
			libovs.ARPSourceProtocolAddress(ip),
		},
	}

	_, err = ovshelpers.AddBridgeFlowBundle(client, bridge, nil, []*libovs.MatchFlow{allowMatch, denyMatch})
	return err
}
