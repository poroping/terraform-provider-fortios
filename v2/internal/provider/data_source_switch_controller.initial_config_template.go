// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure template for auto-generated VLANs.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceSwitchControllerInitialConfigTemplate() *schema.Resource {
	return &schema.Resource{
		Description: "Configure template for auto-generated VLANs.",

		ReadContext: dataSourceSwitchControllerInitialConfigTemplateRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"allowaccess": {
				Type:        schema.TypeString,
				Description: "Permitted types of management access to this interface.",
				Computed:    true,
			},
			"auto_ip": {
				Type:        schema.TypeString,
				Description: "Automatically allocate interface address and subnet block.",
				Computed:    true,
			},
			"dhcp_server": {
				Type:        schema.TypeString,
				Description: "Enable/disable a DHCP server on this interface.",
				Computed:    true,
			},
			"ip": {
				Type:        schema.TypeString,
				Description: "Interface IPv4 address and subnet mask.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Initial config template name.",
				Required:    true,
			},
			"vlanid": {
				Type:        schema.TypeInt,
				Description: "Unique VLAN ID.",
				Computed:    true,
			},
		},
	}
}

func dataSourceSwitchControllerInitialConfigTemplateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSwitchControllerInitialConfigTemplate(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SwitchControllerInitialConfigTemplate dataSource: %v", err)
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

	diags := refreshObjectSwitchControllerInitialConfigTemplate(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
