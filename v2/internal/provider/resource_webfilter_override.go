// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure FortiGuard Web Filter administrative overrides.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/suppressors"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func resourceWebfilterOverride() *schema.Resource {
	return &schema.Resource{
		Description: "Configure FortiGuard Web Filter administrative overrides.",

		CreateContext: resourceWebfilterOverrideCreate,
		ReadContext:   resourceWebfilterOverrideRead,
		UpdateContext: resourceWebfilterOverrideUpdate,
		DeleteContext: resourceWebfilterOverrideDelete,

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
			"expires": {
				Type: schema.TypeString,

				Description: "Override expiration date and time, from 5 minutes to 365 from now (format: yyyy/mm/dd hh:mm:ss).",
				Optional:    true,
				Computed:    true,
			},
			"fosid": {
				Type: schema.TypeInt,

				Description: "Override rule ID.",
				ForceNew:    true,
				Optional:    true,
				Computed:    true,
			},
			"initiator": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 64),

				Description: "Initiating user of override (read-only setting).",
				Optional:    true,
				Computed:    true,
			},
			"ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "IPv4 address which the override applies.",
				Optional:    true,
				Computed:    true,
			},
			"ip6": {
				Type:             schema.TypeString,
				ValidateFunc:     validation.IsIPv6Address,
				DiffSuppressFunc: suppressors.DiffIPEqual,
				Description:      "IPv6 address which the override applies.",
				Optional:         true,
				Computed:         true,
			},
			"new_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of the new web filter profile used by the override.",
				Optional:    true,
				Computed:    true,
			},
			"old_profile": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of the web filter profile which the override applies.",
				Optional:    true,
				Computed:    true,
			},
			"scope": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"user", "user-group", "ip", "ip6"}, false),

				Description: "Override either the specific user, user group, IPv4 address, or IPv6 address.",
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable override rule.",
				Optional:    true,
				Computed:    true,
			},
			"user": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 64),

				Description: "Name of the user which the override applies.",
				Optional:    true,
				Computed:    true,
			},
			"user_group": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Specify the user group for which the override applies.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceWebfilterOverrideCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating WebfilterOverride resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectWebfilterOverride(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateWebfilterOverride(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WebfilterOverride")
	}

	return resourceWebfilterOverrideRead(ctx, d, meta)
}

func resourceWebfilterOverrideUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectWebfilterOverride(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateWebfilterOverride(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating WebfilterOverride resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WebfilterOverride")
	}

	return resourceWebfilterOverrideRead(ctx, d, meta)
}

func resourceWebfilterOverrideDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteWebfilterOverride(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting WebfilterOverride resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWebfilterOverrideRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadWebfilterOverride(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WebfilterOverride resource: %v", err)
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

	diags := refreshObjectWebfilterOverride(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectWebfilterOverride(d *schema.ResourceData, o *models.WebfilterOverride, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Expires != nil {
		v := *o.Expires

		if err = d.Set("expires", v); err != nil {
			return diag.Errorf("error reading expires: %v", err)
		}
	}

	if o.Id != nil {
		v := *o.Id

		if err = d.Set("fosid", v); err != nil {
			return diag.Errorf("error reading fosid: %v", err)
		}
	}

	if o.Initiator != nil {
		v := *o.Initiator

		if err = d.Set("initiator", v); err != nil {
			return diag.Errorf("error reading initiator: %v", err)
		}
	}

	if o.Ip != nil {
		v := *o.Ip

		if err = d.Set("ip", v); err != nil {
			return diag.Errorf("error reading ip: %v", err)
		}
	}

	if o.Ip6 != nil {
		v := *o.Ip6

		if err = d.Set("ip6", v); err != nil {
			return diag.Errorf("error reading ip6: %v", err)
		}
	}

	if o.NewProfile != nil {
		v := *o.NewProfile

		if err = d.Set("new_profile", v); err != nil {
			return diag.Errorf("error reading new_profile: %v", err)
		}
	}

	if o.OldProfile != nil {
		v := *o.OldProfile

		if err = d.Set("old_profile", v); err != nil {
			return diag.Errorf("error reading old_profile: %v", err)
		}
	}

	if o.Scope != nil {
		v := *o.Scope

		if err = d.Set("scope", v); err != nil {
			return diag.Errorf("error reading scope: %v", err)
		}
	}

	if o.Status != nil {
		v := *o.Status

		if err = d.Set("status", v); err != nil {
			return diag.Errorf("error reading status: %v", err)
		}
	}

	if o.User != nil {
		v := *o.User

		if err = d.Set("user", v); err != nil {
			return diag.Errorf("error reading user: %v", err)
		}
	}

	if o.UserGroup != nil {
		v := *o.UserGroup

		if err = d.Set("user_group", v); err != nil {
			return diag.Errorf("error reading user_group: %v", err)
		}
	}

	return nil
}

func getObjectWebfilterOverride(d *schema.ResourceData, sv string) (*models.WebfilterOverride, diag.Diagnostics) {
	obj := models.WebfilterOverride{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("expires"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("expires", sv)
				diags = append(diags, e)
			}
			obj.Expires = &v2
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
	if v1, ok := d.GetOk("initiator"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("initiator", sv)
				diags = append(diags, e)
			}
			obj.Initiator = &v2
		}
	}
	if v1, ok := d.GetOk("ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ip", sv)
				diags = append(diags, e)
			}
			obj.Ip = &v2
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
	if v1, ok := d.GetOk("new_profile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("new_profile", sv)
				diags = append(diags, e)
			}
			obj.NewProfile = &v2
		}
	}
	if v1, ok := d.GetOk("old_profile"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("old_profile", sv)
				diags = append(diags, e)
			}
			obj.OldProfile = &v2
		}
	}
	if v1, ok := d.GetOk("scope"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("scope", sv)
				diags = append(diags, e)
			}
			obj.Scope = &v2
		}
	}
	if v1, ok := d.GetOk("status"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("status", sv)
				diags = append(diags, e)
			}
			obj.Status = &v2
		}
	}
	if v1, ok := d.GetOk("user"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("user", sv)
				diags = append(diags, e)
			}
			obj.User = &v2
		}
	}
	if v1, ok := d.GetOk("user_group"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("user_group", sv)
				diags = append(diags, e)
			}
			obj.UserGroup = &v2
		}
	}
	return &obj, diags
}
