// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: IS-IS interface configuration.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceRouterIsisIsisInterface() *schema.Resource {
	return &schema.Resource{
		Description: "IS-IS interface configuration.",

		ReadContext: dataSourceRouterIsisIsisInterfaceRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
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
				Description: "IS-IS interface's circuit type",
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
				Required:    true,
			},
			"network_type": {
				Type:        schema.TypeString,
				Description: "IS-IS interface's network type",
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
	}
}

func dataSourceRouterIsisIsisInterfaceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadRouterIsisIsisInterface(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading RouterIsisIsisInterface dataSource: %v", err)
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

	diags := refreshObjectRouterIsisIsisInterface(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
