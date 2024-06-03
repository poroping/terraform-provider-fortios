// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure IS-IS.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceRouterIsis() *schema.Resource {
	return &schema.Resource{
		Description: "Configure IS-IS.",

		ReadContext: dataSourceRouterIsisRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"adjacency_check": {
				Type:        schema.TypeString,
				Description: "Enable/disable adjacency check.",
				Computed:    true,
			},
			"adjacency_check6": {
				Type:        schema.TypeString,
				Description: "Enable/disable IPv6 adjacency check.",
				Computed:    true,
			},
			"adv_passive_only": {
				Type:        schema.TypeString,
				Description: "Enable/disable IS-IS advertisement of passive interfaces only.",
				Computed:    true,
			},
			"adv_passive_only6": {
				Type:        schema.TypeString,
				Description: "Enable/disable IPv6 IS-IS advertisement of passive interfaces only.",
				Computed:    true,
			},
			"auth_keychain_l1": {
				Type:        schema.TypeString,
				Description: "Authentication key-chain for level 1 PDUs.",
				Computed:    true,
			},
			"auth_keychain_l2": {
				Type:        schema.TypeString,
				Description: "Authentication key-chain for level 2 PDUs.",
				Computed:    true,
			},
			"auth_mode_l1": {
				Type:        schema.TypeString,
				Description: "Level 1 authentication mode.",
				Computed:    true,
			},
			"auth_mode_l2": {
				Type:        schema.TypeString,
				Description: "Level 2 authentication mode.",
				Computed:    true,
			},
			"auth_password_l1": {
				Type:        schema.TypeString,
				Description: "Authentication password for level 1 PDUs.",
				Computed:    true,
				Sensitive:   true,
			},
			"auth_password_l2": {
				Type:        schema.TypeString,
				Description: "Authentication password for level 2 PDUs.",
				Computed:    true,
				Sensitive:   true,
			},
			"auth_sendonly_l1": {
				Type:        schema.TypeString,
				Description: "Enable/disable level 1 authentication send-only.",
				Computed:    true,
			},
			"auth_sendonly_l2": {
				Type:        schema.TypeString,
				Description: "Enable/disable level 2 authentication send-only.",
				Computed:    true,
			},
			"default_originate": {
				Type:        schema.TypeString,
				Description: "Enable/disable distribution of default route information.",
				Computed:    true,
			},
			"default_originate6": {
				Type:        schema.TypeString,
				Description: "Enable/disable distribution of default IPv6 route information.",
				Computed:    true,
			},
			"dynamic_hostname": {
				Type:        schema.TypeString,
				Description: "Enable/disable dynamic hostname.",
				Computed:    true,
			},
			"ignore_lsp_errors": {
				Type:        schema.TypeString,
				Description: "Enable/disable ignoring of LSP errors with bad checksums.",
				Computed:    true,
			},
			"is_type": {
				Type:        schema.TypeString,
				Description: "IS type.",
				Computed:    true,
			},
			"isis_interface": {
				Type:        schema.TypeList,
				Description: "IS-IS interface configuration.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auth_keychain_l1": {
							Type:        schema.TypeString,
							Description: "Authentication key-chain for level 1 PDUs.",
							Computed:    true,
						},
						"auth_keychain_l2": {
							Type:        schema.TypeString,
							Description: "Authentication key-chain for level 2 PDUs.",
							Computed:    true,
						},
						"auth_mode_l1": {
							Type:        schema.TypeString,
							Description: "Level 1 authentication mode.",
							Computed:    true,
						},
						"auth_mode_l2": {
							Type:        schema.TypeString,
							Description: "Level 2 authentication mode.",
							Computed:    true,
						},
						"auth_password_l1": {
							Type:        schema.TypeString,
							Description: "Authentication password for level 1 PDUs.",
							Computed:    true,
							Sensitive:   true,
						},
						"auth_password_l2": {
							Type:        schema.TypeString,
							Description: "Authentication password for level 2 PDUs.",
							Computed:    true,
							Sensitive:   true,
						},
						"auth_send_only_l1": {
							Type:        schema.TypeString,
							Description: "Enable/disable authentication send-only for level 1 PDUs.",
							Computed:    true,
						},
						"auth_send_only_l2": {
							Type:        schema.TypeString,
							Description: "Enable/disable authentication send-only for level 2 PDUs.",
							Computed:    true,
						},
						"circuit_type": {
							Type:        schema.TypeString,
							Description: "IS-IS interface's circuit type.",
							Computed:    true,
						},
						"csnp_interval_l1": {
							Type:        schema.TypeInt,
							Description: "Level 1 CSNP interval.",
							Computed:    true,
						},
						"csnp_interval_l2": {
							Type:        schema.TypeInt,
							Description: "Level 2 CSNP interval.",
							Computed:    true,
						},
						"hello_interval_l1": {
							Type:        schema.TypeInt,
							Description: "Level 1 hello interval.",
							Computed:    true,
						},
						"hello_interval_l2": {
							Type:        schema.TypeInt,
							Description: "Level 2 hello interval.",
							Computed:    true,
						},
						"hello_multiplier_l1": {
							Type:        schema.TypeInt,
							Description: "Level 1 multiplier for Hello holding time.",
							Computed:    true,
						},
						"hello_multiplier_l2": {
							Type:        schema.TypeInt,
							Description: "Level 2 multiplier for Hello holding time.",
							Computed:    true,
						},
						"hello_padding": {
							Type:        schema.TypeString,
							Description: "Enable/disable padding to IS-IS hello packets.",
							Computed:    true,
						},
						"lsp_interval": {
							Type:        schema.TypeInt,
							Description: "LSP transmission interval (milliseconds).",
							Computed:    true,
						},
						"lsp_retransmit_interval": {
							Type:        schema.TypeInt,
							Description: "LSP retransmission interval (sec).",
							Computed:    true,
						},
						"mesh_group": {
							Type:        schema.TypeString,
							Description: "Enable/disable IS-IS mesh group.",
							Computed:    true,
						},
						"mesh_group_id": {
							Type:        schema.TypeInt,
							Description: "Mesh group ID <0-4294967295>, 0: mesh-group blocked.",
							Computed:    true,
						},
						"metric_l1": {
							Type:        schema.TypeInt,
							Description: "Level 1 metric for interface.",
							Computed:    true,
						},
						"metric_l2": {
							Type:        schema.TypeInt,
							Description: "Level 2 metric for interface.",
							Computed:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "IS-IS interface name.",
							Computed:    true,
						},
						"network_type": {
							Type:        schema.TypeString,
							Description: "IS-IS interface's network type.",
							Computed:    true,
						},
						"priority_l1": {
							Type:        schema.TypeInt,
							Description: "Level 1 priority.",
							Computed:    true,
						},
						"priority_l2": {
							Type:        schema.TypeInt,
							Description: "Level 2 priority.",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Enable/disable interface for IS-IS.",
							Computed:    true,
						},
						"status6": {
							Type:        schema.TypeString,
							Description: "Enable/disable IPv6 interface for IS-IS.",
							Computed:    true,
						},
						"wide_metric_l1": {
							Type:        schema.TypeInt,
							Description: "Level 1 wide metric for interface.",
							Computed:    true,
						},
						"wide_metric_l2": {
							Type:        schema.TypeInt,
							Description: "Level 2 wide metric for interface.",
							Computed:    true,
						},
					},
				},
			},
			"isis_net": {
				Type:        schema.TypeList,
				Description: "IS-IS net configuration.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Description: "ISIS network ID.",
							Computed:    true,
						},
						"net": {
							Type:        schema.TypeString,
							Description: "IS-IS networks (format = xx.xxxx.  .xxxx.xx.).",
							Computed:    true,
						},
					},
				},
			},
			"lsp_gen_interval_l1": {
				Type:        schema.TypeInt,
				Description: "Minimum interval for level 1 LSP regenerating.",
				Computed:    true,
			},
			"lsp_gen_interval_l2": {
				Type:        schema.TypeInt,
				Description: "Minimum interval for level 2 LSP regenerating.",
				Computed:    true,
			},
			"lsp_refresh_interval": {
				Type:        schema.TypeInt,
				Description: "LSP refresh time in seconds.",
				Computed:    true,
			},
			"max_lsp_lifetime": {
				Type:        schema.TypeInt,
				Description: "Maximum LSP lifetime in seconds.",
				Computed:    true,
			},
			"metric_style": {
				Type:        schema.TypeString,
				Description: "Use old-style (ISO 10589) or new-style packet formats.",
				Computed:    true,
			},
			"overload_bit": {
				Type:        schema.TypeString,
				Description: "Enable/disable signal other routers not to use us in SPF.",
				Computed:    true,
			},
			"overload_bit_on_startup": {
				Type:        schema.TypeInt,
				Description: "Overload-bit only temporarily after reboot.",
				Computed:    true,
			},
			"overload_bit_suppress": {
				Type:        schema.TypeString,
				Description: "Suppress overload-bit for the specific prefixes.",
				Computed:    true,
			},
			"redistribute": {
				Type:        schema.TypeList,
				Description: "IS-IS redistribute protocols.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"level": {
							Type:        schema.TypeString,
							Description: "Level.",
							Computed:    true,
						},
						"metric": {
							Type:        schema.TypeInt,
							Description: "Metric.",
							Computed:    true,
						},
						"metric_type": {
							Type:        schema.TypeString,
							Description: "Metric type.",
							Computed:    true,
						},
						"protocol": {
							Type:        schema.TypeString,
							Description: "Protocol name.",
							Computed:    true,
						},
						"routemap": {
							Type:        schema.TypeString,
							Description: "Route map name.",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Status.",
							Computed:    true,
						},
					},
				},
			},
			"redistribute_l1": {
				Type:        schema.TypeString,
				Description: "Enable/disable redistribution of level 1 routes into level 2.",
				Computed:    true,
			},
			"redistribute_l1_list": {
				Type:        schema.TypeString,
				Description: "Access-list for route redistribution from l1 to l2.",
				Computed:    true,
			},
			"redistribute_l2": {
				Type:        schema.TypeString,
				Description: "Enable/disable redistribution of level 2 routes into level 1.",
				Computed:    true,
			},
			"redistribute_l2_list": {
				Type:        schema.TypeString,
				Description: "Access-list for route redistribution from l2 to l1.",
				Computed:    true,
			},
			"redistribute6": {
				Type:        schema.TypeList,
				Description: "IS-IS IPv6 redistribution for routing protocols.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"level": {
							Type:        schema.TypeString,
							Description: "Level.",
							Computed:    true,
						},
						"metric": {
							Type:        schema.TypeInt,
							Description: "Metric.",
							Computed:    true,
						},
						"metric_type": {
							Type:        schema.TypeString,
							Description: "Metric type.",
							Computed:    true,
						},
						"protocol": {
							Type:        schema.TypeString,
							Description: "Protocol name.",
							Computed:    true,
						},
						"routemap": {
							Type:        schema.TypeString,
							Description: "Route map name.",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Enable/disable redistribution.",
							Computed:    true,
						},
					},
				},
			},
			"redistribute6_l1": {
				Type:        schema.TypeString,
				Description: "Enable/disable redistribution of level 1 IPv6 routes into level 2.",
				Computed:    true,
			},
			"redistribute6_l1_list": {
				Type:        schema.TypeString,
				Description: "Access-list for IPv6 route redistribution from l1 to l2.",
				Computed:    true,
			},
			"redistribute6_l2": {
				Type:        schema.TypeString,
				Description: "Enable/disable redistribution of level 2 IPv6 routes into level 1.",
				Computed:    true,
			},
			"redistribute6_l2_list": {
				Type:        schema.TypeString,
				Description: "Access-list for IPv6 route redistribution from l2 to l1.",
				Computed:    true,
			},
			"spf_interval_exp_l1": {
				Type:        schema.TypeString,
				Description: "Level 1 SPF calculation delay.",
				Computed:    true,
			},
			"spf_interval_exp_l2": {
				Type:        schema.TypeString,
				Description: "Level 2 SPF calculation delay.",
				Computed:    true,
			},
			"summary_address": {
				Type:        schema.TypeList,
				Description: "IS-IS summary addresses.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Description: "Summary address entry ID.",
							Computed:    true,
						},
						"level": {
							Type:        schema.TypeString,
							Description: "Level.",
							Computed:    true,
						},
						"prefix": {
							Type:        schema.TypeString,
							Description: "Prefix.",
							Computed:    true,
						},
					},
				},
			},
			"summary_address6": {
				Type:        schema.TypeList,
				Description: "IS-IS IPv6 summary address.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Description: "Prefix entry ID.",
							Computed:    true,
						},
						"level": {
							Type:        schema.TypeString,
							Description: "Level.",
							Computed:    true,
						},
						"prefix6": {
							Type:        schema.TypeString,
							Description: "IPv6 prefix.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceRouterIsisRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "RouterIsis"

	o, err := c.Cmdb.ReadRouterIsis(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading RouterIsis dataSource: %v", err)
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

	diags := refreshObjectRouterIsis(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
