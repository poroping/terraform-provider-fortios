// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Certificate Revocation List as a PEM file.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceCertificateCrl() *schema.Resource {
	return &schema.Resource{
		Description: "Certificate Revocation List as a PEM file.",

		ReadContext: dataSourceCertificateCrlRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"crl": {
				Type:        schema.TypeString,
				Description: "Certificate Revocation List as a PEM file.",
				Computed:    true,
			},
			"http_url": {
				Type:        schema.TypeString,
				Description: "HTTP server URL for CRL auto-update.",
				Computed:    true,
			},
			"ldap_password": {
				Type:        schema.TypeString,
				Description: "LDAP server user password.",
				Computed:    true,
				Sensitive:   true,
			},
			"ldap_server": {
				Type:        schema.TypeString,
				Description: "LDAP server name for CRL auto-update.",
				Computed:    true,
			},
			"ldap_username": {
				Type:        schema.TypeString,
				Description: "LDAP server user name.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Name.",
				Required:    true,
			},
			"range": {
				Type:        schema.TypeString,
				Description: "Either global or VDOM IP address range for the certificate.",
				Computed:    true,
			},
			"scep_cert": {
				Type:        schema.TypeString,
				Description: "Local certificate for SCEP communication for CRL auto-update.",
				Computed:    true,
			},
			"scep_url": {
				Type:        schema.TypeString,
				Description: "SCEP server URL for CRL auto-update.",
				Computed:    true,
			},
			"source": {
				Type:        schema.TypeString,
				Description: "Certificate source type.",
				Computed:    true,
			},
			"source_ip": {
				Type:        schema.TypeString,
				Description: "Source IP address for communications to a HTTP or SCEP CA server.",
				Computed:    true,
			},
			"update_interval": {
				Type:        schema.TypeInt,
				Description: "Time in seconds before the FortiGate checks for an updated CRL. Set to 0 to update only when it expires.",
				Computed:    true,
			},
			"update_vdom": {
				Type:        schema.TypeString,
				Description: "VDOM for CRL update.",
				Computed:    true,
			},
		},
	}
}

func dataSourceCertificateCrlRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadCertificateCrl(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading CertificateCrl dataSource: %v", err)
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

	diags := refreshObjectCertificateCrl(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
