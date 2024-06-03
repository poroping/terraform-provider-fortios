// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: SSL proxy settings.

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

func resourceFirewallSslSetting() *schema.Resource {
	return &schema.Resource{
		Description: "SSL proxy settings.",

		CreateContext: resourceFirewallSslSettingCreate,
		ReadContext:   resourceFirewallSslSettingRead,
		UpdateContext: resourceFirewallSslSettingUpdate,
		DeleteContext: resourceFirewallSslSettingDelete,

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
			"abbreviate_handshake": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable use of SSL abbreviated handshake.",
				Optional:    true,
				Computed:    true,
			},
			"cert_cache_capacity": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 500),

				Description: "Maximum capacity of the host certificate cache (0 - 500, default = 200).",
				Optional:    true,
				Computed:    true,
			},
			"cert_cache_timeout": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 120),

				Description: "Time limit to keep certificate cache (1 - 120 min, default = 10).",
				Optional:    true,
				Computed:    true,
			},
			"kxp_queue_threshold": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 512),

				Description: "Maximum length of the CP KXP queue. When the queue becomes full, the proxy switches cipher functions to the main CPU (0 - 512, default = 16).",
				Optional:    true,
				Computed:    true,
			},
			"no_matching_cipher_action": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"bypass", "drop"}, false),

				Description: "Bypass or drop the connection when no matching cipher is found.",
				Optional:    true,
				Computed:    true,
			},
			"proxy_connect_timeout": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 60),

				Description: "Time limit to make an internal connection to the appropriate proxy process (1 - 60 sec, default = 30).",
				Optional:    true,
				Computed:    true,
			},
			"session_cache_capacity": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 1000),

				Description: "Capacity of the SSL session cache (--Obsolete--) (1 - 1000, default = 500).",
				Optional:    true,
				Computed:    true,
			},
			"session_cache_timeout": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 60),

				Description: "Time limit to keep SSL session state (1 - 60 min, default = 20).",
				Optional:    true,
				Computed:    true,
			},
			"ssl_dh_bits": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"768", "1024", "1536", "2048"}, false),

				Description: "Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048).",
				Optional:    true,
				Computed:    true,
			},
			"ssl_queue_threshold": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 512),

				Description: "Maximum length of the CP SSL queue. When the queue becomes full, the proxy switches cipher functions to the main CPU (0 - 512, default = 32).",
				Optional:    true,
				Computed:    true,
			},
			"ssl_send_empty_frags": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable sending empty fragments to avoid attack on CBC IV (for SSL 3.0 and TLS 1.0 only).",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceFirewallSslSettingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectFirewallSslSetting(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateFirewallSslSetting(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallSslSetting")
	}

	return resourceFirewallSslSettingRead(ctx, d, meta)
}

func resourceFirewallSslSettingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectFirewallSslSetting(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateFirewallSslSetting(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating FirewallSslSetting resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallSslSetting")
	}

	return resourceFirewallSslSettingRead(ctx, d, meta)
}

func resourceFirewallSslSettingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectFirewallSslSetting(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateFirewallSslSetting(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting FirewallSslSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallSslSettingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallSslSetting(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallSslSetting resource: %v", err)
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

	diags := refreshObjectFirewallSslSetting(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectFirewallSslSetting(d *schema.ResourceData, o *models.FirewallSslSetting, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AbbreviateHandshake != nil {
		v := *o.AbbreviateHandshake

		if err = d.Set("abbreviate_handshake", v); err != nil {
			return diag.Errorf("error reading abbreviate_handshake: %v", err)
		}
	}

	if o.CertCacheCapacity != nil {
		v := *o.CertCacheCapacity

		if err = d.Set("cert_cache_capacity", v); err != nil {
			return diag.Errorf("error reading cert_cache_capacity: %v", err)
		}
	}

	if o.CertCacheTimeout != nil {
		v := *o.CertCacheTimeout

		if err = d.Set("cert_cache_timeout", v); err != nil {
			return diag.Errorf("error reading cert_cache_timeout: %v", err)
		}
	}

	if o.KxpQueueThreshold != nil {
		v := *o.KxpQueueThreshold

		if err = d.Set("kxp_queue_threshold", v); err != nil {
			return diag.Errorf("error reading kxp_queue_threshold: %v", err)
		}
	}

	if o.NoMatchingCipherAction != nil {
		v := *o.NoMatchingCipherAction

		if err = d.Set("no_matching_cipher_action", v); err != nil {
			return diag.Errorf("error reading no_matching_cipher_action: %v", err)
		}
	}

	if o.ProxyConnectTimeout != nil {
		v := *o.ProxyConnectTimeout

		if err = d.Set("proxy_connect_timeout", v); err != nil {
			return diag.Errorf("error reading proxy_connect_timeout: %v", err)
		}
	}

	if o.SessionCacheCapacity != nil {
		v := *o.SessionCacheCapacity

		if err = d.Set("session_cache_capacity", v); err != nil {
			return diag.Errorf("error reading session_cache_capacity: %v", err)
		}
	}

	if o.SessionCacheTimeout != nil {
		v := *o.SessionCacheTimeout

		if err = d.Set("session_cache_timeout", v); err != nil {
			return diag.Errorf("error reading session_cache_timeout: %v", err)
		}
	}

	if o.SslDhBits != nil {
		v := *o.SslDhBits

		if err = d.Set("ssl_dh_bits", v); err != nil {
			return diag.Errorf("error reading ssl_dh_bits: %v", err)
		}
	}

	if o.SslQueueThreshold != nil {
		v := *o.SslQueueThreshold

		if err = d.Set("ssl_queue_threshold", v); err != nil {
			return diag.Errorf("error reading ssl_queue_threshold: %v", err)
		}
	}

	if o.SslSendEmptyFrags != nil {
		v := *o.SslSendEmptyFrags

		if err = d.Set("ssl_send_empty_frags", v); err != nil {
			return diag.Errorf("error reading ssl_send_empty_frags: %v", err)
		}
	}

	return nil
}

func getObjectFirewallSslSetting(d *schema.ResourceData, sv string) (*models.FirewallSslSetting, diag.Diagnostics) {
	obj := models.FirewallSslSetting{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("abbreviate_handshake"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("abbreviate_handshake", sv)
				diags = append(diags, e)
			}
			obj.AbbreviateHandshake = &v2
		}
	}
	if v1, ok := d.GetOk("cert_cache_capacity"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("cert_cache_capacity", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.CertCacheCapacity = &tmp
		}
	}
	if v1, ok := d.GetOk("cert_cache_timeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("cert_cache_timeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.CertCacheTimeout = &tmp
		}
	}
	if v1, ok := d.GetOk("kxp_queue_threshold"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "v6.4.0") {
				e := utils.AttributeVersionWarning("kxp_queue_threshold", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.KxpQueueThreshold = &tmp
		}
	}
	if v1, ok := d.GetOk("no_matching_cipher_action"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("no_matching_cipher_action", sv)
				diags = append(diags, e)
			}
			obj.NoMatchingCipherAction = &v2
		}
	}
	if v1, ok := d.GetOk("proxy_connect_timeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("proxy_connect_timeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.ProxyConnectTimeout = &tmp
		}
	}
	if v1, ok := d.GetOk("session_cache_capacity"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("session_cache_capacity", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.SessionCacheCapacity = &tmp
		}
	}
	if v1, ok := d.GetOk("session_cache_timeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("session_cache_timeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.SessionCacheTimeout = &tmp
		}
	}
	if v1, ok := d.GetOk("ssl_dh_bits"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssl_dh_bits", sv)
				diags = append(diags, e)
			}
			obj.SslDhBits = &v2
		}
	}
	if v1, ok := d.GetOk("ssl_queue_threshold"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "v6.4.0") {
				e := utils.AttributeVersionWarning("ssl_queue_threshold", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.SslQueueThreshold = &tmp
		}
	}
	if v1, ok := d.GetOk("ssl_send_empty_frags"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssl_send_empty_frags", sv)
				diags = append(diags, e)
			}
			obj.SslSendEmptyFrags = &v2
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectFirewallSslSetting(d *schema.ResourceData, sv string) (*models.FirewallSslSetting, diag.Diagnostics) {
	obj := models.FirewallSslSetting{}
	diags := diag.Diagnostics{}

	return &obj, diags
}
