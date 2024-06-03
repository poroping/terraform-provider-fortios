// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure MPSK profile.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceWirelessControllerMpskProfile() *schema.Resource {
	return &schema.Resource{
		Description: "Configure MPSK profile.",

		ReadContext: dataSourceWirelessControllerMpskProfileRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"mpsk_concurrent_clients": {
				Type:        schema.TypeInt,
				Description: "Maximum number of concurrent clients that connect using the same passphrase in multiple PSK authentication (0 - 65535, default = 0, meaning no limitation).",
				Computed:    true,
			},
			"mpsk_group": {
				Type:        schema.TypeList,
				Description: "List of multiple PSK groups.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mpsk_key": {
							Type:        schema.TypeList,
							Description: "List of multiple PSK entries.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"comment": {
										Type:        schema.TypeString,
										Description: "Comment.",
										Computed:    true,
									},
									"concurrent_client_limit_type": {
										Type:        schema.TypeString,
										Description: "MPSK client limit type options.",
										Computed:    true,
									},
									"concurrent_clients": {
										Type:        schema.TypeInt,
										Description: "Number of clients that can connect using this pre-shared key (1 - 65535, default is 256).",
										Computed:    true,
									},
									"mac": {
										Type:        schema.TypeString,
										Description: "MAC address.",
										Computed:    true,
									},
									"mpsk_schedules": {
										Type:        schema.TypeList,
										Description: "Firewall schedule for MPSK passphrase. The passphrase will be effective only when at least one schedule is valid.",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:        schema.TypeString,
													Description: "Schedule name.",
													Computed:    true,
												},
											},
										},
									},
									"name": {
										Type:        schema.TypeString,
										Description: "Pre-shared key name.",
										Computed:    true,
									},
									"passphrase": {
										Type:        schema.TypeString,
										Description: "WPA Pre-shared key.",
										Computed:    true,
										Sensitive:   true,
									},
								},
							},
						},
						"name": {
							Type:        schema.TypeString,
							Description: "MPSK group name.",
							Computed:    true,
						},
						"vlan_id": {
							Type:        schema.TypeInt,
							Description: "Optional VLAN ID.",
							Computed:    true,
						},
						"vlan_type": {
							Type:        schema.TypeString,
							Description: "MPSK group VLAN options.",
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:        schema.TypeString,
				Description: "MPSK profile name.",
				Required:    true,
			},
		},
	}
}

func dataSourceWirelessControllerMpskProfileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadWirelessControllerMpskProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WirelessControllerMpskProfile dataSource: %v", err)
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

	diags := refreshObjectWirelessControllerMpskProfile(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
