// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure Fortinet Single Sign On (FSSO) agents.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceUserFsso() *schema.Resource {
	return &schema.Resource{
		Description: "Configure Fortinet Single Sign On (FSSO) agents.",

		ReadContext: dataSourceUserFssoRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"group_poll_interval": {
				Type:        schema.TypeInt,
				Description: "Interval in minutes within to fetch groups from FSSO server, or unset to disable.",
				Computed:    true,
			},
			"interface": {
				Type:        schema.TypeString,
				Description: "Specify outgoing interface to reach server.",
				Computed:    true,
			},
			"interface_select_method": {
				Type:        schema.TypeString,
				Description: "Specify how to select outgoing interface to reach server.",
				Computed:    true,
			},
			"ldap_poll": {
				Type:        schema.TypeString,
				Description: "Enable/disable automatic fetching of groups from LDAP server.",
				Computed:    true,
			},
			"ldap_poll_filter": {
				Type:        schema.TypeString,
				Description: "Filter used to fetch groups.",
				Computed:    true,
			},
			"ldap_poll_interval": {
				Type:        schema.TypeInt,
				Description: "Interval in minutes within to fetch groups from LDAP server.",
				Computed:    true,
			},
			"ldap_server": {
				Type:        schema.TypeString,
				Description: "LDAP server to get group information.",
				Computed:    true,
			},
			"logon_timeout": {
				Type:        schema.TypeInt,
				Description: "Interval in minutes to keep logons after FSSO server down.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Name.",
				Required:    true,
			},
			"password": {
				Type:        schema.TypeString,
				Description: "Password of the first FSSO collector agent.",
				Computed:    true,
				Sensitive:   true,
			},
			"password2": {
				Type:        schema.TypeString,
				Description: "Password of the second FSSO collector agent.",
				Computed:    true,
				Sensitive:   true,
			},
			"password3": {
				Type:        schema.TypeString,
				Description: "Password of the third FSSO collector agent.",
				Computed:    true,
				Sensitive:   true,
			},
			"password4": {
				Type:        schema.TypeString,
				Description: "Password of the fourth FSSO collector agent.",
				Computed:    true,
				Sensitive:   true,
			},
			"password5": {
				Type:        schema.TypeString,
				Description: "Password of the fifth FSSO collector agent.",
				Computed:    true,
				Sensitive:   true,
			},
			"port": {
				Type:        schema.TypeInt,
				Description: "Port of the first FSSO collector agent.",
				Computed:    true,
			},
			"port2": {
				Type:        schema.TypeInt,
				Description: "Port of the second FSSO collector agent.",
				Computed:    true,
			},
			"port3": {
				Type:        schema.TypeInt,
				Description: "Port of the third FSSO collector agent.",
				Computed:    true,
			},
			"port4": {
				Type:        schema.TypeInt,
				Description: "Port of the fourth FSSO collector agent.",
				Computed:    true,
			},
			"port5": {
				Type:        schema.TypeInt,
				Description: "Port of the fifth FSSO collector agent.",
				Computed:    true,
			},
			"server": {
				Type:        schema.TypeString,
				Description: "Domain name or IP address of the first FSSO collector agent.",
				Computed:    true,
			},
			"server2": {
				Type:        schema.TypeString,
				Description: "Domain name or IP address of the second FSSO collector agent.",
				Computed:    true,
			},
			"server3": {
				Type:        schema.TypeString,
				Description: "Domain name or IP address of the third FSSO collector agent.",
				Computed:    true,
			},
			"server4": {
				Type:        schema.TypeString,
				Description: "Domain name or IP address of the fourth FSSO collector agent.",
				Computed:    true,
			},
			"server5": {
				Type:        schema.TypeString,
				Description: "Domain name or IP address of the fifth FSSO collector agent.",
				Computed:    true,
			},
			"sni": {
				Type:        schema.TypeString,
				Description: "Server Name Indication.",
				Computed:    true,
			},
			"source_ip": {
				Type:        schema.TypeString,
				Description: "Source IP for communications to FSSO agent.",
				Computed:    true,
			},
			"source_ip6": {
				Type:        schema.TypeString,
				Description: "IPv6 source for communications to FSSO agent.",
				Computed:    true,
			},
			"ssl": {
				Type:        schema.TypeString,
				Description: "Enable/disable use of SSL.",
				Computed:    true,
			},
			"ssl_server_host_ip_check": {
				Type:        schema.TypeString,
				Description: "Enable/disable server host/IP verification.",
				Computed:    true,
			},
			"ssl_trusted_cert": {
				Type:        schema.TypeString,
				Description: "Trusted server certificate or CA certificate.",
				Computed:    true,
			},
			"type": {
				Type:        schema.TypeString,
				Description: "Server type.",
				Computed:    true,
			},
			"user_info_server": {
				Type:        schema.TypeString,
				Description: "LDAP server to get user information.",
				Computed:    true,
			},
		},
	}
}

func dataSourceUserFssoRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadUserFsso(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading UserFsso dataSource: %v", err)
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

	diags := refreshObjectUserFsso(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
