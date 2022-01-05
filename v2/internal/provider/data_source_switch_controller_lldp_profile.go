// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure FortiSwitch LLDP profiles.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceSwitchControllerLldpProfile() *schema.Resource {
	return &schema.Resource{
		Description: "Configure FortiSwitch LLDP profiles.",

		ReadContext: dataSourceSwitchControllerLldpProfileRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"8021_tlvs": {
				Type:        schema.TypeString,
				Description: "Transmitted IEEE 802.1 TLVs.",
				Computed:    true,
			},
			"8023_tlvs": {
				Type:        schema.TypeString,
				Description: "Transmitted IEEE 802.3 TLVs.",
				Computed:    true,
			},
			"auto_isl": {
				Type:        schema.TypeString,
				Description: "Enable/disable auto inter-switch LAG.",
				Computed:    true,
			},
			"auto_isl_hello_timer": {
				Type:        schema.TypeInt,
				Description: "Auto inter-switch LAG hello timer duration (1 - 30 sec, default = 3).",
				Computed:    true,
			},
			"auto_isl_port_group": {
				Type:        schema.TypeInt,
				Description: "Auto inter-switch LAG port group ID (0 - 9).",
				Computed:    true,
			},
			"auto_isl_receive_timeout": {
				Type:        schema.TypeInt,
				Description: "Auto inter-switch LAG timeout if no response is received (3 - 90 sec, default = 9).",
				Computed:    true,
			},
			"auto_mclag_icl": {
				Type:        schema.TypeString,
				Description: "Enable/disable MCLAG inter chassis link.",
				Computed:    true,
			},
			"custom_tlvs": {
				Type:        schema.TypeList,
				Description: "Configuration method to edit custom TLV entries.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"information_string": {
							Type:        schema.TypeString,
							Description: "Organizationally defined information string (0 - 507 hexadecimal bytes).",
							Computed:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "TLV name (not sent).",
							Computed:    true,
						},
						"oui": {
							Type:        schema.TypeString,
							Description: "Organizationally unique identifier (OUI), a 3-byte hexadecimal number, for this TLV.",
							Computed:    true,
						},
						"subtype": {
							Type:        schema.TypeInt,
							Description: "Organizationally defined subtype (0 - 255).",
							Computed:    true,
						},
					},
				},
			},
			"med_location_service": {
				Type:        schema.TypeList,
				Description: "Configuration method to edit Media Endpoint Discovery (MED) location service type-length-value (TLV) categories.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Location service type name.",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Enable or disable this TLV.",
							Computed:    true,
						},
						"sys_location_id": {
							Type:        schema.TypeString,
							Description: "Location service ID.",
							Computed:    true,
						},
					},
				},
			},
			"med_network_policy": {
				Type:        schema.TypeList,
				Description: "Configuration method to edit Media Endpoint Discovery (MED) network policy type-length-value (TLV) categories.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"assign_vlan": {
							Type:        schema.TypeString,
							Description: "Enable/disable VLAN assignment when this profile is applied on managed FortiSwitch port.",
							Computed:    true,
						},
						"dscp": {
							Type:        schema.TypeInt,
							Description: "Advertised Differentiated Services Code Point (DSCP) value, a packet header value indicating the level of service requested for traffic, such as high priority or best effort delivery.",
							Computed:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "Policy type name.",
							Computed:    true,
						},
						"priority": {
							Type:        schema.TypeInt,
							Description: "Advertised Layer 2 priority (0 - 7; from lowest to highest priority).",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Enable or disable this TLV.",
							Computed:    true,
						},
						"vlan_intf": {
							Type:        schema.TypeString,
							Description: "VLAN interface to advertise; if configured on port.",
							Computed:    true,
						},
					},
				},
			},
			"med_tlvs": {
				Type:        schema.TypeString,
				Description: "Transmitted LLDP-MED TLVs (type-length-value descriptions).",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Profile name.",
				Required:    true,
			},
		},
	}
}

func dataSourceSwitchControllerLldpProfileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	mkey := d.Id()

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

	o, err := c.Cmdb.ReadSwitchControllerLldpProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SwitchControllerLldpProfile dataSource: %v", err)
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

	diags := refreshObjectSwitchControllerLldpProfile(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}
