// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure password policy for locally defined administrator passwords and IPsec VPN pre-shared keys.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceSystemPasswordPolicy() *schema.Resource {
	return &schema.Resource{
		Description: "Configure password policy for locally defined administrator passwords and IPsec VPN pre-shared keys.",

		ReadContext: dataSourceSystemPasswordPolicyRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"apply_to": {
				Type:        schema.TypeString,
				Description: "Apply password policy to administrator passwords or IPsec pre-shared keys or both. Separate entries with a space.",
				Computed:    true,
			},
			"change_4_characters": {
				Type:        schema.TypeString,
				Description: "Enable/disable changing at least 4 characters for a new password (This attribute overrides reuse-password if both are enabled).",
				Computed:    true,
			},
			"expire_day": {
				Type:        schema.TypeInt,
				Description: "Number of days after which passwords expire (1 - 999 days, default = 90).",
				Computed:    true,
			},
			"expire_status": {
				Type:        schema.TypeString,
				Description: "Enable/disable password expiration.",
				Computed:    true,
			},
			"min_change_characters": {
				Type:        schema.TypeInt,
				Description: "Minimum number of unique characters in new password which do not exist in old password (0 - 128, default = 0. This attribute overrides reuse-password if both are enabled).",
				Computed:    true,
			},
			"min_lower_case_letter": {
				Type:        schema.TypeInt,
				Description: "Minimum number of lowercase characters in password (0 - 128, default = 0).",
				Computed:    true,
			},
			"min_non_alphanumeric": {
				Type:        schema.TypeInt,
				Description: "Minimum number of non-alphanumeric characters in password (0 - 128, default = 0).",
				Computed:    true,
			},
			"min_number": {
				Type:        schema.TypeInt,
				Description: "Minimum number of numeric characters in password (0 - 128, default = 0).",
				Computed:    true,
			},
			"min_upper_case_letter": {
				Type:        schema.TypeInt,
				Description: "Minimum number of uppercase characters in password (0 - 128, default = 0).",
				Computed:    true,
			},
			"minimum_length": {
				Type:        schema.TypeInt,
				Description: "Minimum password length (8 - 128, default = 8).",
				Computed:    true,
			},
			"reuse_password": {
				Type:        schema.TypeString,
				Description: "Enable/disable reuse of password. If both reuse-password and min-change-characters are enabled, min-change-characters overrides.",
				Computed:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Enable/disable setting a password policy for locally defined administrator passwords and IPsec VPN pre-shared keys.",
				Computed:    true,
			},
		},
	}
}

func dataSourceSystemPasswordPolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*apiClient).Client
	// c.Retries = 1

	urlparams := &models.CmdbRequestParams{}
	vdomparam := ""
	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}
	urlparams.Vdom = vdomparam

	mkey := "SystemPasswordPolicy"

	o, err := c.Cmdb.ReadSystemPasswordPolicy(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemPasswordPolicy dataSource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] dataSource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	sort := false
	if v, ok := d.GetOk("dynamic_sort_subtable"); ok {
		if b, ok := v.(bool); ok {
			sort = b
		}
	}

	diags := refreshObjectSystemPasswordPolicy(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
