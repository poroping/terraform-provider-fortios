// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure API users.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceSystemApiUser() *schema.Resource {
	return &schema.Resource{
		Description: "Configure API users.",

		ReadContext: dataSourceSystemApiUserRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"accprofile": {
				Type:        schema.TypeString,
				Description: "Admin user access profile.",
				Computed:    true,
			},
			"api_key": {
				Type:        schema.TypeString,
				Description: "Admin user password.",
				Computed:    true,
				Sensitive:   true,
			},
			"comments": {
				Type:        schema.TypeString,
				Description: "Comment.",
				Computed:    true,
			},
			"cors_allow_origin": {
				Type:        schema.TypeString,
				Description: "Value for Access-Control-Allow-Origin on API responses. Avoid using '*' if possible.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "User name.",
				Required:    true,
			},
			"peer_auth": {
				Type:        schema.TypeString,
				Description: "Enable/disable peer authentication.",
				Computed:    true,
			},
			"peer_group": {
				Type:        schema.TypeString,
				Description: "Peer group name.",
				Computed:    true,
			},
			"schedule": {
				Type:        schema.TypeString,
				Description: "Schedule name.",
				Computed:    true,
			},
			"trusthost": {
				Type:        schema.TypeList,
				Description: "Trusthost.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Description: "ID.",
							Computed:    true,
						},
						"ipv4_trusthost": {
							Type:        schema.TypeString,
							Description: "IPv4 trusted host address.",
							Computed:    true,
						},
						"ipv6_trusthost": {
							Type:        schema.TypeString,
							Description: "IPv6 trusted host address.",
							Computed:    true,
						},
						"type": {
							Type:        schema.TypeString,
							Description: "Trusthost type.",
							Computed:    true,
						},
					},
				},
			},
			"vdom": {
				Type:        schema.TypeList,
				Description: "Virtual domains.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Virtual domain name.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceSystemApiUserRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemApiUser(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemApiUser dataSource: %v", err)
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

	diags := refreshObjectSystemApiUser(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
