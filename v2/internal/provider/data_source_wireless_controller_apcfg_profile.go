// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure AP local configuration profiles.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceWirelessControllerApcfgProfile() *schema.Resource {
	return &schema.Resource{
		Description: "Configure AP local configuration profiles.",

		ReadContext: dataSourceWirelessControllerApcfgProfileRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"ac_ip": {
				Type:        schema.TypeString,
				Description: "IP address of the validation controller that AP must be able to join after applying AP local configuration.",
				Computed:    true,
			},
			"ac_port": {
				Type:        schema.TypeInt,
				Description: "Port of the validation controller that AP must be able to join after applying AP local configuration (1024 - 49150, default = 5246).",
				Computed:    true,
			},
			"ac_timer": {
				Type:        schema.TypeInt,
				Description: "Maximum waiting time for the AP to join the validation controller after applying AP local configuration (3 - 30 min, default = 10).",
				Computed:    true,
			},
			"ac_type": {
				Type:        schema.TypeString,
				Description: "Validation controller type (default = default).",
				Computed:    true,
			},
			"ap_family": {
				Type:        schema.TypeString,
				Description: "FortiAP family type (default = fap).",
				Computed:    true,
			},
			"command_list": {
				Type:        schema.TypeList,
				Description: "AP local configuration command list.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Description: "Command ID.",
							Computed:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "AP local configuration command name.",
							Computed:    true,
						},
						"passwd_value": {
							Type:        schema.TypeString,
							Description: "AP local configuration command password value.",
							Computed:    true,
							Sensitive:   true,
						},
						"type": {
							Type:        schema.TypeString,
							Description: "The command type (default = non-password).",
							Computed:    true,
						},
						"value": {
							Type:        schema.TypeString,
							Description: "AP local configuration command value.",
							Computed:    true,
						},
					},
				},
			},
			"comment": {
				Type:        schema.TypeString,
				Description: "Comment.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "AP local configuration profile name.",
				Required:    true,
			},
		},
	}
}

func dataSourceWirelessControllerApcfgProfileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadWirelessControllerApcfgProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WirelessControllerApcfgProfile dataSource: %v", err)
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

	diags := refreshObjectWirelessControllerApcfgProfile(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
