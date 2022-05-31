// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure IPv6 address templates.

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

func resourceFirewallAddress6Template() *schema.Resource {
	return &schema.Resource{
		Description: "Configure IPv6 address templates.",

		CreateContext: resourceFirewallAddress6TemplateCreate,
		ReadContext:   resourceFirewallAddress6TemplateRead,
		UpdateContext: resourceFirewallAddress6TemplateUpdate,
		DeleteContext: resourceFirewallAddress6TemplateDelete,

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
			"fabric_object": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Security Fabric global object setting.",
				Optional:    true,
				Computed:    true,
			},
			"ip6": {
				Type:             schema.TypeString,
				ValidateFunc:     validators.FortiValidateIPv6Network,
				DiffSuppressFunc: suppressors.DiffCidrEqual,
				Description:      "IPv6 address prefix.",
				Optional:         true,
				Computed:         true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "IPv6 address template name.",
				ForceNew:    true,
				Required:    true,
			},
			"subnet_segment": {
				Type:        schema.TypeList,
				Description: "IPv6 subnet segments.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bits": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 16),

							Description: "Number of bits.",
							Optional:    true,
							Computed:    true,
						},
						"exclusive": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable exclusive value.",
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Type: schema.TypeInt,

							Description: "Subnet segment ID.",
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "Subnet segment name.",
							Optional:    true,
							Computed:    true,
						},
						"values": {
							Type:        schema.TypeList,
							Description: "Subnet segment values.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 63),

										Description: "Subnet segment value name.",
										Optional:    true,
										Computed:    true,
									},
									"value": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),

										Description: "Subnet segment value.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
					},
				},
			},
			"subnet_segment_count": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 6),

				Description: "Number of IPv6 subnet segments.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceFirewallAddress6TemplateCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating FirewallAddress6Template resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectFirewallAddress6Template(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateFirewallAddress6Template(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallAddress6Template")
	}

	return resourceFirewallAddress6TemplateRead(ctx, d, meta)
}

func resourceFirewallAddress6TemplateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectFirewallAddress6Template(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateFirewallAddress6Template(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating FirewallAddress6Template resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallAddress6Template")
	}

	return resourceFirewallAddress6TemplateRead(ctx, d, meta)
}

func resourceFirewallAddress6TemplateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteFirewallAddress6Template(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting FirewallAddress6Template resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallAddress6TemplateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallAddress6Template(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallAddress6Template resource: %v", err)
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

	diags := refreshObjectFirewallAddress6Template(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenFirewallAddress6TemplateSubnetSegment(d *schema.ResourceData, v *[]models.FirewallAddress6TemplateSubnetSegment, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Bits; tmp != nil {
				v["bits"] = *tmp
			}

			if tmp := cfg.Exclusive; tmp != nil {
				v["exclusive"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.Values; tmp != nil {
				v["values"] = flattenFirewallAddress6TemplateSubnetSegmentValues(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "values"), sort)
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenFirewallAddress6TemplateSubnetSegmentValues(d *schema.ResourceData, v *[]models.FirewallAddress6TemplateSubnetSegmentValues, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.Value; tmp != nil {
				v["value"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func refreshObjectFirewallAddress6Template(d *schema.ResourceData, o *models.FirewallAddress6Template, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.FabricObject != nil {
		v := *o.FabricObject

		if err = d.Set("fabric_object", v); err != nil {
			return diag.Errorf("error reading fabric_object: %v", err)
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

	if o.SubnetSegment != nil {
		if err = d.Set("subnet_segment", flattenFirewallAddress6TemplateSubnetSegment(d, o.SubnetSegment, "subnet_segment", sort)); err != nil {
			return diag.Errorf("error reading subnet_segment: %v", err)
		}
	}

	if o.SubnetSegmentCount != nil {
		v := *o.SubnetSegmentCount

		if err = d.Set("subnet_segment_count", v); err != nil {
			return diag.Errorf("error reading subnet_segment_count: %v", err)
		}
	}

	return nil
}

func expandFirewallAddress6TemplateSubnetSegment(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallAddress6TemplateSubnetSegment, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallAddress6TemplateSubnetSegment

	for i := range l {
		tmp := models.FirewallAddress6TemplateSubnetSegment{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.bits", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Bits = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.exclusive", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Exclusive = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.values", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandFirewallAddress6TemplateSubnetSegmentValues(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.FirewallAddress6TemplateSubnetSegmentValues
			// 	}
			tmp.Values = v2

		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandFirewallAddress6TemplateSubnetSegmentValues(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallAddress6TemplateSubnetSegmentValues, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallAddress6TemplateSubnetSegmentValues

	for i := range l {
		tmp := models.FirewallAddress6TemplateSubnetSegmentValues{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.value", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Value = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectFirewallAddress6Template(d *schema.ResourceData, sv string) (*models.FirewallAddress6Template, diag.Diagnostics) {
	obj := models.FirewallAddress6Template{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("fabric_object"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("fabric_object", sv)
				diags = append(diags, e)
			}
			obj.FabricObject = &v2
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
	if v, ok := d.GetOk("subnet_segment"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("subnet_segment", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallAddress6TemplateSubnetSegment(d, v, "subnet_segment", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.SubnetSegment = t
		}
	} else if d.HasChange("subnet_segment") {
		old, new := d.GetChange("subnet_segment")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.SubnetSegment = &[]models.FirewallAddress6TemplateSubnetSegment{}
		}
	}
	if v1, ok := d.GetOk("subnet_segment_count"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("subnet_segment_count", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.SubnetSegmentCount = &tmp
		}
	}
	return &obj, diags
}
