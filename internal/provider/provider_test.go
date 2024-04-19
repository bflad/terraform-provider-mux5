package provider

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-mux/tf5muxserver"
)

var protov5ProviderFactories = map[string]func() (tfprotov5.ProviderServer, error){
	"mux5": func() (tfprotov5.ProviderServer, error) {
		ctx := context.Background()
		providers := []func() tfprotov5.ProviderServer{
			New("test")().GRPCProvider,
		}

		muxServer, err := tf5muxserver.NewMuxServer(ctx, providers...)

		if err != nil {
			return nil, err
		}

		return muxServer.ProviderServer(), nil
	},
}

func TestProvider(t *testing.T) {
	if err := New("dev")().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func testAccPreCheck(t *testing.T) {
	// Intentionally blank.
}
