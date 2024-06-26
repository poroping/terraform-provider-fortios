// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure external resource.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceSystemExternalResource() *schema.Resource {
	return &schema.Resource{
		Description: "Configure external resource.",

		ReadContext: dataSourceSystemExternalResourceRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"category": {
				Type:        schema.TypeInt,
				Description: "User resource category.",
				Computed:    true,
			},
			"comments": {
				Type:        schema.TypeString,
				Description: "Comment.",
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
			"name": {
				Type:        schema.TypeString,
				Description: "External resource name.",
				Required:    true,
			},
			"password": {
				Type:        schema.TypeString,
				Description: "HTTP basic authentication password.",
				Computed:    true,
				Sensitive:   true,
			},
			"refresh_rate": {
				Type:        schema.TypeInt,
				Description: "Time interval to refresh external resource (1 - 43200 min, default = 5 min).",
				Computed:    true,
			},
			"resource": {
				Type:        schema.TypeString,
				Description: "URI of external resource.",
				Computed:    true,
			},
			"server_identity_check": {
				Type:        schema.TypeString,
				Description: "Certificate verification option.",
				Computed:    true,
			},
			"source_ip": {
				Type:        schema.TypeString,
				Description: "Source IPv4 address used to communicate with server.",
				Computed:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Enable/disable user resource.",
				Computed:    true,
			},
			"type": {
				Type:        schema.TypeString,
				Description: "User resource type.",
				Computed:    true,
			},
			"update_method": {
				Type:        schema.TypeString,
				Description: "External resource update method.",
				Computed:    true,
			},
			"user_agent": {
				Type:        schema.TypeString,
				Description: "HTTP User-Agent header (default = 'curl/7.58.0').",
				Computed:    true,
			},
			"username": {
				Type:        schema.TypeString,
				Description: "HTTP basic authentication user name.",
				Computed:    true,
			},
			"uuid": {
				Type:        schema.TypeString,
				Description: "Universally Unique Identifier (UUID; automatically assigned but can be manually reset).",
				Computed:    true,
			},
		},
	}
}

func dataSourceSystemExternalResourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemExternalResource(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemExternalResource dataSource: %v", err)
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

	diags := refreshObjectSystemExternalResource(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
