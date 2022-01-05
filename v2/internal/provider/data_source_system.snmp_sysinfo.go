// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: SNMP system info configuration.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceSystemSnmpSysinfo() *schema.Resource {
	return &schema.Resource{
		Description: "SNMP system info configuration.",

		ReadContext: dataSourceSystemSnmpSysinfoRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"contact_info": {
				Type:        schema.TypeString,
				Description: "Contact information.",
				Computed:    true,
			},
			"description": {
				Type:        schema.TypeString,
				Description: "System description.",
				Computed:    true,
			},
			"engine_id": {
				Type:        schema.TypeString,
				Description: "Local SNMP engineID string (maximum 27 characters).",
				Computed:    true,
			},
			"engine_id_type": {
				Type:        schema.TypeString,
				Description: "Local SNMP engineID type (text/hex/mac).",
				Computed:    true,
			},
			"location": {
				Type:        schema.TypeString,
				Description: "System location.",
				Computed:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Enable/disable SNMP.",
				Computed:    true,
			},
			"trap_high_cpu_threshold": {
				Type:        schema.TypeInt,
				Description: "CPU usage when trap is sent.",
				Computed:    true,
			},
			"trap_log_full_threshold": {
				Type:        schema.TypeInt,
				Description: "Log disk usage when trap is sent.",
				Computed:    true,
			},
			"trap_low_memory_threshold": {
				Type:        schema.TypeInt,
				Description: "Memory usage when trap is sent.",
				Computed:    true,
			},
		},
	}
}

func dataSourceSystemSnmpSysinfoRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "SystemSnmpSysinfo"

	o, err := c.Cmdb.ReadSystemSnmpSysinfo(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemSnmpSysinfo dataSource: %v", err)
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

	diags := refreshObjectSystemSnmpSysinfo(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
