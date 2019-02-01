package main

import (
	"github.com/Cloudlock/terraform-provider-librato/librato"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: librato.Provider})
}
