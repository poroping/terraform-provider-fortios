// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure Websense Integrated Services Protocol (WISP) servers.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceWebProxyWisp() *schema.Resource {
	return &schema.Resource{
		Description: "Configure Websense Integrated Services Protocol (WISP) servers.",

		ReadContext: dataSourceWebProxyWispRead,

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
			"max_connections": {
				Type:        schema.TypeInt,
				Description: "Maximum number of web proxy WISP connections (4 - 4096, default = 64).",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Server name.",
				Required:    true,
			},
			"outgoing_ip": {
				Type:        schema.TypeString,
				Description: "WISP outgoing IP address.",
				Computed:    true,
			},
			"server_ip": {
				Type:        schema.TypeString,
				Description: "WISP server IP address.",
				Computed:    true,
			},
			"server_port": {
				Type:        schema.TypeInt,
				Description: "WISP server port (1 - 65535, default = 15868).",
				Computed:    true,
			},
			"timeout": {
				Type:        schema.TypeInt,
				Description: "Period of time before WISP requests time out (1 - 15 sec, default = 5).",
				Computed:    true,
			},
		},
	}
}

func dataSourceWebProxyWispRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadWebProxyWisp(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WebProxyWisp dataSource: %v", err)
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

	diags := refreshObjectWebProxyWisp(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
