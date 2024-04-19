package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccResourceExample(t *testing.T) {
	t.Parallel()

	resource.UnitTest(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: protov5ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `resource "mux5_example" "test" {}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("mux5_example.test", "id", "example"),
				),
			},
		},
	})
}
