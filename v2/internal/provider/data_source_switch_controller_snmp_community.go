// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure FortiSwitch SNMP v1/v2c communities globally.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceSwitchControllerSnmpCommunity() *schema.Resource {
	return &schema.Resource{
		Description: "Configure FortiSwitch SNMP v1/v2c communities globally.",

		ReadContext: dataSourceSwitchControllerSnmpCommunityRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"events": {
				Type:        schema.TypeString,
				Description: "SNMP notifications (traps) to send.",
				Computed:    true,
			},
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
			"fosid": {
				Type:        schema.TypeInt,
				Description: "SNMP community ID.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "SNMP community name.",
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
				Description: "SNMP v2c trap local port (default = 162).",
				Computed:    true,
			},
			"trap_v1_rport": {
				Type:        schema.TypeInt,
				Description: "SNMP v2c trap remote port (default = 162).",
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
		},
	}
}

func dataSourceSwitchControllerSnmpCommunityRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSwitchControllerSnmpCommunity(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SwitchControllerSnmpCommunity dataSource: %v", err)
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

	diags := refreshObjectSwitchControllerSnmpCommunity(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
