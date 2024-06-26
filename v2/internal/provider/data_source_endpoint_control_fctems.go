// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure FortiClient Enterprise Management Server (EMS) entries.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceEndpointControlFctems() *schema.Resource {
	return &schema.Resource{
		Description: "Configure FortiClient Enterprise Management Server (EMS) entries.",

		ReadContext: dataSourceEndpointControlFctemsRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"admin_password": {
				Type:        schema.TypeString,
				Description: "FortiClient EMS admin password.",
				Computed:    true,
				Sensitive:   true,
			},
			"admin_username": {
				Type:        schema.TypeString,
				Description: "FortiClient EMS admin username.",
				Computed:    true,
			},
			"call_timeout": {
				Type:        schema.TypeInt,
				Description: "FortiClient EMS call timeout in seconds (1 - 180 seconds, default = 30).",
				Computed:    true,
			},
			"capabilities": {
				Type:        schema.TypeString,
				Description: "List of EMS capabilities.",
				Computed:    true,
			},
			"certificate": {
				Type:        schema.TypeString,
				Description: "FortiClient EMS certificate.",
				Computed:    true,
			},
			"cloud_server_type": {
				Type:        schema.TypeString,
				Description: "Cloud server type.",
				Computed:    true,
			},
			"dirty_reason": {
				Type:        schema.TypeString,
				Description: "Dirty Reason for FortiClient EMS.",
				Computed:    true,
			},
			"ems_id": {
				Type:        schema.TypeInt,
				Description: "EMS ID in order (1 - 7).",
				Required:    true,
			},
			"fortinetone_cloud_authentication": {
				Type:        schema.TypeString,
				Description: "Enable/disable authentication of FortiClient EMS Cloud through FortiCloud account.",
				Computed:    true,
			},
			"https_port": {
				Type:        schema.TypeInt,
				Description: "FortiClient EMS HTTPS access port number. (1 - 65535, default: 443).",
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
			"name": {
				Type:        schema.TypeString,
				Description: "FortiClient Enterprise Management Server (EMS) name.",
				Computed:    true,
			},
			"out_of_sync_threshold": {
				Type:        schema.TypeInt,
				Description: "Outdated resource threshold in seconds (10 - 3600, default = 180).",
				Computed:    true,
			},
			"preserve_ssl_session": {
				Type:        schema.TypeString,
				Description: "Enable/disable preservation of EMS SSL session connection. Warning, most users should not touch this setting.",
				Computed:    true,
			},
			"pull_avatars": {
				Type:        schema.TypeString,
				Description: "Enable/disable pulling avatars from EMS.",
				Computed:    true,
			},
			"pull_malware_hash": {
				Type:        schema.TypeString,
				Description: "Enable/disable pulling FortiClient malware hash from EMS.",
				Computed:    true,
			},
			"pull_sysinfo": {
				Type:        schema.TypeString,
				Description: "Enable/disable pulling SysInfo from EMS.",
				Computed:    true,
			},
			"pull_tags": {
				Type:        schema.TypeString,
				Description: "Enable/disable pulling FortiClient user tags from EMS.",
				Computed:    true,
			},
			"pull_vulnerabilities": {
				Type:        schema.TypeString,
				Description: "Enable/disable pulling vulnerabilities from EMS.",
				Computed:    true,
			},
			"serial_number": {
				Type:        schema.TypeString,
				Description: "EMS Serial Number.",
				Computed:    true,
			},
			"server": {
				Type:        schema.TypeString,
				Description: "FortiClient EMS FQDN or IPv4 address.",
				Computed:    true,
			},
			"source_ip": {
				Type:        schema.TypeString,
				Description: "REST API call source IP.",
				Computed:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Enable or disable this EMS configuration.",
				Computed:    true,
			},
			"status_check_interval": {
				Type:        schema.TypeInt,
				Description: "FortiClient EMS call timeout in seconds (1 - 120 seconds, default = 5).",
				Computed:    true,
			},
			"tenant_id": {
				Type:        schema.TypeString,
				Description: "EMS Tenant ID.",
				Computed:    true,
			},
			"trust_ca_cn": {
				Type:        schema.TypeString,
				Description: "Enable/disable trust of the EMS certificate issuer(CA) and common name(CN) for certificate auto-renewal.",
				Computed:    true,
			},
			"websocket_override": {
				Type:        schema.TypeString,
				Description: "Enable/disable override behavior for how this FortiGate unit connects to EMS using a WebSocket connection.",
				Computed:    true,
			},
		},
	}
}

func dataSourceEndpointControlFctemsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	i := d.Get("ems_id")
	mkey := utils.ParseMkey(i)

	o, err := c.Cmdb.ReadEndpointControlFctems(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading EndpointControlFctems dataSource: %v", err)
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

	diags := refreshObjectEndpointControlFctems(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
