// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: RIP interface configuration.

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

func resourceRouterRipInterface() *schema.Resource {
	return &schema.Resource{
		Description: "RIP interface configuration.",

		CreateContext: resourceRouterRipInterfaceCreate,
		ReadContext:   resourceRouterRipInterfaceRead,
		UpdateContext: resourceRouterRipInterfaceUpdate,
		DeleteContext: resourceRouterRipInterfaceDelete,

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
			"auth_keychain": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Authentication key-chain name.",
				Optional:    true,
				Computed:    true,
			},
			"auth_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"none", "text", "md5"}, false),

				Description: "Authentication mode.",
				Optional:    true,
				Computed:    true,
			},
			"auth_string": {
				Type: schema.TypeString,

				Description: "Authentication string/password.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"flags": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),

				Description: "Flags.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Interface name.",
				ForceNew:    true,
				Required:    true,
			},
			"receive_version": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Receive version.",
				Optional:         true,
				Computed:         true,
			},
			"send_version": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "Send version.",
				Optional:         true,
				Computed:         true,
			},
			"send_version2_broadcast": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable broadcast version 1 compatible packets.",
				Optional:    true,
				Computed:    true,
			},
			"split_horizon": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"poisoned", "regular"}, false),

				Description: "Enable/disable split horizon.",
				Optional:    true,
				Computed:    true,
			},
			"split_horizon_status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable split horizon.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceRouterRipInterfaceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating RouterRipInterface resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectRouterRipInterface(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateRouterRipInterface(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("RouterRipInterface")
	}

	return resourceRouterRipInterfaceRead(ctx, d, meta)
}

func resourceRouterRipInterfaceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectRouterRipInterface(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateRouterRipInterface(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating RouterRipInterface resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("RouterRipInterface")
	}

	return resourceRouterRipInterfaceRead(ctx, d, meta)
}

func resourceRouterRipInterfaceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteRouterRipInterface(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting RouterRipInterface resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterRipInterfaceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadRouterRipInterface(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading RouterRipInterface resource: %v", err)
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

	diags := refreshObjectRouterRipInterface(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectRouterRipInterface(d *schema.ResourceData, o *models.RouterRipInterface, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AuthKeychain != nil {
		v := *o.AuthKeychain

		if err = d.Set("auth_keychain", v); err != nil {
			return diag.Errorf("error reading auth_keychain: %v", err)
		}
	}

	if o.AuthMode != nil {
		v := *o.AuthMode

		if err = d.Set("auth_mode", v); err != nil {
			return diag.Errorf("error reading auth_mode: %v", err)
		}
	}

	if o.AuthString != nil {
		v := *o.AuthString

		if v == "ENC XXXX" {
		} else if err = d.Set("auth_string", v); err != nil {
			return diag.Errorf("error reading auth_string: %v", err)
		}
	}

	if o.Flags != nil {
		v := *o.Flags

		if err = d.Set("flags", v); err != nil {
			return diag.Errorf("error reading flags: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.ReceiveVersion != nil {
		v := *o.ReceiveVersion

		if err = d.Set("receive_version", v); err != nil {
			return diag.Errorf("error reading receive_version: %v", err)
		}
	}

	if o.SendVersion != nil {
		v := *o.SendVersion

		if err = d.Set("send_version", v); err != nil {
			return diag.Errorf("error reading send_version: %v", err)
		}
	}

	if o.SendVersion2Broadcast != nil {
		v := *o.SendVersion2Broadcast

		if err = d.Set("send_version2_broadcast", v); err != nil {
			return diag.Errorf("error reading send_version2_broadcast: %v", err)
		}
	}

	if o.SplitHorizon != nil {
		v := *o.SplitHorizon

		if err = d.Set("split_horizon", v); err != nil {
			return diag.Errorf("error reading split_horizon: %v", err)
		}
	}

	if o.SplitHorizonStatus != nil {
		v := *o.SplitHorizonStatus

		if err = d.Set("split_horizon_status", v); err != nil {
			return diag.Errorf("error reading split_horizon_status: %v", err)
		}
	}

	return nil
}

func getObjectRouterRipInterface(d *schema.ResourceData, sv string) (*models.RouterRipInterface, diag.Diagnostics) {
	obj := models.RouterRipInterface{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("auth_keychain"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auth_keychain", sv)
				diags = append(diags, e)
			}
			obj.AuthKeychain = &v2
		}
	}
	if v1, ok := d.GetOk("auth_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auth_mode", sv)
				diags = append(diags, e)
			}
			obj.AuthMode = &v2
		}
	}
	if v1, ok := d.GetOk("auth_string"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auth_string", sv)
				diags = append(diags, e)
			}
			obj.AuthString = &v2
		}
	}
	if v1, ok := d.GetOk("flags"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("flags", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Flags = &tmp
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
	if v1, ok := d.GetOk("receive_version"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("receive_version", sv)
				diags = append(diags, e)
			}
			obj.ReceiveVersion = &v2
		}
	}
	if v1, ok := d.GetOk("send_version"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("send_version", sv)
				diags = append(diags, e)
			}
			obj.SendVersion = &v2
		}
	}
	if v1, ok := d.GetOk("send_version2_broadcast"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("send_version2_broadcast", sv)
				diags = append(diags, e)
			}
			obj.SendVersion2Broadcast = &v2
		}
	}
	if v1, ok := d.GetOk("split_horizon"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("split_horizon", sv)
				diags = append(diags, e)
			}
			obj.SplitHorizon = &v2
		}
	}
	if v1, ok := d.GetOk("split_horizon_status"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("split_horizon_status", sv)
				diags = append(diags, e)
			}
			obj.SplitHorizonStatus = &v2
		}
	}
	return &obj, diags
}
