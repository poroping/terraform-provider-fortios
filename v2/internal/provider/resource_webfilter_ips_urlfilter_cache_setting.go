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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func resourceWebfilterIpsUrlfilterCacheSetting() *schema.Resource {
	return &schema.Resource{
		Description: "Configure IPS URL filter cache settings.",

		CreateContext: resourceWebfilterIpsUrlfilterCacheSettingCreate,
		ReadContext:   resourceWebfilterIpsUrlfilterCacheSettingRead,
		UpdateContext: resourceWebfilterIpsUrlfilterCacheSettingUpdate,
		DeleteContext: resourceWebfilterIpsUrlfilterCacheSettingDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"dns_retry_interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 2147483),

				Description: "Retry interval. Refresh DNS faster than TTL to capture multiple IPs for hosts. 0 means use DNS server's TTL only.",
				Optional:    true,
				Computed:    true,
			},
			"extended_ttl": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 2147483),

				Description: "Extend time to live beyond reported by DNS. Use of 0 means use DNS server's TTL.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceWebfilterIpsUrlfilterCacheSettingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*apiClient).Client
	var diags diag.Diagnostics
	var err error
	// c.Retries = 1
	urlparams := &models.CmdbRequestParams{}
	vdomparam := ""
	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}
	urlparams.Vdom = vdomparam

	allow_append := false
	if v, ok := d.GetOk("allow_append"); ok {
		if b, ok := v.(bool); ok {
			allow_append = b
		}
	}
	urlparams.AllowAppend = &allow_append

	obj, diags := getObjectWebfilterIpsUrlfilterCacheSetting(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateWebfilterIpsUrlfilterCacheSetting(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WebfilterIpsUrlfilterCacheSetting")
	}

	return resourceWebfilterIpsUrlfilterCacheSettingRead(ctx, d, meta)
}

func resourceWebfilterIpsUrlfilterCacheSettingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectWebfilterIpsUrlfilterCacheSetting(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateWebfilterIpsUrlfilterCacheSetting(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating WebfilterIpsUrlfilterCacheSetting resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WebfilterIpsUrlfilterCacheSetting")
	}

	return resourceWebfilterIpsUrlfilterCacheSettingRead(ctx, d, meta)
}

func resourceWebfilterIpsUrlfilterCacheSettingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectWebfilterIpsUrlfilterCacheSetting(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateWebfilterIpsUrlfilterCacheSetting(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting WebfilterIpsUrlfilterCacheSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWebfilterIpsUrlfilterCacheSettingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	ptp := true
	urlparams.PlainTextPassword = &ptp

	o, err := c.Cmdb.ReadWebfilterIpsUrlfilterCacheSetting(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WebfilterIpsUrlfilterCacheSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
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
	return nil
}

func refreshObjectWebfilterIpsUrlfilterCacheSetting(d *schema.ResourceData, o *models.WebfilterIpsUrlfilterCacheSetting, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.DnsRetryInterval != nil {
		v := *o.DnsRetryInterval

		if err = d.Set("dns_retry_interval", v); err != nil {
			return diag.Errorf("error reading dns_retry_interval: %v", err)
		}
	}

	if o.ExtendedTtl != nil {
		v := *o.ExtendedTtl

		if err = d.Set("extended_ttl", v); err != nil {
			return diag.Errorf("error reading extended_ttl: %v", err)
		}
	}

	return nil
}

func getObjectWebfilterIpsUrlfilterCacheSetting(d *schema.ResourceData, sv string) (*models.WebfilterIpsUrlfilterCacheSetting, diag.Diagnostics) {
	obj := models.WebfilterIpsUrlfilterCacheSetting{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("dns_retry_interval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dns_retry_interval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.DnsRetryInterval = &tmp
		}
	}
	if v1, ok := d.GetOk("extended_ttl"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("extended_ttl", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.ExtendedTtl = &tmp
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectWebfilterIpsUrlfilterCacheSetting(d *schema.ResourceData, sv string) (*models.WebfilterIpsUrlfilterCacheSetting, diag.Diagnostics) {
	obj := models.WebfilterIpsUrlfilterCacheSetting{}
	diags := diag.Diagnostics{}

	return &obj, diags
}
