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
		Steps: []resource.TestStep{
			{
				Config:      testAccResourceTypeSetConflictsBothAttrs,
				ExpectError: regexp.MustCompile(`cannot define both example_attr1 and example_attr2`),
			},
			{
				Config: testAccResourceTypeSetConflictsAttr1,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckTypeSetElemNestedAttrs("sdk_typeset_conflictswith.test", "example_typeset_block.*", map[string]string{
						"example_attr1": "test1",
					}),
				),
			},
			{
				Config: testAccResourceTypeSetConflictsAttr2,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckTypeSetElemNestedAttrs("sdk_typeset_conflictswith.test", "example_typeset_block.*", map[string]string{
						"example_attr2": "test2",
					}),
				),
			},
			{
				Config: testAccResourceTypeSetConflictsAttr1,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckTypeSetElemNestedAttrs("sdk_typeset_conflictswith.test", "example_typeset_block.*", map[string]string{
						"example_attr1": "test1",
					}),
				),
			},
		},
	})
}

const testAccResourceTypeSetConflictsAttr1 = `
resource "sdk_typeset_conflictswith" "test" {
  example_typeset_block {
    example_attr1 = "test1"
  }
}
`

const testAccResourceTypeSetConflictsAttr2 = `
resource "sdk_typeset_conflictswith" "test" {
  example_typeset_block {
    example_attr2 = "test2"
  }
}
`

const testAccResourceTypeSetConflictsBothAttrs = `
resource "sdk_typeset_conflictswith" "test" {
  example_typeset_block {
    example_attr1 = "test1"
    example_attr2 = "test2"
  }
}
`
