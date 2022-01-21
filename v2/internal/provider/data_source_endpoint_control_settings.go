// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure endpoint control settings.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
)

func dataSourceEndpointControlSettings() *schema.Resource {
	return &schema.Resource{
		Description: "Configure endpoint control settings.",

		ReadContext: dataSourceEndpointControlSettingsRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"forticlient_disconnect_unsupported_client": {
				Type:        schema.TypeString,
				Description: "Enable/disable disconnecting of unsupported FortiClient endpoints.",
				Computed:    true,
			},
			"forticlient_keepalive_interval": {
				Type:        schema.TypeInt,
				Description: "Interval between two KeepAlive messages from FortiClient (20 - 300 sec, default = 60).",
				Computed:    true,
			},
			"forticlient_sys_update_interval": {
				Type:        schema.TypeInt,
				Description: "Interval between two system update messages from FortiClient (30 - 1440 min, default = 720).",
				Computed:    true,
			},
			"forticlient_user_avatar": {
				Type:        schema.TypeString,
				Description: "Enable/disable uploading FortiClient user avatars.",
				Computed:    true,
			},
		},
	}
}

func dataSourceEndpointControlSettingsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := "EndpointControlSettings"

	o, err := c.Cmdb.ReadEndpointControlSettings(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading EndpointControlSettings dataSource: %v", err)
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

	diags := refreshObjectEndpointControlSettings(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
