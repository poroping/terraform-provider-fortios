// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure domain controller entries.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceUserDomainController() *schema.Resource {
	return &schema.Resource{
		Description: "Configure domain controller entries.",

		ReadContext: dataSourceUserDomainControllerRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"ad_mode": {
				Type:        schema.TypeString,
				Description: "Set Active Directory mode.",
				Computed:    true,
			},
			"adlds_dn": {
				Type:        schema.TypeString,
				Description: "AD LDS distinguished name.",
				Computed:    true,
			},
			"adlds_ip_address": {
				Type:        schema.TypeString,
				Description: "AD LDS IPv4 address.",
				Computed:    true,
			},
			"adlds_ip6": {
				Type:        schema.TypeString,
				Description: "AD LDS IPv6 address.",
				Computed:    true,
			},
			"adlds_port": {
				Type:        schema.TypeInt,
				Description: "Port number of AD LDS service (default = 389).",
				Computed:    true,
			},
			"change_detection": {
				Type:        schema.TypeString,
				Description: "Enable/disable detection of a configuration change in the Active Directory server.",
				Computed:    true,
			},
			"change_detection_period": {
				Type:        schema.TypeInt,
				Description: "Minutes to detect a configuration change in the Active Directory server (5 - 10080 minutes (7 days), default = 60).",
				Computed:    true,
			},
			"dns_srv_lookup": {
				Type:        schema.TypeString,
				Description: "Enable/disable DNS service lookup.",
				Computed:    true,
			},
			"domain_name": {
				Type:        schema.TypeString,
				Description: "Domain DNS name.",
				Computed:    true,
			},
			"extra_server": {
				Type:        schema.TypeList,
				Description: "Extra servers.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Description: "Server ID.",
							Computed:    true,
						},
						"ip_address": {
							Type:        schema.TypeString,
							Description: "Domain controller IP address.",
							Computed:    true,
						},
						"port": {
							Type:        schema.TypeInt,
							Description: "Port to be used for communication with the domain controller (default = 445).",
							Computed:    true,
						},
						"source_ip_address": {
							Type:        schema.TypeString,
							Description: "FortiGate IPv4 address to be used for communication with the domain controller.",
							Computed:    true,
						},
						"source_port": {
							Type:        schema.TypeInt,
							Description: "Source port to be used for communication with the domain controller.",
							Computed:    true,
						},
					},
				},
			},
			"hostname": {
				Type:        schema.TypeString,
				Description: "Hostname of the server to connect to.",
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
			"ip_address": {
				Type:        schema.TypeString,
				Description: "Domain controller IPv4 address.",
				Computed:    true,
			},
			"ip6": {
				Type:        schema.TypeString,
				Description: "Domain controller IPv6 address.",
				Computed:    true,
			},
			"ldap_server": {
				Type:        schema.TypeList,
				Description: "LDAP server name(s).",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "LDAP server name.",
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Domain controller entry name.",
				Required:    true,
			},
			"password": {
				Type:        schema.TypeString,
				Description: "Password for specified username.",
				Computed:    true,
				Sensitive:   true,
			},
			"port": {
				Type:        schema.TypeInt,
				Description: "Port to be used for communication with the domain controller (default = 445).",
				Computed:    true,
			},
			"replication_port": {
				Type:        schema.TypeInt,
				Description: "Port to be used for communication with the domain controller for replication service. Port number 0 indicates automatic discovery.",
				Computed:    true,
			},
			"source_ip_address": {
				Type:        schema.TypeString,
				Description: "FortiGate IPv4 address to be used for communication with the domain controller.",
				Computed:    true,
			},
			"source_ip6": {
				Type:        schema.TypeString,
				Description: "FortiGate IPv6 address to be used for communication with the domain controller.",
				Computed:    true,
			},
			"source_port": {
				Type:        schema.TypeInt,
				Description: "Source port to be used for communication with the domain controller.",
				Computed:    true,
			},
			"username": {
				Type:        schema.TypeString,
				Description: "User name to sign in with. Must have proper permissions for service.",
				Computed:    true,
			},
		},
	}
}

func dataSourceUserDomainControllerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadUserDomainController(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading UserDomainController dataSource: %v", err)
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

	diags := refreshObjectUserDomainController(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
