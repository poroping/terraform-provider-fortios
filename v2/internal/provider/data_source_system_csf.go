// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Add this FortiGate to a Security Fabric or set up a new Security Fabric on this FortiGate.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceSystemCsf() *schema.Resource {
	return &schema.Resource{
		Description: "Add this FortiGate to a Security Fabric or set up a new Security Fabric on this FortiGate.",

		ReadContext: dataSourceSystemCsfRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"accept_auth_by_cert": {
				Type:        schema.TypeString,
				Description: "Accept connections with unknown certificates and ask admin for approval.",
				Computed:    true,
			},
			"authorization_request_type": {
				Type:        schema.TypeString,
				Description: "Authorization request type.",
				Computed:    true,
			},
			"certificate": {
				Type:        schema.TypeString,
				Description: "Certificate.",
				Computed:    true,
			},
			"configuration_sync": {
				Type:        schema.TypeString,
				Description: "Configuration sync mode.",
				Computed:    true,
			},
			"downstream_access": {
				Type:        schema.TypeString,
				Description: "Enable/disable downstream device access to this device's configuration and data.",
				Computed:    true,
			},
			"downstream_accprofile": {
				Type:        schema.TypeString,
				Description: "Default access profile for requests from downstream devices.",
				Computed:    true,
			},
			"fabric_connector": {
				Type:        schema.TypeList,
				Description: "Fabric connector configuration.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"accprofile": {
							Type:        schema.TypeString,
							Description: "Override access profile.",
							Computed:    true,
						},
						"configuration_write_access": {
							Type:        schema.TypeString,
							Description: "Enable/disable downstream device write access to configuration.",
							Computed:    true,
						},
						"serial": {
							Type:        schema.TypeString,
							Description: "Serial.",
							Computed:    true,
						},
					},
				},
			},
			"fabric_device": {
				Type:        schema.TypeList,
				Description: "Fabric device configuration.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"access_token": {
							Type:        schema.TypeString,
							Description: "Device access token.",
							Computed:    true,
							Sensitive:   true,
						},
						"device_ip": {
							Type:        schema.TypeString,
							Description: "Device IP.",
							Computed:    true,
						},
						"https_port": {
							Type:        schema.TypeInt,
							Description: "HTTPS port for fabric device.",
							Computed:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "Device name.",
							Computed:    true,
						},
					},
				},
			},
			"fabric_object_unification": {
				Type:        schema.TypeString,
				Description: "Fabric CMDB Object Unification.",
				Computed:    true,
			},
			"fabric_workers": {
				Type:        schema.TypeInt,
				Description: "Number of worker processes for Security Fabric daemon.",
				Computed:    true,
			},
			"group_name": {
				Type:        schema.TypeString,
				Description: "Security Fabric group name. All FortiGates in a Security Fabric must have the same group name.",
				Computed:    true,
			},
			"group_password": {
				Type:        schema.TypeString,
				Description: "Security Fabric group password. All FortiGates in a Security Fabric must have the same group password.",
				Computed:    true,
				Sensitive:   true,
			},
			"log_unification": {
				Type:        schema.TypeString,
				Description: "Enable/disable broadcast of discovery messages for log unification.",
				Computed:    true,
			},
			"management_ip": {
				Type:        schema.TypeString,
				Description: "Management IP address of this FortiGate. Used to log into this FortiGate from another FortiGate in the Security Fabric.",
				Computed:    true,
			},
			"management_port": {
				Type:        schema.TypeInt,
				Description: "Overriding port for management connection (Overrides admin port).",
				Computed:    true,
			},
			"saml_configuration_sync": {
				Type:        schema.TypeString,
				Description: "SAML setting configuration synchronization.",
				Computed:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Enable/disable Security Fabric.",
				Computed:    true,
			},
			"trusted_list": {
				Type:        schema.TypeList,
				Description: "Pre-authorized and blocked security fabric nodes.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:        schema.TypeString,
							Description: "Security fabric authorization action.",
							Computed:    true,
						},
						"authorization_type": {
							Type:        schema.TypeString,
							Description: "Authorization type.",
							Computed:    true,
						},
						"certificate": {
							Type:        schema.TypeString,
							Description: "Certificate.",
							Computed:    true,
						},
						"downstream_authorization": {
							Type:        schema.TypeString,
							Description: "Trust authorizations by this node's administrator.",
							Computed:    true,
						},
						"ha_members": {
							Type:        schema.TypeString,
							Description: "HA members.",
							Computed:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "Name.",
							Computed:    true,
						},
						"serial": {
							Type:        schema.TypeString,
							Description: "Serial.",
							Computed:    true,
						},
					},
				},
			},
			"upstream": {
				Type:        schema.TypeString,
				Description: "IP/FQDN of the FortiGate upstream from this FortiGate in the Security Fabric.",
				Computed:    true,
			},
			"upstream_ip": {
				Type:        schema.TypeString,
				Description: "IP address of the FortiGate upstream from this FortiGate in the Security Fabric.",
				Computed:    true,
			},
			"upstream_port": {
				Type:        schema.TypeInt,
				Description: "The port number to use to communicate with the FortiGate upstream from this FortiGate in the Security Fabric (default = 8013).",
				Computed:    true,
			},
		},
	}
}

func dataSourceSystemCsfRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "SystemCsf"

	o, err := c.Cmdb.ReadSystemCsf(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemCsf dataSource: %v", err)
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

	diags := refreshObjectSystemCsf(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
