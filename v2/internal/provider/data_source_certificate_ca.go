// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: CA certificate.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceCertificateCa() *schema.Resource {
	return &schema.Resource{
		Description: "CA certificate.",

		ReadContext: dataSourceCertificateCaRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"auto_update_days": {
				Type:        schema.TypeInt,
				Description: "Number of days to wait before requesting an updated CA certificate (0 - 4294967295, 0 = disabled).",
				Computed:    true,
			},
			"auto_update_days_warning": {
				Type:        schema.TypeInt,
				Description: "Number of days before an expiry-warning message is generated (0 - 4294967295, 0 = disabled).",
				Computed:    true,
			},
			"ca": {
				Type:        schema.TypeString,
				Description: "CA certificate as a PEM file.",
				Computed:    true,
			},
			"ca_identifier": {
				Type:        schema.TypeString,
				Description: "CA identifier of the SCEP server.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Name.",
				Required:    true,
			},
			"range": {
				Type:        schema.TypeString,
				Description: "Either global or VDOM IP address range for the CA certificate.",
				Computed:    true,
			},
			"scep_url": {
				Type:        schema.TypeString,
				Description: "URL of the SCEP server.",
				Computed:    true,
			},
			"source": {
				Type:        schema.TypeString,
				Description: "CA certificate source type.",
				Computed:    true,
			},
			"source_ip": {
				Type:        schema.TypeString,
				Description: "Source IP address for communications to the SCEP server.",
				Computed:    true,
			},
			"ssl_inspection_trusted": {
				Type:        schema.TypeString,
				Description: "Enable/disable this CA as a trusted CA for SSL inspection.",
				Computed:    true,
			},
		},
	}
}

func dataSourceCertificateCaRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadCertificateCa(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading CertificateCa dataSource: %v", err)
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

	diags := refreshObjectCertificateCa(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
