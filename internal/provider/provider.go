package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	schema.DescriptionKind = schema.StringMarkdown
}

func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			ResourcesMap: map[string]*schema.Resource{
				"sdk_typeset_conflictswith": resourceTypesetConflictsWith(),
			},
		}

		p.ConfigureContextFunc = configure(version, p)

		return p
	}
}

type apiClient struct {
	// Intentionally blank.
}

func configure(providerVersion string, p *schema.Provider) func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
		var diags diag.Diagnostics

		if p.TerraformVersion != "" {
			tfVersion, err := version.NewVersion(p.TerraformVersion)

			if err != nil {
				return nil, diag.FromErr(fmt.Errorf("Unable to parse Terraform CLI version: %w", err))
			}

			tfVersionConstraint, err := version.NewConstraint(">= 1.0")

			if err != nil {
				return nil, diag.FromErr(fmt.Errorf("Unable to parse Terraform CLI version constaint: %w", err))
			}

			if !tfVersionConstraint.Check(tfVersion) {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "Unsupported Terraform CLI Version",
					Detail:   fmt.Sprintf("Current Terraform CLI version is %s. This provider requires Terraform CLI version 1.0 or later.", p.TerraformVersion),
				})
				return nil, diags
			}
		}

		return &apiClient{}, diags
	}
}
