// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Local keys and certificates.

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

func resourceVpnCertificateLocal() *schema.Resource {
	return &schema.Resource{
		Description: "Local keys and certificates.",

		CreateContext: resourceVpnCertificateLocalCreate,
		ReadContext:   resourceVpnCertificateLocalRead,
		UpdateContext: resourceVpnCertificateLocalUpdate,
		DeleteContext: resourceVpnCertificateLocalDelete,

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
			"acme_ca_url": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "The URL for the ACME CA server (Let's Encrypt is the default provider).",
				Optional:    true,
				Computed:    true,
			},
			"acme_domain": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "A valid domain that resolves to this Fortigate.",
				Optional:    true,
				Computed:    true,
			},
			"acme_email": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Contact email address that is required by some CAs like LetsEncrypt.",
				Optional:    true,
				Computed:    true,
			},
			"acme_renew_window": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 60),

				Description: "Beginning of the renewal window (in days before certificate expiration, 30 by default).",
				Optional:    true,
				Computed:    true,
			},
			"acme_rsa_key_size": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(2048, 4096),

				Description: "Length of the RSA private key of the generated cert (Minimum 2048 bits).",
				Optional:    true,
				Computed:    true,
			},
			"auto_regenerate_days": {
				Type: schema.TypeInt,

				Description: "Number of days to wait before expiry of an updated local certificate is requested (0 = disabled).",
				Optional:    true,
				Computed:    true,
			},
			"auto_regenerate_days_warning": {
				Type: schema.TypeInt,

				Description: "Number of days to wait before an expiry warning message is generated (0 = disabled).",
				Optional:    true,
				Computed:    true,
			},
			"ca_identifier": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "CA identifier of the CA server for signing via SCEP.",
				Optional:    true,
				Computed:    true,
			},
			"certificate": {
				Type: schema.TypeString,

				Description: "PEM format certificate.",
				Optional:    true,
				Computed:    true,
			},
			"cmp_path": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Path location inside CMP server.",
				Optional:    true,
				Computed:    true,
			},
			"cmp_regeneration_method": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"keyupate", "renewal"}, false),

				Description: "CMP auto-regeneration method.",
				Optional:    true,
				Computed:    true,
			},
			"cmp_server": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "'ADDRESS:PORT' for CMP server.",
				Optional:    true,
				Computed:    true,
			},
			"cmp_server_cert": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),

				Description: "CMP server certificate.",
				Optional:    true,
				Computed:    true,
			},
			"comments": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 511),

				Description: "Comment.",
				Optional:    true,
				Computed:    true,
			},
			"csr": {
				Type: schema.TypeString,

				Description: "Certificate Signing Request.",
				Optional:    true,
				Computed:    true,
			},
			"enroll_protocol": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"none", "scep", "cmpv2", "acme2"}, false),

				Description: "Certificate enrollment protocol.",
				Optional:    true,
				Computed:    true,
			},
			"ike_localid": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Local ID the FortiGate uses for authentication as a VPN client.",
				Optional:    true,
				Computed:    true,
			},
			"ike_localid_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"asn1dn", "fqdn"}, false),

				Description: "IKE local ID type.",
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
			"name_encoding": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"printable", "utf8"}, false),

				Description: "Name encoding method for auto-regeneration.",
				Optional:    true,
				Computed:    true,
			},
			"password": {
				Type: schema.TypeString,

				Description: "Password as a PEM file.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"private_key": {
				Type: schema.TypeString,

				Description: "PEM format key, encrypted with a password.",
				Optional:    true,
				Computed:    true,
			},
			"range": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"global", "vdom"}, false),

				Description: "Either a global or VDOM IP address range for the certificate.",
				Optional:    true,
				Computed:    true,
			},
			"scep_password": {
				Type: schema.TypeString,

				Description: "SCEP server challenge password for auto-regeneration.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"scep_url": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "SCEP server URL.",
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

				Description: "Source IP address for communications to the SCEP server.",
				Optional:    true,
				Computed:    true,
			},
			"state": {
				Type: schema.TypeString,

				Description: "Certificate Signing Request State.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceVpnCertificateLocalCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating VpnCertificateLocal resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectVpnCertificateLocal(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateVpnCertificateLocal(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("VpnCertificateLocal")
	}

	return resourceVpnCertificateLocalRead(ctx, d, meta)
}

func resourceVpnCertificateLocalUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectVpnCertificateLocal(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateVpnCertificateLocal(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating VpnCertificateLocal resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("VpnCertificateLocal")
	}

	return resourceVpnCertificateLocalRead(ctx, d, meta)
}

func resourceVpnCertificateLocalDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteVpnCertificateLocal(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting VpnCertificateLocal resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnCertificateLocalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadVpnCertificateLocal(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading VpnCertificateLocal resource: %v", err)
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

	diags := refreshObjectVpnCertificateLocal(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectVpnCertificateLocal(d *schema.ResourceData, o *models.VpnCertificateLocal, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AcmeCaUrl != nil {
		v := *o.AcmeCaUrl

		if err = d.Set("acme_ca_url", v); err != nil {
			return diag.Errorf("error reading acme_ca_url: %v", err)
		}
	}

	if o.AcmeDomain != nil {
		v := *o.AcmeDomain

		if err = d.Set("acme_domain", v); err != nil {
			return diag.Errorf("error reading acme_domain: %v", err)
		}
	}

	if o.AcmeEmail != nil {
		v := *o.AcmeEmail

		if err = d.Set("acme_email", v); err != nil {
			return diag.Errorf("error reading acme_email: %v", err)
		}
	}

	if o.AcmeRenewWindow != nil {
		v := *o.AcmeRenewWindow

		if err = d.Set("acme_renew_window", v); err != nil {
			return diag.Errorf("error reading acme_renew_window: %v", err)
		}
	}

	if o.AcmeRsaKeySize != nil {
		v := *o.AcmeRsaKeySize

		if err = d.Set("acme_rsa_key_size", v); err != nil {
			return diag.Errorf("error reading acme_rsa_key_size: %v", err)
		}
	}

	if o.AutoRegenerateDays != nil {
		v := *o.AutoRegenerateDays

		if err = d.Set("auto_regenerate_days", v); err != nil {
			return diag.Errorf("error reading auto_regenerate_days: %v", err)
		}
	}

	if o.AutoRegenerateDaysWarning != nil {
		v := *o.AutoRegenerateDaysWarning

		if err = d.Set("auto_regenerate_days_warning", v); err != nil {
			return diag.Errorf("error reading auto_regenerate_days_warning: %v", err)
		}
	}

	if o.CaIdentifier != nil {
		v := *o.CaIdentifier

		if err = d.Set("ca_identifier", v); err != nil {
			return diag.Errorf("error reading ca_identifier: %v", err)
		}
	}

	if o.Certificate != nil {
		v := *o.Certificate

		if err = d.Set("certificate", v); err != nil {
			return diag.Errorf("error reading certificate: %v", err)
		}
	}

	if o.CmpPath != nil {
		v := *o.CmpPath

		if err = d.Set("cmp_path", v); err != nil {
			return diag.Errorf("error reading cmp_path: %v", err)
		}
	}

	if o.CmpRegenerationMethod != nil {
		v := *o.CmpRegenerationMethod

		if err = d.Set("cmp_regeneration_method", v); err != nil {
			return diag.Errorf("error reading cmp_regeneration_method: %v", err)
		}
	}

	if o.CmpServer != nil {
		v := *o.CmpServer

		if err = d.Set("cmp_server", v); err != nil {
			return diag.Errorf("error reading cmp_server: %v", err)
		}
	}

	if o.CmpServerCert != nil {
		v := *o.CmpServerCert

		if err = d.Set("cmp_server_cert", v); err != nil {
			return diag.Errorf("error reading cmp_server_cert: %v", err)
		}
	}

	if o.Comments != nil {
		v := *o.Comments

		if err = d.Set("comments", v); err != nil {
			return diag.Errorf("error reading comments: %v", err)
		}
	}

	if o.Csr != nil {
		v := *o.Csr

		if err = d.Set("csr", v); err != nil {
			return diag.Errorf("error reading csr: %v", err)
		}
	}

	if o.EnrollProtocol != nil {
		v := *o.EnrollProtocol

		if err = d.Set("enroll_protocol", v); err != nil {
			return diag.Errorf("error reading enroll_protocol: %v", err)
		}
	}

	if o.IkeLocalid != nil {
		v := *o.IkeLocalid

		if err = d.Set("ike_localid", v); err != nil {
			return diag.Errorf("error reading ike_localid: %v", err)
		}
	}

	if o.IkeLocalidType != nil {
		v := *o.IkeLocalidType

		if err = d.Set("ike_localid_type", v); err != nil {
			return diag.Errorf("error reading ike_localid_type: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.NameEncoding != nil {
		v := *o.NameEncoding

		if err = d.Set("name_encoding", v); err != nil {
			return diag.Errorf("error reading name_encoding: %v", err)
		}
	}

	if o.Password != nil {
		v := *o.Password

		if v == "ENC XXXX" {
		} else if err = d.Set("password", v); err != nil {
			return diag.Errorf("error reading password: %v", err)
		}
	}

	if o.PrivateKey != nil {
		v := *o.PrivateKey

		if err = d.Set("private_key", v); err != nil {
			return diag.Errorf("error reading private_key: %v", err)
		}
	}

	if o.Range != nil {
		v := *o.Range

		if err = d.Set("range", v); err != nil {
			return diag.Errorf("error reading range: %v", err)
		}
	}

	if o.ScepPassword != nil {
		v := *o.ScepPassword

		if v == "ENC XXXX" {
		} else if err = d.Set("scep_password", v); err != nil {
			return diag.Errorf("error reading scep_password: %v", err)
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

	if o.State != nil {
		v := *o.State

		if err = d.Set("state", v); err != nil {
			return diag.Errorf("error reading state: %v", err)
		}
	}

	return nil
}

func getObjectVpnCertificateLocal(d *schema.ResourceData, sv string) (*models.VpnCertificateLocal, diag.Diagnostics) {
	obj := models.VpnCertificateLocal{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("acme_ca_url"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("acme_ca_url", sv)
				diags = append(diags, e)
			}
			obj.AcmeCaUrl = &v2
		}
	}
	if v1, ok := d.GetOk("acme_domain"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("acme_domain", sv)
				diags = append(diags, e)
			}
			obj.AcmeDomain = &v2
		}
	}
	if v1, ok := d.GetOk("acme_email"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("acme_email", sv)
				diags = append(diags, e)
			}
			obj.AcmeEmail = &v2
		}
	}
	if v1, ok := d.GetOk("acme_renew_window"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("acme_renew_window", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.AcmeRenewWindow = &tmp
		}
	}
	if v1, ok := d.GetOk("acme_rsa_key_size"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("acme_rsa_key_size", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.AcmeRsaKeySize = &tmp
		}
	}
	if v1, ok := d.GetOk("auto_regenerate_days"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auto_regenerate_days", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.AutoRegenerateDays = &tmp
		}
	}
	if v1, ok := d.GetOk("auto_regenerate_days_warning"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auto_regenerate_days_warning", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.AutoRegenerateDaysWarning = &tmp
		}
	}
	if v1, ok := d.GetOk("ca_identifier"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ca_identifier", sv)
				diags = append(diags, e)
			}
			obj.CaIdentifier = &v2
		}
	}
	if v1, ok := d.GetOk("certificate"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("certificate", sv)
				diags = append(diags, e)
			}
			obj.Certificate = &v2
		}
	}
	if v1, ok := d.GetOk("cmp_path"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("cmp_path", sv)
				diags = append(diags, e)
			}
			obj.CmpPath = &v2
		}
	}
	if v1, ok := d.GetOk("cmp_regeneration_method"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("cmp_regeneration_method", sv)
				diags = append(diags, e)
			}
			obj.CmpRegenerationMethod = &v2
		}
	}
	if v1, ok := d.GetOk("cmp_server"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("cmp_server", sv)
				diags = append(diags, e)
			}
			obj.CmpServer = &v2
		}
	}
	if v1, ok := d.GetOk("cmp_server_cert"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("cmp_server_cert", sv)
				diags = append(diags, e)
			}
			obj.CmpServerCert = &v2
		}
	}
	if v1, ok := d.GetOk("comments"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("comments", sv)
				diags = append(diags, e)
			}
			obj.Comments = &v2
		}
	}
	if v1, ok := d.GetOk("csr"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("csr", sv)
				diags = append(diags, e)
			}
			obj.Csr = &v2
		}
	}
	if v1, ok := d.GetOk("enroll_protocol"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("enroll_protocol", sv)
				diags = append(diags, e)
			}
			obj.EnrollProtocol = &v2
		}
	}
	if v1, ok := d.GetOk("ike_localid"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ike_localid", sv)
				diags = append(diags, e)
			}
			obj.IkeLocalid = &v2
		}
	}
	if v1, ok := d.GetOk("ike_localid_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ike_localid_type", sv)
				diags = append(diags, e)
			}
			obj.IkeLocalidType = &v2
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
	if v1, ok := d.GetOk("name_encoding"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("name_encoding", sv)
				diags = append(diags, e)
			}
			obj.NameEncoding = &v2
		}
	}
	if v1, ok := d.GetOk("password"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("password", sv)
				diags = append(diags, e)
			}
			obj.Password = &v2
		}
	}
	if v1, ok := d.GetOk("private_key"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("private_key", sv)
				diags = append(diags, e)
			}
			obj.PrivateKey = &v2
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
	if v1, ok := d.GetOk("scep_password"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("scep_password", sv)
				diags = append(diags, e)
			}
			obj.ScepPassword = &v2
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
	if v1, ok := d.GetOk("state"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("state", sv)
				diags = append(diags, e)
			}
			obj.State = &v2
		}
	}
	return &obj, diags
}
