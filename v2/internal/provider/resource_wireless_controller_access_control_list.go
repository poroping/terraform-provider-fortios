// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure WiFi bridge access control list.

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

func resourceWirelessControllerAccessControlList() *schema.Resource {
	return &schema.Resource{
		Description: "Configure WiFi bridge access control list.",

		CreateContext: resourceWirelessControllerAccessControlListCreate,
		ReadContext:   resourceWirelessControllerAccessControlListRead,
		UpdateContext: resourceWirelessControllerAccessControlListUpdate,
		DeleteContext: resourceWirelessControllerAccessControlListDelete,

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
			"comment": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Description.",
				Optional:    true,
				Computed:    true,
			},
			"layer3_ipv4_rules": {
				Type:        schema.TypeList,
				Description: "AP ACL layer3 ipv4 rule list.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "deny"}, false),

							Description: "Policy action (allow | deny).",
							Optional:    true,
							Computed:    true,
						},
						"comment": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "Description.",
							Optional:    true,
							Computed:    true,
						},
						"dstaddr": {
							Type: schema.TypeString,

							Description: "Destination IP address (any | local-LAN | IPv4 address[/<network mask | mask length>], default = any).",
							Optional:    true,
							Computed:    true,
						},
						"dstport": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),

							Description: "Destination port (0 - 65535, default = 0, meaning any).",
							Optional:    true,
							Computed:    true,
						},
						"protocol": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),

							Description: "Protocol type as defined by IANA (0 - 255, default = 255, meaning any).",
							Optional:    true,
							Computed:    true,
						},
						"rule_id": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),

							Description: "Rule ID (1 - 65535).",
							Optional:    true,
							Computed:    true,
						},
						"srcaddr": {
							Type: schema.TypeString,

							Description: "Source IP address (any | local-LAN | IPv4 address[/<network mask | mask length>], default = any).",
							Optional:    true,
							Computed:    true,
						},
						"srcport": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),

							Description: "Source port (0 - 65535, default = 0, meaning any).",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"layer3_ipv6_rules": {
				Type:        schema.TypeList,
				Description: "AP ACL layer3 ipv6 rule list.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"allow", "deny"}, false),

							Description: "Policy action (allow | deny).",
							Optional:    true,
							Computed:    true,
						},
						"comment": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "Description.",
							Optional:    true,
							Computed:    true,
						},
						"dstaddr": {
							Type: schema.TypeString,

							Description: "Destination IPv6 address (any | local-LAN | IPv6 address[/prefix length]), default = any.",
							Optional:    true,
							Computed:    true,
						},
						"dstport": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),

							Description: "Destination port (0 - 65535, default = 0, meaning any).",
							Optional:    true,
							Computed:    true,
						},
						"protocol": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),

							Description: "Protocol type as defined by IANA (0 - 255, default = 255, meaning any).",
							Optional:    true,
							Computed:    true,
						},
						"rule_id": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),

							Description: "Rule ID (1 - 65535).",
							Optional:    true,
							Computed:    true,
						},
						"srcaddr": {
							Type: schema.TypeString,

							Description: "Source IPv6 address (any | local-LAN | IPv6 address[/prefix length]), default = any.",
							Optional:    true,
							Computed:    true,
						},
						"srcport": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),

							Description: "Source port (0 - 65535, default = 0, meaning any).",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "AP access control list name.",
				ForceNew:    true,
				Required:    true,
			},
		},
	}
}

func resourceWirelessControllerAccessControlListCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating WirelessControllerAccessControlList resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectWirelessControllerAccessControlList(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateWirelessControllerAccessControlList(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WirelessControllerAccessControlList")
	}

	return resourceWirelessControllerAccessControlListRead(ctx, d, meta)
}

func resourceWirelessControllerAccessControlListUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectWirelessControllerAccessControlList(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateWirelessControllerAccessControlList(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating WirelessControllerAccessControlList resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WirelessControllerAccessControlList")
	}

	return resourceWirelessControllerAccessControlListRead(ctx, d, meta)
}

func resourceWirelessControllerAccessControlListDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteWirelessControllerAccessControlList(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting WirelessControllerAccessControlList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerAccessControlListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadWirelessControllerAccessControlList(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WirelessControllerAccessControlList resource: %v", err)
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

	diags := refreshObjectWirelessControllerAccessControlList(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenWirelessControllerAccessControlListLayer3Ipv4Rules(d *schema.ResourceData, v *[]models.WirelessControllerAccessControlListLayer3Ipv4Rules, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Action; tmp != nil {
				v["action"] = *tmp
			}

			if tmp := cfg.Comment; tmp != nil {
				v["comment"] = *tmp
			}

			if tmp := cfg.Dstaddr; tmp != nil {
				v["dstaddr"] = *tmp
			}

			if tmp := cfg.Dstport; tmp != nil {
				v["dstport"] = *tmp
			}

			if tmp := cfg.Protocol; tmp != nil {
				v["protocol"] = *tmp
			}

			if tmp := cfg.RuleId; tmp != nil {
				v["rule_id"] = *tmp
			}

			if tmp := cfg.Srcaddr; tmp != nil {
				v["srcaddr"] = *tmp
			}

			if tmp := cfg.Srcport; tmp != nil {
				v["srcport"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "rule_id")
	}

	return flat
}

func flattenWirelessControllerAccessControlListLayer3Ipv6Rules(d *schema.ResourceData, v *[]models.WirelessControllerAccessControlListLayer3Ipv6Rules, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Action; tmp != nil {
				v["action"] = *tmp
			}

			if tmp := cfg.Comment; tmp != nil {
				v["comment"] = *tmp
			}

			if tmp := cfg.Dstaddr; tmp != nil {
				v["dstaddr"] = *tmp
			}

			if tmp := cfg.Dstport; tmp != nil {
				v["dstport"] = *tmp
			}

			if tmp := cfg.Protocol; tmp != nil {
				v["protocol"] = *tmp
			}

			if tmp := cfg.RuleId; tmp != nil {
				v["rule_id"] = *tmp
			}

			if tmp := cfg.Srcaddr; tmp != nil {
				v["srcaddr"] = *tmp
			}

			if tmp := cfg.Srcport; tmp != nil {
				v["srcport"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "rule_id")
	}

	return flat
}

func refreshObjectWirelessControllerAccessControlList(d *schema.ResourceData, o *models.WirelessControllerAccessControlList, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Comment != nil {
		v := *o.Comment

		if err = d.Set("comment", v); err != nil {
			return diag.Errorf("error reading comment: %v", err)
		}
	}

	if o.Layer3Ipv4Rules != nil {
		if err = d.Set("layer3_ipv4_rules", flattenWirelessControllerAccessControlListLayer3Ipv4Rules(d, o.Layer3Ipv4Rules, "layer3_ipv4_rules", sort)); err != nil {
			return diag.Errorf("error reading layer3_ipv4_rules: %v", err)
		}
	}

	if o.Layer3Ipv6Rules != nil {
		if err = d.Set("layer3_ipv6_rules", flattenWirelessControllerAccessControlListLayer3Ipv6Rules(d, o.Layer3Ipv6Rules, "layer3_ipv6_rules", sort)); err != nil {
			return diag.Errorf("error reading layer3_ipv6_rules: %v", err)
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

func expandWirelessControllerAccessControlListLayer3Ipv4Rules(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WirelessControllerAccessControlListLayer3Ipv4Rules, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WirelessControllerAccessControlListLayer3Ipv4Rules

	for i := range l {
		tmp := models.WirelessControllerAccessControlListLayer3Ipv4Rules{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.action", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Action = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.comment", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Comment = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dstaddr", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dstaddr = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dstport", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Dstport = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.protocol", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Protocol = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.rule_id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.RuleId = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.srcaddr", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Srcaddr = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.srcport", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Srcport = &v3
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandWirelessControllerAccessControlListLayer3Ipv6Rules(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WirelessControllerAccessControlListLayer3Ipv6Rules, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WirelessControllerAccessControlListLayer3Ipv6Rules

	for i := range l {
		tmp := models.WirelessControllerAccessControlListLayer3Ipv6Rules{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.action", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Action = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.comment", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Comment = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dstaddr", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dstaddr = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dstport", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Dstport = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.protocol", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Protocol = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.rule_id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.RuleId = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.srcaddr", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Srcaddr = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.srcport", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Srcport = &v3
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectWirelessControllerAccessControlList(d *schema.ResourceData, sv string) (*models.WirelessControllerAccessControlList, diag.Diagnostics) {
	obj := models.WirelessControllerAccessControlList{}
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
	if v, ok := d.GetOk("layer3_ipv4_rules"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("layer3_ipv4_rules", sv)
			diags = append(diags, e)
		}
		t, err := expandWirelessControllerAccessControlListLayer3Ipv4Rules(d, v, "layer3_ipv4_rules", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Layer3Ipv4Rules = t
		}
	} else if d.HasChange("layer3_ipv4_rules") {
		old, new := d.GetChange("layer3_ipv4_rules")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Layer3Ipv4Rules = &[]models.WirelessControllerAccessControlListLayer3Ipv4Rules{}
		}
	}
	if v, ok := d.GetOk("layer3_ipv6_rules"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("layer3_ipv6_rules", sv)
			diags = append(diags, e)
		}
		t, err := expandWirelessControllerAccessControlListLayer3Ipv6Rules(d, v, "layer3_ipv6_rules", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Layer3Ipv6Rules = t
		}
	} else if d.HasChange("layer3_ipv6_rules") {
		old, new := d.GetChange("layer3_ipv6_rules")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Layer3Ipv6Rules = &[]models.WirelessControllerAccessControlListLayer3Ipv6Rules{}
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
