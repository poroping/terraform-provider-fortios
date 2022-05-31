// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: SNMP user configuration.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceSystemSnmpUser() *schema.Resource {
	return &schema.Resource{
		Description: "SNMP user configuration.",

		ReadContext: dataSourceSystemSnmpUserRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"auth_proto": {
				Type:        schema.TypeString,
				Description: "Authentication protocol.",
				Computed:    true,
			},
			"auth_pwd": {
				Type:        schema.TypeString,
				Description: "Password for authentication protocol.",
				Computed:    true,
				Sensitive:   true,
			},
			"events": {
				Type:        schema.TypeString,
				Description: "SNMP notifications (traps) to send.",
				Computed:    true,
			},
			"ha_direct": {
				Type:        schema.TypeString,
				Description: "Enable/disable direct management of HA cluster members.",
				Computed:    true,
			},
			"mib_view": {
				Type:        schema.TypeString,
				Description: "SNMP access control MIB view.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "SNMP user name.",
				Required:    true,
			},
			"notify_hosts": {
				Type:        schema.TypeString,
				Description: "SNMP managers to send notifications (traps) to.",
				Computed:    true,
			},
			"notify_hosts6": {
				Type:        schema.TypeString,
				Description: "IPv6 SNMP managers to send notifications (traps) to.",
				Computed:    true,
			},
			"priv_proto": {
				Type:        schema.TypeString,
				Description: "Privacy (encryption) protocol.",
				Computed:    true,
			},
			"priv_pwd": {
				Type:        schema.TypeString,
				Description: "Password for privacy (encryption) protocol.",
				Computed:    true,
				Sensitive:   true,
			},
			"queries": {
				Type:        schema.TypeString,
				Description: "Enable/disable SNMP queries for this user.",
				Computed:    true,
			},
			"query_port": {
				Type:        schema.TypeInt,
				Description: "SNMPv3 query port (default = 161).",
				Computed:    true,
			},
			"security_level": {
				Type:        schema.TypeString,
				Description: "Security level for message authentication and encryption.",
				Computed:    true,
			},
			"source_ip": {
				Type:        schema.TypeString,
				Description: "Source IP for SNMP trap.",
				Computed:    true,
			},
			"source_ipv6": {
				Type:        schema.TypeString,
				Description: "Source IPv6 for SNMP trap.",
				Computed:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Enable/disable this SNMP user.",
				Computed:    true,
			},
			"trap_lport": {
				Type:        schema.TypeInt,
				Description: "SNMPv3 local trap port (default = 162).",
				Computed:    true,
			},
			"trap_rport": {
				Type:        schema.TypeInt,
				Description: "SNMPv3 trap remote port (default = 162).",
				Computed:    true,
			},
			"trap_status": {
				Type:        schema.TypeString,
				Description: "Enable/disable traps for this SNMP user.",
				Computed:    true,
			},
			"vdoms": {
				Type:        schema.TypeList,
				Description: "SNMP access control VDOMs.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "VDOM name",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceSystemSnmpUserRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemSnmpUser(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemSnmpUser dataSource: %v", err)
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

	diags := refreshObjectSystemSnmpUser(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
