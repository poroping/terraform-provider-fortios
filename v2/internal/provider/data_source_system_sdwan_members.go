// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: FortiGate interfaces added to the SD-WAN.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceSystemSdwanMembers() *schema.Resource {
	return &schema.Resource{
		Description: "FortiGate interfaces added to the SD-WAN.",

		ReadContext: dataSourceSystemSdwanMembersRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"comment": {
				Type:        schema.TypeString,
				Description: "Comments.",
				Computed:    true,
			},
			"cost": {
				Type:        schema.TypeInt,
				Description: "Cost of this interface for services in SLA mode (0 - 4294967295, default = 0).",
				Computed:    true,
			},
			"gateway": {
				Type:        schema.TypeString,
				Description: "The default gateway for this interface. Usually the default gateway of the Internet service provider that this interface is connected to.",
				Computed:    true,
			},
			"gateway6": {
				Type:        schema.TypeString,
				Description: "IPv6 gateway.",
				Computed:    true,
			},
			"ingress_spillover_threshold": {
				Type:        schema.TypeInt,
				Description: "Ingress spillover threshold for this interface (0 - 16776000 kbit/s). When this traffic volume threshold is reached, new sessions spill over to other interfaces in the SD-WAN.",
				Computed:    true,
			},
			"interface": {
				Type:        schema.TypeString,
				Description: "Interface name.",
				Computed:    true,
			},
			"priority": {
				Type:        schema.TypeInt,
				Description: "Priority of the interface for IPv4 (1 - 65535, default = 1). Used for SD-WAN rules or priority rules.",
				Computed:    true,
			},
			"priority6": {
				Type:        schema.TypeInt,
				Description: "Priority of the interface for IPv6 (1 - 65535, default = 1024). Used for SD-WAN rules or priority rules.",
				Computed:    true,
			},
			"seq_num": {
				Type:        schema.TypeInt,
				Description: "Sequence number(1-512).",
				Computed:    true,
			},
			"source": {
				Type:        schema.TypeString,
				Description: "Source IP address used in the health-check packet to the server.",
				Computed:    true,
			},
			"source6": {
				Type:        schema.TypeString,
				Description: "Source IPv6 address used in the health-check packet to the server.",
				Computed:    true,
			},
			"spillover_threshold": {
				Type:        schema.TypeInt,
				Description: "Egress spillover threshold for this interface (0 - 16776000 kbit/s). When this traffic volume threshold is reached, new sessions spill over to other interfaces in the SD-WAN.",
				Computed:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Enable/disable this interface in the SD-WAN.",
				Computed:    true,
			},
			"volume_ratio": {
				Type:        schema.TypeInt,
				Description: "Measured volume ratio (this value / sum of all values = percentage of link volume, 1 - 255).",
				Computed:    true,
			},
			"weight": {
				Type:        schema.TypeInt,
				Description: "Weight of this interface for weighted load balancing. (1 - 255) More traffic is directed to interfaces with higher weights.",
				Computed:    true,
			},
			"zone": {
				Type:        schema.TypeString,
				Description: "Zone name.",
				Computed:    true,
			},
		},
	}
}

func dataSourceSystemSdwanMembersRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemSdwanMembers(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemSdwanMembers dataSource: %v", err)
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

	diags := refreshObjectSystemSdwanMembers(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
