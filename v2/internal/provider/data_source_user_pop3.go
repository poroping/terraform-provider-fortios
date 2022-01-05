// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: POP3 server entry configuration.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceUserPop3() *schema.Resource {
	return &schema.Resource{
		Description: "POP3 server entry configuration.",

		ReadContext: dataSourceUserPop3Read,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "POP3 server entry name.",
				Required:    true,
			},
			"port": {
				Type:        schema.TypeInt,
				Description: "POP3 service port number.",
				Computed:    true,
			},
			"secure": {
				Type:        schema.TypeString,
				Description: "SSL connection.",
				Computed:    true,
			},
			"server": {
				Type:        schema.TypeString,
				Description: "{<name_str|ip_str>} server domain name or IP.",
				Computed:    true,
			},
			"ssl_min_proto_version": {
				Type:        schema.TypeString,
				Description: "Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).",
				Computed:    true,
			},
		},
	}
}

func dataSourceUserPop3Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadUserPop3(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading UserPop3 dataSource: %v", err)
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

	diags := refreshObjectUserPop3(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}
