// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure IPv6 multicast address.

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
	"github.com/poroping/terraform-provider-fortios/v2/validators"
)

func resourceFirewallMulticastAddress6() *schema.Resource {
	return &schema.Resource{
		Description: "Configure IPv6 multicast address.",

		CreateContext: resourceFirewallMulticastAddress6Create,
		ReadContext:   resourceFirewallMulticastAddress6Read,
		UpdateContext: resourceFirewallMulticastAddress6Update,
		DeleteContext: resourceFirewallMulticastAddress6Delete,

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
				Description: "If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution!",
				Optional:    true,
				Default:     false,
			},
			"dynamic_sort_subtable": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"color": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 32),

				Description: "Color of icon on the GUI.",
				Optional:    true,
				Computed:    true,
			},
			"comment": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Comment.",
				Optional:    true,
				Computed:    true,
			},
			"ip6": {
				Type:             schema.TypeString,
				ValidateFunc:     validators.FortiValidateIPv6Network,
				DiffSuppressFunc: suppressors.DiffCidrEqual,
				Description:      "IPv6 address prefix (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx).",
				Optional:         true,
				Computed:         true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),

				Description: "IPv6 multicast address name.",
				ForceNew:    true,
				Required:    true,
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
			"visibility": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable visibility of the IPv6 multicast address on the GUI.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceFirewallMulticastAddress6Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating FirewallMulticastAddress6 resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectFirewallMulticastAddress6(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateFirewallMulticastAddress6(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallMulticastAddress6")
	}

	return resourceFirewallMulticastAddress6Read(ctx, d, meta)
}

func resourceFirewallMulticastAddress6Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectFirewallMulticastAddress6(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateFirewallMulticastAddress6(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating FirewallMulticastAddress6 resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallMulticastAddress6")
	}

	return resourceFirewallMulticastAddress6Read(ctx, d, meta)
}

func resourceFirewallMulticastAddress6Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteFirewallMulticastAddress6(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting FirewallMulticastAddress6 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallMulticastAddress6Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallMulticastAddress6(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallMulticastAddress6 resource: %v", err)
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

	diags := refreshObjectFirewallMulticastAddress6(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenFirewallMulticastAddress6Tagging(v *[]models.FirewallMulticastAddress6Tagging, sort bool) interface{} {
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
				v["tags"] = flattenFirewallMulticastAddress6TaggingTags(tmp, sort)
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenFirewallMulticastAddress6TaggingTags(v *[]models.FirewallMulticastAddress6TaggingTags, sort bool) interface{} {
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

func refreshObjectFirewallMulticastAddress6(d *schema.ResourceData, o *models.FirewallMulticastAddress6, sv string, sort bool) diag.Diagnostics {
	var err error

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

	if o.Ip6 != nil {
		v := *o.Ip6

		if err = d.Set("ip6", v); err != nil {
			return diag.Errorf("error reading ip6: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.Tagging != nil {
		if err = d.Set("tagging", flattenFirewallMulticastAddress6Tagging(o.Tagging, sort)); err != nil {
			return diag.Errorf("error reading tagging: %v", err)
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

func expandFirewallMulticastAddress6Tagging(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallMulticastAddress6Tagging, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallMulticastAddress6Tagging

	for i := range l {
		tmp := models.FirewallMulticastAddress6Tagging{}
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
			v2, _ := expandFirewallMulticastAddress6TaggingTags(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.FirewallMulticastAddress6TaggingTags
			// 	}
			tmp.Tags = v2

		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandFirewallMulticastAddress6TaggingTags(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallMulticastAddress6TaggingTags, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallMulticastAddress6TaggingTags

	for i := range l {
		tmp := models.FirewallMulticastAddress6TaggingTags{}
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

func getObjectFirewallMulticastAddress6(d *schema.ResourceData, sv string) (*models.FirewallMulticastAddress6, diag.Diagnostics) {
	obj := models.FirewallMulticastAddress6{}
	diags := diag.Diagnostics{}

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
	if v1, ok := d.GetOk("ip6"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ip6", sv)
				diags = append(diags, e)
			}
			obj.Ip6 = &v2
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
	if v, ok := d.GetOk("tagging"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("tagging", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallMulticastAddress6Tagging(d, v, "tagging", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Tagging = t
		}
	} else if d.HasChange("tagging") {
		old, new := d.GetChange("tagging")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Tagging = &[]models.FirewallMulticastAddress6Tagging{}
		}
	}
	if v1, ok := d.GetOk("visibility"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("visibility", sv)
				diags = append(diags, e)
			}
			obj.Visibility = &v2
		}
	}
	return &obj, diags
}
