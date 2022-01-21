// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure object tagging.

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

func resourceSystemObjectTagging() *schema.Resource {
	return &schema.Resource{
		Description: "Configure object tagging.",

		CreateContext: resourceSystemObjectTaggingCreate,
		ReadContext:   resourceSystemObjectTaggingRead,
		UpdateContext: resourceSystemObjectTaggingUpdate,
		DeleteContext: resourceSystemObjectTaggingDelete,

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
			"address": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "mandatory", "optional"}, false),

				Description: "Address.",
				Optional:    true,
				Computed:    true,
			},
			"category": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Tag Category.",
				ForceNew:    true,
				Required:    true,
			},
			"color": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 32),

				Description: "Color of icon on the GUI.",
				Optional:    true,
				Computed:    true,
			},
			"device": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "mandatory", "optional"}, false),

				Description: "Device.",
				Optional:    true,
				Computed:    true,
			},
			"interface": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "mandatory", "optional"}, false),

				Description: "Interface.",
				Optional:    true,
				Computed:    true,
			},
			"multiple": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Allow multiple tag selection.",
				Optional:    true,
				Computed:    true,
			},
			"tags": {
				Type:        schema.TypeList,
				Description: "Tags.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Tag name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func resourceSystemObjectTaggingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
	key := "category"
	if v, ok := d.GetOk(key); ok {
		mkey = utils.ParseMkey(v)
		if mkey == "" && allow_append {
			return diag.Errorf("error creating SystemObjectTagging resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectSystemObjectTagging(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemObjectTagging(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemObjectTagging")
	}

	return resourceSystemObjectTaggingRead(ctx, d, meta)
}

func resourceSystemObjectTaggingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemObjectTagging(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemObjectTagging(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemObjectTagging resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemObjectTagging")
	}

	return resourceSystemObjectTaggingRead(ctx, d, meta)
}

func resourceSystemObjectTaggingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteSystemObjectTagging(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemObjectTagging resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemObjectTaggingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemObjectTagging(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemObjectTagging resource: %v", err)
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

	diags := refreshObjectSystemObjectTagging(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenSystemObjectTaggingTags(v *[]models.SystemObjectTaggingTags, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
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

func refreshObjectSystemObjectTagging(d *schema.ResourceData, o *models.SystemObjectTagging, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Address != nil {
		v := *o.Address

		if err = d.Set("address", v); err != nil {
			return diag.Errorf("error reading address: %v", err)
		}
	}

	if o.Category != nil {
		v := *o.Category

		if err = d.Set("category", v); err != nil {
			return diag.Errorf("error reading category: %v", err)
		}
	}

	if o.Color != nil {
		v := *o.Color

		if err = d.Set("color", v); err != nil {
			return diag.Errorf("error reading color: %v", err)
		}
	}

	if o.Device != nil {
		v := *o.Device

		if err = d.Set("device", v); err != nil {
			return diag.Errorf("error reading device: %v", err)
		}
	}

	if o.Interface != nil {
		v := *o.Interface

		if err = d.Set("interface", v); err != nil {
			return diag.Errorf("error reading interface: %v", err)
		}
	}

	if o.Multiple != nil {
		v := *o.Multiple

		if err = d.Set("multiple", v); err != nil {
			return diag.Errorf("error reading multiple: %v", err)
		}
	}

	if o.Tags != nil {
		if err = d.Set("tags", flattenSystemObjectTaggingTags(o.Tags, sort)); err != nil {
			return diag.Errorf("error reading tags: %v", err)
		}
	}

	return nil
}

func expandSystemObjectTaggingTags(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemObjectTaggingTags, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemObjectTaggingTags

	for i := range l {
		tmp := models.SystemObjectTaggingTags{}
		var pre_append string

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

func getObjectSystemObjectTagging(d *schema.ResourceData, sv string) (*models.SystemObjectTagging, diag.Diagnostics) {
	obj := models.SystemObjectTagging{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("address"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("address", sv)
				diags = append(diags, e)
			}
			obj.Address = &v2
		}
	}
	if v1, ok := d.GetOk("category"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("category", sv)
				diags = append(diags, e)
			}
			obj.Category = &v2
		}
	}
	if v1, ok := d.GetOk("color"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("color", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Color = &tmp
		}
	}
	if v1, ok := d.GetOk("device"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("device", sv)
				diags = append(diags, e)
			}
			obj.Device = &v2
		}
	}
	if v1, ok := d.GetOk("interface"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("interface", sv)
				diags = append(diags, e)
			}
			obj.Interface = &v2
		}
	}
	if v1, ok := d.GetOk("multiple"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("multiple", sv)
				diags = append(diags, e)
			}
			obj.Multiple = &v2
		}
	}
	if v, ok := d.GetOk("tags"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("tags", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemObjectTaggingTags(d, v, "tags", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Tags = t
		}
	} else if d.HasChange("tags") {
		old, new := d.GetChange("tags")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Tags = &[]models.SystemObjectTaggingTags{}
		}
	}
	return &obj, diags
}
