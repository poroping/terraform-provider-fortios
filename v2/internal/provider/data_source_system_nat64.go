// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure NAT64.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceSystemNat64() *schema.Resource {
	return &schema.Resource{
		Description: "Configure NAT64.",

		ReadContext: dataSourceSystemNat64Read,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"always_synthesize_aaaa_record": {
				Type:        schema.TypeString,
				Description: "Enable/disable AAAA record synthesis (default = enable).",
				Computed:    true,
			},
			"generate_ipv6_fragment_header": {
				Type:        schema.TypeString,
				Description: "Enable/disable IPv6 fragment header generation.",
				Computed:    true,
			},
			"nat46_force_ipv4_packet_forwarding": {
				Type:        schema.TypeString,
				Description: "Enable/disable mandatory IPv4 packet forwarding in nat46.",
				Computed:    true,
			},
			"nat64_prefix": {
				Type:        schema.TypeString,
				Description: "NAT64 prefix must be ::/96 (default = 64:ff9b::/96).",
				Computed:    true,
			},
			"secondary_prefix": {
				Type:        schema.TypeList,
				Description: "Secondary NAT64 prefix.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "NAT64 prefix name.",
							Computed:    true,
						},
						"nat64_prefix": {
							Type:        schema.TypeString,
							Description: "NAT64 prefix.",
							Computed:    true,
						},
					},
				},
			},
			"secondary_prefix_status": {
				Type:        schema.TypeString,
				Description: "Enable/disable secondary NAT64 prefix.",
				Computed:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Enable/disable NAT64 (default = disable).",
				Computed:    true,
			},
		},
	}
}

func dataSourceSystemNat64Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemNat64(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemNat64 dataSource: %v", err)
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

	diags := refreshObjectSystemNat64(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}
