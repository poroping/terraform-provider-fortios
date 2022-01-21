// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure VXLAN devices.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceSystemVxlan() *schema.Resource {
	return &schema.Resource{
		Description: "Configure VXLAN devices.",

		ReadContext: dataSourceSystemVxlanRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"dstport": {
				Type:        schema.TypeInt,
				Description: "VXLAN destination port (1 - 65535, default = 4789).",
				Computed:    true,
			},
			"interface": {
				Type:        schema.TypeString,
				Description: "Outgoing interface for VXLAN encapsulated traffic.",
				Computed:    true,
			},
			"ip_version": {
				Type:        schema.TypeString,
				Description: "IP version to use for the VXLAN interface and so for communication over the VXLAN. IPv4 or IPv6 unicast or multicast.",
				Computed:    true,
			},
			"multicast_ttl": {
				Type:        schema.TypeInt,
				Description: "VXLAN multicast TTL (1-255, default = 0).",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "VXLAN device or interface name. Must be a unique interface name.",
				Required:    true,
			},
			"remote_ip": {
				Type:        schema.TypeList,
				Description: "IPv4 address of the VXLAN interface on the device at the remote end of the VXLAN.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip": {
							Type:        schema.TypeString,
							Description: "IPv4 address.",
							Computed:    true,
						},
					},
				},
			},
			"remote_ip6": {
				Type:        schema.TypeList,
				Description: "IPv6 IP address of the VXLAN interface on the device at the remote end of the VXLAN.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip6": {
							Type:        schema.TypeString,
							Description: "IPv6 address.",
							Computed:    true,
						},
					},
				},
			},
			"vni": {
				Type:        schema.TypeInt,
				Description: "VXLAN network ID.",
				Computed:    true,
			},
		},
	}
}

func dataSourceSystemVxlanRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemVxlan(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemVxlan dataSource: %v", err)
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

	diags := refreshObjectSystemVxlan(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
