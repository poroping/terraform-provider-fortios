package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/poroping/terraform-provider-fortios/fortios"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: fortios.Provider})
}
