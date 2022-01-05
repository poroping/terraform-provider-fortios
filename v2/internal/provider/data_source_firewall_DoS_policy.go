// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure IPv4 DoS policies.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceFirewallDoSPolicy() *schema.Resource {
	return &schema.Resource{
		Description: "Configure IPv4 DoS policies.",

		ReadContext: dataSourceFirewallDoSPolicyRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"anomaly": {
				Type:        schema.TypeList,
				Description: "Anomaly name.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:        schema.TypeString,
							Description: "Action taken when the threshold is reached.",
							Computed:    true,
						},
						"log": {
							Type:        schema.TypeString,
							Description: "Enable/disable anomaly logging.",
							Computed:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "Anomaly name.",
							Computed:    true,
						},
						"quarantine": {
							Type:        schema.TypeString,
							Description: "Quarantine method.",
							Computed:    true,
						},
						"quarantine_expiry": {
							Type:        schema.TypeString,
							Description: "Duration of quarantine. (Format ###d##h##m, minimum 1m, maximum 364d23h59m, default = 5m). Requires quarantine set to attacker.",
							Computed:    true,
						},
						"quarantine_log": {
							Type:        schema.TypeString,
							Description: "Enable/disable quarantine logging.",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Enable/disable this anomaly.",
							Computed:    true,
						},
						"threshold": {
							Type:        schema.TypeInt,
							Description: "Anomaly threshold. Number of detected instances per minute that triggers the anomaly action.",
							Computed:    true,
						},
						"thresholddefault": {
							Type:        schema.TypeInt,
							Description: "Number of detected instances per minute which triggers action (1 - 2147483647, default = 1000). Note that each anomaly has a different threshold value assigned to it.",
							Computed:    true,
						},
					},
				},
			},
			"comments": {
				Type:        schema.TypeString,
				Description: "Comment.",
				Computed:    true,
			},
			"dstaddr": {
				Type:        schema.TypeList,
				Description: "Destination address name from available addresses.",
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
			"interface": {
				Type:        schema.TypeString,
				Description: "Incoming interface name from available interfaces.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Policy name.",
				Computed:    true,
			},
			"policyid": {
				Type:        schema.TypeInt,
				Description: "Policy ID.",
				Required:    true,
			},
			"service": {
				Type:        schema.TypeList,
				Description: "Service object from available options.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Service name.",
							Computed:    true,
						},
					},
				},
			},
			"srcaddr": {
				Type:        schema.TypeList,
				Description: "Source address name from available addresses.",
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
		},
	}
}

func dataSourceFirewallDoSPolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallDoSPolicy(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallDoSPolicy dataSource: %v", err)
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

	diags := refreshObjectFirewallDoSPolicy(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}
