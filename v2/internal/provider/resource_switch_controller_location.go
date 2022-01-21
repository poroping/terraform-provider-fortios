// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure FortiSwitch location services.

package provider

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func resourceSwitchControllerLocation() *schema.Resource {
	return &schema.Resource{
		Description: "Configure FortiSwitch location services.",

		CreateContext: resourceSwitchControllerLocationCreate,
		ReadContext:   resourceSwitchControllerLocationRead,
		UpdateContext: resourceSwitchControllerLocationUpdate,
		DeleteContext: resourceSwitchControllerLocationDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:        schema.TypeString,
				Description: "Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. If you want to inherit the VDOM configuration of the provider, do not set this parameter.",
				Optional:    true,
				ForceNew:    true,
			},
			"allow_append": {
				Type:        schema.TypeBool,
				Description: "If set the provider will overwrite existing resources with the same mkey instead of erroring. Useful for brownfield implementations. Use with caution!",
				Optional:    true,
				Default:     false,
			},
			"address_civic": {
				Type:        schema.TypeList,
				Description: "Configure location civic address.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),

							Description: "Location additional details.",
							Optional:    true,
							Computed:    true,
						},
						"additional_code": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),

							Description: "Location additional code details.",
							Optional:    true,
							Computed:    true,
						},
						"block": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),

							Description: "Location block details.",
							Optional:    true,
							Computed:    true,
						},
						"branch_road": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),

							Description: "Location branch road details.",
							Optional:    true,
							Computed:    true,
						},
						"building": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),

							Description: "Location building details.",
							Optional:    true,
							Computed:    true,
						},
						"city": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),

							Description: "Location city details.",
							Optional:    true,
							Computed:    true,
						},
						"city_division": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),

							Description: "Location city division details.",
							Optional:    true,
							Computed:    true,
						},
						"country": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),

							Description: "The two-letter ISO 3166 country code in capital ASCII letters eg. US, CA, DK, DE.",
							Optional:    true,
							Computed:    true,
						},
						"country_subdivision": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),

							Description: "National subdivisions (state, canton, region, province, or prefecture).",
							Optional:    true,
							Computed:    true,
						},
						"county": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),

							Description: "County, parish, gun (JP), or district (IN).",
							Optional:    true,
							Computed:    true,
						},
						"direction": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),

							Description: "Leading street direction.",
							Optional:    true,
							Computed:    true,
						},
						"floor": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),

							Description: "Floor.",
							Optional:    true,
							Computed:    true,
						},
						"landmark": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),

							Description: "Landmark or vanity address.",
							Optional:    true,
							Computed:    true,
						},
						"language": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),

							Description: "Language.",
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),

							Description: "Name (residence and office occupant).",
							Optional:    true,
							Computed:    true,
						},
						"number": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),

							Description: "House number.",
							Optional:    true,
							Computed:    true,
						},
						"number_suffix": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),

							Description: "House number suffix.",
							Optional:    true,
							Computed:    true,
						},
						"parent_key": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "Parent key name.",
							Optional:    true,
							Computed:    true,
						},
						"place_type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),

							Description: "Placetype.",
							Optional:    true,
							Computed:    true,
						},
						"post_office_box": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),

							Description: "Post office box (P.O. box).",
							Optional:    true,
							Computed:    true,
						},
						"postal_community": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),

							Description: "Postal community name.",
							Optional:    true,
							Computed:    true,
						},
						"primary_road": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),

							Description: "Primary road name.",
							Optional:    true,
							Computed:    true,
						},
						"road_section": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),

							Description: "Road section.",
							Optional:    true,
							Computed:    true,
						},
						"room": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),

							Description: "Room number.",
							Optional:    true,
							Computed:    true,
						},
						"script": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),

							Description: "Script used to present the address information.",
							Optional:    true,
							Computed:    true,
						},
						"seat": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),

							Description: "Seat number.",
							Optional:    true,
							Computed:    true,
						},
						"street": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),

							Description: "Street.",
							Optional:    true,
							Computed:    true,
						},
						"street_name_post_mod": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),

							Description: "Street name post modifier.",
							Optional:    true,
							Computed:    true,
						},
						"street_name_pre_mod": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),

							Description: "Street name pre modifier.",
							Optional:    true,
							Computed:    true,
						},
						"street_suffix": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),

							Description: "Street suffix.",
							Optional:    true,
							Computed:    true,
						},
						"sub_branch_road": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),

							Description: "Sub branch road name.",
							Optional:    true,
							Computed:    true,
						},
						"trailing_str_suffix": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),

							Description: "Trailing street suffix.",
							Optional:    true,
							Computed:    true,
						},
						"unit": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),

							Description: "Unit (apartment, suite).",
							Optional:    true,
							Computed:    true,
						},
						"zip": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 47),

							Description: "Postal/zip code.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"coordinates": {
				Type:        schema.TypeList,
				Description: "Configure location GPS coordinates.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"altitude": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),

							Description: "+/- Floating point no. eg. 117.47.",
							Optional:    true,
							Computed:    true,
						},
						"altitude_unit": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"m", "f"}, false),

							Description: "m ( meters), f ( floors).",
							Optional:    true,
							Computed:    true,
						},
						"datum": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"WGS84", "NAD83", "NAD83/MLLW"}, false),

							Description: "WGS84, NAD83, NAD83/MLLW.",
							Optional:    true,
							Computed:    true,
						},
						"latitude": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),

							Description: "Floating point start with ( +/- )  or end with ( N or S ) eg. +/-16.67 or 16.67N.",
							Optional:    true,
							Computed:    true,
						},
						"longitude": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),

							Description: "Floating point start with ( +/- )  or end with ( E or W ) eg. +/-26.789 or 26.789E.",
							Optional:    true,
							Computed:    true,
						},
						"parent_key": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "Parent key name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"elin_number": {
				Type:        schema.TypeList,
				Description: "Configure location ELIN number.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"elin_num": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 31),

							Description: "Configure ELIN callback number.",
							Optional:    true,
							Computed:    true,
						},
						"parent_key": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "Parent key name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Unique location item name.",
				ForceNew:    true,
				Required:    true,
			},
		},
	}
}

func resourceSwitchControllerLocationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*apiClient).Client
	var diags diag.Diagnostics
	var err error
	// c.Retries = 1
	urlparams := &models.CmdbRequestParams{}
	vdomparam := ""
	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}
	urlparams.Vdom = vdomparam

	allow_append := false
	if v, ok := d.GetOk("allow_append"); ok {
		if b, ok := v.(bool); ok {
			allow_append = b
		}
	}
	urlparams.AllowAppend = &allow_append

	mkey := ""
	key := "name"
	if v, ok := d.GetOk(key); ok {
		mkey = utils.ParseMkey(v)
		if mkey == "" && allow_append {
			return diag.Errorf("error creating SwitchControllerLocation resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectSwitchControllerLocation(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSwitchControllerLocation(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SwitchControllerLocation")
	}

	return resourceSwitchControllerLocationRead(ctx, d, meta)
}

func resourceSwitchControllerLocationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	mkey := d.Id()
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

	obj, diags := getObjectSwitchControllerLocation(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSwitchControllerLocation(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SwitchControllerLocation resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SwitchControllerLocation")
	}

	return resourceSwitchControllerLocationRead(ctx, d, meta)
}

func resourceSwitchControllerLocationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	mkey := d.Id()

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

	err := c.Cmdb.DeleteSwitchControllerLocation(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SwitchControllerLocation resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerLocationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	mkey := d.Id()

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

	ptp := true
	urlparams.PlainTextPassword = &ptp

	o, err := c.Cmdb.ReadSwitchControllerLocation(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SwitchControllerLocation resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
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
	return nil
}

func flattenSwitchControllerLocationAddressCivic(v *[]models.SwitchControllerLocationAddressCivic, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Additional; tmp != nil {
				v["additional"] = *tmp
			}

			if tmp := cfg.AdditionalCode; tmp != nil {
				v["additional_code"] = *tmp
			}

			if tmp := cfg.Block; tmp != nil {
				v["block"] = *tmp
			}

			if tmp := cfg.BranchRoad; tmp != nil {
				v["branch_road"] = *tmp
			}

			if tmp := cfg.Building; tmp != nil {
				v["building"] = *tmp
			}

			if tmp := cfg.City; tmp != nil {
				v["city"] = *tmp
			}

			if tmp := cfg.CityDivision; tmp != nil {
				v["city_division"] = *tmp
			}

			if tmp := cfg.Country; tmp != nil {
				v["country"] = *tmp
			}

			if tmp := cfg.CountrySubdivision; tmp != nil {
				v["country_subdivision"] = *tmp
			}

			if tmp := cfg.County; tmp != nil {
				v["county"] = *tmp
			}

			if tmp := cfg.Direction; tmp != nil {
				v["direction"] = *tmp
			}

			if tmp := cfg.Floor; tmp != nil {
				v["floor"] = *tmp
			}

			if tmp := cfg.Landmark; tmp != nil {
				v["landmark"] = *tmp
			}

			if tmp := cfg.Language; tmp != nil {
				v["language"] = *tmp
			}

			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.Number; tmp != nil {
				v["number"] = *tmp
			}

			if tmp := cfg.NumberSuffix; tmp != nil {
				v["number_suffix"] = *tmp
			}

			if tmp := cfg.ParentKey; tmp != nil {
				v["parent_key"] = *tmp
			}

			if tmp := cfg.PlaceType; tmp != nil {
				v["place_type"] = *tmp
			}

			if tmp := cfg.PostOfficeBox; tmp != nil {
				v["post_office_box"] = *tmp
			}

			if tmp := cfg.PostalCommunity; tmp != nil {
				v["postal_community"] = *tmp
			}

			if tmp := cfg.PrimaryRoad; tmp != nil {
				v["primary_road"] = *tmp
			}

			if tmp := cfg.RoadSection; tmp != nil {
				v["road_section"] = *tmp
			}

			if tmp := cfg.Room; tmp != nil {
				v["room"] = *tmp
			}

			if tmp := cfg.Script; tmp != nil {
				v["script"] = *tmp
			}

			if tmp := cfg.Seat; tmp != nil {
				v["seat"] = *tmp
			}

			if tmp := cfg.Street; tmp != nil {
				v["street"] = *tmp
			}

			if tmp := cfg.StreetNamePostMod; tmp != nil {
				v["street_name_post_mod"] = *tmp
			}

			if tmp := cfg.StreetNamePreMod; tmp != nil {
				v["street_name_pre_mod"] = *tmp
			}

			if tmp := cfg.StreetSuffix; tmp != nil {
				v["street_suffix"] = *tmp
			}

			if tmp := cfg.SubBranchRoad; tmp != nil {
				v["sub_branch_road"] = *tmp
			}

			if tmp := cfg.TrailingStrSuffix; tmp != nil {
				v["trailing_str_suffix"] = *tmp
			}

			if tmp := cfg.Unit; tmp != nil {
				v["unit"] = *tmp
			}

			if tmp := cfg.Zip; tmp != nil {
				v["zip"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenSwitchControllerLocationCoordinates(v *[]models.SwitchControllerLocationCoordinates, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Altitude; tmp != nil {
				v["altitude"] = *tmp
			}

			if tmp := cfg.AltitudeUnit; tmp != nil {
				v["altitude_unit"] = *tmp
			}

			if tmp := cfg.Datum; tmp != nil {
				v["datum"] = *tmp
			}

			if tmp := cfg.Latitude; tmp != nil {
				v["latitude"] = *tmp
			}

			if tmp := cfg.Longitude; tmp != nil {
				v["longitude"] = *tmp
			}

			if tmp := cfg.ParentKey; tmp != nil {
				v["parent_key"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenSwitchControllerLocationElinNumber(v *[]models.SwitchControllerLocationElinNumber, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.ElinNum; tmp != nil {
				v["elin_num"] = *tmp
			}

			if tmp := cfg.ParentKey; tmp != nil {
				v["parent_key"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func refreshObjectSwitchControllerLocation(d *schema.ResourceData, o *models.SwitchControllerLocation, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AddressCivic != nil {
		if err = d.Set("address_civic", flattenSwitchControllerLocationAddressCivic(o.AddressCivic, sort)); err != nil {
			return diag.Errorf("error reading address_civic: %v", err)
		}
	}

	if o.Coordinates != nil {
		if err = d.Set("coordinates", flattenSwitchControllerLocationCoordinates(o.Coordinates, sort)); err != nil {
			return diag.Errorf("error reading coordinates: %v", err)
		}
	}

	if o.ElinNumber != nil {
		if err = d.Set("elin_number", flattenSwitchControllerLocationElinNumber(o.ElinNumber, sort)); err != nil {
			return diag.Errorf("error reading elin_number: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	return nil
}

func expandSwitchControllerLocationAddressCivic(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SwitchControllerLocationAddressCivic, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SwitchControllerLocationAddressCivic

	for i := range l {
		tmp := models.SwitchControllerLocationAddressCivic{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.additional", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Additional = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.additional_code", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AdditionalCode = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.block", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Block = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.branch_road", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.BranchRoad = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.building", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Building = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.city", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.City = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.city_division", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.CityDivision = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.country", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Country = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.country_subdivision", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.CountrySubdivision = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.county", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.County = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.direction", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Direction = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.floor", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Floor = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.landmark", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Landmark = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.language", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Language = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.number", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Number = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.number_suffix", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.NumberSuffix = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.parent_key", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ParentKey = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.place_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PlaceType = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.post_office_box", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PostOfficeBox = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.postal_community", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PostalCommunity = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.primary_road", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PrimaryRoad = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.road_section", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.RoadSection = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.room", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Room = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.script", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Script = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.seat", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Seat = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.street", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Street = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.street_name_post_mod", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.StreetNamePostMod = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.street_name_pre_mod", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.StreetNamePreMod = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.street_suffix", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.StreetSuffix = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.sub_branch_road", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SubBranchRoad = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.trailing_str_suffix", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.TrailingStrSuffix = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.unit", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Unit = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.zip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Zip = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSwitchControllerLocationCoordinates(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SwitchControllerLocationCoordinates, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SwitchControllerLocationCoordinates

	for i := range l {
		tmp := models.SwitchControllerLocationCoordinates{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.altitude", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Altitude = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.altitude_unit", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AltitudeUnit = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.datum", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Datum = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.latitude", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Latitude = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.longitude", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Longitude = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.parent_key", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ParentKey = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSwitchControllerLocationElinNumber(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SwitchControllerLocationElinNumber, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SwitchControllerLocationElinNumber

	for i := range l {
		tmp := models.SwitchControllerLocationElinNumber{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.elin_num", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ElinNum = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.parent_key", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.ParentKey = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectSwitchControllerLocation(d *schema.ResourceData, sv string) (*models.SwitchControllerLocation, diag.Diagnostics) {
	obj := models.SwitchControllerLocation{}
	diags := diag.Diagnostics{}

	if v, ok := d.GetOk("address_civic"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("address_civic", sv)
			diags = append(diags, e)
		}
		t, err := expandSwitchControllerLocationAddressCivic(d, v, "address_civic", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.AddressCivic = t
		}
	} else if d.HasChange("address_civic") {
		old, new := d.GetChange("address_civic")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.AddressCivic = &[]models.SwitchControllerLocationAddressCivic{}
		}
	}
	if v, ok := d.GetOk("coordinates"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("coordinates", sv)
			diags = append(diags, e)
		}
		t, err := expandSwitchControllerLocationCoordinates(d, v, "coordinates", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Coordinates = t
		}
	} else if d.HasChange("coordinates") {
		old, new := d.GetChange("coordinates")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Coordinates = &[]models.SwitchControllerLocationCoordinates{}
		}
	}
	if v, ok := d.GetOk("elin_number"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("elin_number", sv)
			diags = append(diags, e)
		}
		t, err := expandSwitchControllerLocationElinNumber(d, v, "elin_number", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.ElinNumber = t
		}
	} else if d.HasChange("elin_number") {
		old, new := d.GetChange("elin_number")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.ElinNumber = &[]models.SwitchControllerLocationElinNumber{}
		}
	}
	if v1, ok := d.GetOk("name"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("name", sv)
				diags = append(diags, e)
			}
			obj.Name = &v2
		}
	}
	return &obj, diags
}
