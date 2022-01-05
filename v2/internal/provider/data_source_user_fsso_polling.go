// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure FSSO active directory servers for polling mode.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceUserFssoPolling() *schema.Resource {
	return &schema.Resource{
		Description: "Configure FSSO active directory servers for polling mode.",

		ReadContext: dataSourceUserFssoPollingRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"adgrp": {
				Type:        schema.TypeList,
				Description: "LDAP Group Info.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Name.",
							Computed:    true,
						},
					},
				},
			},
			"default_domain": {
				Type:        schema.TypeString,
				Description: "Default domain managed by this Active Directory server.",
				Computed:    true,
			},
			"fosid": {
				Type:        schema.TypeInt,
				Description: "Active Directory server ID.",
				Computed:    true,
			},
			"ldap_server": {
				Type:        schema.TypeString,
				Description: "LDAP server name used in LDAP connection strings.",
				Computed:    true,
			},
			"logon_history": {
				Type:        schema.TypeInt,
				Description: "Number of hours of logon history to keep, 0 means keep all history.",
				Computed:    true,
			},
			"password": {
				Type:        schema.TypeString,
				Description: "Password required to log into this Active Directory server",
				Computed:    true,
				Sensitive:   true,
			},
			"polling_frequency": {
				Type:        schema.TypeInt,
				Description: "Polling frequency (every 1 to 30 seconds).",
				Computed:    true,
			},
			"port": {
				Type:        schema.TypeInt,
				Description: "Port to communicate with this Active Directory server.",
				Computed:    true,
			},
			"server": {
				Type:        schema.TypeString,
				Description: "Host name or IP address of the Active Directory server.",
				Computed:    true,
			},
			"smb_ntlmv1_auth": {
				Type:        schema.TypeString,
				Description: "Enable/disable support of NTLMv1 for Samba authentication.",
				Computed:    true,
			},
			"smbv1": {
				Type:        schema.TypeString,
				Description: "Enable/disable support of SMBv1 for Samba.",
				Computed:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Enable/disable polling for the status of this Active Directory server.",
				Computed:    true,
			},
			"user": {
				Type:        schema.TypeString,
				Description: "User name required to log into this Active Directory server.",
				Computed:    true,
			},
		},
	}
}

func dataSourceUserFssoPollingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadUserFssoPolling(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading UserFssoPolling dataSource: %v", err)
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

	diags := refreshObjectUserFssoPolling(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
