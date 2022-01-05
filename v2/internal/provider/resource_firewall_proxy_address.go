// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure web proxy address.

package provider

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/suppressors"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func resourceFirewallProxyAddress() *schema.Resource {
	return &schema.Resource{
		Description: "Configure web proxy address.",

		CreateContext: resourceFirewallProxyAddressCreate,
		ReadContext:   resourceFirewallProxyAddressRead,
		UpdateContext: resourceFirewallProxyAddressUpdate,
		DeleteContext: resourceFirewallProxyAddressDelete,

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
			"case_sensitivity": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable to make the pattern case sensitive.",
				Optional:    true,
				Computed:    true,
			},
			"category": {
				Type:        schema.TypeList,
				Description: "FortiGuard category ID.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type: schema.TypeInt,

							Description: "Fortiguard category id.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"color": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 32),

				Description: "Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets value to 1).",
				Optional:    true,
				Computed:    true,
			},
			"comment": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Optional comments.",
				Optional:    true,
				Computed:    true,
			},
			"header": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "HTTP header name as a regular expression.",
				Optional:    true,
				Computed:    true,
			},
			"header_group": {
				Type:        schema.TypeList,
				Description: "HTTP header group.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"case_sensitivity": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Case sensitivity in pattern.",
							Optional:    true,
							Computed:    true,
						},
						"header": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),

							Description: "HTTP header regular expression.",
							Optional:    true,
							Computed:    true,
						},
						"header_name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "HTTP header.",
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Type: schema.TypeInt,

							Description: "ID.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"header_name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),

				Description: "Name of HTTP header.",
				Optional:    true,
				Computed:    true,
			},
			"host": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),

				Description: "Address object for the host.",
				Optional:    true,
				Computed:    true,
			},
			"host_regex": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Host name as a regular expression.",
				Optional:    true,
				Computed:    true,
			},
			"method": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "HTTP request methods to be used.",
				Optional:         true,
				Computed:         true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Address name.",
				ForceNew:    true,
				Required:    true,
			},
			"path": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "URL path as a regular expression.",
				Optional:    true,
				Computed:    true,
			},
			"query": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Match the query part of the URL as a regular expression.",
				Optional:    true,
				Computed:    true,
			},
			"referrer": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable use of referrer field in the HTTP header to match the address.",
				Optional:    true,
				Computed:    true,
			},
			"tagging": {
				Type:        schema.TypeList,
				Description: "Config object tagging.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"category": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "Tag category.",
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "Tagging entry name.",
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
				},
			},
			"type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"host-regex", "url", "category", "method", "ua", "header", "src-advanced", "dst-advanced"}, false),

				Description: "Proxy address type.",
				Optional:    true,
				Computed:    true,
			},
			"ua": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Names of browsers to be used as user agent.",
				Optional:         true,
				Computed:         true,
			},
			"uuid": {
				Type: schema.TypeString,

				Description: "Universally Unique Identifier (UUID; automatically assigned but can be manually reset).",
				Optional:    true,
				Computed:    true,
			},
			"visibility": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable visibility of the object in the GUI.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceFirewallProxyAddressCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating FirewallProxyAddress resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectFirewallProxyAddress(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateFirewallProxyAddress(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallProxyAddress")
	}

	return resourceFirewallProxyAddressRead(ctx, d, meta)
}

func resourceFirewallProxyAddressUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectFirewallProxyAddress(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateFirewallProxyAddress(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating FirewallProxyAddress resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallProxyAddress")
	}

	return resourceFirewallProxyAddressRead(ctx, d, meta)
}

func resourceFirewallProxyAddressDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteFirewallProxyAddress(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting FirewallProxyAddress resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallProxyAddressRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallProxyAddress(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallProxyAddress resource: %v", err)
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

	diags := refreshObjectFirewallProxyAddress(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenFirewallProxyAddressCategory(v *[]models.FirewallProxyAddressCategory, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenFirewallProxyAddressHeaderGroup(v *[]models.FirewallProxyAddressHeaderGroup, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.CaseSensitivity; tmp != nil {
				v["case_sensitivity"] = *tmp
			}

			if tmp := cfg.Header; tmp != nil {
				v["header"] = *tmp
			}

			if tmp := cfg.HeaderName; tmp != nil {
				v["header_name"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenFirewallProxyAddressTagging(v *[]models.FirewallProxyAddressTagging, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Category; tmp != nil {
				v["category"] = *tmp
			}

			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.Tags; tmp != nil {
				v["tags"] = flattenFirewallProxyAddressTaggingTags(tmp, sort)
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenFirewallProxyAddressTaggingTags(v *[]models.FirewallProxyAddressTaggingTags, sort bool) interface{} {
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

func refreshObjectFirewallProxyAddress(d *schema.ResourceData, o *models.FirewallProxyAddress, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.CaseSensitivity != nil {
		v := *o.CaseSensitivity

		if err = d.Set("case_sensitivity", v); err != nil {
			return diag.Errorf("error reading case_sensitivity: %v", err)
		}
	}

	if o.Category != nil {
		if err = d.Set("category", flattenFirewallProxyAddressCategory(o.Category, sort)); err != nil {
			return diag.Errorf("error reading category: %v", err)
		}
	}

	if o.Color != nil {
		v := *o.Color

		if err = d.Set("color", v); err != nil {
			return diag.Errorf("error reading color: %v", err)
		}
	}

	if o.Comment != nil {
		v := *o.Comment

		if err = d.Set("comment", v); err != nil {
			return diag.Errorf("error reading comment: %v", err)
		}
	}

	if o.Header != nil {
		v := *o.Header

		if err = d.Set("header", v); err != nil {
			return diag.Errorf("error reading header: %v", err)
		}
	}

	if o.HeaderGroup != nil {
		if err = d.Set("header_group", flattenFirewallProxyAddressHeaderGroup(o.HeaderGroup, sort)); err != nil {
			return diag.Errorf("error reading header_group: %v", err)
		}
	}

	if o.HeaderName != nil {
		v := *o.HeaderName

		if err = d.Set("header_name", v); err != nil {
			return diag.Errorf("error reading header_name: %v", err)
		}
	}

	if o.Host != nil {
		v := *o.Host

		if err = d.Set("host", v); err != nil {
			return diag.Errorf("error reading host: %v", err)
		}
	}

	if o.HostRegex != nil {
		v := *o.HostRegex

		if err = d.Set("host_regex", v); err != nil {
			return diag.Errorf("error reading host_regex: %v", err)
		}
	}

	if o.Method != nil {
		v := *o.Method

		if err = d.Set("method", v); err != nil {
			return diag.Errorf("error reading method: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.Path != nil {
		v := *o.Path

		if err = d.Set("path", v); err != nil {
			return diag.Errorf("error reading path: %v", err)
		}
	}

	if o.Query != nil {
		v := *o.Query

		if err = d.Set("query", v); err != nil {
			return diag.Errorf("error reading query: %v", err)
		}
	}

	if o.Referrer != nil {
		v := *o.Referrer

		if err = d.Set("referrer", v); err != nil {
			return diag.Errorf("error reading referrer: %v", err)
		}
	}

	if o.Tagging != nil {
		if err = d.Set("tagging", flattenFirewallProxyAddressTagging(o.Tagging, sort)); err != nil {
			return diag.Errorf("error reading tagging: %v", err)
		}
	}

	if o.Type != nil {
		v := *o.Type

		if err = d.Set("type", v); err != nil {
			return diag.Errorf("error reading type: %v", err)
		}
	}

	if o.Ua != nil {
		v := *o.Ua

		if err = d.Set("ua", v); err != nil {
			return diag.Errorf("error reading ua: %v", err)
		}
	}

	if o.Uuid != nil {
		v := *o.Uuid

		if err = d.Set("uuid", v); err != nil {
			return diag.Errorf("error reading uuid: %v", err)
		}
	}

	if o.Visibility != nil {
		v := *o.Visibility

		if err = d.Set("visibility", v); err != nil {
			return diag.Errorf("error reading visibility: %v", err)
		}
	}

	return nil
}

func expandFirewallProxyAddressCategory(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallProxyAddressCategory, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallProxyAddressCategory

	for i := range l {
		tmp := models.FirewallProxyAddressCategory{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandFirewallProxyAddressHeaderGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallProxyAddressHeaderGroup, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallProxyAddressHeaderGroup

	for i := range l {
		tmp := models.FirewallProxyAddressHeaderGroup{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.case_sensitivity", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.CaseSensitivity = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.header", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Header = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.header_name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.HeaderName = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandFirewallProxyAddressTagging(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallProxyAddressTagging, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallProxyAddressTagging

	for i := range l {
		tmp := models.FirewallProxyAddressTagging{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.category", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Category = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.tags", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandFirewallProxyAddressTaggingTags(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.FirewallProxyAddressTaggingTags
			// 	}
			tmp.Tags = v2

		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandFirewallProxyAddressTaggingTags(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallProxyAddressTaggingTags, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallProxyAddressTaggingTags

	for i := range l {
		tmp := models.FirewallProxyAddressTaggingTags{}
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

func getObjectFirewallProxyAddress(d *schema.ResourceData, sv string) (*models.FirewallProxyAddress, diag.Diagnostics) {
	obj := models.FirewallProxyAddress{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("case_sensitivity"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("case_sensitivity", sv)
				diags = append(diags, e)
			}
			obj.CaseSensitivity = &v2
		}
	}
	if v, ok := d.GetOk("category"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("category", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallProxyAddressCategory(d, v, "category", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Category = t
		}
	} else if d.HasChange("category") {
		old, new := d.GetChange("category")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Category = &[]models.FirewallProxyAddressCategory{}
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
	if v1, ok := d.GetOk("comment"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("comment", sv)
				diags = append(diags, e)
			}
			obj.Comment = &v2
		}
	}
	if v1, ok := d.GetOk("header"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("header", sv)
				diags = append(diags, e)
			}
			obj.Header = &v2
		}
	}
	if v, ok := d.GetOk("header_group"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("header_group", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallProxyAddressHeaderGroup(d, v, "header_group", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.HeaderGroup = t
		}
	} else if d.HasChange("header_group") {
		old, new := d.GetChange("header_group")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.HeaderGroup = &[]models.FirewallProxyAddressHeaderGroup{}
		}
	}
	if v1, ok := d.GetOk("header_name"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("header_name", sv)
				diags = append(diags, e)
			}
			obj.HeaderName = &v2
		}
	}
	if v1, ok := d.GetOk("host"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("host", sv)
				diags = append(diags, e)
			}
			obj.Host = &v2
		}
	}
	if v1, ok := d.GetOk("host_regex"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("host_regex", sv)
				diags = append(diags, e)
			}
			obj.HostRegex = &v2
		}
	}
	if v1, ok := d.GetOk("method"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("method", sv)
				diags = append(diags, e)
			}
			obj.Method = &v2
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
	if v1, ok := d.GetOk("path"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("path", sv)
				diags = append(diags, e)
			}
			obj.Path = &v2
		}
	}
	if v1, ok := d.GetOk("query"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("query", sv)
				diags = append(diags, e)
			}
			obj.Query = &v2
		}
	}
	if v1, ok := d.GetOk("referrer"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("referrer", sv)
				diags = append(diags, e)
			}
			obj.Referrer = &v2
		}
	}
	if v, ok := d.GetOk("tagging"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("tagging", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallProxyAddressTagging(d, v, "tagging", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Tagging = t
		}
	} else if d.HasChange("tagging") {
		old, new := d.GetChange("tagging")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Tagging = &[]models.FirewallProxyAddressTagging{}
		}
	}
	if v1, ok := d.GetOk("type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("type", sv)
				diags = append(diags, e)
			}
			obj.Type = &v2
		}
	}
	if v1, ok := d.GetOk("ua"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ua", sv)
				diags = append(diags, e)
			}
			obj.Ua = &v2
		}
	}
	if v1, ok := d.GetOk("uuid"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("uuid", sv)
				diags = append(diags, e)
			}
			obj.Uuid = &v2
		}
	}
	if v1, ok := d.GetOk("visibility"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.0") {
				e := utils.AttributeVersionWarning("visibility", sv)
				diags = append(diags, e)
			}
			obj.Visibility = &v2
		}
	}
	return &obj, diags
}
