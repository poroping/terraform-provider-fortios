// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure SNMP.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceWirelessControllerSnmp() *schema.Resource {
	return &schema.Resource{
		Description: "Configure SNMP.",

		ReadContext: dataSourceWirelessControllerSnmpRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"community": {
				Type:        schema.TypeList,
				Description: "SNMP Community Configuration.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hosts": {
							Type:        schema.TypeList,
							Description: "Configure IPv4 SNMP managers (hosts).",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeInt,
										Description: "Host entry ID.",
										Computed:    true,
									},
									"ip": {
										Type:        schema.TypeString,
										Description: "IPv4 address of the SNMP manager (host).",
										Computed:    true,
									},
								},
							},
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "Community ID.",
							Computed:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "Community name.",
							Computed:    true,
						},
						"query_v1_status": {
							Type:        schema.TypeString,
							Description: "Enable/disable SNMP v1 queries.",
							Computed:    true,
						},
						"query_v2c_status": {
							Type:        schema.TypeString,
							Description: "Enable/disable SNMP v2c queries.",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Enable/disable this SNMP community.",
							Computed:    true,
						},
						"trap_v1_status": {
							Type:        schema.TypeString,
							Description: "Enable/disable SNMP v1 traps.",
							Computed:    true,
						},
						"trap_v2c_status": {
							Type:        schema.TypeString,
							Description: "Enable/disable SNMP v2c traps.",
							Computed:    true,
						},
					},
				},
			},
			"contact_info": {
				Type:        schema.TypeString,
				Description: "Contact Information.",
				Computed:    true,
			},
			"engine_id": {
				Type:        schema.TypeString,
				Description: "AC SNMP engineID string (maximum 24 characters).",
				Computed:    true,
			},
			"trap_high_cpu_threshold": {
				Type:        schema.TypeInt,
				Description: "CPU usage when trap is sent.",
				Computed:    true,
			},
			"trap_high_mem_threshold": {
				Type:        schema.TypeInt,
				Description: "Memory usage when trap is sent.",
				Computed:    true,
			},
			"user": {
				Type:        schema.TypeList,
				Description: "SNMP User Configuration.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
						"name": {
							Type:        schema.TypeString,
							Description: "SNMP user name.",
							Computed:    true,
						},
						"notify_hosts": {
							Type:        schema.TypeString,
							Description: "Configure SNMP User Notify Hosts.",
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
						"security_level": {
							Type:        schema.TypeString,
							Description: "Security level for message authentication and encryption.",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "SNMP user enable.",
							Computed:    true,
						},
						"trap_status": {
							Type:        schema.TypeString,
							Description: "Enable/disable traps for this SNMP user.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceWirelessControllerSnmpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "WirelessControllerSnmp"

	o, err := c.Cmdb.ReadWirelessControllerSnmp(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WirelessControllerSnmp dataSource: %v", err)
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

	diags := refreshObjectWirelessControllerSnmp(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
