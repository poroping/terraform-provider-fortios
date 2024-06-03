// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure Forward Error Correction (FEC) mapping profiles.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceVpnIpsecFec() *schema.Resource {
	return &schema.Resource{
		Description: "Configure Forward Error Correction (FEC) mapping profiles.",

		ReadContext: dataSourceVpnIpsecFecRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"mappings": {
				Type:        schema.TypeList,
				Description: "FEC redundancy mapping table.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bandwidth_bi_threshold": {
							Type:        schema.TypeInt,
							Description: "Apply FEC parameters when available bi-bandwidth is >= threshold (kbps, 0 means no threshold).",
							Computed:    true,
						},
						"bandwidth_down_threshold": {
							Type:        schema.TypeInt,
							Description: "Apply FEC parameters when available down bandwidth is >= threshold (kbps, 0 means no threshold).",
							Computed:    true,
						},
						"bandwidth_up_threshold": {
							Type:        schema.TypeInt,
							Description: "Apply FEC parameters when available up bandwidth is >= threshold (kbps, 0 means no threshold).",
							Computed:    true,
						},
						"base": {
							Type:        schema.TypeInt,
							Description: "Number of base FEC packets (1 - 20).",
							Computed:    true,
						},
						"latency_threshold": {
							Type:        schema.TypeInt,
							Description: "Apply FEC parameters when latency is <= threshold (0 means no threshold).",
							Computed:    true,
						},
						"packet_loss_threshold": {
							Type:        schema.TypeInt,
							Description: "Apply FEC parameters when packet loss is >= threshold (0 - 100, 0 means no threshold).",
							Computed:    true,
						},
						"redundant": {
							Type:        schema.TypeInt,
							Description: "Number of redundant FEC packets (1 - 5).",
							Computed:    true,
						},
						"seqno": {
							Type:        schema.TypeInt,
							Description: "Sequence number (1 - 64).",
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Profile name.",
				Required:    true,
			},
		},
	}
}

func dataSourceVpnIpsecFecRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadVpnIpsecFec(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading VpnIpsecFec dataSource: %v", err)
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

	diags := refreshObjectVpnIpsecFec(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
