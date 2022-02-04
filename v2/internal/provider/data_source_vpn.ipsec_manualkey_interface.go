// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
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

func dataSourceVpnIpsecManualkeyInterface() *schema.Resource {
	return &schema.Resource{
		Description: "Configure IPsec manual keys.",

		ReadContext: dataSourceVpnIpsecManualkeyInterfaceRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"addr_type": {
				Type:        schema.TypeString,
				Description: "IP version to use for IP packets.",
				Computed:    true,
			},
			"auth_alg": {
				Type:        schema.TypeString,
				Description: "Authentication algorithm. Must be the same for both ends of the tunnel.",
				Computed:    true,
			},
			"auth_key": {
				Type:        schema.TypeString,
				Description: "Hexadecimal authentication key in 16-digit (8-byte) segments separated by hyphens.",
				Computed:    true,
			},
			"enc_alg": {
				Type:        schema.TypeString,
				Description: "Encryption algorithm. Must be the same for both ends of the tunnel.",
				Computed:    true,
			},
			"enc_key": {
				Type:        schema.TypeString,
				Description: "Hexadecimal encryption key in 16-digit (8-byte) segments separated by hyphens.",
				Computed:    true,
			},
			"interface": {
				Type:        schema.TypeString,
				Description: "Name of the physical, aggregate, or VLAN interface.",
				Computed:    true,
			},
			"ip_version": {
				Type:        schema.TypeString,
				Description: "IP version to use for VPN interface.",
				Computed:    true,
			},
			"local_gw": {
				Type:        schema.TypeString,
				Description: "IPv4 address of the local gateway's external interface.",
				Computed:    true,
			},
			"local_gw6": {
				Type:        schema.TypeString,
				Description: "Local IPv6 address of VPN gateway.",
				Computed:    true,
			},
			"local_spi": {
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
				Description: "Enable/disable offloading IPsec VPN manual key sessions to NPUs.",
				Computed:    true,
			},
			"remote_gw": {
				Type:        schema.TypeString,
				Description: "IPv4 address of the remote gateway's external interface.",
				Computed:    true,
			},
			"remote_gw6": {
				Type:        schema.TypeString,
				Description: "Remote IPv6 address of VPN gateway.",
				Computed:    true,
			},
			"remote_spi": {
				Type:        schema.TypeString,
				Description: "Remote SPI, a hexadecimal 8-digit (4-byte) tag. Discerns between two traffic streams with different encryption rules.",
				Computed:    true,
			},
		},
	}
}

func dataSourceVpnIpsecManualkeyInterfaceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadVpnIpsecManualkeyInterface(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading VpnIpsecManualkeyInterface dataSource: %v", err)
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

	diags := refreshObjectVpnIpsecManualkeyInterface(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
