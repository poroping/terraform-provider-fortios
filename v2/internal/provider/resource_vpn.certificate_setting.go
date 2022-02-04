// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: VPN certificate setting.

package provider

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func resourceVpnCertificateSetting() *schema.Resource {
	return &schema.Resource{
		Description: "VPN certificate setting.",

		CreateContext: resourceVpnCertificateSettingCreate,
		ReadContext:   resourceVpnCertificateSettingRead,
		UpdateContext: resourceVpnCertificateSettingUpdate,
		DeleteContext: resourceVpnCertificateSettingDelete,

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
			"certname_dsa1024": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "1024 bit DSA key certificate for re-signing server certificates for SSL inspection.",
				Optional:    true,
				Computed:    true,
			},
			"certname_dsa2048": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "2048 bit DSA key certificate for re-signing server certificates for SSL inspection.",
				Optional:    true,
				Computed:    true,
			},
			"certname_ecdsa256": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "256 bit ECDSA key certificate for re-signing server certificates for SSL inspection.",
				Optional:    true,
				Computed:    true,
			},
			"certname_ecdsa384": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "384 bit ECDSA key certificate for re-signing server certificates for SSL inspection.",
				Optional:    true,
				Computed:    true,
			},
			"certname_ecdsa521": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "521 bit ECDSA key certificate for re-signing server certificates for SSL inspection.",
				Optional:    true,
				Computed:    true,
			},
			"certname_ed25519": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "253 bit EdDSA key certificate for re-signing server certificates for SSL inspection.",
				Optional:    true,
				Computed:    true,
			},
			"certname_ed448": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "456 bit EdDSA key certificate for re-signing server certificates for SSL inspection.",
				Optional:    true,
				Computed:    true,
			},
			"certname_rsa1024": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "1024 bit RSA key certificate for re-signing server certificates for SSL inspection.",
				Optional:    true,
				Computed:    true,
			},
			"certname_rsa2048": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "2048 bit RSA key certificate for re-signing server certificates for SSL inspection.",
				Optional:    true,
				Computed:    true,
			},
			"certname_rsa4096": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "4096 bit RSA key certificate for re-signing server certificates for SSL inspection.",
				Optional:    true,
				Computed:    true,
			},
			"check_ca_cert": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable verification of the user certificate and pass authentication if any CA in the chain is trusted (default = enable).",
				Optional:    true,
				Computed:    true,
			},
			"check_ca_chain": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable verification of the entire certificate chain and pass authentication only if the chain is complete and all of the CAs in the chain are trusted (default = disable).",
				Optional:    true,
				Computed:    true,
			},
			"cmp_key_usage_checking": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable server certificate key usage checking in CMP mode (default = enable).",
				Optional:    true,
				Computed:    true,
			},
			"cmp_save_extra_certs": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable saving extra certificates in CMP mode (default = disable).",
				Optional:    true,
				Computed:    true,
			},
			"cn_allow_multi": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "When searching for a matching certificate, allow multiple CN fields in certificate subject name (default = enable).",
				Optional:    true,
				Computed:    true,
			},
			"cn_match": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"substring", "value"}, false),

				Description: "When searching for a matching certificate, control how to do CN value matching with certificate subject name (default = substring).",
				Optional:    true,
				Computed:    true,
			},
			"crl_verification": {
				Type:        schema.TypeList,
				Description: "CRL verification options.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"chain_crl_absence": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"ignore", "revoke"}, false),

							Description: "CRL verification option when CRL of any certificate in chain is absent (default = ignore).",
							Optional:    true,
							Computed:    true,
						},
						"expiry": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"ignore", "revoke"}, false),

							Description: "CRL verification option when CRL is expired (default = ignore).",
							Optional:    true,
							Computed:    true,
						},
						"leaf_crl_absence": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"ignore", "revoke"}, false),

							Description: "CRL verification option when leaf CRL is absent (default = ignore).",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"interface": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "Specify outgoing interface to reach server.",
				Optional:    true,
				Computed:    true,
			},
			"interface_select_method": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"auto", "sdwan", "specify"}, false),

				Description: "Specify how to select outgoing interface to reach server.",
				Optional:    true,
				Computed:    true,
			},
			"ocsp_default_server": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Default OCSP server.",
				Optional:    true,
				Computed:    true,
			},
			"ocsp_option": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"certificate", "server"}, false),

				Description: "Specify whether the OCSP URL is from certificate or configured OCSP server.",
				Optional:    true,
				Computed:    true,
			},
			"ocsp_status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable receiving certificates using the OCSP.",
				Optional:    true,
				Computed:    true,
			},
			"ssl_min_proto_version": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"default", "SSLv3", "TLSv1", "TLSv1-1", "TLSv1-2"}, false),

				Description: "Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).",
				Optional:    true,
				Computed:    true,
			},
			"ssl_ocsp_source_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "Source IP address to use to communicate with the OCSP server.",
				Optional:    true,
				Computed:    true,
			},
			"strict_crl_check": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable strict mode CRL checking.",
				Optional:    true,
				Computed:    true,
			},
			"strict_ocsp_check": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable strict mode OCSP checking.",
				Optional:    true,
				Computed:    true,
			},
			"subject_match": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"substring", "value"}, false),

				Description: "When searching for a matching certificate, control how to do RDN value matching with certificate subject name (default = substring).",
				Optional:    true,
				Computed:    true,
			},
			"subject_set": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"subset", "superset"}, false),

				Description: "When searching for a matching certificate, control how to do RDN set matching with certificate subject name (default = subset).",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceVpnCertificateSettingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectVpnCertificateSetting(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateVpnCertificateSetting(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("VpnCertificateSetting")
	}

	return resourceVpnCertificateSettingRead(ctx, d, meta)
}

func resourceVpnCertificateSettingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectVpnCertificateSetting(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateVpnCertificateSetting(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating VpnCertificateSetting resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("VpnCertificateSetting")
	}

	return resourceVpnCertificateSettingRead(ctx, d, meta)
}

func resourceVpnCertificateSettingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectVpnCertificateSetting(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateVpnCertificateSetting(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting VpnCertificateSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnCertificateSettingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadVpnCertificateSetting(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading VpnCertificateSetting resource: %v", err)
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

	diags := refreshObjectVpnCertificateSetting(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenVpnCertificateSettingCrlVerification(v *[]models.VpnCertificateSettingCrlVerification, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.ChainCrlAbsence; tmp != nil {
				v["chain_crl_absence"] = *tmp
			}

			if tmp := cfg.Expiry; tmp != nil {
				v["expiry"] = *tmp
			}

			if tmp := cfg.LeafCrlAbsence; tmp != nil {
				v["leaf_crl_absence"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func refreshObjectVpnCertificateSetting(d *schema.ResourceData, o *models.VpnCertificateSetting, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.CertnameDsa1024 != nil {
		v := *o.CertnameDsa1024

		if err = d.Set("certname_dsa1024", v); err != nil {
			return diag.Errorf("error reading certname_dsa1024: %v", err)
		}
	}

	if o.CertnameDsa2048 != nil {
		v := *o.CertnameDsa2048

		if err = d.Set("certname_dsa2048", v); err != nil {
			return diag.Errorf("error reading certname_dsa2048: %v", err)
		}
	}

	if o.CertnameEcdsa256 != nil {
		v := *o.CertnameEcdsa256

		if err = d.Set("certname_ecdsa256", v); err != nil {
			return diag.Errorf("error reading certname_ecdsa256: %v", err)
		}
	}

	if o.CertnameEcdsa384 != nil {
		v := *o.CertnameEcdsa384

		if err = d.Set("certname_ecdsa384", v); err != nil {
			return diag.Errorf("error reading certname_ecdsa384: %v", err)
		}
	}

	if o.CertnameEcdsa521 != nil {
		v := *o.CertnameEcdsa521

		if err = d.Set("certname_ecdsa521", v); err != nil {
			return diag.Errorf("error reading certname_ecdsa521: %v", err)
		}
	}

	if o.CertnameEd25519 != nil {
		v := *o.CertnameEd25519

		if err = d.Set("certname_ed25519", v); err != nil {
			return diag.Errorf("error reading certname_ed25519: %v", err)
		}
	}

	if o.CertnameEd448 != nil {
		v := *o.CertnameEd448

		if err = d.Set("certname_ed448", v); err != nil {
			return diag.Errorf("error reading certname_ed448: %v", err)
		}
	}

	if o.CertnameRsa1024 != nil {
		v := *o.CertnameRsa1024

		if err = d.Set("certname_rsa1024", v); err != nil {
			return diag.Errorf("error reading certname_rsa1024: %v", err)
		}
	}

	if o.CertnameRsa2048 != nil {
		v := *o.CertnameRsa2048

		if err = d.Set("certname_rsa2048", v); err != nil {
			return diag.Errorf("error reading certname_rsa2048: %v", err)
		}
	}

	if o.CertnameRsa4096 != nil {
		v := *o.CertnameRsa4096

		if err = d.Set("certname_rsa4096", v); err != nil {
			return diag.Errorf("error reading certname_rsa4096: %v", err)
		}
	}

	if o.CheckCaCert != nil {
		v := *o.CheckCaCert

		if err = d.Set("check_ca_cert", v); err != nil {
			return diag.Errorf("error reading check_ca_cert: %v", err)
		}
	}

	if o.CheckCaChain != nil {
		v := *o.CheckCaChain

		if err = d.Set("check_ca_chain", v); err != nil {
			return diag.Errorf("error reading check_ca_chain: %v", err)
		}
	}

	if o.CmpKeyUsageChecking != nil {
		v := *o.CmpKeyUsageChecking

		if err = d.Set("cmp_key_usage_checking", v); err != nil {
			return diag.Errorf("error reading cmp_key_usage_checking: %v", err)
		}
	}

	if o.CmpSaveExtraCerts != nil {
		v := *o.CmpSaveExtraCerts

		if err = d.Set("cmp_save_extra_certs", v); err != nil {
			return diag.Errorf("error reading cmp_save_extra_certs: %v", err)
		}
	}

	if o.CnAllowMulti != nil {
		v := *o.CnAllowMulti

		if err = d.Set("cn_allow_multi", v); err != nil {
			return diag.Errorf("error reading cn_allow_multi: %v", err)
		}
	}

	if o.CnMatch != nil {
		v := *o.CnMatch

		if err = d.Set("cn_match", v); err != nil {
			return diag.Errorf("error reading cn_match: %v", err)
		}
	}

	if o.CrlVerification != nil {
		if err = d.Set("crl_verification", flattenVpnCertificateSettingCrlVerification(o.CrlVerification, sort)); err != nil {
			return diag.Errorf("error reading crl_verification: %v", err)
		}
	}

	if o.Interface != nil {
		v := *o.Interface

		if err = d.Set("interface", v); err != nil {
			return diag.Errorf("error reading interface: %v", err)
		}
	}

	if o.InterfaceSelectMethod != nil {
		v := *o.InterfaceSelectMethod

		if err = d.Set("interface_select_method", v); err != nil {
			return diag.Errorf("error reading interface_select_method: %v", err)
		}
	}

	if o.OcspDefaultServer != nil {
		v := *o.OcspDefaultServer

		if err = d.Set("ocsp_default_server", v); err != nil {
			return diag.Errorf("error reading ocsp_default_server: %v", err)
		}
	}

	if o.OcspOption != nil {
		v := *o.OcspOption

		if err = d.Set("ocsp_option", v); err != nil {
			return diag.Errorf("error reading ocsp_option: %v", err)
		}
	}

	if o.OcspStatus != nil {
		v := *o.OcspStatus

		if err = d.Set("ocsp_status", v); err != nil {
			return diag.Errorf("error reading ocsp_status: %v", err)
		}
	}

	if o.SslMinProtoVersion != nil {
		v := *o.SslMinProtoVersion

		if err = d.Set("ssl_min_proto_version", v); err != nil {
			return diag.Errorf("error reading ssl_min_proto_version: %v", err)
		}
	}

	if o.SslOcspSourceIp != nil {
		v := *o.SslOcspSourceIp

		if err = d.Set("ssl_ocsp_source_ip", v); err != nil {
			return diag.Errorf("error reading ssl_ocsp_source_ip: %v", err)
		}
	}

	if o.StrictCrlCheck != nil {
		v := *o.StrictCrlCheck

		if err = d.Set("strict_crl_check", v); err != nil {
			return diag.Errorf("error reading strict_crl_check: %v", err)
		}
	}

	if o.StrictOcspCheck != nil {
		v := *o.StrictOcspCheck

		if err = d.Set("strict_ocsp_check", v); err != nil {
			return diag.Errorf("error reading strict_ocsp_check: %v", err)
		}
	}

	if o.SubjectMatch != nil {
		v := *o.SubjectMatch

		if err = d.Set("subject_match", v); err != nil {
			return diag.Errorf("error reading subject_match: %v", err)
		}
	}

	if o.SubjectSet != nil {
		v := *o.SubjectSet

		if err = d.Set("subject_set", v); err != nil {
			return diag.Errorf("error reading subject_set: %v", err)
		}
	}

	return nil
}

func expandVpnCertificateSettingCrlVerification(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.VpnCertificateSettingCrlVerification, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.VpnCertificateSettingCrlVerification

	for i := range l {
		tmp := models.VpnCertificateSettingCrlVerification{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.chain_crl_absence", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ChainCrlAbsence = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.expiry", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Expiry = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.leaf_crl_absence", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.LeafCrlAbsence = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectVpnCertificateSetting(d *schema.ResourceData, sv string) (*models.VpnCertificateSetting, diag.Diagnostics) {
	obj := models.VpnCertificateSetting{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("certname_dsa1024"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("certname_dsa1024", sv)
				diags = append(diags, e)
			}
			obj.CertnameDsa1024 = &v2
		}
	}
	if v1, ok := d.GetOk("certname_dsa2048"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("certname_dsa2048", sv)
				diags = append(diags, e)
			}
			obj.CertnameDsa2048 = &v2
		}
	}
	if v1, ok := d.GetOk("certname_ecdsa256"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("certname_ecdsa256", sv)
				diags = append(diags, e)
			}
			obj.CertnameEcdsa256 = &v2
		}
	}
	if v1, ok := d.GetOk("certname_ecdsa384"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("certname_ecdsa384", sv)
				diags = append(diags, e)
			}
			obj.CertnameEcdsa384 = &v2
		}
	}
	if v1, ok := d.GetOk("certname_ecdsa521"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("certname_ecdsa521", sv)
				diags = append(diags, e)
			}
			obj.CertnameEcdsa521 = &v2
		}
	}
	if v1, ok := d.GetOk("certname_ed25519"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("certname_ed25519", sv)
				diags = append(diags, e)
			}
			obj.CertnameEd25519 = &v2
		}
	}
	if v1, ok := d.GetOk("certname_ed448"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("certname_ed448", sv)
				diags = append(diags, e)
			}
			obj.CertnameEd448 = &v2
		}
	}
	if v1, ok := d.GetOk("certname_rsa1024"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("certname_rsa1024", sv)
				diags = append(diags, e)
			}
			obj.CertnameRsa1024 = &v2
		}
	}
	if v1, ok := d.GetOk("certname_rsa2048"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("certname_rsa2048", sv)
				diags = append(diags, e)
			}
			obj.CertnameRsa2048 = &v2
		}
	}
	if v1, ok := d.GetOk("certname_rsa4096"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("certname_rsa4096", sv)
				diags = append(diags, e)
			}
			obj.CertnameRsa4096 = &v2
		}
	}
	if v1, ok := d.GetOk("check_ca_cert"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("check_ca_cert", sv)
				diags = append(diags, e)
			}
			obj.CheckCaCert = &v2
		}
	}
	if v1, ok := d.GetOk("check_ca_chain"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("check_ca_chain", sv)
				diags = append(diags, e)
			}
			obj.CheckCaChain = &v2
		}
	}
	if v1, ok := d.GetOk("cmp_key_usage_checking"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.0") {
				e := utils.AttributeVersionWarning("cmp_key_usage_checking", sv)
				diags = append(diags, e)
			}
			obj.CmpKeyUsageChecking = &v2
		}
	}
	if v1, ok := d.GetOk("cmp_save_extra_certs"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("cmp_save_extra_certs", sv)
				diags = append(diags, e)
			}
			obj.CmpSaveExtraCerts = &v2
		}
	}
	if v1, ok := d.GetOk("cn_allow_multi"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.1", "") {
				e := utils.AttributeVersionWarning("cn_allow_multi", sv)
				diags = append(diags, e)
			}
			obj.CnAllowMulti = &v2
		}
	}
	if v1, ok := d.GetOk("cn_match"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("cn_match", sv)
				diags = append(diags, e)
			}
			obj.CnMatch = &v2
		}
	}
	if v, ok := d.GetOk("crl_verification"); ok {
		if !utils.CheckVer(sv, "v7.0.1", "") {
			e := utils.AttributeVersionWarning("crl_verification", sv)
			diags = append(diags, e)
		}
		t, err := expandVpnCertificateSettingCrlVerification(d, v, "crl_verification", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.CrlVerification = t
		}
	} else if d.HasChange("crl_verification") {
		old, new := d.GetChange("crl_verification")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.CrlVerification = &[]models.VpnCertificateSettingCrlVerification{}
		}
	}
	if v1, ok := d.GetOk("interface"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.0") {
				e := utils.AttributeVersionWarning("interface", sv)
				diags = append(diags, e)
			}
			obj.Interface = &v2
		}
	}
	if v1, ok := d.GetOk("interface_select_method"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.0") {
				e := utils.AttributeVersionWarning("interface_select_method", sv)
				diags = append(diags, e)
			}
			obj.InterfaceSelectMethod = &v2
		}
	}
	if v1, ok := d.GetOk("ocsp_default_server"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ocsp_default_server", sv)
				diags = append(diags, e)
			}
			obj.OcspDefaultServer = &v2
		}
	}
	if v1, ok := d.GetOk("ocsp_option"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ocsp_option", sv)
				diags = append(diags, e)
			}
			obj.OcspOption = &v2
		}
	}
	if v1, ok := d.GetOk("ocsp_status"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ocsp_status", sv)
				diags = append(diags, e)
			}
			obj.OcspStatus = &v2
		}
	}
	if v1, ok := d.GetOk("ssl_min_proto_version"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssl_min_proto_version", sv)
				diags = append(diags, e)
			}
			obj.SslMinProtoVersion = &v2
		}
	}
	if v1, ok := d.GetOk("ssl_ocsp_source_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssl_ocsp_source_ip", sv)
				diags = append(diags, e)
			}
			obj.SslOcspSourceIp = &v2
		}
	}
	if v1, ok := d.GetOk("strict_crl_check"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v7.0.1") {
				e := utils.AttributeVersionWarning("strict_crl_check", sv)
				diags = append(diags, e)
			}
			obj.StrictCrlCheck = &v2
		}
	}
	if v1, ok := d.GetOk("strict_ocsp_check"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("strict_ocsp_check", sv)
				diags = append(diags, e)
			}
			obj.StrictOcspCheck = &v2
		}
	}
	if v1, ok := d.GetOk("subject_match"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("subject_match", sv)
				diags = append(diags, e)
			}
			obj.SubjectMatch = &v2
		}
	}
	if v1, ok := d.GetOk("subject_set"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.1", "") {
				e := utils.AttributeVersionWarning("subject_set", sv)
				diags = append(diags, e)
			}
			obj.SubjectSet = &v2
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectVpnCertificateSetting(d *schema.ResourceData, sv string) (*models.VpnCertificateSetting, diag.Diagnostics) {
	obj := models.VpnCertificateSetting{}
	diags := diag.Diagnostics{}

	obj.CrlVerification = &[]models.VpnCertificateSettingCrlVerification{}

	return &obj, diags
}
