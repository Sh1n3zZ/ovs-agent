package bootstrap

import (
	libovs "github.com/digitalocean/go-openvswitch/ovs"
)

type Application struct {
	Env       *Env
	OVSClient *libovs.Client
}

func App() Application {
	app := &Application{}
	app.Env = NewConfig()
	app.OVSClient = NewOVSClient()
	return *app
}
