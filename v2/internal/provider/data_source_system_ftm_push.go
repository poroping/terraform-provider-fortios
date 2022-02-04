// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure FortiToken Mobile push services.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceSystemFtmPush() *schema.Resource {
	return &schema.Resource{
		Description: "Configure FortiToken Mobile push services.",

		ReadContext: dataSourceSystemFtmPushRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"server": {
				Type:        schema.TypeString,
				Description: "IPv4 address or domain name of FortiToken Mobile push services server.",
				Computed:    true,
			},
			"server_cert": {
				Type:        schema.TypeString,
				Description: "Name of the server certificate to be used for SSL (default = Fortinet_Factory).",
				Computed:    true,
			},
			"server_ip": {
				Type:        schema.TypeString,
				Description: "IPv4 address of FortiToken Mobile push services server (format: xxx.xxx.xxx.xxx).",
				Computed:    true,
			},
			"server_port": {
				Type:        schema.TypeInt,
				Description: "Port to communicate with FortiToken Mobile push services server (1 - 65535, default = 4433).",
				Computed:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Enable/disable the use of FortiToken Mobile push services.",
				Computed:    true,
			},
		},
	}
}

func dataSourceSystemFtmPushRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "SystemFtmPush"

	o, err := c.Cmdb.ReadSystemFtmPush(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemFtmPush dataSource: %v", err)
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

	diags := refreshObjectSystemFtmPush(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
