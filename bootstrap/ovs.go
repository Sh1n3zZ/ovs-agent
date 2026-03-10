package bootstrap

import (
	libovs "github.com/digitalocean/go-openvswitch/ovs"
	ovsclient "github.com/Sh1n3zZ/ovs-agent/server/ovs"
)

// NewOVSClient initializes and returns a new Open vSwitch client for use at application startup.
// It uses the helper functions defined in the server/ovs package to construct the client.
func NewOVSClient() *libovs.Client {
	return ovsclient.NewSudoClient()
}

