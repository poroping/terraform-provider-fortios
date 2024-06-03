// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure LLDP network policy.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceSystemLldpNetworkPolicy() *schema.Resource {
	return &schema.Resource{
		Description: "Configure LLDP network policy.",

		ReadContext: dataSourceSystemLldpNetworkPolicyRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"comment": {
				Type:        schema.TypeString,
				Description: "Comment.",
				Computed:    true,
			},
			"guest": {
				Type:        schema.TypeList,
				Description: "Guest.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dscp": {
							Type:        schema.TypeInt,
							Description: "Differentiated Services Code Point (DSCP) value to advertise.",
							Computed:    true,
						},
						"priority": {
							Type:        schema.TypeInt,
							Description: "802.1P CoS/PCP to advertise (0 - 7; from lowest to highest priority).",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Enable/disable advertising this policy.",
							Computed:    true,
						},
						"tag": {
							Type:        schema.TypeString,
							Description: "Advertise tagged or untagged traffic.",
							Computed:    true,
						},
						"vlan": {
							Type:        schema.TypeInt,
							Description: "802.1Q VLAN ID to advertise (1 - 4094).",
							Computed:    true,
						},
					},
				},
			},
			"guest_voice_signaling": {
				Type:        schema.TypeList,
				Description: "Guest Voice Signaling.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dscp": {
							Type:        schema.TypeInt,
							Description: "Differentiated Services Code Point (DSCP) value to advertise.",
							Computed:    true,
						},
						"priority": {
							Type:        schema.TypeInt,
							Description: "802.1P CoS/PCP to advertise (0 - 7; from lowest to highest priority).",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Enable/disable advertising this policy.",
							Computed:    true,
						},
						"tag": {
							Type:        schema.TypeString,
							Description: "Advertise tagged or untagged traffic.",
							Computed:    true,
						},
						"vlan": {
							Type:        schema.TypeInt,
							Description: "802.1Q VLAN ID to advertise (1 - 4094).",
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:        schema.TypeString,
				Description: "LLDP network policy name.",
				Required:    true,
			},
			"softphone": {
				Type:        schema.TypeList,
				Description: "Softphone.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dscp": {
							Type:        schema.TypeInt,
							Description: "Differentiated Services Code Point (DSCP) value to advertise.",
							Computed:    true,
						},
						"priority": {
							Type:        schema.TypeInt,
							Description: "802.1P CoS/PCP to advertise (0 - 7; from lowest to highest priority).",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Enable/disable advertising this policy.",
							Computed:    true,
						},
						"tag": {
							Type:        schema.TypeString,
							Description: "Advertise tagged or untagged traffic.",
							Computed:    true,
						},
						"vlan": {
							Type:        schema.TypeInt,
							Description: "802.1Q VLAN ID to advertise (1 - 4094).",
							Computed:    true,
						},
					},
				},
			},
			"streaming_video": {
				Type:        schema.TypeList,
				Description: "Streaming Video.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dscp": {
							Type:        schema.TypeInt,
							Description: "Differentiated Services Code Point (DSCP) value to advertise.",
							Computed:    true,
						},
						"priority": {
							Type:        schema.TypeInt,
							Description: "802.1P CoS/PCP to advertise (0 - 7; from lowest to highest priority).",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Enable/disable advertising this policy.",
							Computed:    true,
						},
						"tag": {
							Type:        schema.TypeString,
							Description: "Advertise tagged or untagged traffic.",
							Computed:    true,
						},
						"vlan": {
							Type:        schema.TypeInt,
							Description: "802.1Q VLAN ID to advertise (1 - 4094).",
							Computed:    true,
						},
					},
				},
			},
			"video_conferencing": {
				Type:        schema.TypeList,
				Description: "Video Conferencing.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dscp": {
							Type:        schema.TypeInt,
							Description: "Differentiated Services Code Point (DSCP) value to advertise.",
							Computed:    true,
						},
						"priority": {
							Type:        schema.TypeInt,
							Description: "802.1P CoS/PCP to advertise (0 - 7; from lowest to highest priority).",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Enable/disable advertising this policy.",
							Computed:    true,
						},
						"tag": {
							Type:        schema.TypeString,
							Description: "Advertise tagged or untagged traffic.",
							Computed:    true,
						},
						"vlan": {
							Type:        schema.TypeInt,
							Description: "802.1Q VLAN ID to advertise (1 - 4094).",
							Computed:    true,
						},
					},
				},
			},
			"video_signaling": {
				Type:        schema.TypeList,
				Description: "Video Signaling.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dscp": {
							Type:        schema.TypeInt,
							Description: "Differentiated Services Code Point (DSCP) value to advertise.",
							Computed:    true,
						},
						"priority": {
							Type:        schema.TypeInt,
							Description: "802.1P CoS/PCP to advertise (0 - 7; from lowest to highest priority).",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Enable/disable advertising this policy.",
							Computed:    true,
						},
						"tag": {
							Type:        schema.TypeString,
							Description: "Advertise tagged or untagged traffic.",
							Computed:    true,
						},
						"vlan": {
							Type:        schema.TypeInt,
							Description: "802.1Q VLAN ID to advertise (1 - 4094).",
							Computed:    true,
						},
					},
				},
			},
			"voice": {
				Type:        schema.TypeList,
				Description: "Voice.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dscp": {
							Type:        schema.TypeInt,
							Description: "Differentiated Services Code Point (DSCP) value to advertise.",
							Computed:    true,
						},
						"priority": {
							Type:        schema.TypeInt,
							Description: "802.1P CoS/PCP to advertise (0 - 7; from lowest to highest priority).",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Enable/disable advertising this policy.",
							Computed:    true,
						},
						"tag": {
							Type:        schema.TypeString,
							Description: "Advertise tagged or untagged traffic.",
							Computed:    true,
						},
						"vlan": {
							Type:        schema.TypeInt,
							Description: "802.1Q VLAN ID to advertise (1 - 4094).",
							Computed:    true,
						},
					},
				},
			},
			"voice_signaling": {
				Type:        schema.TypeList,
				Description: "Voice signaling.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dscp": {
							Type:        schema.TypeInt,
							Description: "Differentiated Services Code Point (DSCP) value to advertise.",
							Computed:    true,
						},
						"priority": {
							Type:        schema.TypeInt,
							Description: "802.1P CoS/PCP to advertise (0 - 7; from lowest to highest priority).",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Enable/disable advertising this policy.",
							Computed:    true,
						},
						"tag": {
							Type:        schema.TypeString,
							Description: "Advertise tagged or untagged traffic.",
							Computed:    true,
						},
						"vlan": {
							Type:        schema.TypeInt,
							Description: "802.1Q VLAN ID to advertise (1 - 4094).",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceSystemLldpNetworkPolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemLldpNetworkPolicy(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemLldpNetworkPolicy dataSource: %v", err)
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

	diags := refreshObjectSystemLldpNetworkPolicy(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
