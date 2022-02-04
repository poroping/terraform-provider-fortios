// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure user password policy.

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

func resourceUserPasswordPolicy() *schema.Resource {
	return &schema.Resource{
		Description: "Configure user password policy.",

		CreateContext: resourceUserPasswordPolicyCreate,
		ReadContext:   resourceUserPasswordPolicyRead,
		UpdateContext: resourceUserPasswordPolicyUpdate,
		DeleteContext: resourceUserPasswordPolicyDelete,

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
			"expire_days": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 999),

				Description: "Time in days before the user's password expires.",
				Optional:    true,
				Computed:    true,
			},
			"expired_password_renewal": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable renewal of a password that already is expired.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Password policy name.",
				ForceNew:    true,
				Required:    true,
			},
			"warn_days": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 30),

				Description: "Time in days before a password expiration warning message is displayed to the user upon login.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceUserPasswordPolicyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating UserPasswordPolicy resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectUserPasswordPolicy(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateUserPasswordPolicy(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("UserPasswordPolicy")
	}

	return resourceUserPasswordPolicyRead(ctx, d, meta)
}

func resourceUserPasswordPolicyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectUserPasswordPolicy(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateUserPasswordPolicy(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating UserPasswordPolicy resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("UserPasswordPolicy")
	}

	return resourceUserPasswordPolicyRead(ctx, d, meta)
}

func resourceUserPasswordPolicyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteUserPasswordPolicy(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting UserPasswordPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserPasswordPolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadUserPasswordPolicy(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading UserPasswordPolicy resource: %v", err)
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

	diags := refreshObjectUserPasswordPolicy(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectUserPasswordPolicy(d *schema.ResourceData, o *models.UserPasswordPolicy, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.ExpireDays != nil {
		v := *o.ExpireDays

		if err = d.Set("expire_days", v); err != nil {
			return diag.Errorf("error reading expire_days: %v", err)
		}
	}

	if o.ExpiredPasswordRenewal != nil {
		v := *o.ExpiredPasswordRenewal

		if err = d.Set("expired_password_renewal", v); err != nil {
			return diag.Errorf("error reading expired_password_renewal: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.WarnDays != nil {
		v := *o.WarnDays

		if err = d.Set("warn_days", v); err != nil {
			return diag.Errorf("error reading warn_days: %v", err)
		}
	}

	return nil
}

func getObjectUserPasswordPolicy(d *schema.ResourceData, sv string) (*models.UserPasswordPolicy, diag.Diagnostics) {
	obj := models.UserPasswordPolicy{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("expire_days"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("expire_days", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.ExpireDays = &tmp
		}
	}
	if v1, ok := d.GetOk("expired_password_renewal"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("expired_password_renewal", sv)
				diags = append(diags, e)
			}
			obj.ExpiredPasswordRenewal = &v2
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
	if v1, ok := d.GetOk("warn_days"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("warn_days", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.WarnDays = &tmp
		}
	}
	return &obj, diags
}
