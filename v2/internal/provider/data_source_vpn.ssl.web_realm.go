// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Realm.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceVpnSslWebRealm() *schema.Resource {
	return &schema.Resource{
		Description: "Realm.",

		ReadContext: dataSourceVpnSslWebRealmRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"login_page": {
				Type:        schema.TypeString,
				Description: "Replacement HTML for SSL-VPN login page.",
				Computed:    true,
			},
			"max_concurrent_user": {
				Type:        schema.TypeInt,
				Description: "Maximum concurrent users (0 - 65535, 0 means unlimited).",
				Computed:    true,
			},
			"nas_ip": {
				Type:        schema.TypeString,
				Description: "IP address used as a NAS-IP to communicate with the RADIUS server.",
				Computed:    true,
			},
			"radius_port": {
				Type:        schema.TypeInt,
				Description: "RADIUS service port number (0 - 65535, 0 means user.radius.radius-port).",
				Computed:    true,
			},
			"radius_server": {
				Type:        schema.TypeString,
				Description: "RADIUS server associated with realm.",
				Computed:    true,
			},
			"url_path": {
				Type:        schema.TypeString,
				Description: "URL path to access SSL-VPN login page.",
				Required:    true,
			},
			"virtual_host": {
				Type:        schema.TypeString,
				Description: "Virtual host name for realm.",
				Computed:    true,
			},
			"virtual_host_only": {
				Type:        schema.TypeString,
				Description: "Enable/disable enforcement of virtual host method for SSL-VPN client access.",
				Computed:    true,
			},
			"virtual_host_server_cert": {
				Type:        schema.TypeString,
				Description: "Name of the server certificate to used for this realm.",
				Computed:    true,
			},
		},
	}
}

func dataSourceVpnSslWebRealmRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	i := d.Get("url_path")
	mkey := utils.ParseMkey(i)

	o, err := c.Cmdb.ReadVpnSslWebRealm(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading VpnSslWebRealm dataSource: %v", err)
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

	diags := refreshObjectVpnSslWebRealm(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
