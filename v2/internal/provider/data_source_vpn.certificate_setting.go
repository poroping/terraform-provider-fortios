// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: VPN certificate setting.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceVpnCertificateSetting() *schema.Resource {
	return &schema.Resource{
		Description: "VPN certificate setting.",

		ReadContext: dataSourceVpnCertificateSettingRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"certname_dsa1024": {
				Type:        schema.TypeString,
				Description: "1024 bit DSA key certificate for re-signing server certificates for SSL inspection.",
				Computed:    true,
			},
			"certname_dsa2048": {
				Type:        schema.TypeString,
				Description: "2048 bit DSA key certificate for re-signing server certificates for SSL inspection.",
				Computed:    true,
			},
			"certname_ecdsa256": {
				Type:        schema.TypeString,
				Description: "256 bit ECDSA key certificate for re-signing server certificates for SSL inspection.",
				Computed:    true,
			},
			"certname_ecdsa384": {
				Type:        schema.TypeString,
				Description: "384 bit ECDSA key certificate for re-signing server certificates for SSL inspection.",
				Computed:    true,
			},
			"certname_ecdsa521": {
				Type:        schema.TypeString,
				Description: "521 bit ECDSA key certificate for re-signing server certificates for SSL inspection.",
				Computed:    true,
			},
			"certname_ed25519": {
				Type:        schema.TypeString,
				Description: "253 bit EdDSA key certificate for re-signing server certificates for SSL inspection.",
				Computed:    true,
			},
			"certname_ed448": {
				Type:        schema.TypeString,
				Description: "456 bit EdDSA key certificate for re-signing server certificates for SSL inspection.",
				Computed:    true,
			},
			"certname_rsa1024": {
				Type:        schema.TypeString,
				Description: "1024 bit RSA key certificate for re-signing server certificates for SSL inspection.",
				Computed:    true,
			},
			"certname_rsa2048": {
				Type:        schema.TypeString,
				Description: "2048 bit RSA key certificate for re-signing server certificates for SSL inspection.",
				Computed:    true,
			},
			"certname_rsa4096": {
				Type:        schema.TypeString,
				Description: "4096 bit RSA key certificate for re-signing server certificates for SSL inspection.",
				Computed:    true,
			},
			"check_ca_cert": {
				Type:        schema.TypeString,
				Description: "Enable/disable verification of the user certificate and pass authentication if any CA in the chain is trusted (default = enable).",
				Computed:    true,
			},
			"check_ca_chain": {
				Type:        schema.TypeString,
				Description: "Enable/disable verification of the entire certificate chain and pass authentication only if the chain is complete and all of the CAs in the chain are trusted (default = disable).",
				Computed:    true,
			},
			"cmp_key_usage_checking": {
				Type:        schema.TypeString,
				Description: "Enable/disable server certificate key usage checking in CMP mode (default = enable).",
				Computed:    true,
			},
			"cmp_save_extra_certs": {
				Type:        schema.TypeString,
				Description: "Enable/disable saving extra certificates in CMP mode (default = disable).",
				Computed:    true,
			},
			"cn_allow_multi": {
				Type:        schema.TypeString,
				Description: "When searching for a matching certificate, allow mutliple CN fields in certificate subject name (default = enable).",
				Computed:    true,
			},
			"cn_match": {
				Type:        schema.TypeString,
				Description: "When searching for a matching certificate, control how to do CN value matching with certificate subject name (default = substring).",
				Computed:    true,
			},
			"crl_verification": {
				Type:        schema.TypeList,
				Description: "CRL verification options.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"chain_crl_absence": {
							Type:        schema.TypeString,
							Description: "CRL verification option when CRL of any certificate in chain is absent (default = ignore).",
							Computed:    true,
						},
						"expiry": {
							Type:        schema.TypeString,
							Description: "CRL verification option when CRL is expired (default = ignore).",
							Computed:    true,
						},
						"leaf_crl_absence": {
							Type:        schema.TypeString,
							Description: "CRL verification option when leaf CRL is absent (default = ignore).",
							Computed:    true,
						},
					},
				},
			},
			"interface": {
				Type:        schema.TypeString,
				Description: "Specify outgoing interface to reach server.",
				Computed:    true,
			},
			"interface_select_method": {
				Type:        schema.TypeString,
				Description: "Specify how to select outgoing interface to reach server.",
				Computed:    true,
			},
			"ocsp_default_server": {
				Type:        schema.TypeString,
				Description: "Default OCSP server.",
				Computed:    true,
			},
			"ocsp_option": {
				Type:        schema.TypeString,
				Description: "Specify whether the OCSP URL is from certificate or configured OCSP server.",
				Computed:    true,
			},
			"ocsp_status": {
				Type:        schema.TypeString,
				Description: "Enable/disable receiving certificates using the OCSP.",
				Computed:    true,
			},
			"ssl_min_proto_version": {
				Type:        schema.TypeString,
				Description: "Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).",
				Computed:    true,
			},
			"ssl_ocsp_source_ip": {
				Type:        schema.TypeString,
				Description: "Source IP address to use to communicate with the OCSP server.",
				Computed:    true,
			},
			"strict_crl_check": {
				Type:        schema.TypeString,
				Description: "Enable/disable strict mode CRL checking.",
				Computed:    true,
			},
			"strict_ocsp_check": {
				Type:        schema.TypeString,
				Description: "Enable/disable strict mode OCSP checking.",
				Computed:    true,
			},
			"subject_match": {
				Type:        schema.TypeString,
				Description: "When searching for a matching certificate, control how to do RDN value matching with certificate subject name (default = substring).",
				Computed:    true,
			},
			"subject_set": {
				Type:        schema.TypeString,
				Description: "When searching for a matching certificate, control how to do RDN set matching with certificate subject name (default = subset).",
				Computed:    true,
			},
		},
	}
}

func dataSourceVpnCertificateSettingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadVpnCertificateSetting(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading VpnCertificateSetting dataSource: %v", err)
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

	diags := refreshObjectVpnCertificateSetting(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}
