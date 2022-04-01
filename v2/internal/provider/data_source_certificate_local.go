// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Local keys and certificates.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceCertificateLocal() *schema.Resource {
	return &schema.Resource{
		Description: "Local keys and certificates.",

		ReadContext: dataSourceCertificateLocalRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"acme_ca_url": {
				Type:        schema.TypeString,
				Description: "The URL for the ACME CA server (Let's Encrypt is the default provider).",
				Computed:    true,
			},
			"acme_domain": {
				Type:        schema.TypeString,
				Description: "A valid domain that resolves to this FortiGate unit.",
				Computed:    true,
			},
			"acme_email": {
				Type:        schema.TypeString,
				Description: "Contact email address that is required by some CAs like LetsEncrypt.",
				Computed:    true,
			},
			"acme_renew_window": {
				Type:        schema.TypeInt,
				Description: "Beginning of the renewal window (in days before certificate expiration, 30 by default).",
				Computed:    true,
			},
			"acme_rsa_key_size": {
				Type:        schema.TypeInt,
				Description: "Length of the RSA private key of the generated cert (Minimum 2048 bits).",
				Computed:    true,
			},
			"auto_regenerate_days": {
				Type:        schema.TypeInt,
				Description: "Number of days to wait before expiry of an updated local certificate is requested (0 = disabled).",
				Computed:    true,
			},
			"auto_regenerate_days_warning": {
				Type:        schema.TypeInt,
				Description: "Number of days to wait before an expiry warning message is generated (0 = disabled).",
				Computed:    true,
			},
			"ca_identifier": {
				Type:        schema.TypeString,
				Description: "CA identifier of the CA server for signing via SCEP.",
				Computed:    true,
			},
			"certificate": {
				Type:        schema.TypeString,
				Description: "PEM format certificate.",
				Computed:    true,
			},
			"cmp_path": {
				Type:        schema.TypeString,
				Description: "Path location inside CMP server.",
				Computed:    true,
			},
			"cmp_regeneration_method": {
				Type:        schema.TypeString,
				Description: "CMP auto-regeneration method.",
				Computed:    true,
			},
			"cmp_server": {
				Type:        schema.TypeString,
				Description: "Address and port for CMP server (format = address:port).",
				Computed:    true,
			},
			"cmp_server_cert": {
				Type:        schema.TypeString,
				Description: "CMP server certificate.",
				Computed:    true,
			},
			"comments": {
				Type:        schema.TypeString,
				Description: "Comment.",
				Computed:    true,
			},
			"csr": {
				Type:        schema.TypeString,
				Description: "Certificate Signing Request.",
				Computed:    true,
			},
			"enroll_protocol": {
				Type:        schema.TypeString,
				Description: "Certificate enrollment protocol.",
				Computed:    true,
			},
			"ike_localid": {
				Type:        schema.TypeString,
				Description: "Local ID the FortiGate uses for authentication as a VPN client.",
				Computed:    true,
			},
			"ike_localid_type": {
				Type:        schema.TypeString,
				Description: "IKE local ID type.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Name.",
				Required:    true,
			},
			"name_encoding": {
				Type:        schema.TypeString,
				Description: "Name encoding method for auto-regeneration.",
				Computed:    true,
			},
			"password": {
				Type:        schema.TypeString,
				Description: "Password as a PEM file.",
				Computed:    true,
				Sensitive:   true,
			},
			"private_key": {
				Type:        schema.TypeString,
				Description: "PEM format key encrypted with a password.",
				Computed:    true,
			},
			"range": {
				Type:        schema.TypeString,
				Description: "Either a global or VDOM IP address range for the certificate.",
				Computed:    true,
			},
			"scep_password": {
				Type:        schema.TypeString,
				Description: "SCEP server challenge password for auto-regeneration.",
				Computed:    true,
				Sensitive:   true,
			},
			"scep_url": {
				Type:        schema.TypeString,
				Description: "SCEP server URL.",
				Computed:    true,
			},
			"source": {
				Type:        schema.TypeString,
				Description: "Certificate source type.",
				Computed:    true,
			},
			"source_ip": {
				Type:        schema.TypeString,
				Description: "Source IP address for communications to the SCEP server.",
				Computed:    true,
			},
			"state": {
				Type:        schema.TypeString,
				Description: "Certificate Signing Request State.",
				Computed:    true,
			},
		},
	}
}

func dataSourceCertificateLocalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	i := d.Get("name")
	mkey := utils.ParseMkey(i)

	o, err := c.Cmdb.ReadCertificateLocal(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading CertificateLocal dataSource: %v", err)
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

	diags := refreshObjectCertificateLocal(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
