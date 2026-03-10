package ovs

import ovs "github.com/digitalocean/go-openvswitch/ovs"

// NewClient creates a new Open vSwitch client with the given options.
// Callers can pass option functions such as ovs.Sudo() to customize behavior.
func NewClient(options ...ovs.OptionFunc) *ovs.Client {
	return ovs.New(options...)
}

// NewSudoClient creates a new Open vSwitch client that runs commands via sudo.
// This is a convenience wrapper around NewClient(ovs.Sudo()).
func NewSudoClient() *ovs.Client {
	return NewClient(ovs.Sudo())
}

