// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure Access Proxy SSH client certificate.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceFirewallAccessProxySshClientCert() *schema.Resource {
	return &schema.Resource{
		Description: "Configure Access Proxy SSH client certificate.",

		ReadContext: dataSourceFirewallAccessProxySshClientCertRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"auth_ca": {
				Type:        schema.TypeString,
				Description: "Name of the SSH server public key authentication CA.",
				Computed:    true,
			},
			"cert_extension": {
				Type:        schema.TypeList,
				Description: "Configure certificate extension for user certificate.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"critical": {
							Type:        schema.TypeString,
							Description: "Critical option.",
							Computed:    true,
						},
						"data": {
							Type:        schema.TypeString,
							Description: "Data of certificate extension.",
							Computed:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "Name of certificate extension.",
							Computed:    true,
						},
						"type": {
							Type:        schema.TypeString,
							Description: "Type of certificate extension.",
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:        schema.TypeString,
				Description: "SSH client certificate name.",
				Required:    true,
			},
			"permit_agent_forwarding": {
				Type:        schema.TypeString,
				Description: "Enable/disable appending permit-agent-forwarding certificate extension.",
				Computed:    true,
			},
			"permit_port_forwarding": {
				Type:        schema.TypeString,
				Description: "Enable/disable appending permit-port-forwarding certificate extension.",
				Computed:    true,
			},
			"permit_pty": {
				Type:        schema.TypeString,
				Description: "Enable/disable appending permit-pty certificate extension.",
				Computed:    true,
			},
			"permit_user_rc": {
				Type:        schema.TypeString,
				Description: "Enable/disable appending permit-user-rc certificate extension.",
				Computed:    true,
			},
			"permit_x11_forwarding": {
				Type:        schema.TypeString,
				Description: "Enable/disable appending permit-x11-forwarding certificate extension.",
				Computed:    true,
			},
			"source_address": {
				Type:        schema.TypeString,
				Description: "Enable/disable appending source-address certificate critical option. This option ensure certificate only accepted from FortiGate source address.",
				Computed:    true,
			},
		},
	}
}

func dataSourceFirewallAccessProxySshClientCertRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallAccessProxySshClientCert(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallAccessProxySshClientCert dataSource: %v", err)
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

	diags := refreshObjectFirewallAccessProxySshClientCert(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
