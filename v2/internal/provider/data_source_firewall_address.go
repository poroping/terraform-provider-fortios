// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure IPv4 addresses.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceFirewallAddress() *schema.Resource {
	return &schema.Resource{
		Description: "Configure IPv4 addresses.",

		ReadContext: dataSourceFirewallAddressRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"allow_routing": {
				Type:        schema.TypeString,
				Description: "Enable/disable use of this address in the static route configuration.",
				Computed:    true,
			},
			"associated_interface": {
				Type:        schema.TypeString,
				Description: "Network interface associated with address.",
				Computed:    true,
			},
			"cache_ttl": {
				Type:        schema.TypeInt,
				Description: "Defines the minimal TTL of individual IP addresses in FQDN cache measured in seconds.",
				Computed:    true,
			},
			"clearpass_spt": {
				Type:        schema.TypeString,
				Description: "SPT (System Posture Token) value.",
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
			"country": {
				Type:        schema.TypeString,
				Description: "IP addresses associated to a specific country.",
				Computed:    true,
			},
			"end_ip": {
				Type:        schema.TypeString,
				Description: "Final IP address (inclusive) in the range for the address.",
				Computed:    true,
			},
			"end_mac": {
				Type:        schema.TypeString,
				Description: "Last MAC address in the range.",
				Computed:    true,
			},
			"epg_name": {
				Type:        schema.TypeString,
				Description: "Endpoint group name.",
				Computed:    true,
			},
			"fabric_object": {
				Type:        schema.TypeString,
				Description: "Security Fabric global object setting.",
				Computed:    true,
			},
			"filter": {
				Type:        schema.TypeString,
				Description: "Match criteria filter.",
				Computed:    true,
			},
			"fqdn": {
				Type:        schema.TypeString,
				Description: "Fully Qualified Domain Name address.",
				Computed:    true,
			},
			"fsso_group": {
				Type:        schema.TypeList,
				Description: "FSSO group(s).",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "FSSO group name.",
							Computed:    true,
						},
					},
				},
			},
			"interface": {
				Type:        schema.TypeString,
				Description: "Name of interface whose IP address is to be used.",
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
			"node_ip_only": {
				Type:        schema.TypeString,
				Description: "Enable/disable collection of node addresses only in Kubernetes.",
				Computed:    true,
			},
			"obj_id": {
				Type:        schema.TypeString,
				Description: "Object ID for NSX.",
				Computed:    true,
			},
			"obj_tag": {
				Type:        schema.TypeString,
				Description: "Tag of dynamic address object.",
				Computed:    true,
			},
			"obj_type": {
				Type:        schema.TypeString,
				Description: "Object type.",
				Computed:    true,
			},
			"organization": {
				Type:        schema.TypeString,
				Description: "Organization domain name (Syntax: organization/domain).",
				Computed:    true,
			},
			"policy_group": {
				Type:        schema.TypeString,
				Description: "Policy group name.",
				Computed:    true,
			},
			"sdn": {
				Type:        schema.TypeString,
				Description: "SDN.",
				Computed:    true,
			},
			"sdn_addr_type": {
				Type:        schema.TypeString,
				Description: "Type of addresses to collect.",
				Computed:    true,
			},
			"sdn_tag": {
				Type:        schema.TypeString,
				Description: "SDN Tag.",
				Computed:    true,
			},
			"start_ip": {
				Type:        schema.TypeString,
				Description: "First IP address (inclusive) in the range for the address.",
				Computed:    true,
			},
			"start_mac": {
				Type:        schema.TypeString,
				Description: "First MAC address in the range.",
				Computed:    true,
			},
			"sub_type": {
				Type:        schema.TypeString,
				Description: "Sub-type of address.",
				Computed:    true,
			},
			"subnet": {
				Type:        schema.TypeString,
				Description: "IP address and subnet mask of address.",
				Computed:    true,
			},
			"subnet_name": {
				Type:        schema.TypeString,
				Description: "Subnet name.",
				Computed:    true,
			},
			"tag_detection_level": {
				Type:        schema.TypeString,
				Description: "Tag detection level of dynamic address object.",
				Computed:    true,
			},
			"tag_type": {
				Type:        schema.TypeString,
				Description: "Tag type of dynamic address object.",
				Computed:    true,
			},
			"tagging": {
				Type:        schema.TypeList,
				Description: "Config object tagging.",
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
			"tenant": {
				Type:        schema.TypeString,
				Description: "Tenant.",
				Computed:    true,
			},
			"type": {
				Type:        schema.TypeString,
				Description: "Type of address.",
				Computed:    true,
			},
			"uuid": {
				Type:        schema.TypeString,
				Description: "Universally Unique Identifier (UUID; automatically assigned but can be manually reset).",
				Computed:    true,
			},
			"visibility": {
				Type:        schema.TypeString,
				Description: "Enable/disable address visibility in the GUI.",
				Computed:    true,
			},
			"wildcard": {
				Type:        schema.TypeString,
				Description: "IP address and wildcard netmask.",
				Computed:    true,
			},
			"wildcard_fqdn": {
				Type:        schema.TypeString,
				Description: "Fully Qualified Domain Name with wildcard characters.",
				Computed:    true,
			},
		},
	}
}

func dataSourceFirewallAddressRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallAddress(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallAddress dataSource: %v", err)
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

	diags := refreshObjectFirewallAddress(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
