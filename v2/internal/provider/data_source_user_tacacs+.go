// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure TACACS+ server entries.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceUserTacacs() *schema.Resource {
	return &schema.Resource{
		Description: "Configure TACACS+ server entries.",

		ReadContext: dataSourceUserTacacsRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"authen_type": {
				Type:        schema.TypeString,
				Description: "Allowed authentication protocols/methods.",
				Computed:    true,
			},
			"authorization": {
				Type:        schema.TypeString,
				Description: "Enable/disable TACACS+ authorization.",
				Computed:    true,
			},
			"interface": {
				Type:        schema.TypeString,
				Description: "Specify outgoing interface to reach server.",
				Computed:    true,
			},
			"interface_select_method": {
				Type:        schema.TypeString,
				Description: "Specify how to select outgoing interface to reach server.",
				Computed:    true,
			},
			"key": {
				Type:        schema.TypeString,
				Description: "Key to access the primary server.",
				Computed:    true,
				Sensitive:   true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "TACACS+ server entry name.",
				Required:    true,
			},
			"port": {
				Type:        schema.TypeInt,
				Description: "Port number of the TACACS+ server.",
				Computed:    true,
			},
			"secondary_key": {
				Type:        schema.TypeString,
				Description: "Key to access the secondary server.",
				Computed:    true,
				Sensitive:   true,
			},
			"secondary_server": {
				Type:        schema.TypeString,
				Description: "Secondary TACACS+ server CN domain name or IP address.",
				Computed:    true,
			},
			"server": {
				Type:        schema.TypeString,
				Description: "Primary TACACS+ server CN domain name or IP address.",
				Computed:    true,
			},
			"source_ip": {
				Type:        schema.TypeString,
				Description: "Source IP address for communications to TACACS+ server.",
				Computed:    true,
			},
			"tertiary_key": {
				Type:        schema.TypeString,
				Description: "Key to access the tertiary server.",
				Computed:    true,
				Sensitive:   true,
			},
			"tertiary_server": {
				Type:        schema.TypeString,
				Description: "Tertiary TACACS+ server CN domain name or IP address.",
				Computed:    true,
			},
		},
	}
}

func dataSourceUserTacacsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadUserTacacs(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading UserTacacs dataSource: %v", err)
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

	diags := refreshObjectUserTacacs(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
