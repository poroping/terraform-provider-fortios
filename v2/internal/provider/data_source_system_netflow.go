// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure NetFlow.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceSystemNetflow() *schema.Resource {
	return &schema.Resource{
		Description: "Configure NetFlow.",

		ReadContext: dataSourceSystemNetflowRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"active_flow_timeout": {
				Type:        schema.TypeInt,
				Description: "Timeout to report active flows (60 - 3600 sec, default = 1800).",
				Computed:    true,
			},
			"collector_ip": {
				Type:        schema.TypeString,
				Description: "Collector IP.",
				Computed:    true,
			},
			"collector_port": {
				Type:        schema.TypeInt,
				Description: "NetFlow collector port number.",
				Computed:    true,
			},
			"inactive_flow_timeout": {
				Type:        schema.TypeInt,
				Description: "Timeout for periodic report of finished flows (10 - 600 sec, default = 15).",
				Computed:    true,
			},
			"interface": {
				Type:        schema.TypeString,
				Description: "Specify outgoing interface to reach server.",
				Computed:    true,
			},
			"interface_select_method": {
				Type:        schema.TypeString,
				Description: "Specify how to select outgoing interface to reach server.",
				Computed:    true,
			},
			"source_ip": {
				Type:        schema.TypeString,
				Description: "Source IP address for communication with the NetFlow agent.",
				Computed:    true,
			},
			"template_tx_counter": {
				Type:        schema.TypeInt,
				Description: "Counter of flowset records before resending a template flowset record.",
				Computed:    true,
			},
			"template_tx_timeout": {
				Type:        schema.TypeInt,
				Description: "Timeout for periodic template flowset transmission (60 - 86400 sec, default = 1800).",
				Computed:    true,
			},
		},
	}
}

func dataSourceSystemNetflowRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "SystemNetflow"

	o, err := c.Cmdb.ReadSystemNetflow(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemNetflow dataSource: %v", err)
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

	diags := refreshObjectSystemNetflow(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
