// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure Internet Services Extension.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceFirewallInternetServiceExtension() *schema.Resource {
	return &schema.Resource{
		Description: "Configure Internet Services Extension.",

		ReadContext: dataSourceFirewallInternetServiceExtensionRead,

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
			"disable_entry": {
				Type:        schema.TypeList,
				Description: "Disable entries in the Internet Service database.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"addr_mode": {
							Type:        schema.TypeString,
							Description: "Address mode (IPv4 or IPv6).",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "Disable entry ID.",
							Computed:    true,
						},
						"ip_range": {
							Type:        schema.TypeList,
							Description: "IPv4 ranges in the disable entry.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"end_ip": {
										Type:        schema.TypeString,
										Description: "End IPv4 address.",
										Computed:    true,
									},
									"id": {
										Type:        schema.TypeInt,
										Description: "Disable entry range ID.",
										Computed:    true,
									},
									"start_ip": {
										Type:        schema.TypeString,
										Description: "Start IPv4 address.",
										Computed:    true,
									},
								},
							},
						},
						"ip6_range": {
							Type:        schema.TypeList,
							Description: "IPv6 ranges in the disable entry.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"end_ip6": {
										Type:        schema.TypeString,
										Description: "End IPv6 address.",
										Computed:    true,
									},
									"id": {
										Type:        schema.TypeInt,
										Description: "Disable entry range ID.",
										Computed:    true,
									},
									"start_ip6": {
										Type:        schema.TypeString,
										Description: "Start IPv6 address.",
										Computed:    true,
									},
								},
							},
						},
						"port_range": {
							Type:        schema.TypeList,
							Description: "Port ranges in the disable entry.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"end_port": {
										Type:        schema.TypeInt,
										Description: "Ending TCP/UDP/SCTP destination port (0 to 65535).",
										Computed:    true,
									},
									"id": {
										Type:        schema.TypeInt,
										Description: "Custom entry port range ID.",
										Computed:    true,
									},
									"start_port": {
										Type:        schema.TypeInt,
										Description: "Starting TCP/UDP/SCTP destination port (0 to 65535).",
										Computed:    true,
									},
								},
							},
						},
						"protocol": {
							Type:        schema.TypeInt,
							Description: "Integer value for the protocol type as defined by IANA (0 - 255).",
							Computed:    true,
						},
					},
				},
			},
			"entry": {
				Type:        schema.TypeList,
				Description: "Entries added to the Internet Service extension database.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"addr_mode": {
							Type:        schema.TypeString,
							Description: "Address mode (IPv4 or IPv6).",
							Computed:    true,
						},
						"dst": {
							Type:        schema.TypeList,
							Description: "Destination address or address group name.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Description: "Select the destination address or address group object from available options.",
										Computed:    true,
									},
								},
							},
						},
						"dst6": {
							Type:        schema.TypeList,
							Description: "Destination address6 or address6 group name.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Description: "Select the destination address6 or address group object from available options.",
										Computed:    true,
									},
								},
							},
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "Entry ID(1-255).",
							Computed:    true,
						},
						"port_range": {
							Type:        schema.TypeList,
							Description: "Port ranges in the custom entry.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"end_port": {
										Type:        schema.TypeInt,
										Description: "Integer value for ending TCP/UDP/SCTP destination port in range (0 to 65535).",
										Computed:    true,
									},
									"id": {
										Type:        schema.TypeInt,
										Description: "Custom entry port range ID.",
										Computed:    true,
									},
									"start_port": {
										Type:        schema.TypeInt,
										Description: "Integer value for starting TCP/UDP/SCTP destination port in range (0 to 65535).",
										Computed:    true,
									},
								},
							},
						},
						"protocol": {
							Type:        schema.TypeInt,
							Description: "Integer value for the protocol type as defined by IANA (0 - 255).",
							Computed:    true,
						},
					},
				},
			},
			"fosid": {
				Type:        schema.TypeInt,
				Description: "Internet Service ID in the Internet Service database.",
				Computed:    true,
			},
		},
	}
}

func dataSourceFirewallInternetServiceExtensionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallInternetServiceExtension(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallInternetServiceExtension dataSource: %v", err)
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

	diags := refreshObjectFirewallInternetServiceExtension(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
