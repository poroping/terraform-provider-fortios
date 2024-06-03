// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure software switch interfaces by grouping physical and WiFi interfaces.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceSystemSwitchInterface() *schema.Resource {
	return &schema.Resource{
		Description: "Configure software switch interfaces by grouping physical and WiFi interfaces.",

		ReadContext: dataSourceSystemSwitchInterfaceRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"intra_switch_policy": {
				Type:        schema.TypeString,
				Description: "Allow any traffic between switch interfaces or require firewall policies to allow traffic between switch interfaces.",
				Computed:    true,
			},
			"mac_ttl": {
				Type:        schema.TypeInt,
				Description: "Duration for which MAC addresses are held in the ARP table (300 - 8640000 sec, default = 300).",
				Computed:    true,
			},
			"member": {
				Type:        schema.TypeList,
				Description: "Names of the interfaces that belong to the virtual switch.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interface_name": {
							Type:        schema.TypeString,
							Description: "Physical interface name.",
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Interface name (name cannot be in use by any other interfaces, VLANs, or inter-VDOM links).",
				Required:    true,
			},
			"span": {
				Type:        schema.TypeString,
				Description: "Enable/disable port spanning. Port spanning echoes traffic received by the software switch to the span destination port.",
				Computed:    true,
			},
			"span_dest_port": {
				Type:        schema.TypeString,
				Description: "SPAN destination port name. All traffic on the SPAN source ports is echoed to the SPAN destination port.",
				Computed:    true,
			},
			"span_direction": {
				Type:        schema.TypeString,
				Description: "The direction in which the SPAN port operates, either: rx, tx, or both.",
				Computed:    true,
			},
			"span_source_port": {
				Type:        schema.TypeList,
				Description: "Physical interface name. Port spanning echoes all traffic on the SPAN source ports to the SPAN destination port.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interface_name": {
							Type:        schema.TypeString,
							Description: "Physical interface name.",
							Computed:    true,
						},
					},
				},
			},
			"type": {
				Type:        schema.TypeString,
				Description: "Type of switch based on functionality: switch for normal functionality, or hub to duplicate packets to all port members.",
				Computed:    true,
			},
			"vdom": {
				Type:        schema.TypeString,
				Description: "VDOM that the software switch belongs to.",
				Computed:    true,
			},
		},
	}
}

func dataSourceSystemSwitchInterfaceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemSwitchInterface(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemSwitchInterface dataSource: %v", err)
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

	diags := refreshObjectSystemSwitchInterface(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
