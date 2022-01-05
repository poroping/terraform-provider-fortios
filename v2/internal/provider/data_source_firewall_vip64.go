// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure IPv6 to IPv4 virtual IPs.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceFirewallVip64() *schema.Resource {
	return &schema.Resource{
		Description: "Configure IPv6 to IPv4 virtual IPs.",

		ReadContext: dataSourceFirewallVip64Read,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"arp_reply": {
				Type:        schema.TypeString,
				Description: "Enable ARP reply.",
				Computed:    true,
			},
			"color": {
				Type:        schema.TypeInt,
				Description: "Color of icon on the GUI.",
				Computed:    true,
			},
			"comment": {
				Type:        schema.TypeString,
				Description: "Comment.",
				Computed:    true,
			},
			"extip": {
				Type:        schema.TypeString,
				Description: "Start-external-IPv6-address [-end-external-IPv6-address].",
				Computed:    true,
			},
			"extport": {
				Type:        schema.TypeString,
				Description: "External service port.",
				Computed:    true,
			},
			"fosid": {
				Type:        schema.TypeInt,
				Description: "Custom defined id.",
				Computed:    true,
			},
			"ldb_method": {
				Type:        schema.TypeString,
				Description: "Load balance method.",
				Computed:    true,
			},
			"mappedip": {
				Type:        schema.TypeString,
				Description: "Start-mapped-IP [-end-mapped-IP].",
				Computed:    true,
			},
			"mappedport": {
				Type:        schema.TypeString,
				Description: "Mapped service port.",
				Computed:    true,
			},
			"monitor": {
				Type:        schema.TypeList,
				Description: "Health monitors.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Health monitor name.",
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:        schema.TypeString,
				Description: "VIP64 name.",
				Required:    true,
			},
			"portforward": {
				Type:        schema.TypeString,
				Description: "Enable port forwarding.",
				Computed:    true,
			},
			"protocol": {
				Type:        schema.TypeString,
				Description: "Mapped port protocol.",
				Computed:    true,
			},
			"realservers": {
				Type:        schema.TypeList,
				Description: "Real servers.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"client_ip": {
							Type:        schema.TypeString,
							Description: "Restrict server to a client IP in this range.",
							Computed:    true,
						},
						"healthcheck": {
							Type:        schema.TypeString,
							Description: "Per server health check.",
							Computed:    true,
						},
						"holddown_interval": {
							Type:        schema.TypeInt,
							Description: "Hold down interval.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "Real server ID.",
							Computed:    true,
						},
						"ip": {
							Type:        schema.TypeString,
							Description: "Mapped server IP.",
							Computed:    true,
						},
						"max_connections": {
							Type:        schema.TypeInt,
							Description: "Maximum number of connections allowed to server.",
							Computed:    true,
						},
						"monitor": {
							Type:        schema.TypeList,
							Description: "Health monitors.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Description: "Health monitor name.",
										Computed:    true,
									},
								},
							},
						},
						"port": {
							Type:        schema.TypeInt,
							Description: "Mapped server port.",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Server administrative status.",
							Computed:    true,
						},
						"weight": {
							Type:        schema.TypeInt,
							Description: "weight",
							Computed:    true,
						},
					},
				},
			},
			"server_type": {
				Type:        schema.TypeString,
				Description: "Server type.",
				Computed:    true,
			},
			"src_filter": {
				Type:        schema.TypeList,
				Description: "Source IP6 filter (x:x:x:x:x:x:x:x/x).",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"range": {
							Type:        schema.TypeString,
							Description: "Src-filter range.",
							Computed:    true,
						},
					},
				},
			},
			"type": {
				Type:        schema.TypeString,
				Description: "VIP type: static NAT or server load balance.",
				Computed:    true,
			},
			"uuid": {
				Type:        schema.TypeString,
				Description: "Universally Unique Identifier (UUID; automatically assigned but can be manually reset).",
				Computed:    true,
			},
		},
	}
}

func dataSourceFirewallVip64Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallVip64(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallVip64 dataSource: %v", err)
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

	diags := refreshObjectFirewallVip64(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
