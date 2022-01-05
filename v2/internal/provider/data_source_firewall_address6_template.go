// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure IPv6 address templates.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceFirewallAddress6Template() *schema.Resource {
	return &schema.Resource{
		Description: "Configure IPv6 address templates.",

		ReadContext: dataSourceFirewallAddress6TemplateRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"fabric_object": {
				Type:        schema.TypeString,
				Description: "Security Fabric global object setting.",
				Computed:    true,
			},
			"ip6": {
				Type:        schema.TypeString,
				Description: "IPv6 address prefix.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "IPv6 address template name.",
				Required:    true,
			},
			"subnet_segment": {
				Type:        schema.TypeList,
				Description: "IPv6 subnet segments.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bits": {
							Type:        schema.TypeInt,
							Description: "Number of bits.",
							Computed:    true,
						},
						"exclusive": {
							Type:        schema.TypeString,
							Description: "Enable/disable exclusive value.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "Subnet segment ID.",
							Computed:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "Subnet segment name.",
							Computed:    true,
						},
						"values": {
							Type:        schema.TypeList,
							Description: "Subnet segment values.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Description: "Subnet segment value name.",
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
					},
				},
			},
			"subnet_segment_count": {
				Type:        schema.TypeInt,
				Description: "Number of IPv6 subnet segments.",
				Computed:    true,
			},
		},
	}
}

func dataSourceFirewallAddress6TemplateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallAddress6Template(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallAddress6Template dataSource: %v", err)
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

	diags := refreshObjectFirewallAddress6Template(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}
