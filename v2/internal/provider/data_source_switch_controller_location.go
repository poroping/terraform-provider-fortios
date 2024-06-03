// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure FortiSwitch location services.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func dataSourceSwitchControllerLocation() *schema.Resource {
	return &schema.Resource{
		Description: "Configure FortiSwitch location services.",

		ReadContext: dataSourceSwitchControllerLocationRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the dataSource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"address_civic": {
				Type:        schema.TypeList,
				Description: "Configure location civic address.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional": {
							Type:        schema.TypeString,
							Description: "Location additional details.",
							Computed:    true,
						},
						"additional_code": {
							Type:        schema.TypeString,
							Description: "Location additional code details.",
							Computed:    true,
						},
						"block": {
							Type:        schema.TypeString,
							Description: "Location block details.",
							Computed:    true,
						},
						"branch_road": {
							Type:        schema.TypeString,
							Description: "Location branch road details.",
							Computed:    true,
						},
						"building": {
							Type:        schema.TypeString,
							Description: "Location building details.",
							Computed:    true,
						},
						"city": {
							Type:        schema.TypeString,
							Description: "Location city details.",
							Computed:    true,
						},
						"city_division": {
							Type:        schema.TypeString,
							Description: "Location city division details.",
							Computed:    true,
						},
						"country": {
							Type:        schema.TypeString,
							Description: "The two-letter ISO 3166 country code in capital ASCII letters eg. US, CA, DK, DE.",
							Computed:    true,
						},
						"country_subdivision": {
							Type:        schema.TypeString,
							Description: "National subdivisions (state, canton, region, province, or prefecture).",
							Computed:    true,
						},
						"county": {
							Type:        schema.TypeString,
							Description: "County, parish, gun (JP), or district (IN).",
							Computed:    true,
						},
						"direction": {
							Type:        schema.TypeString,
							Description: "Leading street direction.",
							Computed:    true,
						},
						"floor": {
							Type:        schema.TypeString,
							Description: "Floor.",
							Computed:    true,
						},
						"landmark": {
							Type:        schema.TypeString,
							Description: "Landmark or vanity address.",
							Computed:    true,
						},
						"language": {
							Type:        schema.TypeString,
							Description: "Language.",
							Computed:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "Name (residence and office occupant).",
							Computed:    true,
						},
						"number": {
							Type:        schema.TypeString,
							Description: "House number.",
							Computed:    true,
						},
						"number_suffix": {
							Type:        schema.TypeString,
							Description: "House number suffix.",
							Computed:    true,
						},
						"parent_key": {
							Type:        schema.TypeString,
							Description: "Parent key name.",
							Computed:    true,
						},
						"place_type": {
							Type:        schema.TypeString,
							Description: "Place type.",
							Computed:    true,
						},
						"post_office_box": {
							Type:        schema.TypeString,
							Description: "Post office box.",
							Computed:    true,
						},
						"postal_community": {
							Type:        schema.TypeString,
							Description: "Postal community name.",
							Computed:    true,
						},
						"primary_road": {
							Type:        schema.TypeString,
							Description: "Primary road name.",
							Computed:    true,
						},
						"road_section": {
							Type:        schema.TypeString,
							Description: "Road section.",
							Computed:    true,
						},
						"room": {
							Type:        schema.TypeString,
							Description: "Room number.",
							Computed:    true,
						},
						"script": {
							Type:        schema.TypeString,
							Description: "Script used to present the address information.",
							Computed:    true,
						},
						"seat": {
							Type:        schema.TypeString,
							Description: "Seat number.",
							Computed:    true,
						},
						"street": {
							Type:        schema.TypeString,
							Description: "Street.",
							Computed:    true,
						},
						"street_name_post_mod": {
							Type:        schema.TypeString,
							Description: "Street name post modifier.",
							Computed:    true,
						},
						"street_name_pre_mod": {
							Type:        schema.TypeString,
							Description: "Street name pre modifier.",
							Computed:    true,
						},
						"street_suffix": {
							Type:        schema.TypeString,
							Description: "Street suffix.",
							Computed:    true,
						},
						"sub_branch_road": {
							Type:        schema.TypeString,
							Description: "Sub branch road name.",
							Computed:    true,
						},
						"trailing_str_suffix": {
							Type:        schema.TypeString,
							Description: "Trailing street suffix.",
							Computed:    true,
						},
						"unit": {
							Type:        schema.TypeString,
							Description: "Unit (apartment, suite).",
							Computed:    true,
						},
						"zip": {
							Type:        schema.TypeString,
							Description: "Postal/zip code.",
							Computed:    true,
						},
					},
				},
			},
			"coordinates": {
				Type:        schema.TypeList,
				Description: "Configure location GPS coordinates.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"altitude": {
							Type:        schema.TypeString,
							Description: "Plus or minus floating point number. For example, 117.47.",
							Computed:    true,
						},
						"altitude_unit": {
							Type:        schema.TypeString,
							Description: "Configure the unit for which the altitude is to (m = meters, f = floors of a building).",
							Computed:    true,
						},
						"datum": {
							Type:        schema.TypeString,
							Description: "WGS84, NAD83, NAD83/MLLW.",
							Computed:    true,
						},
						"latitude": {
							Type:        schema.TypeString,
							Description: "Floating point starting with +/- or ending with (N or S). For example, +/-16.67 or 16.67N.",
							Computed:    true,
						},
						"longitude": {
							Type:        schema.TypeString,
							Description: "Floating point starting with +/- or ending with (N or S). For example, +/-26.789 or 26.789E.",
							Computed:    true,
						},
						"parent_key": {
							Type:        schema.TypeString,
							Description: "Parent key name.",
							Computed:    true,
						},
					},
				},
			},
			"elin_number": {
				Type:        schema.TypeList,
				Description: "Configure location ELIN number.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"elin_num": {
							Type:        schema.TypeString,
							Description: "Configure ELIN callback number.",
							Computed:    true,
						},
						"parent_key": {
							Type:        schema.TypeString,
							Description: "Parent key name.",
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Unique location item name.",
				Required:    true,
			},
		},
	}
}

func dataSourceSwitchControllerLocationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSwitchControllerLocation(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SwitchControllerLocation dataSource: %v", err)
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

	diags := refreshObjectSwitchControllerLocation(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}

	d.SetId(mkey)

	return nil
}
