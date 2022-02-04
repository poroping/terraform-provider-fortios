// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure certificate users.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceUserCertificate() *schema.Resource {
	return &schema.Resource{
		Description: "Configure certificate users.",

		ReadContext: dataSourceUserCertificateRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"common_name": {
				Type:        schema.TypeString,
				Description: "Certificate common name.",
				Computed:    true,
			},
			"fosid": {
				Type:        schema.TypeInt,
				Description: "User ID.",
				Computed:    true,
			},
			"issuer": {
				Type:        schema.TypeString,
				Description: "CA certificate used for client certificate verification.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "User name.",
				Required:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Enable/disable allowing the certificate user to authenticate with the FortiGate unit.",
				Computed:    true,
			},
			"type": {
				Type:        schema.TypeString,
				Description: "Type of certificate authentication method.",
				Computed:    true,
			},
		},
	}
}

func dataSourceUserCertificateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadUserCertificate(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading UserCertificate dataSource: %v", err)
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

	diags := refreshObjectUserCertificate(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
