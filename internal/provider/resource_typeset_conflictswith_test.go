package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceTypeSetConflictsWith(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		ExternalProviders: map[string]resource.ExternalProvider{
			"random": {
				Source:            "hashicorp/random",
				VersionConstraint: "3.1.0",
			},
		},
		Steps: []resource.TestStep{
			{
				Config:      testAccResourceTypeSetConflictsBothAttrs,
				ExpectError: regexp.MustCompile(`cannot define both example_attr1 and example_attr2`),
			},
			{
				Config: testAccResourceTypeSetConflictsAttr1,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckTypeSetElemAttrPair(
						"sdk_typeset_conflictswith.test", "example_typeset_block.*.example_attr1",
						"random_pet.test", "id",
					),
				),
			},
			{
				Config: testAccResourceTypeSetConflictsAttr2,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckTypeSetElemAttrPair(
						"sdk_typeset_conflictswith.test", "example_typeset_block.*.example_attr2",
						"random_pet.test", "id",
					),
				),
			},
			{
				Config: testAccResourceTypeSetConflictsAttr1,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckTypeSetElemAttrPair(
						"sdk_typeset_conflictswith.test", "example_typeset_block.*.example_attr1",
						"random_pet.test", "id",
					),
				),
			},
		},
	})
}

const testAccResourceTypeSetConflictsAttr1 = `
resource "random_pet" "test" {}

resource "sdk_typeset_conflictswith" "test" {
  example_typeset_block {
    example_attr1 = random_pet.test.id
  }
}
`

const testAccResourceTypeSetConflictsAttr2 = `
resource "random_pet" "test" {}

resource "sdk_typeset_conflictswith" "test" {
  example_typeset_block {
    example_attr2 = random_pet.test.id
  }
}
`

const testAccResourceTypeSetConflictsBothAttrs = `
resource "random_pet" "test" {}

resource "sdk_typeset_conflictswith" "test" {
  example_typeset_block {
    example_attr1 = random_pet.test.id
    example_attr2 = random_pet.test.id
  }
}
`
