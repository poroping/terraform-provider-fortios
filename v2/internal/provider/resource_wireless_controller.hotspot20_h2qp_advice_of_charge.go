// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure advice of charge.

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

func resourceWirelessControllerHotspot20H2qpAdviceOfCharge() *schema.Resource {
	return &schema.Resource{
		Description: "Configure advice of charge.",

		CreateContext: resourceWirelessControllerHotspot20H2qpAdviceOfChargeCreate,
		ReadContext:   resourceWirelessControllerHotspot20H2qpAdviceOfChargeRead,
		UpdateContext: resourceWirelessControllerHotspot20H2qpAdviceOfChargeUpdate,
		DeleteContext: resourceWirelessControllerHotspot20H2qpAdviceOfChargeDelete,

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
			"dynamic_sort_subtable": {
				Type:        schema.TypeBool,
				Description: "If set will sort table response by mkey",
				Optional:    true,
				Default:     false,
			},
			"aoc_list": {
				Type:        schema.TypeList,
				Description: "AOC list.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"nai_realm": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),

							Description: "NAI realm list name.",
							Optional:    true,
							Computed:    true,
						},
						"nai_realm_encoding": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 1),

							Description: "NAI realm encoding.",
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Advice of charge ID.",
							Optional:    true,
							Computed:    true,
						},
						"plan_info": {
							Type:        schema.TypeList,
							Description: "Plan info.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"currency": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 3),

										Description: "Currency code.",
										Optional:    true,
										Computed:    true,
									},
									"info_file": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 64),

										Description: "Info file.",
										Optional:    true,
										Computed:    true,
									},
									"lang": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 3),

										Description: "Language code.",
										Optional:    true,
										Computed:    true,
									},
									"name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),

										Description: "Plan name.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"time-based", "volume-based", "time-and-volume-based", "unlimited"}, false),

							Description: "Usage charge type.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Plan name.",
				ForceNew:    true,
				Required:    true,
			},
		},
	}
}

func resourceWirelessControllerHotspot20H2qpAdviceOfChargeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating WirelessControllerHotspot20H2qpAdviceOfCharge resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectWirelessControllerHotspot20H2qpAdviceOfCharge(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateWirelessControllerHotspot20H2qpAdviceOfCharge(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WirelessControllerHotspot20H2qpAdviceOfCharge")
	}

	return resourceWirelessControllerHotspot20H2qpAdviceOfChargeRead(ctx, d, meta)
}

func resourceWirelessControllerHotspot20H2qpAdviceOfChargeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectWirelessControllerHotspot20H2qpAdviceOfCharge(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateWirelessControllerHotspot20H2qpAdviceOfCharge(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating WirelessControllerHotspot20H2qpAdviceOfCharge resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WirelessControllerHotspot20H2qpAdviceOfCharge")
	}

	return resourceWirelessControllerHotspot20H2qpAdviceOfChargeRead(ctx, d, meta)
}

func resourceWirelessControllerHotspot20H2qpAdviceOfChargeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteWirelessControllerHotspot20H2qpAdviceOfCharge(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting WirelessControllerHotspot20H2qpAdviceOfCharge resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerHotspot20H2qpAdviceOfChargeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadWirelessControllerHotspot20H2qpAdviceOfCharge(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WirelessControllerHotspot20H2qpAdviceOfCharge resource: %v", err)
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

	diags := refreshObjectWirelessControllerHotspot20H2qpAdviceOfCharge(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenWirelessControllerHotspot20H2qpAdviceOfChargeAocList(d *schema.ResourceData, v *[]models.WirelessControllerHotspot20H2qpAdviceOfChargeAocList, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.NaiRealm; tmp != nil {
				v["nai_realm"] = *tmp
			}

			if tmp := cfg.NaiRealmEncoding; tmp != nil {
				v["nai_realm_encoding"] = *tmp
			}

			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.PlanInfo; tmp != nil {
				v["plan_info"] = flattenWirelessControllerHotspot20H2qpAdviceOfChargeAocListPlanInfo(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "plan_info"), sort)
			}

			if tmp := cfg.Type; tmp != nil {
				v["type"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenWirelessControllerHotspot20H2qpAdviceOfChargeAocListPlanInfo(d *schema.ResourceData, v *[]models.WirelessControllerHotspot20H2qpAdviceOfChargeAocListPlanInfo, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Currency; tmp != nil {
				v["currency"] = *tmp
			}

			if tmp := cfg.InfoFile; tmp != nil {
				v["info_file"] = *tmp
			}

			if tmp := cfg.Lang; tmp != nil {
				v["lang"] = *tmp
			}

			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func refreshObjectWirelessControllerHotspot20H2qpAdviceOfCharge(d *schema.ResourceData, o *models.WirelessControllerHotspot20H2qpAdviceOfCharge, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AocList != nil {
		if err = d.Set("aoc_list", flattenWirelessControllerHotspot20H2qpAdviceOfChargeAocList(d, o.AocList, "aoc_list", sort)); err != nil {
			return diag.Errorf("error reading aoc_list: %v", err)
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

func expandWirelessControllerHotspot20H2qpAdviceOfChargeAocList(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WirelessControllerHotspot20H2qpAdviceOfChargeAocList, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WirelessControllerHotspot20H2qpAdviceOfChargeAocList

	for i := range l {
		tmp := models.WirelessControllerHotspot20H2qpAdviceOfChargeAocList{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.nai_realm", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.NaiRealm = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.nai_realm_encoding", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.NaiRealmEncoding = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.plan_info", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandWirelessControllerHotspot20H2qpAdviceOfChargeAocListPlanInfo(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.WirelessControllerHotspot20H2qpAdviceOfChargeAocListPlanInfo
			// 	}
			tmp.PlanInfo = v2

		}

		pre_append = fmt.Sprintf("%s.%d.type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Type = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandWirelessControllerHotspot20H2qpAdviceOfChargeAocListPlanInfo(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WirelessControllerHotspot20H2qpAdviceOfChargeAocListPlanInfo, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WirelessControllerHotspot20H2qpAdviceOfChargeAocListPlanInfo

	for i := range l {
		tmp := models.WirelessControllerHotspot20H2qpAdviceOfChargeAocListPlanInfo{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.currency", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Currency = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.info_file", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.InfoFile = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.lang", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Lang = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectWirelessControllerHotspot20H2qpAdviceOfCharge(d *schema.ResourceData, sv string) (*models.WirelessControllerHotspot20H2qpAdviceOfCharge, diag.Diagnostics) {
	obj := models.WirelessControllerHotspot20H2qpAdviceOfCharge{}
	diags := diag.Diagnostics{}

	if v, ok := d.GetOk("aoc_list"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("aoc_list", sv)
			diags = append(diags, e)
		}
		t, err := expandWirelessControllerHotspot20H2qpAdviceOfChargeAocList(d, v, "aoc_list", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.AocList = t
		}
	} else if d.HasChange("aoc_list") {
		old, new := d.GetChange("aoc_list")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.AocList = &[]models.WirelessControllerHotspot20H2qpAdviceOfChargeAocList{}
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
