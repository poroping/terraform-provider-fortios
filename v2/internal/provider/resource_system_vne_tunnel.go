// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure virtual network enabler tunnel.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
	"github.com/poroping/terraform-provider-fortios/v2/validators"
)

func resourceSystemVneTunnel() *schema.Resource {
	return &schema.Resource{
		Description: "Configure virtual network enabler tunnel.",

		CreateContext: resourceSystemVneTunnelCreate,
		ReadContext:   resourceSystemVneTunnelRead,
		UpdateContext: resourceSystemVneTunnelUpdate,
		DeleteContext: resourceSystemVneTunnelDelete,

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
			"bmr_hostname": {
				Type: schema.TypeString,

				Description: "BMR hostname.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"br": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "IPv6 address or FQDN of the border relay.",
				Optional:    true,
				Computed:    true,
			},
			"http_password": {
				Type: schema.TypeString,

				Description: "HTTP authentication password.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"http_username": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 64),

				Description: "HTTP authentication user name.",
				Optional:    true,
				Computed:    true,
			},
			"interface": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "Interface name.",
				Optional:    true,
				Computed:    true,
			},
			"ipv4_address": {
				Type:         schema.TypeString,
				ValidateFunc: validators.FortiValidateIPv4ClassnetHost,

				Description: "Tunnel IPv4 address and netmask.",
				Optional:    true,
				Computed:    true,
			},
			"mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"map-e", "fixed-ip", "ds-lite"}, false),

				Description: "VNE tunnel mode.",
				Optional:    true,
				Computed:    true,
			},
			"ssl_certificate": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of local certificate for SSL connections.",
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable VNE tunnel.",
				Optional:    true,
				Computed:    true,
			},
			"update_url": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 511),

				Description: "URL of provisioning server.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSystemVneTunnelCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemVneTunnel(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemVneTunnel(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemVneTunnel")
	}

	return resourceSystemVneTunnelRead(ctx, d, meta)
}

func resourceSystemVneTunnelUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemVneTunnel(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemVneTunnel(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemVneTunnel resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemVneTunnel")
	}

	return resourceSystemVneTunnelRead(ctx, d, meta)
}

func resourceSystemVneTunnelDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectSystemVneTunnel(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateSystemVneTunnel(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemVneTunnel resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemVneTunnelRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemVneTunnel(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemVneTunnel resource: %v", err)
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

	diags := refreshObjectSystemVneTunnel(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectSystemVneTunnel(d *schema.ResourceData, o *models.SystemVneTunnel, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.BmrHostname != nil {
		v := *o.BmrHostname

		if v == "ENC XXXX" {
		} else if err = d.Set("bmr_hostname", v); err != nil {
			return diag.Errorf("error reading bmr_hostname: %v", err)
		}
	}

	if o.Br != nil {
		v := *o.Br

		if err = d.Set("br", v); err != nil {
			return diag.Errorf("error reading br: %v", err)
		}
	}

	if o.HttpPassword != nil {
		v := *o.HttpPassword

		if v == "ENC XXXX" {
		} else if err = d.Set("http_password", v); err != nil {
			return diag.Errorf("error reading http_password: %v", err)
		}
	}

	if o.HttpUsername != nil {
		v := *o.HttpUsername

		if err = d.Set("http_username", v); err != nil {
			return diag.Errorf("error reading http_username: %v", err)
		}
	}

	if o.Interface != nil {
		v := *o.Interface

		if err = d.Set("interface", v); err != nil {
			return diag.Errorf("error reading interface: %v", err)
		}
	}

	if o.Ipv4Address != nil {
		v := *o.Ipv4Address
		if current, ok := d.GetOk("ipv4_address"); ok {
			if s, ok := current.(string); ok {
				v = utils.ValidateConvIPMask2CIDR(s, v)
			}
		}

		if err = d.Set("ipv4_address", v); err != nil {
			return diag.Errorf("error reading ipv4_address: %v", err)
		}
	}

	if o.Mode != nil {
		v := *o.Mode

		if err = d.Set("mode", v); err != nil {
			return diag.Errorf("error reading mode: %v", err)
		}
	}

	if o.SslCertificate != nil {
		v := *o.SslCertificate

		if err = d.Set("ssl_certificate", v); err != nil {
			return diag.Errorf("error reading ssl_certificate: %v", err)
		}
	}

	if o.Status != nil {
		v := *o.Status

		if err = d.Set("status", v); err != nil {
			return diag.Errorf("error reading status: %v", err)
		}
	}

	if o.UpdateUrl != nil {
		v := *o.UpdateUrl

		if err = d.Set("update_url", v); err != nil {
			return diag.Errorf("error reading update_url: %v", err)
		}
	}

	return nil
}

func getObjectSystemVneTunnel(d *schema.ResourceData, sv string) (*models.SystemVneTunnel, diag.Diagnostics) {
	obj := models.SystemVneTunnel{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("bmr_hostname"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("bmr_hostname", sv)
				diags = append(diags, e)
			}
			obj.BmrHostname = &v2
		}
	}
	if v1, ok := d.GetOk("br"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("br", sv)
				diags = append(diags, e)
			}
			obj.Br = &v2
		}
	}
	if v1, ok := d.GetOk("http_password"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.2.0", "") {
				e := utils.AttributeVersionWarning("http_password", sv)
				diags = append(diags, e)
			}
			obj.HttpPassword = &v2
		}
	}
	if v1, ok := d.GetOk("http_username"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.2.0", "") {
				e := utils.AttributeVersionWarning("http_username", sv)
				diags = append(diags, e)
			}
			obj.HttpUsername = &v2
		}
	}
	if v1, ok := d.GetOk("interface"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("interface", sv)
				diags = append(diags, e)
			}
			obj.Interface = &v2
		}
	}
	if v1, ok := d.GetOk("ipv4_address"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ipv4_address", sv)
				diags = append(diags, e)
			}
			obj.Ipv4Address = &v2
		}
	}
	if v1, ok := d.GetOk("mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("mode", sv)
				diags = append(diags, e)
			}
			obj.Mode = &v2
		}
	}
	if v1, ok := d.GetOk("ssl_certificate"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssl_certificate", sv)
				diags = append(diags, e)
			}
			obj.SslCertificate = &v2
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
	if v1, ok := d.GetOk("update_url"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("update_url", sv)
				diags = append(diags, e)
			}
			obj.UpdateUrl = &v2
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectSystemVneTunnel(d *schema.ResourceData, sv string) (*models.SystemVneTunnel, diag.Diagnostics) {
	obj := models.SystemVneTunnel{}
	diags := diag.Diagnostics{}

	return &obj, diags
}
