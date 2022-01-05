// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure VideoFilter profile.

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

func resourceVideofilterProfile() *schema.Resource {
	return &schema.Resource{
		Description: "Configure VideoFilter profile.",

		CreateContext: resourceVideofilterProfileCreate,
		ReadContext:   resourceVideofilterProfileRead,
		UpdateContext: resourceVideofilterProfileUpdate,
		DeleteContext: resourceVideofilterProfileDelete,

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
			"comment": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Comment.",
				Optional:    true,
				Computed:    true,
			},
			"dailymotion": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable Dailymotion video source.",
				Optional:    true,
				Computed:    true,
			},
			"fortiguard_category": {
				Type:        schema.TypeList,
				Description: "Configure FortiGuard categories.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"filters": {
							Type:        schema.TypeList,
							Description: "Configure VideoFilter FortiGuard category.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"allow", "monitor", "block"}, false),

										Description: "VideoFilter action.",
										Optional:    true,
										Computed:    true,
									},
									"category_id": {
										Type: schema.TypeInt,

										Description: "Category ID.",
										Optional:    true,
										Computed:    true,
									},
									"id": {
										Type: schema.TypeInt,

										Description: "ID.",
										Optional:    true,
										Computed:    true,
									},
									"log": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

										Description: "Enable/disable logging.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
					},
				},
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name.",
				ForceNew:    true,
				Required:    true,
			},
			"replacemsg_group": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Replacement message group.",
				Optional:    true,
				Computed:    true,
			},
			"vimeo": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable Vimeo video source.",
				Optional:    true,
				Computed:    true,
			},
			"vimeo_restrict": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Set Vimeo-restrict (\"7\" = don't show mature content, \"134\" = don't show unrated and mature content). A value of cookie \"content_rating\".",
				Optional:    true,
				Computed:    true,
			},
			"youtube": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable YouTube video source.",
				Optional:    true,
				Computed:    true,
			},
			"youtube_channel_filter": {
				Type: schema.TypeInt,

				Description: "Set YouTube channel filter.",
				Optional:    true,
				Computed:    true,
			},
			"youtube_restrict": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"none", "strict", "moderate"}, false),

				Description: "Set YouTube-restrict mode.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceVideofilterProfileCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating VideofilterProfile resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectVideofilterProfile(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateVideofilterProfile(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("VideofilterProfile")
	}

	return resourceVideofilterProfileRead(ctx, d, meta)
}

func resourceVideofilterProfileUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectVideofilterProfile(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateVideofilterProfile(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating VideofilterProfile resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("VideofilterProfile")
	}

	return resourceVideofilterProfileRead(ctx, d, meta)
}

func resourceVideofilterProfileDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteVideofilterProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting VideofilterProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVideofilterProfileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadVideofilterProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading VideofilterProfile resource: %v", err)
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

	diags := refreshObjectVideofilterProfile(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenVideofilterProfileFortiguardCategory(v *[]models.VideofilterProfileFortiguardCategory, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Filters; tmp != nil {
				v["filters"] = flattenVideofilterProfileFortiguardCategoryFilters(tmp, sort)
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenVideofilterProfileFortiguardCategoryFilters(v *[]models.VideofilterProfileFortiguardCategoryFilters, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Action; tmp != nil {
				v["action"] = *tmp
			}

			if tmp := cfg.CategoryId; tmp != nil {
				v["category_id"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Log; tmp != nil {
				v["log"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func refreshObjectVideofilterProfile(d *schema.ResourceData, o *models.VideofilterProfile, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Comment != nil {
		v := *o.Comment

		if err = d.Set("comment", v); err != nil {
			return diag.Errorf("error reading comment: %v", err)
		}
	}

	if o.Dailymotion != nil {
		v := *o.Dailymotion

		if err = d.Set("dailymotion", v); err != nil {
			return diag.Errorf("error reading dailymotion: %v", err)
		}
	}

	if o.FortiguardCategory != nil {
		if err = d.Set("fortiguard_category", flattenVideofilterProfileFortiguardCategory(o.FortiguardCategory, sort)); err != nil {
			return diag.Errorf("error reading fortiguard_category: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.ReplacemsgGroup != nil {
		v := *o.ReplacemsgGroup

		if err = d.Set("replacemsg_group", v); err != nil {
			return diag.Errorf("error reading replacemsg_group: %v", err)
		}
	}

	if o.Vimeo != nil {
		v := *o.Vimeo

		if err = d.Set("vimeo", v); err != nil {
			return diag.Errorf("error reading vimeo: %v", err)
		}
	}

	if o.VimeoRestrict != nil {
		v := *o.VimeoRestrict

		if err = d.Set("vimeo_restrict", v); err != nil {
			return diag.Errorf("error reading vimeo_restrict: %v", err)
		}
	}

	if o.Youtube != nil {
		v := *o.Youtube

		if err = d.Set("youtube", v); err != nil {
			return diag.Errorf("error reading youtube: %v", err)
		}
	}

	if o.YoutubeChannelFilter != nil {
		v := *o.YoutubeChannelFilter

		if err = d.Set("youtube_channel_filter", v); err != nil {
			return diag.Errorf("error reading youtube_channel_filter: %v", err)
		}
	}

	if o.YoutubeRestrict != nil {
		v := *o.YoutubeRestrict

		if err = d.Set("youtube_restrict", v); err != nil {
			return diag.Errorf("error reading youtube_restrict: %v", err)
		}
	}

	return nil
}

func expandVideofilterProfileFortiguardCategory(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.VideofilterProfileFortiguardCategory, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.VideofilterProfileFortiguardCategory

	for i := range l {
		tmp := models.VideofilterProfileFortiguardCategory{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.filters", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandVideofilterProfileFortiguardCategoryFilters(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.VideofilterProfileFortiguardCategoryFilters
			// 	}
			tmp.Filters = v2

		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandVideofilterProfileFortiguardCategoryFilters(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.VideofilterProfileFortiguardCategoryFilters, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.VideofilterProfileFortiguardCategoryFilters

	for i := range l {
		tmp := models.VideofilterProfileFortiguardCategoryFilters{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.action", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Action = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.category_id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.CategoryId = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.log", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Log = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectVideofilterProfile(d *schema.ResourceData, sv string) (*models.VideofilterProfile, diag.Diagnostics) {
	obj := models.VideofilterProfile{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("comment"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("comment", sv)
				diags = append(diags, e)
			}
			obj.Comment = &v2
		}
	}
	if v1, ok := d.GetOk("dailymotion"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dailymotion", sv)
				diags = append(diags, e)
			}
			obj.Dailymotion = &v2
		}
	}
	if v, ok := d.GetOk("fortiguard_category"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("fortiguard_category", sv)
			diags = append(diags, e)
		}
		t, err := expandVideofilterProfileFortiguardCategory(d, v, "fortiguard_category", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.FortiguardCategory = t
		}
	} else if d.HasChange("fortiguard_category") {
		old, new := d.GetChange("fortiguard_category")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.FortiguardCategory = &[]models.VideofilterProfileFortiguardCategory{}
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
	if v1, ok := d.GetOk("replacemsg_group"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.1", "") {
				e := utils.AttributeVersionWarning("replacemsg_group", sv)
				diags = append(diags, e)
			}
			obj.ReplacemsgGroup = &v2
		}
	}
	if v1, ok := d.GetOk("vimeo"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("vimeo", sv)
				diags = append(diags, e)
			}
			obj.Vimeo = &v2
		}
	}
	if v1, ok := d.GetOk("vimeo_restrict"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v7.0.1") {
				e := utils.AttributeVersionWarning("vimeo_restrict", sv)
				diags = append(diags, e)
			}
			obj.VimeoRestrict = &v2
		}
	}
	if v1, ok := d.GetOk("youtube"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("youtube", sv)
				diags = append(diags, e)
			}
			obj.Youtube = &v2
		}
	}
	if v1, ok := d.GetOk("youtube_channel_filter"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("youtube_channel_filter", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.YoutubeChannelFilter = &tmp
		}
	}
	if v1, ok := d.GetOk("youtube_restrict"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v7.0.1") {
				e := utils.AttributeVersionWarning("youtube_restrict", sv)
				diags = append(diags, e)
			}
			obj.YoutubeRestrict = &v2
		}
	}
	return &obj, diags
}
