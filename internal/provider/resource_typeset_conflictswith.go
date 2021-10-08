package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTypesetConflictsWith() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceTypesetConflictsWithCreate,
		ReadContext:   resourceTypesetConflictsWithRead,
		UpdateContext: resourceTypesetConflictsWithUpdate,
		DeleteContext: resourceTypesetConflictsWithDelete,

		Schema: map[string]*schema.Schema{
			"example_typeset_block": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"example_attr1": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"example_attr2": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},

		CustomizeDiff: func(ctx context.Context, diff *schema.ResourceDiff, meta interface{}) error {
			blocksRaw, ok := diff.GetOk("example_typeset_block")

			if !ok {
				return nil
			}

			blocks, ok := blocksRaw.(*schema.Set)

			if !ok {
				return nil
			}

			for _, blockRaw := range blocks.List() {
				block, ok := blockRaw.(map[string]interface{})

				if !ok {
					continue
				}

				attr1, attr1Ok := block["example_attr1"].(string)
				attr2, attr2Ok := block["example_attr2"].(string)

				if attr1Ok && attr2Ok && attr1 != "" && attr2 != "" {
					return fmt.Errorf("cannot define both example_attr1 and example_attr2")
				}
			}

			return nil
		},
	}
}

func resourceTypesetConflictsWithCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Configuration will be copied to state

	d.SetId("example")

	return nil
}

func resourceTypesetConflictsWithRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Intentionally blank.

	return nil
}

func resourceTypesetConflictsWithUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Intentionally blank.

	return nil
}

func resourceTypesetConflictsWithDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Intentionally blank.

	return nil
}
