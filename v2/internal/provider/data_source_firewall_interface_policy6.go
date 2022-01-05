// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure IPv6 interface policies.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceFirewallInterfacePolicy6() *schema.Resource {
	return &schema.Resource{
		Description: "Configure IPv6 interface policies.",

		ReadContext: dataSourceFirewallInterfacePolicy6Read,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"application_list": {
				Type:        schema.TypeString,
				Description: "Application list name.",
				Computed:    true,
			},
			"application_list_status": {
				Type:        schema.TypeString,
				Description: "Enable/disable application control.",
				Computed:    true,
			},
			"av_profile": {
				Type:        schema.TypeString,
				Description: "Antivirus profile.",
				Computed:    true,
			},
			"av_profile_status": {
				Type:        schema.TypeString,
				Description: "Enable/disable antivirus.",
				Computed:    true,
			},
			"comments": {
				Type:        schema.TypeString,
				Description: "Comments.",
				Computed:    true,
			},
			"dlp_sensor": {
				Type:        schema.TypeString,
				Description: "DLP sensor name.",
				Computed:    true,
			},
			"dlp_sensor_status": {
				Type:        schema.TypeString,
				Description: "Enable/disable DLP.",
				Computed:    true,
			},
			"dsri": {
				Type:        schema.TypeString,
				Description: "Enable/disable DSRI.",
				Computed:    true,
			},
			"dstaddr6": {
				Type:        schema.TypeList,
				Description: "IPv6 address object to limit traffic monitoring to network traffic sent to the specified address or range.",
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
			"emailfilter_profile": {
				Type:        schema.TypeString,
				Description: "Email filter profile.",
				Computed:    true,
			},
			"emailfilter_profile_status": {
				Type:        schema.TypeString,
				Description: "Enable/disable email filter.",
				Computed:    true,
			},
			"interface": {
				Type:        schema.TypeString,
				Description: "Monitored interface name from available interfaces.",
				Computed:    true,
			},
			"ips_sensor": {
				Type:        schema.TypeString,
				Description: "IPS sensor name.",
				Computed:    true,
			},
			"ips_sensor_status": {
				Type:        schema.TypeString,
				Description: "Enable/disable IPS.",
				Computed:    true,
			},
			"logtraffic": {
				Type:        schema.TypeString,
				Description: "Logging type to be used in this policy (Options: all | utm | disable, Default: utm).",
				Computed:    true,
			},
			"policyid": {
				Type:        schema.TypeInt,
				Description: "Policy ID (0 - 4294967295).",
				Required:    true,
			},
			"service6": {
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
			"srcaddr6": {
				Type:        schema.TypeList,
				Description: "IPv6 address object to limit traffic monitoring to network traffic sent from the specified address or range.",
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
			"status": {
				Type:        schema.TypeString,
				Description: "Enable/disable this policy.",
				Computed:    true,
			},
			"webfilter_profile": {
				Type:        schema.TypeString,
				Description: "Web filter profile.",
				Computed:    true,
			},
			"webfilter_profile_status": {
				Type:        schema.TypeString,
				Description: "Enable/disable web filtering.",
				Computed:    true,
			},
		},
	}
}

func dataSourceFirewallInterfacePolicy6Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallInterfacePolicy6(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallInterfacePolicy6 dataSource: %v", err)
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

	diags := refreshObjectFirewallInterfacePolicy6(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}
