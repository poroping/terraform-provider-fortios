// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure IPv6 firewall addresses.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceFirewallAddress6() *schema.Resource {
	return &schema.Resource{
		Description: "Configure IPv6 firewall addresses.",

		ReadContext: dataSourceFirewallAddress6Read,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"cache_ttl": {
				Type:        schema.TypeInt,
				Description: "Minimal TTL of individual IPv6 addresses in FQDN cache.",
				Computed:    true,
			},
			"color": {
				Type:        schema.TypeInt,
				Description: "Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).",
				Computed:    true,
			},
			"comment": {
				Type:        schema.TypeString,
				Description: "Comment.",
				Computed:    true,
			},
			"country": {
				Type:        schema.TypeString,
				Description: "IPv6 addresses associated to a specific country.",
				Computed:    true,
			},
			"end_ip": {
				Type:        schema.TypeString,
				Description: "Final IP address (inclusive) in the range for the address (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).",
				Computed:    true,
			},
			"end_mac": {
				Type:        schema.TypeString,
				Description: "Last MAC address in the range.",
				Computed:    true,
			},
			"fabric_object": {
				Type:        schema.TypeString,
				Description: "Security Fabric global object setting.",
				Computed:    true,
			},
			"fqdn": {
				Type:        schema.TypeString,
				Description: "Fully qualified domain name.",
				Computed:    true,
			},
			"host": {
				Type:        schema.TypeString,
				Description: "Host Address.",
				Computed:    true,
			},
			"host_type": {
				Type:        schema.TypeString,
				Description: "Host type.",
				Computed:    true,
			},
			"ip6": {
				Type:        schema.TypeString,
				Description: "IPv6 address prefix (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx).",
				Computed:    true,
			},
			"list": {
				Type:        schema.TypeList,
				Description: "IP address list.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip": {
							Type:        schema.TypeString,
							Description: "IP.",
							Computed:    true,
						},
					},
				},
			},
			"macaddr": {
				Type:        schema.TypeList,
				Description: "Multiple MAC address ranges.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"macaddr": {
							Type:        schema.TypeString,
							Description: "MAC address ranges <start>[-<end>] separated by space.",
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Address name.",
				Required:    true,
			},
			"obj_id": {
				Type:        schema.TypeString,
				Description: "Object ID for NSX.",
				Computed:    true,
			},
			"sdn": {
				Type:        schema.TypeString,
				Description: "SDN.",
				Computed:    true,
			},
			"start_ip": {
				Type:        schema.TypeString,
				Description: "First IP address (inclusive) in the range for the address (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).",
				Computed:    true,
			},
			"start_mac": {
				Type:        schema.TypeString,
				Description: "First MAC address in the range.",
				Computed:    true,
			},
			"subnet_segment": {
				Type:        schema.TypeList,
				Description: "IPv6 subnet segments.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Name.",
							Computed:    true,
						},
						"type": {
							Type:        schema.TypeString,
							Description: "Subnet segment type.",
							Computed:    true,
						},
						"value": {
							Type:        schema.TypeString,
							Description: "Subnet segment value.",
							Computed:    true,
						},
					},
				},
			},
			"tagging": {
				Type:        schema.TypeList,
				Description: "Config object tagging",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"category": {
							Type:        schema.TypeString,
							Description: "Tag category.",
							Computed:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "Tagging entry name.",
							Computed:    true,
						},
						"tags": {
							Type:        schema.TypeList,
							Description: "Tags.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Description: "Tag name.",
										Computed:    true,
									},
								},
							},
						},
					},
				},
			},
			"template": {
				Type:        schema.TypeString,
				Description: "IPv6 address template.",
				Computed:    true,
			},
			"type": {
				Type:        schema.TypeString,
				Description: "Type of IPv6 address object (default = ipprefix).",
				Computed:    true,
			},
			"uuid": {
				Type:        schema.TypeString,
				Description: "Universally Unique Identifier (UUID; automatically assigned but can be manually reset).",
				Computed:    true,
			},
			"visibility": {
				Type:        schema.TypeString,
				Description: "Enable/disable the visibility of the object in the GUI.",
				Computed:    true,
			},
		},
	}
}

func dataSourceFirewallAddress6Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	mkey := d.Id()

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

	o, err := c.Cmdb.ReadFirewallAddress6(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallAddress6 dataSource: %v", err)
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

	diags := refreshObjectFirewallAddress6(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}
