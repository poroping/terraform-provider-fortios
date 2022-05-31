// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: SNMP community configuration.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceSystemSnmpCommunity() *schema.Resource {
	return &schema.Resource{
		Description: "SNMP community configuration.",

		ReadContext: dataSourceSystemSnmpCommunityRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"events": {
				Type:        schema.TypeString,
				Description: "SNMP trap events.",
				Computed:    true,
			},
			"hosts": {
				Type:        schema.TypeList,
				Description: "Configure IPv4 SNMP managers (hosts).",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ha_direct": {
							Type:        schema.TypeString,
							Description: "Enable/disable direct management of HA cluster members.",
							Computed:    true,
						},
						"host_type": {
							Type:        schema.TypeString,
							Description: "Control whether the SNMP manager sends SNMP queries, receives SNMP traps, or both. No traps will be sent when IP type is subnet.",
							Computed:    true,
						},
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
						"source_ip": {
							Type:        schema.TypeString,
							Description: "Source IPv4 address for SNMP traps.",
							Computed:    true,
						},
					},
				},
			},
			"hosts6": {
				Type:        schema.TypeList,
				Description: "Configure IPv6 SNMP managers.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ha_direct": {
							Type:        schema.TypeString,
							Description: "Enable/disable direct management of HA cluster members.",
							Computed:    true,
						},
						"host_type": {
							Type:        schema.TypeString,
							Description: "Control whether the SNMP manager sends SNMP queries, receives SNMP traps, or both.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "Host6 entry ID.",
							Computed:    true,
						},
						"ipv6": {
							Type:        schema.TypeString,
							Description: "SNMP manager IPv6 address prefix.",
							Computed:    true,
						},
						"source_ipv6": {
							Type:        schema.TypeString,
							Description: "Source IPv6 address for SNMP traps.",
							Computed:    true,
						},
					},
				},
			},
			"fosid": {
				Type:        schema.TypeInt,
				Description: "Community ID.",
				Computed:    true,
			},
			"mib_view": {
				Type:        schema.TypeString,
				Description: "SNMP access control MIB view.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Community name.",
				Computed:    true,
			},
			"query_v1_port": {
				Type:        schema.TypeInt,
				Description: "SNMP v1 query port (default = 161).",
				Computed:    true,
			},
			"query_v1_status": {
				Type:        schema.TypeString,
				Description: "Enable/disable SNMP v1 queries.",
				Computed:    true,
			},
			"query_v2c_port": {
				Type:        schema.TypeInt,
				Description: "SNMP v2c query port (default = 161).",
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
			"trap_v1_lport": {
				Type:        schema.TypeInt,
				Description: "SNMP v1 trap local port (default = 162).",
				Computed:    true,
			},
			"trap_v1_rport": {
				Type:        schema.TypeInt,
				Description: "SNMP v1 trap remote port (default = 162).",
				Computed:    true,
			},
			"trap_v1_status": {
				Type:        schema.TypeString,
				Description: "Enable/disable SNMP v1 traps.",
				Computed:    true,
			},
			"trap_v2c_lport": {
				Type:        schema.TypeInt,
				Description: "SNMP v2c trap local port (default = 162).",
				Computed:    true,
			},
			"trap_v2c_rport": {
				Type:        schema.TypeInt,
				Description: "SNMP v2c trap remote port (default = 162).",
				Computed:    true,
			},
			"trap_v2c_status": {
				Type:        schema.TypeString,
				Description: "Enable/disable SNMP v2c traps.",
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

func dataSourceSystemSnmpCommunityRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	i := d.Get("fosid")
	mkey := utils.ParseMkey(i)

	o, err := c.Cmdb.ReadSystemSnmpCommunity(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemSnmpCommunity dataSource: %v", err)
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

	diags := refreshObjectSystemSnmpCommunity(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
