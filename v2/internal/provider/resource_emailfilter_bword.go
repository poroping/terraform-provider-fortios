// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure AntiSpam banned word list.

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

func resourceEmailfilterBword() *schema.Resource {
	return &schema.Resource{
		Description: "Configure AntiSpam banned word list.",

		CreateContext: resourceEmailfilterBwordCreate,
		ReadContext:   resourceEmailfilterBwordRead,
		UpdateContext: resourceEmailfilterBwordUpdate,
		DeleteContext: resourceEmailfilterBwordDelete,

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
				Description:  "If set the provider will overwrite existing resources with the same mkey instead of erroring. Useful for brownfield implementations. Use with caution!",
				Optional:     true,
				Default:      false,
				RequiredWith: []string{"fosid"},
			},
			"dynamic_sort_subtable": {
				Type:        schema.TypeBool,
				Description: "If set will sort table response by mkey",
				Optional:    true,
				Default:     false,
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
				Description: "Spam filter banned word.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"spam", "clear"}, false),

							Description: "Mark spam or good.",
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Type: schema.TypeInt,

							Description: "Banned word entry ID.",
							Optional:    true,
							Computed:    true,
						},
						"language": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"western", "simch", "trach", "japanese", "korean", "french", "thai", "spanish"}, false),

							Description: "Language for the banned word.",
							Optional:    true,
							Computed:    true,
						},
						"pattern": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 127),

							Description: "Pattern for the banned word.",
							Optional:    true,
							Computed:    true,
						},
						"pattern_type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"wildcard", "regexp"}, false),

							Description: "Wildcard pattern or regular expression.",
							Optional:    true,
							Computed:    true,
						},
						"score": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 99999),

							Description: "Score value.",
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable status.",
							Optional:    true,
							Computed:    true,
						},
						"where": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"subject", "body", "all"}, false),

							Description: "Component of the email to be scanned.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"fosid": {
				Type: schema.TypeInt,

				Description: "ID.",
				ForceNew:    true,
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

func resourceEmailfilterBwordCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating EmailfilterBword resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectEmailfilterBword(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateEmailfilterBword(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("EmailfilterBword")
	}

	return resourceEmailfilterBwordRead(ctx, d, meta)
}

func resourceEmailfilterBwordUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectEmailfilterBword(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateEmailfilterBword(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating EmailfilterBword resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("EmailfilterBword")
	}

	return resourceEmailfilterBwordRead(ctx, d, meta)
}

func resourceEmailfilterBwordDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteEmailfilterBword(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting EmailfilterBword resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceEmailfilterBwordRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadEmailfilterBword(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading EmailfilterBword resource: %v", err)
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

	diags := refreshObjectEmailfilterBword(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenEmailfilterBwordEntries(d *schema.ResourceData, v *[]models.EmailfilterBwordEntries, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Action; tmp != nil {
				v["action"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Language; tmp != nil {
				v["language"] = *tmp
			}

			if tmp := cfg.Pattern; tmp != nil {
				v["pattern"] = *tmp
			}

			if tmp := cfg.PatternType; tmp != nil {
				v["pattern_type"] = *tmp
			}

			if tmp := cfg.Score; tmp != nil {
				v["score"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			if tmp := cfg.Where; tmp != nil {
				v["where"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func refreshObjectEmailfilterBword(d *schema.ResourceData, o *models.EmailfilterBword, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Comment != nil {
		v := *o.Comment

		if err = d.Set("comment", v); err != nil {
			return diag.Errorf("error reading comment: %v", err)
		}
	}

	if o.Entries != nil {
		if err = d.Set("entries", flattenEmailfilterBwordEntries(d, o.Entries, "entries", sort)); err != nil {
			return diag.Errorf("error reading entries: %v", err)
		}
	}

	if o.Id != nil {
		v := *o.Id

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

func expandEmailfilterBwordEntries(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.EmailfilterBwordEntries, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.EmailfilterBwordEntries

	for i := range l {
		tmp := models.EmailfilterBwordEntries{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.action", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Action = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.language", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Language = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.pattern", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Pattern = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.pattern_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PatternType = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.score", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Score = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.where", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Where = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectEmailfilterBword(d *schema.ResourceData, sv string) (*models.EmailfilterBword, diag.Diagnostics) {
	obj := models.EmailfilterBword{}
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
		t, err := expandEmailfilterBwordEntries(d, v, "entries", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Entries = t
		}
	} else if d.HasChange("entries") {
		old, new := d.GetChange("entries")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Entries = &[]models.EmailfilterBwordEntries{}
		}
	}
	if v1, ok := d.GetOk("fosid"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fosid", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Id = &tmp
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
