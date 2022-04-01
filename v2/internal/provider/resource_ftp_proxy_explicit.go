// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure explicit FTP proxy settings.

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

func resourceFtpProxyExplicit() *schema.Resource {
	return &schema.Resource{
		Description: "Configure explicit FTP proxy settings.",

		CreateContext: resourceFtpProxyExplicitCreate,
		ReadContext:   resourceFtpProxyExplicitRead,
		UpdateContext: resourceFtpProxyExplicitUpdate,
		DeleteContext: resourceFtpProxyExplicitDelete,

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
			"incoming_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "Accept incoming FTP requests from this IP address. An interface must have this IP address.",
				Optional:    true,
				Computed:    true,
			},
			"incoming_port": {
				Type: schema.TypeString,

				Description: "Accept incoming FTP requests on one or more ports.",
				Optional:    true,
				Computed:    true,
			},
			"outgoing_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "Outgoing FTP requests will leave from this IP address. An interface must have this IP address.",
				Optional:    true,
				Computed:    true,
			},
			"sec_default_action": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"accept", "deny"}, false),

				Description: "Accept or deny explicit FTP proxy sessions when no FTP proxy firewall policy exists.",
				Optional:    true,
				Computed:    true,
			},
			"ssl": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable the explicit FTPS proxy.",
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
			"ssl_dh_bits": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"768", "1024", "1536", "2048"}, false),

				Description: "Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048).",
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable the explicit FTP proxy.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceFtpProxyExplicitCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectFtpProxyExplicit(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateFtpProxyExplicit(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FtpProxyExplicit")
	}

	return resourceFtpProxyExplicitRead(ctx, d, meta)
}

func resourceFtpProxyExplicitUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectFtpProxyExplicit(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateFtpProxyExplicit(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating FtpProxyExplicit resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FtpProxyExplicit")
	}

	return resourceFtpProxyExplicitRead(ctx, d, meta)
}

func resourceFtpProxyExplicitDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectFtpProxyExplicit(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateFtpProxyExplicit(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting FtpProxyExplicit resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFtpProxyExplicitRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFtpProxyExplicit(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FtpProxyExplicit resource: %v", err)
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

	diags := refreshObjectFtpProxyExplicit(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectFtpProxyExplicit(d *schema.ResourceData, o *models.FtpProxyExplicit, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.IncomingIp != nil {
		v := *o.IncomingIp

		if err = d.Set("incoming_ip", v); err != nil {
			return diag.Errorf("error reading incoming_ip: %v", err)
		}
	}

	if o.IncomingPort != nil {
		v := *o.IncomingPort

		if err = d.Set("incoming_port", v); err != nil {
			return diag.Errorf("error reading incoming_port: %v", err)
		}
	}

	if o.OutgoingIp != nil {
		v := *o.OutgoingIp

		if err = d.Set("outgoing_ip", v); err != nil {
			return diag.Errorf("error reading outgoing_ip: %v", err)
		}
	}

	if o.SecDefaultAction != nil {
		v := *o.SecDefaultAction

		if err = d.Set("sec_default_action", v); err != nil {
			return diag.Errorf("error reading sec_default_action: %v", err)
		}
	}

	if o.Ssl != nil {
		v := *o.Ssl

		if err = d.Set("ssl", v); err != nil {
			return diag.Errorf("error reading ssl: %v", err)
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

	if o.SslDhBits != nil {
		v := *o.SslDhBits

		if err = d.Set("ssl_dh_bits", v); err != nil {
			return diag.Errorf("error reading ssl_dh_bits: %v", err)
		}
	}

	if o.Status != nil {
		v := *o.Status

		if err = d.Set("status", v); err != nil {
			return diag.Errorf("error reading status: %v", err)
		}
	}

	return nil
}

func getObjectFtpProxyExplicit(d *schema.ResourceData, sv string) (*models.FtpProxyExplicit, diag.Diagnostics) {
	obj := models.FtpProxyExplicit{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("incoming_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("incoming_ip", sv)
				diags = append(diags, e)
			}
			obj.IncomingIp = &v2
		}
	}
	if v1, ok := d.GetOk("incoming_port"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("incoming_port", sv)
				diags = append(diags, e)
			}
			obj.IncomingPort = &v2
		}
	}
	if v1, ok := d.GetOk("outgoing_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("outgoing_ip", sv)
				diags = append(diags, e)
			}
			obj.OutgoingIp = &v2
		}
	}
	if v1, ok := d.GetOk("sec_default_action"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("sec_default_action", sv)
				diags = append(diags, e)
			}
			obj.SecDefaultAction = &v2
		}
	}
	if v1, ok := d.GetOk("ssl"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssl", sv)
				diags = append(diags, e)
			}
			obj.Ssl = &v2
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
	if v1, ok := d.GetOk("ssl_dh_bits"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssl_dh_bits", sv)
				diags = append(diags, e)
			}
			obj.SslDhBits = &v2
		}
	}
	if v1, ok := d.GetOk("status"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("status", sv)
				diags = append(diags, e)
			}
			obj.Status = &v2
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectFtpProxyExplicit(d *schema.ResourceData, sv string) (*models.FtpProxyExplicit, diag.Diagnostics) {
	obj := models.FtpProxyExplicit{}
	diags := diag.Diagnostics{}

	return &obj, diags
}
