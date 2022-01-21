// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Replacement messages.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceSystemReplacemsgDeviceDetectionPortal() *schema.Resource {
	return &schema.Resource{
		Description: "Replacement messages.",

		ReadContext: dataSourceSystemReplacemsgDeviceDetectionPortalRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"buffer": {
				Type:        schema.TypeString,
				Description: "Message string.",
				Computed:    true,
			},
			"format": {
				Type:        schema.TypeString,
				Description: "Format flag.",
				Computed:    true,
			},
			"header": {
				Type:        schema.TypeString,
				Description: "Header flag.",
				Computed:    true,
			},
			"msg_type": {
				Type:        schema.TypeString,
				Description: "Message type.",
				Required:    true,
			},
		},
	}
}

func dataSourceSystemReplacemsgDeviceDetectionPortalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	i := d.Get("msg_type")
	mkey := utils.ParseMkey(i)

	o, err := c.Cmdb.ReadSystemReplacemsgDeviceDetectionPortal(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemReplacemsgDeviceDetectionPortal dataSource: %v", err)
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

	diags := refreshObjectSystemReplacemsgDeviceDetectionPortal(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
