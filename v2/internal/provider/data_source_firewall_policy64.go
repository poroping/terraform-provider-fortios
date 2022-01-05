// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure IPv6 to IPv4 policies.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceFirewallPolicy64() *schema.Resource {
	return &schema.Resource{
		Description: "Configure IPv6 to IPv4 policies.",

		ReadContext: dataSourceFirewallPolicy64Read,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"action": {
				Type:        schema.TypeString,
				Description: "Policy action.",
				Computed:    true,
			},
			"comments": {
				Type:        schema.TypeString,
				Description: "Comment.",
				Computed:    true,
			},
			"dstaddr": {
				Type:        schema.TypeList,
				Description: "Destination address name.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Address name.",
							Computed:    true,
						},
					},
				},
			},
			"dstintf": {
				Type:        schema.TypeString,
				Description: "Destination interface name.",
				Computed:    true,
			},
			"fixedport": {
				Type:        schema.TypeString,
				Description: "Enable/disable policy fixed port.",
				Computed:    true,
			},
			"ippool": {
				Type:        schema.TypeString,
				Description: "Enable/disable policy64 IP pool.",
				Computed:    true,
			},
			"logtraffic": {
				Type:        schema.TypeString,
				Description: "Enable/disable policy log traffic.",
				Computed:    true,
			},
			"logtraffic_start": {
				Type:        schema.TypeString,
				Description: "Record logs when a session starts and ends.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Policy name.",
				Computed:    true,
			},
			"per_ip_shaper": {
				Type:        schema.TypeString,
				Description: "Per-IP traffic shaper.",
				Computed:    true,
			},
			"permit_any_host": {
				Type:        schema.TypeString,
				Description: "Enable/disable permit any host in.",
				Computed:    true,
			},
			"policyid": {
				Type:        schema.TypeInt,
				Description: "Policy ID (0 - 4294967294).",
				Required:    true,
			},
			"poolname": {
				Type:        schema.TypeList,
				Description: "Policy IP pool names.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "IP pool name.",
							Computed:    true,
						},
					},
				},
			},
			"schedule": {
				Type:        schema.TypeString,
				Description: "Schedule name.",
				Computed:    true,
			},
			"service": {
				Type:        schema.TypeList,
				Description: "Service name.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Address name.",
							Computed:    true,
						},
					},
				},
			},
			"srcaddr": {
				Type:        schema.TypeList,
				Description: "Source address name.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Address name.",
							Computed:    true,
						},
					},
				},
			},
			"srcintf": {
				Type:        schema.TypeString,
				Description: "Source interface name.",
				Computed:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Enable/disable policy status.",
				Computed:    true,
			},
			"tcp_mss_receiver": {
				Type:        schema.TypeInt,
				Description: "TCP MSS value of receiver.",
				Computed:    true,
			},
			"tcp_mss_sender": {
				Type:        schema.TypeInt,
				Description: "TCP MSS value of sender.",
				Computed:    true,
			},
			"traffic_shaper": {
				Type:        schema.TypeString,
				Description: "Traffic shaper.",
				Computed:    true,
			},
			"traffic_shaper_reverse": {
				Type:        schema.TypeString,
				Description: "Reverse traffic shaper.",
				Computed:    true,
			},
			"uuid": {
				Type:        schema.TypeString,
				Description: "Universally Unique Identifier (UUID; automatically assigned but can be manually reset).",
				Computed:    true,
			},
		},
	}
}

func dataSourceFirewallPolicy64Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallPolicy64(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallPolicy64 dataSource: %v", err)
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

	diags := refreshObjectFirewallPolicy64(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}
