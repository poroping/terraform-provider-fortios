// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure speed test server list.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceSystemSpeedTestServer() *schema.Resource {
	return &schema.Resource{
		Description: "Configure speed test server list.",

		ReadContext: dataSourceSystemSpeedTestServerRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"host": {
				Type:        schema.TypeList,
				Description: "Hosts of the server.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"distance": {
							Type:        schema.TypeInt,
							Description: "Speed test host distance.",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "Server host ID.",
							Computed:    true,
						},
						"ip": {
							Type:        schema.TypeString,
							Description: "Server host IPv4 address.",
							Computed:    true,
						},
						"latitude": {
							Type:        schema.TypeString,
							Description: "Speed test host latitude.",
							Computed:    true,
						},
						"longitude": {
							Type:        schema.TypeString,
							Description: "Speed test host longitude.",
							Computed:    true,
						},
						"password": {
							Type:        schema.TypeString,
							Description: "Speed test host password.",
							Computed:    true,
							Sensitive:   true,
						},
						"port": {
							Type:        schema.TypeInt,
							Description: "Server host port number to communicate with client.",
							Computed:    true,
						},
						"user": {
							Type:        schema.TypeString,
							Description: "Speed test host user name.",
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Speed test server name.",
				Required:    true,
			},
			"timestamp": {
				Type:        schema.TypeInt,
				Description: "Speed test server timestamp.",
				Computed:    true,
			},
		},
	}
}

func dataSourceSystemSpeedTestServerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemSpeedTestServer(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemSpeedTestServer dataSource: %v", err)
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

	diags := refreshObjectSystemSpeedTestServer(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
