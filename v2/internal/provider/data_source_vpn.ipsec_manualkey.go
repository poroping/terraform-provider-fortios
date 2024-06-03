// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure IPsec manual keys.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceVpnIpsecManualkey() *schema.Resource {
	return &schema.Resource{
		Description: "Configure IPsec manual keys.",

		ReadContext: dataSourceVpnIpsecManualkeyRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"authentication": {
				Type:        schema.TypeString,
				Description: "Authentication algorithm. Must be the same for both ends of the tunnel.",
				Computed:    true,
			},
			"authkey": {
				Type:        schema.TypeString,
				Description: "Hexadecimal authentication key in 16-digit (8-byte) segments separated by hyphens.",
				Computed:    true,
			},
			"enckey": {
				Type:        schema.TypeString,
				Description: "Hexadecimal encryption key in 16-digit (8-byte) segments separated by hyphens.",
				Computed:    true,
			},
			"encryption": {
				Type:        schema.TypeString,
				Description: "Encryption algorithm. Must be the same for both ends of the tunnel.",
				Computed:    true,
			},
			"interface": {
				Type:        schema.TypeString,
				Description: "Name of the physical, aggregate, or VLAN interface.",
				Computed:    true,
			},
			"local_gw": {
				Type:        schema.TypeString,
				Description: "Local gateway.",
				Computed:    true,
			},
			"localspi": {
				Type:        schema.TypeString,
				Description: "Local SPI, a hexadecimal 8-digit (4-byte) tag. Discerns between two traffic streams with different encryption rules.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "IPsec tunnel name.",
				Required:    true,
			},
			"npu_offload": {
				Type:        schema.TypeString,
				Description: "Enable/disable NPU offloading.",
				Computed:    true,
			},
			"remote_gw": {
				Type:        schema.TypeString,
				Description: "Peer gateway.",
				Computed:    true,
			},
			"remotespi": {
				Type:        schema.TypeString,
				Description: "Remote SPI, a hexadecimal 8-digit (4-byte) tag. Discerns between two traffic streams with different encryption rules.",
				Computed:    true,
			},
		},
	}
}

func dataSourceVpnIpsecManualkeyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadVpnIpsecManualkey(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading VpnIpsecManualkey dataSource: %v", err)
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

	diags := refreshObjectVpnIpsecManualkey(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
