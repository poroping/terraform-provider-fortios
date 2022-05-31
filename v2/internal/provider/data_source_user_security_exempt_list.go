// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure security exemption list.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceUserSecurityExemptList() *schema.Resource {
	return &schema.Resource{
		Description: "Configure security exemption list.",

		ReadContext: dataSourceUserSecurityExemptListRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"description": {
				Type:        schema.TypeString,
				Description: "Description.",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Name of the exempt list.",
				Required:    true,
			},
			"rule": {
				Type:        schema.TypeList,
				Description: "Configure rules for exempting users from captive portal authentication.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dstaddr": {
							Type:        schema.TypeList,
							Description: "Destination addresses or address groups.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Description: "Address or group name.",
										Computed:    true,
									},
								},
							},
						},
						"id": {
							Type:        schema.TypeInt,
							Description: "ID.",
							Computed:    true,
						},
						"service": {
							Type:        schema.TypeList,
							Description: "Destination services.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Description: "Service name.",
										Computed:    true,
									},
								},
							},
						},
						"srcaddr": {
							Type:        schema.TypeList,
							Description: "Source addresses or address groups.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Description: "Address or group name.",
										Computed:    true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceUserSecurityExemptListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadUserSecurityExemptList(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading UserSecurityExemptList dataSource: %v", err)
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

	diags := refreshObjectUserSecurityExemptList(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
