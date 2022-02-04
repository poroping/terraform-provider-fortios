// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure auto script.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceSystemAutoScript() *schema.Resource {
	return &schema.Resource{
		Description: "Configure auto script.",

		ReadContext: dataSourceSystemAutoScriptRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"interval": {
				Type:        schema.TypeInt,
				Description: "Repeat interval in seconds.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Auto script name.",
				Required:    true,
			},
			"output_size": {
				Type:        schema.TypeInt,
				Description: "Number of megabytes to limit script output to (10 - 1024, default = 10).",
				Computed:    true,
			},
			"repeat": {
				Type:        schema.TypeInt,
				Description: "Number of times to repeat this script (0 = infinite).",
				Computed:    true,
			},
			"script": {
				Type:        schema.TypeString,
				Description: "List of FortiOS CLI commands to repeat.",
				Computed:    true,
			},
			"start": {
				Type:        schema.TypeString,
				Description: "Script starting mode.",
				Computed:    true,
			},
			"timeout": {
				Type:        schema.TypeInt,
				Description: "Maximum running time for this script in seconds (0 = no timeout).",
				Computed:    true,
			},
		},
	}
}

func dataSourceSystemAutoScriptRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	i := d.Get("name")
	mkey := utils.ParseMkey(i)

	o, err := c.Cmdb.ReadSystemAutoScript(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemAutoScript dataSource: %v", err)
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

	diags := refreshObjectSystemAutoScript(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
