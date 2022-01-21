// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure Wireless Termination Points (WTP) system log server profile.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceWirelessControllerSyslogProfile() *schema.Resource {
	return &schema.Resource{
		Description: "Configure Wireless Termination Points (WTP) system log server profile.",

		ReadContext: dataSourceWirelessControllerSyslogProfileRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"comment": {
				Type:        schema.TypeString,
				Description: "Comment.",
				Computed:    true,
			},
			"log_level": {
				Type:        schema.TypeString,
				Description: "Lowest level of log messages that FortiAP units send to this server (default = information)",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "WTP system log server profile name.",
				Required:    true,
			},
			"server_addr_type": {
				Type:        schema.TypeString,
				Description: "Syslog server address type (default = IP)",
				Computed:    true,
			},
			"server_fqdn": {
				Type:        schema.TypeString,
				Description: "FQDN of syslog server that FortiAP units send log messages to.",
				Computed:    true,
			},
			"server_ip": {
				Type:        schema.TypeString,
				Description: "IP address of syslog server that FortiAP units send log messages to.",
				Computed:    true,
			},
			"server_port": {
				Type:        schema.TypeInt,
				Description: "Port number of syslog server that FortiAP units send log messages to (default = 514).",
				Computed:    true,
			},
			"server_status": {
				Type:        schema.TypeString,
				Description: "Enable/disable FortiAP units to send log messages to a syslog server (default = enable).",
				Computed:    true,
			},
		},
	}
}

func dataSourceWirelessControllerSyslogProfileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadWirelessControllerSyslogProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WirelessControllerSyslogProfile dataSource: %v", err)
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

	diags := refreshObjectWirelessControllerSyslogProfile(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
