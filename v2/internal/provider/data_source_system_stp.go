// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.4.6,v7.0.2 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure Spanning Tree Protocol (STP).

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceSystemStp() *schema.Resource {
	return &schema.Resource{
		Description: "Configure Spanning Tree Protocol (STP).",

		ReadContext: dataSourceSystemStpRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"forward_delay": {
				Type:        schema.TypeInt,
				Description: "Forward delay (4 - 30 sec, default = 15).",
				Computed:    true,
			},
			"hello_time": {
				Type:        schema.TypeInt,
				Description: "Hello time (1 - 10 sec, default = 2).",
				Computed:    true,
			},
			"max_age": {
				Type:        schema.TypeInt,
				Description: "Maximum packet age (6 - 40 sec, default = 20).",
				Computed:    true,
			},
			"max_hops": {
				Type:        schema.TypeInt,
				Description: "Maximum number of hops (1 - 40, default = 20).",
				Computed:    true,
			},
			"switch_priority": {
				Type:        schema.TypeString,
				Description: "STP switch priority; the lower the number the higher the priority (select from 0, 4096, 8192, 12288, 16384, 20480, 24576, 28672, 32768, 36864, 40960, 45056, 49152, 53248, and 57344).",
				Computed:    true,
			},
		},
	}
}

func dataSourceSystemStpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "SystemStp"

	o, err := c.Cmdb.ReadSystemStp(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemStp dataSource: %v", err)
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

	diags := refreshObjectSystemStp(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
