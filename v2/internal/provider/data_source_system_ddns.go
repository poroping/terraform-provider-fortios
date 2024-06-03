// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure DDNS.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceSystemDdns() *schema.Resource {
	return &schema.Resource{
		Description: "Configure DDNS.",

		ReadContext: dataSourceSystemDdnsRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"addr_type": {
				Type:        schema.TypeString,
				Description: "Address type of interface address in DDNS update.",
				Computed:    true,
			},
			"bound_ip": {
				Type:        schema.TypeString,
				Description: "Bound IP address.",
				Computed:    true,
			},
			"clear_text": {
				Type:        schema.TypeString,
				Description: "Enable/disable use of clear text connections.",
				Computed:    true,
			},
			"ddns_auth": {
				Type:        schema.TypeString,
				Description: "Enable/disable TSIG authentication for your DDNS server.",
				Computed:    true,
			},
			"ddns_domain": {
				Type:        schema.TypeString,
				Description: "Your fully qualified domain name. For example, yourname.ddns.com.",
				Computed:    true,
			},
			"ddns_key": {
				Type:        schema.TypeString,
				Description: "DDNS update key (base 64 encoding).",
				Computed:    true,
				Sensitive:   true,
			},
			"ddns_keyname": {
				Type:        schema.TypeString,
				Description: "DDNS update key name.",
				Computed:    true,
			},
			"ddns_password": {
				Type:        schema.TypeString,
				Description: "DDNS password.",
				Computed:    true,
				Sensitive:   true,
			},
			"ddns_server": {
				Type:        schema.TypeString,
				Description: "Select a DDNS service provider.",
				Computed:    true,
			},
			"ddns_server_addr": {
				Type:        schema.TypeList,
				Description: "Generic DDNS server IP/FQDN list.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"addr": {
							Type:        schema.TypeString,
							Description: "IP address or FQDN of the server.",
							Computed:    true,
						},
					},
				},
			},
			"ddns_server_ip": {
				Type:        schema.TypeString,
				Description: "Generic DDNS server IP.",
				Computed:    true,
			},
			"ddns_sn": {
				Type:        schema.TypeString,
				Description: "DDNS Serial Number.",
				Computed:    true,
			},
			"ddns_ttl": {
				Type:        schema.TypeInt,
				Description: "Time-to-live for DDNS packets.",
				Computed:    true,
			},
			"ddns_username": {
				Type:        schema.TypeString,
				Description: "DDNS user name.",
				Computed:    true,
			},
			"ddns_zone": {
				Type:        schema.TypeString,
				Description: "Zone of your domain name (for example, DDNS.com).",
				Computed:    true,
			},
			"ddnsid": {
				Type:        schema.TypeInt,
				Description: "DDNS ID.",
				Required:    true,
			},
			"monitor_interface": {
				Type:        schema.TypeList,
				Description: "Monitored interface.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interface_name": {
							Type:        schema.TypeString,
							Description: "Interface name.",
							Computed:    true,
						},
					},
				},
			},
			"server_type": {
				Type:        schema.TypeString,
				Description: "Address type of the DDNS server.",
				Computed:    true,
			},
			"ssl_certificate": {
				Type:        schema.TypeString,
				Description: "Name of local certificate for SSL connections.",
				Computed:    true,
			},
			"update_interval": {
				Type:        schema.TypeInt,
				Description: "DDNS update interval (60 - 2592000 sec, 0 means default).",
				Computed:    true,
			},
			"use_public_ip": {
				Type:        schema.TypeString,
				Description: "Enable/disable use of public IP address.",
				Computed:    true,
			},
		},
	}
}

func dataSourceSystemDdnsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	i := d.Get("ddnsid")
	mkey := utils.ParseMkey(i)

	o, err := c.Cmdb.ReadSystemDdns(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemDdns dataSource: %v", err)
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

	diags := refreshObjectSystemDdns(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
