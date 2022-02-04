// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure IPS custom signature.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceIpsCustom() *schema.Resource {
	return &schema.Resource{
		Description: "Configure IPS custom signature.",

		ReadContext: dataSourceIpsCustomRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"action": {
				Type:        schema.TypeString,
				Description: "Default action (pass or block) for this signature.",
				Computed:    true,
			},
			"application": {
				Type:        schema.TypeString,
				Description: "Applications to be protected. Blank for all applications.",
				Computed:    true,
			},
			"comment": {
				Type:        schema.TypeString,
				Description: "Comment.",
				Computed:    true,
			},
			"location": {
				Type:        schema.TypeString,
				Description: "Protect client or server traffic.",
				Computed:    true,
			},
			"log": {
				Type:        schema.TypeString,
				Description: "Enable/disable logging.",
				Computed:    true,
			},
			"log_packet": {
				Type:        schema.TypeString,
				Description: "Enable/disable packet logging.",
				Computed:    true,
			},
			"os": {
				Type:        schema.TypeString,
				Description: "Operating system(s) that the signature protects. Blank for all operating systems.",
				Computed:    true,
			},
			"protocol": {
				Type:        schema.TypeString,
				Description: "Protocol(s) that the signature scans. Blank for all protocols.",
				Computed:    true,
			},
			"rule_id": {
				Type:        schema.TypeInt,
				Description: "Signature ID.",
				Computed:    true,
			},
			"severity": {
				Type:        schema.TypeString,
				Description: "Relative severity of the signature, from info to critical. Log messages generated by the signature include the severity.",
				Computed:    true,
			},
			"signature": {
				Type:        schema.TypeString,
				Description: "Custom signature enclosed in single quotes.",
				Computed:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Enable/disable this signature.",
				Computed:    true,
			},
			"tag": {
				Type:        schema.TypeString,
				Description: "Signature tag.",
				Required:    true,
			},
		},
	}
}

func dataSourceIpsCustomRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	i := d.Get("tag")
	mkey := utils.ParseMkey(i)

	o, err := c.Cmdb.ReadIpsCustom(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading IpsCustom dataSource: %v", err)
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

	diags := refreshObjectIpsCustom(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
