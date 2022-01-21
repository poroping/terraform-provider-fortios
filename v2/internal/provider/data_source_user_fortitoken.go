// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure FortiToken.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceUserFortitoken() *schema.Resource {
	return &schema.Resource{
		Description: "Configure FortiToken.",

		ReadContext: dataSourceUserFortitokenRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"activation_code": {
				Type:        schema.TypeString,
				Description: "Mobile token user activation-code.",
				Computed:    true,
			},
			"activation_expire": {
				Type:        schema.TypeInt,
				Description: "Mobile token user activation-code expire time.",
				Computed:    true,
			},
			"comments": {
				Type:        schema.TypeString,
				Description: "Comment.",
				Computed:    true,
			},
			"license": {
				Type:        schema.TypeString,
				Description: "Mobile token license.",
				Computed:    true,
			},
			"os_ver": {
				Type:        schema.TypeString,
				Description: "Device Mobile Version.",
				Computed:    true,
			},
			"reg_id": {
				Type:        schema.TypeString,
				Description: "Device Reg ID.",
				Computed:    true,
			},
			"serial_number": {
				Type:        schema.TypeString,
				Description: "Serial number.",
				Required:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Status",
				Computed:    true,
			},
		},
	}
}

func dataSourceUserFortitokenRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	i := d.Get("serial_number")
	mkey := utils.ParseMkey(i)

	o, err := c.Cmdb.ReadUserFortitoken(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading UserFortitoken dataSource: %v", err)
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

	diags := refreshObjectUserFortitoken(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
