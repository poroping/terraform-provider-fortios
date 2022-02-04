// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure IPv6 static routing tables.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceRouterStatic6() *schema.Resource {
	return &schema.Resource{
		Description: "Configure IPv6 static routing tables.",

		ReadContext: dataSourceRouterStatic6Read,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"bfd": {
				Type:        schema.TypeString,
				Description: "Enable/disable Bidirectional Forwarding Detection (BFD).",
				Computed:    true,
			},
			"blackhole": {
				Type:        schema.TypeString,
				Description: "Enable/disable black hole.",
				Computed:    true,
			},
			"comment": {
				Type:        schema.TypeString,
				Description: "Optional comments.",
				Computed:    true,
			},
			"device": {
				Type:        schema.TypeString,
				Description: "Gateway out interface or tunnel.",
				Computed:    true,
			},
			"devindex": {
				Type:        schema.TypeInt,
				Description: "Device index (0 - 4294967295).",
				Computed:    true,
			},
			"distance": {
				Type:        schema.TypeInt,
				Description: "Administrative distance (1 - 255).",
				Computed:    true,
			},
			"dst": {
				Type:        schema.TypeString,
				Description: "Destination IPv6 prefix.",
				Computed:    true,
			},
			"dynamic_gateway": {
				Type:        schema.TypeString,
				Description: "Enable use of dynamic gateway retrieved from Router Advertisement (RA).",
				Computed:    true,
			},
			"gateway": {
				Type:        schema.TypeString,
				Description: "IPv6 address of the gateway.",
				Computed:    true,
			},
			"link_monitor_exempt": {
				Type:        schema.TypeString,
				Description: "Enable/disable withdrawal of this static route when link monitor or health check is down.",
				Computed:    true,
			},
			"priority": {
				Type:        schema.TypeInt,
				Description: "Administrative priority (1 - 65535).",
				Computed:    true,
			},
			"sdwan": {
				Type:        schema.TypeString,
				Description: "Enable/disable egress through the SD-WAN.",
				Computed:    true,
			},
			"sdwan_zone": {
				Type:        schema.TypeList,
				Description: "Choose SD-WAN Zone.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "SD-WAN zone name.",
							Computed:    true,
						},
					},
				},
			},
			"seq_num": {
				Type:        schema.TypeInt,
				Description: "Sequence number.",
				Computed:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Enable/disable this static route.",
				Computed:    true,
			},
			"virtual_wan_link": {
				Type:        schema.TypeString,
				Description: "Enable/disable egress through the virtual-wan-link.",
				Computed:    true,
			},
			"vrf": {
				Type:        schema.TypeInt,
				Description: "Virtual Routing Forwarding ID.",
				Computed:    true,
			},
		},
	}
}

func dataSourceRouterStatic6Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	i := d.Get("seq_num")
	mkey := utils.ParseMkey(i)

	o, err := c.Cmdb.ReadRouterStatic6(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading RouterStatic6 dataSource: %v", err)
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

	diags := refreshObjectRouterStatic6(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
