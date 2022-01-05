// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure custom services.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceFirewallServiceCustom() *schema.Resource {
	return &schema.Resource{
		Description: "Configure custom services.",

		ReadContext: dataSourceFirewallServiceCustomRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"app_category": {
				Type:        schema.TypeList,
				Description: "Application category ID.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Description: "Application category id.",
							Computed:    true,
						},
					},
				},
			},
			"app_service_type": {
				Type:        schema.TypeString,
				Description: "Application service type.",
				Computed:    true,
			},
			"application": {
				Type:        schema.TypeList,
				Description: "Application ID.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Description: "Application id.",
							Computed:    true,
						},
					},
				},
			},
			"category": {
				Type:        schema.TypeString,
				Description: "Service category.",
				Computed:    true,
			},
			"check_reset_range": {
				Type:        schema.TypeString,
				Description: "Configure the type of ICMP error message verification.",
				Computed:    true,
			},
			"color": {
				Type:        schema.TypeInt,
				Description: "Color of icon on the GUI.",
				Computed:    true,
			},
			"comment": {
				Type:        schema.TypeString,
				Description: "Comment.",
				Computed:    true,
			},
			"fabric_object": {
				Type:        schema.TypeString,
				Description: "Security Fabric global object setting.",
				Computed:    true,
			},
			"fqdn": {
				Type:        schema.TypeString,
				Description: "Fully qualified domain name.",
				Computed:    true,
			},
			"helper": {
				Type:        schema.TypeString,
				Description: "Helper name.",
				Computed:    true,
			},
			"icmpcode": {
				Type:        schema.TypeInt,
				Description: "ICMP code.",
				Computed:    true,
			},
			"icmptype": {
				Type:        schema.TypeInt,
				Description: "ICMP type.",
				Computed:    true,
			},
			"iprange": {
				Type:        schema.TypeString,
				Description: "Start and end of the IP range associated with service.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Custom service name.",
				Required:    true,
			},
			"protocol": {
				Type:        schema.TypeString,
				Description: "Protocol type based on IANA numbers.",
				Computed:    true,
			},
			"protocol_number": {
				Type:        schema.TypeInt,
				Description: "IP protocol number.",
				Computed:    true,
			},
			"proxy": {
				Type:        schema.TypeString,
				Description: "Enable/disable web proxy service.",
				Computed:    true,
			},
			"sctp_portrange": {
				Type:        schema.TypeString,
				Description: "Multiple SCTP port ranges.",
				Computed:    true,
			},
			"session_ttl": {
				Type:        schema.TypeString,
				Description: "Session TTL (300 - 2764800, 0 = default).",
				Computed:    true,
			},
			"tcp_halfclose_timer": {
				Type:        schema.TypeInt,
				Description: "Wait time to close a TCP session waiting for an unanswered FIN packet (1 - 86400 sec, 0 = default).",
				Computed:    true,
			},
			"tcp_halfopen_timer": {
				Type:        schema.TypeInt,
				Description: "Wait time to close a TCP session waiting for an unanswered open session packet (1 - 86400 sec, 0 = default).",
				Computed:    true,
			},
			"tcp_portrange": {
				Type:        schema.TypeString,
				Description: "Multiple TCP port ranges.",
				Computed:    true,
			},
			"tcp_rst_timer": {
				Type:        schema.TypeInt,
				Description: "Set the length of the TCP CLOSE state in seconds (5 - 300 sec, 0 = default).",
				Computed:    true,
			},
			"tcp_timewait_timer": {
				Type:        schema.TypeInt,
				Description: "Set the length of the TCP TIME-WAIT state in seconds (1 - 300 sec, 0 = default).",
				Computed:    true,
			},
			"udp_idle_timer": {
				Type:        schema.TypeInt,
				Description: "UDP half close timeout (0 - 86400 sec, 0 = default).",
				Computed:    true,
			},
			"udp_portrange": {
				Type:        schema.TypeString,
				Description: "Multiple UDP port ranges.",
				Computed:    true,
			},
			"visibility": {
				Type:        schema.TypeString,
				Description: "Enable/disable the visibility of the service on the GUI.",
				Computed:    true,
			},
		},
	}
}

func dataSourceFirewallServiceCustomRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallServiceCustom(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallServiceCustom dataSource: %v", err)
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

	diags := refreshObjectFirewallServiceCustom(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}
