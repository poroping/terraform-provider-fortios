// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure FortiToken.

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

func resourceUserFortitoken() *schema.Resource {
	return &schema.Resource{
		Description: "Configure FortiToken.",

		CreateContext: resourceUserFortitokenCreate,
		ReadContext:   resourceUserFortitokenRead,
		UpdateContext: resourceUserFortitokenUpdate,
		DeleteContext: resourceUserFortitokenDelete,

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
			"activation_code": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 32),

				Description: "Mobile token user activation-code.",
				Optional:    true,
				Computed:    true,
			},
			"activation_expire": {
				Type: schema.TypeInt,

				Description: "Mobile token user activation-code expire time.",
				Optional:    true,
				Computed:    true,
			},
			"comments": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Comment.",
				Optional:    true,
				Computed:    true,
			},
			"license": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),

				Description: "Mobile token license.",
				Optional:    true,
				Computed:    true,
			},
			"os_ver": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "Device Mobile Version.",
				Optional:    true,
				Computed:    true,
			},
			"reg_id": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 256),

				Description: "Device Reg ID.",
				Optional:    true,
				Computed:    true,
			},
			"serial_number": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 16),

				Description: "Serial number.",
				ForceNew:    true,
				Required:    true,
			},
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"active", "lock"}, false),

				Description: "Status",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceUserFortitokenCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
	key := "serial_number"
	if v, ok := d.GetOk(key); ok {
		mkey = utils.ParseMkey(v)
		if mkey == "" && allow_append {
			return diag.Errorf("error creating UserFortitoken resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectUserFortitoken(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateUserFortitoken(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("UserFortitoken")
	}

	return resourceUserFortitokenRead(ctx, d, meta)
}

func resourceUserFortitokenUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectUserFortitoken(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateUserFortitoken(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating UserFortitoken resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("UserFortitoken")
	}

	return resourceUserFortitokenRead(ctx, d, meta)
}

func resourceUserFortitokenDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteUserFortitoken(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting UserFortitoken resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserFortitokenRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadUserFortitoken(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading UserFortitoken resource: %v", err)
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

	diags := refreshObjectUserFortitoken(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectUserFortitoken(d *schema.ResourceData, o *models.UserFortitoken, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.ActivationCode != nil {
		v := *o.ActivationCode

		if err = d.Set("activation_code", v); err != nil {
			return diag.Errorf("error reading activation_code: %v", err)
		}
	}

	if o.ActivationExpire != nil {
		v := *o.ActivationExpire

		if err = d.Set("activation_expire", v); err != nil {
			return diag.Errorf("error reading activation_expire: %v", err)
		}
	}

	if o.Comments != nil {
		v := *o.Comments

		if err = d.Set("comments", v); err != nil {
			return diag.Errorf("error reading comments: %v", err)
		}
	}

	if o.License != nil {
		v := *o.License

		if err = d.Set("license", v); err != nil {
			return diag.Errorf("error reading license: %v", err)
		}
	}

	if o.OsVer != nil {
		v := *o.OsVer

		if err = d.Set("os_ver", v); err != nil {
			return diag.Errorf("error reading os_ver: %v", err)
		}
	}

	if o.RegId != nil {
		v := *o.RegId

		if err = d.Set("reg_id", v); err != nil {
			return diag.Errorf("error reading reg_id: %v", err)
		}
	}

	if o.SerialNumber != nil {
		v := *o.SerialNumber

		if err = d.Set("serial_number", v); err != nil {
			return diag.Errorf("error reading serial_number: %v", err)
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

func getObjectUserFortitoken(d *schema.ResourceData, sv string) (*models.UserFortitoken, diag.Diagnostics) {
	obj := models.UserFortitoken{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("activation_code"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("activation_code", sv)
				diags = append(diags, e)
			}
			obj.ActivationCode = &v2
		}
	}
	if v1, ok := d.GetOk("activation_expire"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("activation_expire", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.ActivationExpire = &tmp
		}
	}
	if v1, ok := d.GetOk("comments"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("comments", sv)
				diags = append(diags, e)
			}
			obj.Comments = &v2
		}
	}
	if v1, ok := d.GetOk("license"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("license", sv)
				diags = append(diags, e)
			}
			obj.License = &v2
		}
	}
	if v1, ok := d.GetOk("os_ver"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("os_ver", sv)
				diags = append(diags, e)
			}
			obj.OsVer = &v2
		}
	}
	if v1, ok := d.GetOk("reg_id"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("reg_id", sv)
				diags = append(diags, e)
			}
			obj.RegId = &v2
		}
	}
	if v1, ok := d.GetOk("serial_number"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("serial_number", sv)
				diags = append(diags, e)
			}
			obj.SerialNumber = &v2
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
