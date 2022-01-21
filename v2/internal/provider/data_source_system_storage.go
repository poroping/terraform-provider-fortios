// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure logical storage.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceSystemStorage() *schema.Resource {
	return &schema.Resource{
		Description: "Configure logical storage.",

		ReadContext: dataSourceSystemStorageRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"device": {
				Type:        schema.TypeString,
				Description: "Partition device.",
				Computed:    true,
			},
			"media_status": {
				Type:        schema.TypeString,
				Description: "The physical status of current media.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Storage name.",
				Required:    true,
			},
			"order": {
				Type:        schema.TypeInt,
				Description: "Set storage order.",
				Computed:    true,
			},
			"partition": {
				Type:        schema.TypeString,
				Description: "Label of underlying partition.",
				Computed:    true,
			},
			"size": {
				Type:        schema.TypeInt,
				Description: "Partition size.",
				Computed:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Enable/disable storage.",
				Computed:    true,
			},
			"usage": {
				Type:        schema.TypeString,
				Description: "Use hard disk for logging or WAN Optimization (default = log).",
				Computed:    true,
			},
			"wanopt_mode": {
				Type:        schema.TypeString,
				Description: "WAN Optimization mode (default = mix).",
				Computed:    true,
			},
		},
	}
}

func dataSourceSystemStorageRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemStorage(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemStorage dataSource: %v", err)
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

	diags := refreshObjectSystemStorage(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
