// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure WiFi bridge access control list.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceWirelessControllerAccessControlList() *schema.Resource {
	return &schema.Resource{
		Description: "Configure WiFi bridge access control list.",

		ReadContext: dataSourceWirelessControllerAccessControlListRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"comment": {
				Type:        schema.TypeString,
				Description: "Description.",
				Computed:    true,
			},
			"layer3_ipv4_rules": {
				Type:        schema.TypeList,
				Description: "AP ACL layer3 ipv4 rule list.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:        schema.TypeString,
							Description: "Policy action (allow | deny).",
							Computed:    true,
						},
						"comment": {
							Type:        schema.TypeString,
							Description: "Description.",
							Computed:    true,
						},
						"dstaddr": {
							Type:        schema.TypeString,
							Description: "Destination IP address (any | local-LAN | IPv4 address[/<network mask | mask length>], default = any).",
							Computed:    true,
						},
						"dstport": {
							Type:        schema.TypeInt,
							Description: "Destination port (0 - 65535, default = 0, meaning any).",
							Computed:    true,
						},
						"protocol": {
							Type:        schema.TypeInt,
							Description: "Protocol type as defined by IANA (0 - 255, default = 255, meaning any).",
							Computed:    true,
						},
						"rule_id": {
							Type:        schema.TypeInt,
							Description: "Rule ID (1 - 65535).",
							Computed:    true,
						},
						"srcaddr": {
							Type:        schema.TypeString,
							Description: "Source IP address (any | local-LAN | IPv4 address[/<network mask | mask length>], default = any).",
							Computed:    true,
						},
						"srcport": {
							Type:        schema.TypeInt,
							Description: "Source port (0 - 65535, default = 0, meaning any).",
							Computed:    true,
						},
					},
				},
			},
			"layer3_ipv6_rules": {
				Type:        schema.TypeList,
				Description: "AP ACL layer3 ipv6 rule list.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:        schema.TypeString,
							Description: "Policy action (allow | deny).",
							Computed:    true,
						},
						"comment": {
							Type:        schema.TypeString,
							Description: "Description.",
							Computed:    true,
						},
						"dstaddr": {
							Type:        schema.TypeString,
							Description: "Destination IPv6 address (any | local-LAN | IPv6 address[/prefix length]), default = any.",
							Computed:    true,
						},
						"dstport": {
							Type:        schema.TypeInt,
							Description: "Destination port (0 - 65535, default = 0, meaning any).",
							Computed:    true,
						},
						"protocol": {
							Type:        schema.TypeInt,
							Description: "Protocol type as defined by IANA (0 - 255, default = 255, meaning any).",
							Computed:    true,
						},
						"rule_id": {
							Type:        schema.TypeInt,
							Description: "Rule ID (1 - 65535).",
							Computed:    true,
						},
						"srcaddr": {
							Type:        schema.TypeString,
							Description: "Source IPv6 address (any | local-LAN | IPv6 address[/prefix length]), default = any.",
							Computed:    true,
						},
						"srcport": {
							Type:        schema.TypeInt,
							Description: "Source port (0 - 65535, default = 0, meaning any).",
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:        schema.TypeString,
				Description: "AP access control list name.",
				Required:    true,
			},
		},
	}
}

func dataSourceWirelessControllerAccessControlListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadWirelessControllerAccessControlList(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WirelessControllerAccessControlList dataSource: %v", err)
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

	diags := refreshObjectWirelessControllerAccessControlList(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
