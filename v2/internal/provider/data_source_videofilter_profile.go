// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure VideoFilter profile.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceVideofilterProfile() *schema.Resource {
	return &schema.Resource{
		Description: "Configure VideoFilter profile.",

		ReadContext: dataSourceVideofilterProfileRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"comment": {
				Type:        schema.TypeString,
				Description: "Comment.",
				Computed:    true,
			},
			"dailymotion": {
				Type:        schema.TypeString,
				Description: "Enable/disable Dailymotion video source.",
				Computed:    true,
			},
			"fortiguard_category": {
				Type:        schema.TypeList,
				Description: "Configure FortiGuard categories.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"filters": {
							Type:        schema.TypeList,
							Description: "Configure VideoFilter FortiGuard category.",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": {
										Type:        schema.TypeString,
										Description: "VideoFilter action.",
										Computed:    true,
									},
									"category_id": {
										Type:        schema.TypeInt,
										Description: "Category ID.",
										Computed:    true,
									},
									"id": {
										Type:        schema.TypeInt,
										Description: "ID.",
										Computed:    true,
									},
									"log": {
										Type:        schema.TypeString,
										Description: "Enable/disable logging.",
										Computed:    true,
									},
								},
							},
						},
					},
				},
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Name.",
				Required:    true,
			},
			"replacemsg_group": {
				Type:        schema.TypeString,
				Description: "Replacement message group.",
				Computed:    true,
			},
			"vimeo": {
				Type:        schema.TypeString,
				Description: "Enable/disable Vimeo video source.",
				Computed:    true,
			},
			"vimeo_restrict": {
				Type:        schema.TypeString,
				Description: "Set Vimeo-restrict (\"7\" = don't show mature content, \"134\" = don't show unrated and mature content). A value of cookie \"content_rating\".",
				Computed:    true,
			},
			"youtube": {
				Type:        schema.TypeString,
				Description: "Enable/disable YouTube video source.",
				Computed:    true,
			},
			"youtube_channel_filter": {
				Type:        schema.TypeInt,
				Description: "Set YouTube channel filter.",
				Computed:    true,
			},
			"youtube_restrict": {
				Type:        schema.TypeString,
				Description: "Set YouTube-restrict mode.",
				Computed:    true,
			},
		},
	}
}

func dataSourceVideofilterProfileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadVideofilterProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading VideofilterProfile dataSource: %v", err)
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

	diags := refreshObjectVideofilterProfile(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
