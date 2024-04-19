package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceExample() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceExampleCreate,
		ReadContext:   resourceExampleRead,
		UpdateContext: resourceExampleUpdate,
		DeleteContext: resourceExampleDelete,

		Description: "Example resource",

		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Example identifier",
			},
			"optional": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Example optional attribute",
			},
		},
	}
}

func resourceExampleCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	d.SetId("example")

	return nil
}

func resourceExampleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func resourceExampleUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func resourceExampleDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}
