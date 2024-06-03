// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Certificate Revocation List as a PEM file.

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

func resourceCertificateCrl() *schema.Resource {
	return &schema.Resource{
		Description: "Certificate Revocation List as a PEM file.",

		CreateContext: resourceCertificateCrlCreate,
		ReadContext:   resourceCertificateCrlRead,
		UpdateContext: resourceCertificateCrlUpdate,
		DeleteContext: resourceCertificateCrlDelete,

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
			"crl": {
				Type: schema.TypeString,

				Description: "Certificate Revocation List as a PEM file.",
				Optional:    true,
				Computed:    true,
			},
			"http_url": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "HTTP server URL for CRL auto-update.",
				Optional:    true,
				Computed:    true,
			},
			"ldap_password": {
				Type: schema.TypeString,

				Description: "LDAP server user password.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"ldap_server": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "LDAP server name for CRL auto-update.",
				Optional:    true,
				Computed:    true,
			},
			"ldap_username": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "LDAP server user name.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name.",
				ForceNew:    true,
				Required:    true,
			},
			"range": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"global", "vdom"}, false),

				Description: "Either global or VDOM IP address range for the certificate.",
				Optional:    true,
				Computed:    true,
			},
			"scep_cert": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Local certificate for SCEP communication for CRL auto-update.",
				Optional:    true,
				Computed:    true,
			},
			"scep_url": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "SCEP server URL for CRL auto-update.",
				Optional:    true,
				Computed:    true,
			},
			"source": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"factory", "user", "bundle"}, false),

				Description: "Certificate source type.",
				Optional:    true,
				Computed:    true,
			},
			"source_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "Source IP address for communications to a HTTP or SCEP CA server.",
				Optional:    true,
				Computed:    true,
			},
			"update_interval": {
				Type: schema.TypeInt,

				Description: "Time in seconds before the FortiGate checks for an updated CRL. Set to 0 to update only when it expires.",
				Optional:    true,
				Computed:    true,
			},
			"update_vdom": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),

				Description: "VDOM for CRL update.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceCertificateCrlCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating CertificateCrl resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectCertificateCrl(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateCertificateCrl(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("CertificateCrl")
	}

	return resourceCertificateCrlRead(ctx, d, meta)
}

func resourceCertificateCrlUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectCertificateCrl(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateCertificateCrl(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating CertificateCrl resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("CertificateCrl")
	}

	return resourceCertificateCrlRead(ctx, d, meta)
}

func resourceCertificateCrlDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteCertificateCrl(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting CertificateCrl resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceCertificateCrlRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadCertificateCrl(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading CertificateCrl resource: %v", err)
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

	diags := refreshObjectCertificateCrl(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectCertificateCrl(d *schema.ResourceData, o *models.CertificateCrl, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Crl != nil {
		v := *o.Crl

		if err = d.Set("crl", v); err != nil {
			return diag.Errorf("error reading crl: %v", err)
		}
	}

	if o.HttpUrl != nil {
		v := *o.HttpUrl

		if err = d.Set("http_url", v); err != nil {
			return diag.Errorf("error reading http_url: %v", err)
		}
	}

	if o.LdapPassword != nil {
		v := *o.LdapPassword

		if v == "ENC XXXX" {
		} else if err = d.Set("ldap_password", v); err != nil {
			return diag.Errorf("error reading ldap_password: %v", err)
		}
	}

	if o.LdapServer != nil {
		v := *o.LdapServer

		if err = d.Set("ldap_server", v); err != nil {
			return diag.Errorf("error reading ldap_server: %v", err)
		}
	}

	if o.LdapUsername != nil {
		v := *o.LdapUsername

		if err = d.Set("ldap_username", v); err != nil {
			return diag.Errorf("error reading ldap_username: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.Range != nil {
		v := *o.Range

		if err = d.Set("range", v); err != nil {
			return diag.Errorf("error reading range: %v", err)
		}
	}

	if o.ScepCert != nil {
		v := *o.ScepCert

		if err = d.Set("scep_cert", v); err != nil {
			return diag.Errorf("error reading scep_cert: %v", err)
		}
	}

	if o.ScepUrl != nil {
		v := *o.ScepUrl

		if err = d.Set("scep_url", v); err != nil {
			return diag.Errorf("error reading scep_url: %v", err)
		}
	}

	if o.Source != nil {
		v := *o.Source

		if err = d.Set("source", v); err != nil {
			return diag.Errorf("error reading source: %v", err)
		}
	}

	if o.SourceIp != nil {
		v := *o.SourceIp

		if err = d.Set("source_ip", v); err != nil {
			return diag.Errorf("error reading source_ip: %v", err)
		}
	}

	if o.UpdateInterval != nil {
		v := *o.UpdateInterval

		if err = d.Set("update_interval", v); err != nil {
			return diag.Errorf("error reading update_interval: %v", err)
		}
	}

	if o.UpdateVdom != nil {
		v := *o.UpdateVdom

		if err = d.Set("update_vdom", v); err != nil {
			return diag.Errorf("error reading update_vdom: %v", err)
		}
	}

	return nil
}

func getObjectCertificateCrl(d *schema.ResourceData, sv string) (*models.CertificateCrl, diag.Diagnostics) {
	obj := models.CertificateCrl{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("crl"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("crl", sv)
				diags = append(diags, e)
			}
			obj.Crl = &v2
		}
	}
	if v1, ok := d.GetOk("http_url"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("http_url", sv)
				diags = append(diags, e)
			}
			obj.HttpUrl = &v2
		}
	}
	if v1, ok := d.GetOk("ldap_password"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ldap_password", sv)
				diags = append(diags, e)
			}
			obj.LdapPassword = &v2
		}
	}
	if v1, ok := d.GetOk("ldap_server"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ldap_server", sv)
				diags = append(diags, e)
			}
			obj.LdapServer = &v2
		}
	}
	if v1, ok := d.GetOk("ldap_username"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ldap_username", sv)
				diags = append(diags, e)
			}
			obj.LdapUsername = &v2
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
	if v1, ok := d.GetOk("range"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("range", sv)
				diags = append(diags, e)
			}
			obj.Range = &v2
		}
	}
	if v1, ok := d.GetOk("scep_cert"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("scep_cert", sv)
				diags = append(diags, e)
			}
			obj.ScepCert = &v2
		}
	}
	if v1, ok := d.GetOk("scep_url"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("scep_url", sv)
				diags = append(diags, e)
			}
			obj.ScepUrl = &v2
		}
	}
	if v1, ok := d.GetOk("source"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("source", sv)
				diags = append(diags, e)
			}
			obj.Source = &v2
		}
	}
	if v1, ok := d.GetOk("source_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("source_ip", sv)
				diags = append(diags, e)
			}
			obj.SourceIp = &v2
		}
	}
	if v1, ok := d.GetOk("update_interval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("update_interval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.UpdateInterval = &tmp
		}
	}
	if v1, ok := d.GetOk("update_vdom"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("update_vdom", sv)
				diags = append(diags, e)
			}
			obj.UpdateVdom = &v2
		}
	}
	return &obj, diags
}
