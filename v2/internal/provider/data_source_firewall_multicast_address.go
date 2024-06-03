// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure multicast addresses.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceFirewallMulticastAddress() *schema.Resource {
	return &schema.Resource{
		Description: "Configure multicast addresses.",

		ReadContext: dataSourceFirewallMulticastAddressRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"associated_interface": {
				Type:        schema.TypeString,
				Description: "Interface associated with the address object. When setting up a policy, only addresses associated with this interface are available.",
				Computed:    true,
			},
			"color": {
				Type:        schema.TypeInt,
				Description: "Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets value to 1).",
				Computed:    true,
			},
			"comment": {
				Type:        schema.TypeString,
				Description: "Comment.",
				Computed:    true,
			},
			"end_ip": {
				Type:        schema.TypeString,
				Description: "Final IPv4 address (inclusive) in the range for the address.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Multicast address name.",
				Required:    true,
			},
			"start_ip": {
				Type:        schema.TypeString,
				Description: "First IPv4 address (inclusive) in the range for the address.",
				Computed:    true,
			},
			"subnet": {
				Type:        schema.TypeString,
				Description: "Broadcast address and subnet.",
				Computed:    true,
			},
			"tagging": {
				Type:        schema.TypeList,
				Description: "Config object tagging.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"category": {
							Type:        schema.TypeString,
							Description: "Tag category.",
							Computed:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "Tagging entry name.",
							Computed:    true,
						},
						"tags": {
							Type:        schema.TypeList,
							Description: "Tags.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Description: "Tag name.",
										Computed:    true,
									},
								},
							},
						},
					},
				},
			},
			"type": {
				Type:        schema.TypeString,
				Description: "Type of address object: multicast IP address range or broadcast IP/mask to be treated as a multicast address.",
				Computed:    true,
			},
			"visibility": {
				Type:        schema.TypeString,
				Description: "Enable/disable visibility of the multicast address on the GUI.",
				Computed:    true,
			},
		},
	}
}

func dataSourceFirewallMulticastAddressRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallMulticastAddress(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallMulticastAddress dataSource: %v", err)
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

	diags := refreshObjectFirewallMulticastAddress(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
