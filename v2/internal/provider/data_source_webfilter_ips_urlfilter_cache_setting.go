// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure IPS URL filter cache settings.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceWebfilterIpsUrlfilterCacheSetting() *schema.Resource {
	return &schema.Resource{
		Description: "Configure IPS URL filter cache settings.",

		ReadContext: dataSourceWebfilterIpsUrlfilterCacheSettingRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"dns_retry_interval": {
				Type:        schema.TypeInt,
				Description: "Retry interval. Refresh DNS faster than TTL to capture multiple IPs for hosts. 0 means use DNS server's TTL only.",
				Computed:    true,
			},
			"extended_ttl": {
				Type:        schema.TypeInt,
				Description: "Extend time to live beyond reported by DNS. Use of 0 means use DNS server's TTL.",
				Computed:    true,
			},
		},
	}
}

func dataSourceWebfilterIpsUrlfilterCacheSettingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "WebfilterIpsUrlfilterCacheSetting"

	o, err := c.Cmdb.ReadWebfilterIpsUrlfilterCacheSetting(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WebfilterIpsUrlfilterCacheSetting dataSource: %v", err)
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

	diags := refreshObjectWebfilterIpsUrlfilterCacheSetting(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
