// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: RIP interface configuration.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceRouterRipInterface() *schema.Resource {
	return &schema.Resource{
		Description: "RIP interface configuration.",

		ReadContext: dataSourceRouterRipInterfaceRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"auth_keychain": {
				Type:        schema.TypeString,
				Description: "Authentication key-chain name.",
				Computed:    true,
			},
			"auth_mode": {
				Type:        schema.TypeString,
				Description: "Authentication mode.",
				Computed:    true,
			},
			"auth_string": {
				Type:        schema.TypeString,
				Description: "Authentication string/password.",
				Computed:    true,
				Sensitive:   true,
			},
			"flags": {
				Type:        schema.TypeInt,
				Description: "Flags.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Interface name.",
				Required:    true,
			},
			"receive_version": {
				Type:        schema.TypeString,
				Description: "Receive version.",
				Computed:    true,
			},
			"send_version": {
				Type:        schema.TypeString,
				Description: "Send version.",
				Computed:    true,
			},
			"send_version2_broadcast": {
				Type:        schema.TypeString,
				Description: "Enable/disable broadcast version 1 compatible packets.",
				Computed:    true,
			},
			"split_horizon": {
				Type:        schema.TypeString,
				Description: "Enable/disable split horizon.",
				Computed:    true,
			},
			"split_horizon_status": {
				Type:        schema.TypeString,
				Description: "Enable/disable split horizon.",
				Computed:    true,
			},
		},
	}
}

func dataSourceRouterRipInterfaceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadRouterRipInterface(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading RouterRipInterface dataSource: %v", err)
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

	diags := refreshObjectRouterRipInterface(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
