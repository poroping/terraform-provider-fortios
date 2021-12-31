// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: IPv4 prefix list rule.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func resourceRouterprefixListRule() *schema.Resource {
	return &schema.Resource{
		Description: "IPv4 prefix list rule.",

		CreateContext: resourceRouterprefixListRuleCreate,
		ReadContext:   resourceRouterprefixListRuleRead,
		UpdateContext: resourceRouterprefixListRuleUpdate,
		DeleteContext: resourceRouterprefixListRuleDelete,

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
			"fosid": {
				Type: schema.TypeInt,

				Description: "Rule ID.",
				ForceNew:    true,
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

				Description: "IPv4 prefix to define regular filter criteria, such as \"any\" or subnets.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceRouterprefixListRuleCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating RouterprefixListRule resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectRouterprefixListRule(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateRouterprefixListRule(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("RouterprefixListRule")
	}

	return resourceRouterprefixListRuleRead(ctx, d, meta)
}

func resourceRouterprefixListRuleUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectRouterprefixListRule(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateRouterprefixListRule(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating RouterprefixListRule resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("RouterprefixListRule")
	}

	return resourceRouterprefixListRuleRead(ctx, d, meta)
}

func resourceRouterprefixListRuleDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteRouterprefixListRule(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting RouterprefixListRule resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterprefixListRuleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadRouterprefixListRule(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading RouterprefixListRule resource: %v", err)
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

	diags := refreshObjectRouterprefixListRule(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectRouterprefixListRule(d *schema.ResourceData, o *models.RouterprefixListRule, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Action != nil {
		v := *o.Action

		if err = d.Set("action", v); err != nil {
			return diag.Errorf("error reading action: %v", err)
		}
	}

	if o.Flags != nil {
		v := *o.Flags

		if err = d.Set("flags", v); err != nil {
			return diag.Errorf("error reading flags: %v", err)
		}
	}

	if o.Ge != nil {
		v := *o.Ge

		if err = d.Set("ge", v); err != nil {
			return diag.Errorf("error reading ge: %v", err)
		}
	}

	if o.Id != nil {
		v := *o.Id

		if err = d.Set("fosid", v); err != nil {
			return diag.Errorf("error reading fosid: %v", err)
		}
	}

	if o.Le != nil {
		v := *o.Le

		if err = d.Set("le", v); err != nil {
			return diag.Errorf("error reading le: %v", err)
		}
	}

	if o.Prefix != nil {
		v := *o.Prefix

		if err = d.Set("prefix", v); err != nil {
			return diag.Errorf("error reading prefix: %v", err)
		}
	}

	return nil
}

func getObjectRouterprefixListRule(d *schema.ResourceData, sv string) (*models.RouterprefixListRule, diag.Diagnostics) {
	obj := models.RouterprefixListRule{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("action"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("action", sv)
				diags = append(diags, e)
			}
			obj.Action = &v2
		}
	}
	if v1, ok := d.GetOk("flags"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "v7.0.0") {
				e := utils.AttributeVersionWarning("flags", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Flags = &tmp
		}
	}
	if v1, ok := d.GetOk("ge"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ge", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Ge = &tmp
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
	if v1, ok := d.GetOk("le"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("le", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Le = &tmp
		}
	}
	if v1, ok := d.GetOk("prefix"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("prefix", sv)
				diags = append(diags, e)
			}
			obj.Prefix = &v2
		}
	}
	return &obj, diags
}
