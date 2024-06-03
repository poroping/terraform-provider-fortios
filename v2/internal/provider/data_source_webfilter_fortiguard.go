// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure FortiGuard Web Filter service.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceWebfilterFortiguard() *schema.Resource {
	return &schema.Resource{
		Description: "Configure FortiGuard Web Filter service.",

		ReadContext: dataSourceWebfilterFortiguardRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"cache_mem_percent": {
				Type:        schema.TypeInt,
				Description: "Maximum percentage of available memory allocated to caching (1 - 15).",
				Computed:    true,
			},
			"cache_mode": {
				Type:        schema.TypeString,
				Description: "Cache entry expiration mode.",
				Computed:    true,
			},
			"cache_prefix_match": {
				Type:        schema.TypeString,
				Description: "Enable/disable prefix matching in the cache.",
				Computed:    true,
			},
			"close_ports": {
				Type:        schema.TypeString,
				Description: "Close ports used for HTTP/HTTPS override authentication and disable user overrides.",
				Computed:    true,
			},
			"embed_image": {
				Type:        schema.TypeString,
				Description: "Enable/disable embedding images into replacement messages (default = enable).",
				Computed:    true,
			},
			"ovrd_auth_https": {
				Type:        schema.TypeString,
				Description: "Enable/disable use of HTTPS for override authentication.",
				Computed:    true,
			},
			"ovrd_auth_port_http": {
				Type:        schema.TypeInt,
				Description: "Port to use for FortiGuard Web Filter HTTP override authentication.",
				Computed:    true,
			},
			"ovrd_auth_port_https": {
				Type:        schema.TypeInt,
				Description: "Port to use for FortiGuard Web Filter HTTPS override authentication in proxy mode.",
				Computed:    true,
			},
			"ovrd_auth_port_https_flow": {
				Type:        schema.TypeInt,
				Description: "Port to use for FortiGuard Web Filter HTTPS override authentication in flow mode.",
				Computed:    true,
			},
			"ovrd_auth_port_warning": {
				Type:        schema.TypeInt,
				Description: "Port to use for FortiGuard Web Filter Warning override authentication.",
				Computed:    true,
			},
			"request_packet_size_limit": {
				Type:        schema.TypeInt,
				Description: "Limit size of URL request packets sent to FortiGuard server (0 for default).",
				Computed:    true,
			},
			"warn_auth_https": {
				Type:        schema.TypeString,
				Description: "Enable/disable use of HTTPS for warning and authentication.",
				Computed:    true,
			},
		},
	}
}

func dataSourceWebfilterFortiguardRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "WebfilterFortiguard"

	o, err := c.Cmdb.ReadWebfilterFortiguard(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WebfilterFortiguard dataSource: %v", err)
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

	diags := refreshObjectWebfilterFortiguard(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
