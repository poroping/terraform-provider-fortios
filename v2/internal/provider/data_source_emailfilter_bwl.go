// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure anti-spam black/white list.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceEmailfilterBwl() *schema.Resource {
	return &schema.Resource{
		Description: "Configure anti-spam black/white list.",

		ReadContext: dataSourceEmailfilterBwlRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"comment": {
				Type:        schema.TypeString,
				Description: "Optional comments.",
				Computed:    true,
			},
			"entries": {
				Type:        schema.TypeList,
				Description: "Anti-spam black/white list entries.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:        schema.TypeString,
							Description: "Reject, mark as spam or good email.",
							Computed:    true,
						},
						"addr_type": {
							Type:        schema.TypeString,
							Description: "IP address type.",
							Computed:    true,
						},
						"email_pattern": {
							Type:        schema.TypeString,
							Description: "Email address pattern.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "Entry ID.",
							Computed:    true,
						},
						"ip4_subnet": {
							Type:        schema.TypeString,
							Description: "IPv4 network address/subnet mask bits.",
							Computed:    true,
						},
						"ip6_subnet": {
							Type:        schema.TypeString,
							Description: "IPv6 network address/subnet mask bits.",
							Computed:    true,
						},
						"pattern_type": {
							Type:        schema.TypeString,
							Description: "Wildcard pattern or regular expression.",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Enable/disable status.",
							Computed:    true,
						},
						"type": {
							Type:        schema.TypeString,
							Description: "Entry type.",
							Computed:    true,
						},
					},
				},
			},
			"fosid": {
				Type:        schema.TypeInt,
				Description: "ID.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Name of table.",
				Computed:    true,
			},
		},
	}
}

func dataSourceEmailfilterBwlRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadEmailfilterBwl(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading EmailfilterBwl dataSource: %v", err)
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

	diags := refreshObjectEmailfilterBwl(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
