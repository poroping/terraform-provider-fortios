// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure the password policy for guest administrators.

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

func resourceSystemPasswordPolicyGuestAdmin() *schema.Resource {
	return &schema.Resource{
		Description: "Configure the password policy for guest administrators.",

		CreateContext: resourceSystemPasswordPolicyGuestAdminCreate,
		ReadContext:   resourceSystemPasswordPolicyGuestAdminRead,
		UpdateContext: resourceSystemPasswordPolicyGuestAdminUpdate,
		DeleteContext: resourceSystemPasswordPolicyGuestAdminDelete,

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
			"apply_to": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Guest administrator to which this password policy applies.",
				Optional:         true,
				Computed:         true,
			},
			"change_4_characters": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable changing at least 4 characters for a new password (This attribute overrides reuse-password if both are enabled).",
				Optional:    true,
				Computed:    true,
			},
			"expire_day": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 999),

				Description: "Number of days after which passwords expire (1 - 999 days, default = 90).",
				Optional:    true,
				Computed:    true,
			},
			"expire_status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable password expiration.",
				Optional:    true,
				Computed:    true,
			},
			"min_change_characters": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 128),

				Description: "Minimum number of unique characters in new password which do not exist in old password (0 - 128, default = 0. This attribute overrides reuse-password if both are enabled).",
				Optional:    true,
				Computed:    true,
			},
			"min_lower_case_letter": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 128),

				Description: "Minimum number of lowercase characters in password (0 - 128, default = 0).",
				Optional:    true,
				Computed:    true,
			},
			"min_non_alphanumeric": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 128),

				Description: "Minimum number of non-alphanumeric characters in password (0 - 128, default = 0).",
				Optional:    true,
				Computed:    true,
			},
			"min_number": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 128),

				Description: "Minimum number of numeric characters in password (0 - 128, default = 0).",
				Optional:    true,
				Computed:    true,
			},
			"min_upper_case_letter": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 128),

				Description: "Minimum number of uppercase characters in password (0 - 128, default = 0).",
				Optional:    true,
				Computed:    true,
			},
			"minimum_length": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(8, 128),

				Description: "Minimum password length (8 - 128, default = 8).",
				Optional:    true,
				Computed:    true,
			},
			"reuse_password": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable reuse of password. If both reuse-password and min-change-characters are enabled, min-change-characters overrides.",
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable setting a password policy for locally defined administrator passwords and IPsec VPN pre-shared keys.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSystemPasswordPolicyGuestAdminCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemPasswordPolicyGuestAdmin(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemPasswordPolicyGuestAdmin(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemPasswordPolicyGuestAdmin")
	}

	return resourceSystemPasswordPolicyGuestAdminRead(ctx, d, meta)
}

func resourceSystemPasswordPolicyGuestAdminUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemPasswordPolicyGuestAdmin(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemPasswordPolicyGuestAdmin(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemPasswordPolicyGuestAdmin resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemPasswordPolicyGuestAdmin")
	}

	return resourceSystemPasswordPolicyGuestAdminRead(ctx, d, meta)
}

func resourceSystemPasswordPolicyGuestAdminDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectSystemPasswordPolicyGuestAdmin(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateSystemPasswordPolicyGuestAdmin(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemPasswordPolicyGuestAdmin resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemPasswordPolicyGuestAdminRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemPasswordPolicyGuestAdmin(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemPasswordPolicyGuestAdmin resource: %v", err)
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

	diags := refreshObjectSystemPasswordPolicyGuestAdmin(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectSystemPasswordPolicyGuestAdmin(d *schema.ResourceData, o *models.SystemPasswordPolicyGuestAdmin, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.ApplyTo != nil {
		v := *o.ApplyTo

		if err = d.Set("apply_to", v); err != nil {
			return diag.Errorf("error reading apply_to: %v", err)
		}
	}

	if o.Change4Characters != nil {
		v := *o.Change4Characters

		if err = d.Set("change_4_characters", v); err != nil {
			return diag.Errorf("error reading change_4_characters: %v", err)
		}
	}

	if o.ExpireDay != nil {
		v := *o.ExpireDay

		if err = d.Set("expire_day", v); err != nil {
			return diag.Errorf("error reading expire_day: %v", err)
		}
	}

	if o.ExpireStatus != nil {
		v := *o.ExpireStatus

		if err = d.Set("expire_status", v); err != nil {
			return diag.Errorf("error reading expire_status: %v", err)
		}
	}

	if o.MinChangeCharacters != nil {
		v := *o.MinChangeCharacters

		if err = d.Set("min_change_characters", v); err != nil {
			return diag.Errorf("error reading min_change_characters: %v", err)
		}
	}

	if o.MinLowerCaseLetter != nil {
		v := *o.MinLowerCaseLetter

		if err = d.Set("min_lower_case_letter", v); err != nil {
			return diag.Errorf("error reading min_lower_case_letter: %v", err)
		}
	}

	if o.MinNonAlphanumeric != nil {
		v := *o.MinNonAlphanumeric

		if err = d.Set("min_non_alphanumeric", v); err != nil {
			return diag.Errorf("error reading min_non_alphanumeric: %v", err)
		}
	}

	if o.MinNumber != nil {
		v := *o.MinNumber

		if err = d.Set("min_number", v); err != nil {
			return diag.Errorf("error reading min_number: %v", err)
		}
	}

	if o.MinUpperCaseLetter != nil {
		v := *o.MinUpperCaseLetter

		if err = d.Set("min_upper_case_letter", v); err != nil {
			return diag.Errorf("error reading min_upper_case_letter: %v", err)
		}
	}

	if o.MinimumLength != nil {
		v := *o.MinimumLength

		if err = d.Set("minimum_length", v); err != nil {
			return diag.Errorf("error reading minimum_length: %v", err)
		}
	}

	if o.ReusePassword != nil {
		v := *o.ReusePassword

		if err = d.Set("reuse_password", v); err != nil {
			return diag.Errorf("error reading reuse_password: %v", err)
		}
	}

	if o.Status != nil {
		v := *o.Status

		if err = d.Set("status", v); err != nil {
			return diag.Errorf("error reading status: %v", err)
		}
	}

	return nil
}

func getObjectSystemPasswordPolicyGuestAdmin(d *schema.ResourceData, sv string) (*models.SystemPasswordPolicyGuestAdmin, diag.Diagnostics) {
	obj := models.SystemPasswordPolicyGuestAdmin{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("apply_to"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("apply_to", sv)
				diags = append(diags, e)
			}
			obj.ApplyTo = &v2
		}
	}
	if v1, ok := d.GetOk("change_4_characters"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v7.0.0") {
				e := utils.AttributeVersionWarning("change_4_characters", sv)
				diags = append(diags, e)
			}
			obj.Change4Characters = &v2
		}
	}
	if v1, ok := d.GetOk("expire_day"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("expire_day", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.ExpireDay = &tmp
		}
	}
	if v1, ok := d.GetOk("expire_status"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("expire_status", sv)
				diags = append(diags, e)
			}
			obj.ExpireStatus = &v2
		}
	}
	if v1, ok := d.GetOk("min_change_characters"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("min_change_characters", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MinChangeCharacters = &tmp
		}
	}
	if v1, ok := d.GetOk("min_lower_case_letter"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("min_lower_case_letter", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MinLowerCaseLetter = &tmp
		}
	}
	if v1, ok := d.GetOk("min_non_alphanumeric"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("min_non_alphanumeric", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MinNonAlphanumeric = &tmp
		}
	}
	if v1, ok := d.GetOk("min_number"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("min_number", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MinNumber = &tmp
		}
	}
	if v1, ok := d.GetOk("min_upper_case_letter"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("min_upper_case_letter", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MinUpperCaseLetter = &tmp
		}
	}
	if v1, ok := d.GetOk("minimum_length"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("minimum_length", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.MinimumLength = &tmp
		}
	}
	if v1, ok := d.GetOk("reuse_password"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("reuse_password", sv)
				diags = append(diags, e)
			}
			obj.ReusePassword = &v2
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
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectSystemPasswordPolicyGuestAdmin(d *schema.ResourceData, sv string) (*models.SystemPasswordPolicyGuestAdmin, diag.Diagnostics) {
	obj := models.SystemPasswordPolicyGuestAdmin{}
	diags := diag.Diagnostics{}

	return &obj, diags
}
