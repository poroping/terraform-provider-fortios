// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure SSL servers.

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

func resourceFirewallSslServer() *schema.Resource {
	return &schema.Resource{
		Description: "Configure SSL servers.",

		CreateContext: resourceFirewallSslServerCreate,
		ReadContext:   resourceFirewallSslServerRead,
		UpdateContext: resourceFirewallSslServerUpdate,
		DeleteContext: resourceFirewallSslServerDelete,

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
			"allow_append": {
				Type:        schema.TypeBool,
				Description: "If set the provider will overwrite existing resources with the same mkey instead of erroring. Useful for brownfield implementations. Use with caution!",
				Optional:    true,
				Default:     false,
			},
			"add_header_x_forwarded_proto": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable adding an X-Forwarded-Proto header to forwarded requests.",
				Optional:    true,
				Computed:    true,
			},
			"ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "IPv4 address of the SSL server.",
				Optional:    true,
				Computed:    true,
			},
			"mapped_port": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),

				Description: "Mapped server service port (1 - 65535, default = 80).",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Server name.",
				ForceNew:    true,
				Required:    true,
			},
			"port": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),

				Description: "Server service port (1 - 65535, default = 443).",
				Optional:    true,
				Computed:    true,
			},
			"ssl_algorithm": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"high", "medium", "low"}, false),

				Description: "Relative strength of encryption algorithms accepted in negotiation.",
				Optional:    true,
				Computed:    true,
			},
			"ssl_cert": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of certificate for SSL connections to this server (default = \"Fortinet_CA_SSL\").",
				Optional:    true,
				Computed:    true,
			},
			"ssl_client_renegotiation": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"allow", "deny", "secure"}, false),

				Description: "Allow or block client renegotiation by server.",
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
			"ssl_max_version": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"tls-1.0", "tls-1.1", "tls-1.2", "tls-1.3"}, false),

				Description: "Highest SSL/TLS version to negotiate.",
				Optional:    true,
				Computed:    true,
			},
			"ssl_min_version": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"tls-1.0", "tls-1.1", "tls-1.2", "tls-1.3"}, false),

				Description: "Lowest SSL/TLS version to negotiate.",
				Optional:    true,
				Computed:    true,
			},
			"ssl_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"half", "full"}, false),

				Description: "SSL/TLS mode for encryption and decryption of traffic.",
				Optional:    true,
				Computed:    true,
			},
			"ssl_send_empty_frags": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable sending empty fragments to avoid attack on CBC IV.",
				Optional:    true,
				Computed:    true,
			},
			"url_rewrite": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable rewriting the URL.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceFirewallSslServerCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := ""
	key := "name"
	if v, ok := d.GetOk(key); ok {
		mkey = utils.ParseMkey(v)
		if mkey == "" && allow_append {
			return diag.Errorf("error creating FirewallSslServer resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectFirewallSslServer(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateFirewallSslServer(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallSslServer")
	}

	return resourceFirewallSslServerRead(ctx, d, meta)
}

func resourceFirewallSslServerUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectFirewallSslServer(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateFirewallSslServer(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating FirewallSslServer resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallSslServer")
	}

	return resourceFirewallSslServerRead(ctx, d, meta)
}

func resourceFirewallSslServerDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteFirewallSslServer(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting FirewallSslServer resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallSslServerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallSslServer(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallSslServer resource: %v", err)
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

	diags := refreshObjectFirewallSslServer(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectFirewallSslServer(d *schema.ResourceData, o *models.FirewallSslServer, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AddHeaderXForwardedProto != nil {
		v := *o.AddHeaderXForwardedProto

		if err = d.Set("add_header_x_forwarded_proto", v); err != nil {
			return diag.Errorf("error reading add_header_x_forwarded_proto: %v", err)
		}
	}

	if o.Ip != nil {
		v := *o.Ip

		if err = d.Set("ip", v); err != nil {
			return diag.Errorf("error reading ip: %v", err)
		}
	}

	if o.MappedPort != nil {
		v := *o.MappedPort

		if err = d.Set("mapped_port", v); err != nil {
			return diag.Errorf("error reading mapped_port: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.Port != nil {
		v := *o.Port

		if err = d.Set("port", v); err != nil {
			return diag.Errorf("error reading port: %v", err)
		}
	}

	if o.SslAlgorithm != nil {
		v := *o.SslAlgorithm

		if err = d.Set("ssl_algorithm", v); err != nil {
			return diag.Errorf("error reading ssl_algorithm: %v", err)
		}
	}

	if o.SslCert != nil {
		v := *o.SslCert

		if err = d.Set("ssl_cert", v); err != nil {
			return diag.Errorf("error reading ssl_cert: %v", err)
		}
	}

	if o.SslClientRenegotiation != nil {
		v := *o.SslClientRenegotiation

		if err = d.Set("ssl_client_renegotiation", v); err != nil {
			return diag.Errorf("error reading ssl_client_renegotiation: %v", err)
		}
	}

	if o.SslDhBits != nil {
		v := *o.SslDhBits

		if err = d.Set("ssl_dh_bits", v); err != nil {
			return diag.Errorf("error reading ssl_dh_bits: %v", err)
		}
	}

	if o.SslMaxVersion != nil {
		v := *o.SslMaxVersion

		if err = d.Set("ssl_max_version", v); err != nil {
			return diag.Errorf("error reading ssl_max_version: %v", err)
		}
	}

	if o.SslMinVersion != nil {
		v := *o.SslMinVersion

		if err = d.Set("ssl_min_version", v); err != nil {
			return diag.Errorf("error reading ssl_min_version: %v", err)
		}
	}

	if o.SslMode != nil {
		v := *o.SslMode

		if err = d.Set("ssl_mode", v); err != nil {
			return diag.Errorf("error reading ssl_mode: %v", err)
		}
	}

	if o.SslSendEmptyFrags != nil {
		v := *o.SslSendEmptyFrags

		if err = d.Set("ssl_send_empty_frags", v); err != nil {
			return diag.Errorf("error reading ssl_send_empty_frags: %v", err)
		}
	}

	if o.UrlRewrite != nil {
		v := *o.UrlRewrite

		if err = d.Set("url_rewrite", v); err != nil {
			return diag.Errorf("error reading url_rewrite: %v", err)
		}
	}

	return nil
}

func getObjectFirewallSslServer(d *schema.ResourceData, sv string) (*models.FirewallSslServer, diag.Diagnostics) {
	obj := models.FirewallSslServer{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("add_header_x_forwarded_proto"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("add_header_x_forwarded_proto", sv)
				diags = append(diags, e)
			}
			obj.AddHeaderXForwardedProto = &v2
		}
	}
	if v1, ok := d.GetOk("ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ip", sv)
				diags = append(diags, e)
			}
			obj.Ip = &v2
		}
	}
	if v1, ok := d.GetOk("mapped_port"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("mapped_port", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MappedPort = &tmp
		}
	}
	if v1, ok := d.GetOk("name"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("name", sv)
				diags = append(diags, e)
			}
			obj.Name = &v2
		}
	}
	if v1, ok := d.GetOk("port"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("port", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Port = &tmp
		}
	}
	if v1, ok := d.GetOk("ssl_algorithm"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssl_algorithm", sv)
				diags = append(diags, e)
			}
			obj.SslAlgorithm = &v2
		}
	}
	if v1, ok := d.GetOk("ssl_cert"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssl_cert", sv)
				diags = append(diags, e)
			}
			obj.SslCert = &v2
		}
	}
	if v1, ok := d.GetOk("ssl_client_renegotiation"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssl_client_renegotiation", sv)
				diags = append(diags, e)
			}
			obj.SslClientRenegotiation = &v2
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
	if v1, ok := d.GetOk("ssl_max_version"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssl_max_version", sv)
				diags = append(diags, e)
			}
			obj.SslMaxVersion = &v2
		}
	}
	if v1, ok := d.GetOk("ssl_min_version"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssl_min_version", sv)
				diags = append(diags, e)
			}
			obj.SslMinVersion = &v2
		}
	}
	if v1, ok := d.GetOk("ssl_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssl_mode", sv)
				diags = append(diags, e)
			}
			obj.SslMode = &v2
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
	if v1, ok := d.GetOk("url_rewrite"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("url_rewrite", sv)
				diags = append(diags, e)
			}
			obj.UrlRewrite = &v2
		}
	}
	return &obj, diags
}
