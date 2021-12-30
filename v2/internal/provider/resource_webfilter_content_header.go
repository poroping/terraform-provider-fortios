// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure content types used by Web filter.

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

func resourceWebfilterContentHeader() *schema.Resource {
	return &schema.Resource{
		Description: "Configure content types used by Web filter.",

		CreateContext: resourceWebfilterContentHeaderCreate,
		ReadContext:   resourceWebfilterContentHeaderRead,
		UpdateContext: resourceWebfilterContentHeaderUpdate,
		DeleteContext: resourceWebfilterContentHeaderDelete,

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
				Type:         schema.TypeBool,
				Description:  "If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution!",
				Optional:     true,
				Default:      false,
				RequiredWith: []string{"fosid"},
			},
			"dynamic_sort_subtable": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"comment": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Optional comments.",
				Optional:    true,
				Computed:    true,
			},
			"entries": {
				Type:        schema.TypeList,
				Description: "Configure content types used by web filter.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"block", "allow", "exempt"}, false),

							Description: "Action to take for this content type.",
							Optional:    true,
							Computed:    true,
						},
						"category": {
							Type: schema.TypeString,

							Description: "Categories that this content type applies to.",
							Optional:    true,
							Computed:    true,
						},
						"pattern": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 31),

							Description: "Content type (regular expression).",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"fosid": {
				Type: schema.TypeInt,

				Description: "ID.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Name of table.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceWebfilterContentHeaderCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
	key := "id"
	if v, ok := d.GetOk(key); ok {
		mkey = utils.ParseMkey(v)
		if mkey == "" && allow_append {
			return diag.Errorf("error creating WebfilterContentHeader resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectWebfilterContentHeader(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateWebfilterContentHeader(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WebfilterContentHeader")
	}

	return resourceWebfilterContentHeaderRead(ctx, d, meta)
}

func resourceWebfilterContentHeaderUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectWebfilterContentHeader(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateWebfilterContentHeader(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating WebfilterContentHeader resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WebfilterContentHeader")
	}

	return resourceWebfilterContentHeaderRead(ctx, d, meta)
}

func resourceWebfilterContentHeaderDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteWebfilterContentHeader(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting WebfilterContentHeader resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWebfilterContentHeaderRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadWebfilterContentHeader(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WebfilterContentHeader resource: %v", err)
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

	diags := refreshObjectWebfilterContentHeader(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenWebfilterContentHeaderEntries(v *[]models.WebfilterContentHeaderEntries, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Action; tmp != nil {
				v["action"] = *tmp
			}

			if tmp := cfg.Category; tmp != nil {
				v["category"] = *tmp
			}

			if tmp := cfg.Pattern; tmp != nil {
				v["pattern"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "pattern")
	}

	return flat
}

func refreshObjectWebfilterContentHeader(d *schema.ResourceData, o *models.WebfilterContentHeader, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Comment != nil {
		v := *o.Comment

		if err = d.Set("comment", v); err != nil {
			return diag.Errorf("error reading comment: %v", err)
		}
	}

	if o.Entries != nil {
		if err = d.Set("entries", flattenWebfilterContentHeaderEntries(o.Entries, sort)); err != nil {
			return diag.Errorf("error reading entries: %v", err)
		}
	}

	if o.Fosid != nil {
		v := *o.Fosid

		if err = d.Set("fosid", v); err != nil {
			return diag.Errorf("error reading fosid: %v", err)
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

func expandWebfilterContentHeaderEntries(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WebfilterContentHeaderEntries, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WebfilterContentHeaderEntries

	for i := range l {
		tmp := models.WebfilterContentHeaderEntries{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.action", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Action = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.category", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Category = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.pattern", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Pattern = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectWebfilterContentHeader(d *schema.ResourceData, sv string) (*models.WebfilterContentHeader, diag.Diagnostics) {
	obj := models.WebfilterContentHeader{}
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
	if v, ok := d.GetOk("entries"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("entries", sv)
			diags = append(diags, e)
		}
		t, err := expandWebfilterContentHeaderEntries(d, v, "entries", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Entries = t
		}
	} else if d.HasChange("entries") {
		old, new := d.GetChange("entries")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Entries = &[]models.WebfilterContentHeaderEntries{}
		}
	}
	if v1, ok := d.GetOk("fosid"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fosid", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Fosid = &tmp
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
