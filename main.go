package main

import (
	"context"
	"flag"
	"log"

	"github.com/bflad/terraform-provider-mux5/internal/provider"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/tf5server"
	"github.com/hashicorp/terraform-plugin-mux/tf5muxserver"
)

// Run "go generate" to format example terraform files and generate the docs for the registry/website

// If you do not have terraform installed, you can remove the formatting command, but its suggested to
// ensure the documentation is formatted properly.
//go:generate terraform fmt -recursive ./examples/

// Run the docs generation tool, check its repository for more information on how it works and how docs
// can be customized.
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

var (
	// these will be set by the goreleaser configuration
	// to appropriate values for the compiled binary
	version string = "dev"

	// goreleaser can also pass the specific commit if you want
	// commit  string = ""
)

func main() {
	var debugFlag bool

	flag.BoolVar(&debugFlag, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	ctx := context.Background()
	providers := []func() tfprotov5.ProviderServer{
		provider.New(version)().GRPCProvider,
	}

	muxServer, err := tf5muxserver.NewMuxServer(ctx, providers...)

	if err != nil {
		log.Fatal(err)
	}

	var serveOpts []tf5server.ServeOpt

	if debugFlag {
		serveOpts = append(serveOpts, tf5server.WithManagedDebug())
	}

	err = tf5server.Serve(
		"registry.terraform.io/bflad/mux5",
		muxServer.ProviderServer,
		serveOpts...,
	)

	if err != nil {
		log.Fatal(err)
	}
}
