// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure IPv4 and IPv6 central SNAT policies.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceFirewallCentralSnatMap() *schema.Resource {
	return &schema.Resource{
		Description: "Configure IPv4 and IPv6 central SNAT policies.",

		ReadContext: dataSourceFirewallCentralSnatMapRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"comments": {
				Type:        schema.TypeString,
				Description: "Comment.",
				Computed:    true,
			},
			"dst_addr": {
				Type:        schema.TypeList,
				Description: "IPv4 Destination address.",
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
			"dst_addr6": {
				Type:        schema.TypeList,
				Description: "IPv6 Destination address.",
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
				Type:        schema.TypeList,
				Description: "Destination interface name from available interfaces.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Interface name.",
							Computed:    true,
						},
					},
				},
			},
			"nat": {
				Type:        schema.TypeString,
				Description: "Enable/disable source NAT.",
				Computed:    true,
			},
			"nat_ippool": {
				Type:        schema.TypeList,
				Description: "Name of the IP pools to be used to translate addresses from available IP Pools.",
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
			"nat_ippool6": {
				Type:        schema.TypeList,
				Description: "IPv6 pools to be used for source NAT.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "IPv6 pool name.",
							Computed:    true,
						},
					},
				},
			},
			"nat_port": {
				Type:        schema.TypeString,
				Description: "Translated port or port range (1 to 65535, 0 means any port).",
				Computed:    true,
			},
			"nat46": {
				Type:        schema.TypeString,
				Description: "Enable/disable NAT46.",
				Computed:    true,
			},
			"nat64": {
				Type:        schema.TypeString,
				Description: "Enable/disable NAT64.",
				Computed:    true,
			},
			"orig_addr": {
				Type:        schema.TypeList,
				Description: "IPv4 Original address.",
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
			"orig_addr6": {
				Type:        schema.TypeList,
				Description: "IPv6 Original address.",
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
			"orig_port": {
				Type:        schema.TypeString,
				Description: "Original TCP port (1 to 65535, 0 means any port).",
				Computed:    true,
			},
			"policyid": {
				Type:        schema.TypeInt,
				Description: "Policy ID.",
				Required:    true,
			},
			"protocol": {
				Type:        schema.TypeInt,
				Description: "Integer value for the protocol type (0 - 255).",
				Computed:    true,
			},
			"srcintf": {
				Type:        schema.TypeList,
				Description: "Source interface name from available interfaces.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Interface name.",
							Computed:    true,
						},
					},
				},
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Enable/disable the active status of this policy.",
				Computed:    true,
			},
			"type": {
				Type:        schema.TypeString,
				Description: "IPv4/IPv6 source NAT.",
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

func dataSourceFirewallCentralSnatMapRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallCentralSnatMap(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallCentralSnatMap dataSource: %v", err)
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

	diags := refreshObjectFirewallCentralSnatMap(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}
