package main

import (
	"github.com/hashicorp/terraform/builtin/providers/packet"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: packet.Provider,
	})
}
