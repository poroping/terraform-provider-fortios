// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Settings for TACACS+ accounting events filter.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceLogTacacsaccounting2Filter() *schema.Resource {
	return &schema.Resource{
		Description: "Settings for TACACS+ accounting events filter.",

		ReadContext: dataSourceLogTacacsaccounting2FilterRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"cli_cmd_audit": {
				Type:        schema.TypeString,
				Description: "Enable/disable TACACS+ accounting for CLI commands audit.",
				Computed:    true,
			},
			"config_change_audit": {
				Type:        schema.TypeString,
				Description: "Enable/disable TACACS+ accounting for configuration change events audit.",
				Computed:    true,
			},
			"login_audit": {
				Type:        schema.TypeString,
				Description: "Enable/disable TACACS+ accounting for login events audit.",
				Computed:    true,
			},
		},
	}
}

func dataSourceLogTacacsaccounting2FilterRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "LogTacacsaccounting2Filter"

	o, err := c.Cmdb.ReadLogTacacsaccounting2Filter(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading LogTacacsaccounting2Filter dataSource: %v", err)
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

	diags := refreshObjectLogTacacsaccounting2Filter(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
