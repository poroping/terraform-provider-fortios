// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Settings for TACACS+ accounting.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceLogTacacsaccountingSetting() *schema.Resource {
	return &schema.Resource{
		Description: "Settings for TACACS+ accounting.",

		ReadContext: dataSourceLogTacacsaccountingSettingRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"server": {
				Type:        schema.TypeString,
				Description: "Address of TACACS+ server.",
				Computed:    true,
			},
			"server_key": {
				Type:        schema.TypeString,
				Description: "Key to access the TACACS+ server.",
				Computed:    true,
				Sensitive:   true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Enable/disable TACACS+ accounting.",
				Computed:    true,
			},
		},
	}
}

func dataSourceLogTacacsaccountingSettingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "LogTacacsaccountingSetting"

	o, err := c.Cmdb.ReadLogTacacsaccountingSetting(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading LogTacacsaccountingSetting dataSource: %v", err)
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

	diags := refreshObjectLogTacacsaccountingSetting(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
