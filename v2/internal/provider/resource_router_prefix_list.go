// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure IPv4 prefix lists.

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

func resourceRouterPrefixList() *schema.Resource {
	return &schema.Resource{
		Description: "Configure IPv4 prefix lists.",

		CreateContext: resourceRouterPrefixListCreate,
		ReadContext:   resourceRouterPrefixListRead,
		UpdateContext: resourceRouterPrefixListUpdate,
		DeleteContext: resourceRouterPrefixListDelete,

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
			"comments": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),

				Description: "Comment.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name.",
				ForceNew:    true,
				Required:    true,
			},
			"rule": {
				Type:        schema.TypeList,
				Description: "IPv4 prefix list rule.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"permit", "deny"}, false),

							Description: "Permit or deny this IP address and netmask prefix.",
							Optional:    true,
							Computed:    true,
						},
						"flags": {
							Type: schema.TypeInt,

							Description: "Flags.",
							Optional:    true,
							Computed:    true,
						},
						"ge": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 32),

							Description: "Minimum prefix length to be matched (0 - 32).",
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Type: schema.TypeInt,

							Description: "Rule ID.",
							Optional:    true,
							Computed:    true,
						},
						"le": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 32),

							Description: "Maximum prefix length to be matched (0 - 32).",
							Optional:    true,
							Computed:    true,
						},
						"prefix": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffCidrEqual,
							Description:      "IPv4 prefix to define regular filter criteria, such as \"any\" or subnets.",
							Optional:         true,
							Computed:         true,
						},
					},
				},
			},
		},
	}
}

func resourceRouterPrefixListCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating RouterPrefixList resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectRouterPrefixList(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateRouterPrefixList(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("RouterPrefixList")
	}

	return resourceRouterPrefixListRead(ctx, d, meta)
}

func resourceRouterPrefixListUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectRouterPrefixList(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateRouterPrefixList(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating RouterPrefixList resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("RouterPrefixList")
	}

	return resourceRouterPrefixListRead(ctx, d, meta)
}

func resourceRouterPrefixListDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteRouterPrefixList(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting RouterPrefixList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterPrefixListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadRouterPrefixList(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading RouterPrefixList resource: %v", err)
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

	diags := refreshObjectRouterPrefixList(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenRouterPrefixListRule(d *schema.ResourceData, v *[]models.RouterPrefixListRule, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Action; tmp != nil {
				v["action"] = *tmp
			}

			if tmp := cfg.Flags; tmp != nil {
				v["flags"] = *tmp
			}

			if tmp := cfg.Ge; tmp != nil {
				v["ge"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Le; tmp != nil {
				v["le"] = *tmp
			}

			if tmp := cfg.Prefix; tmp != nil {
				v["prefix"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func refreshObjectRouterPrefixList(d *schema.ResourceData, o *models.RouterPrefixList, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Comments != nil {
		v := *o.Comments

		if err = d.Set("comments", v); err != nil {
			return diag.Errorf("error reading comments: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.Rule != nil {
		if err = d.Set("rule", flattenRouterPrefixListRule(d, o.Rule, "rule", sort)); err != nil {
			return diag.Errorf("error reading rule: %v", err)
		}
	}

	return nil
}

func expandRouterPrefixListRule(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterPrefixListRule, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterPrefixListRule

	for i := range l {
		tmp := models.RouterPrefixListRule{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.action", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Action = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.flags", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Flags = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ge", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Ge = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.le", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Le = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.prefix", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Prefix = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectRouterPrefixList(d *schema.ResourceData, sv string) (*models.RouterPrefixList, diag.Diagnostics) {
	obj := models.RouterPrefixList{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("comments"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("comments", sv)
				diags = append(diags, e)
			}
			obj.Comments = &v2
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
	if v, ok := d.GetOk("rule"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("rule", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterPrefixListRule(d, v, "rule", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Rule = t
		}
	} else if d.HasChange("rule") {
		old, new := d.GetChange("rule")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Rule = &[]models.RouterPrefixListRule{}
		}
	}
	return &obj, diags
}
