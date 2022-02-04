// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure TACACS+ server entries.

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

func resourceUserTacacs() *schema.Resource {
	return &schema.Resource{
		Description: "Configure TACACS+ server entries.",

		CreateContext: resourceUserTacacsCreate,
		ReadContext:   resourceUserTacacsRead,
		UpdateContext: resourceUserTacacsUpdate,
		DeleteContext: resourceUserTacacsDelete,

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
			"authen_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"mschap", "chap", "pap", "ascii", "auto"}, false),

				Description: "Allowed authentication protocols/methods.",
				Optional:    true,
				Computed:    true,
			},
			"authorization": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable TACACS+ authorization.",
				Optional:    true,
				Computed:    true,
			},
			"interface": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "Specify outgoing interface to reach server.",
				Optional:    true,
				Computed:    true,
			},
			"interface_select_method": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"auto", "sdwan", "specify"}, false),

				Description: "Specify how to select outgoing interface to reach server.",
				Optional:    true,
				Computed:    true,
			},
			"key": {
				Type: schema.TypeString,

				Description: "Key to access the primary server.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "TACACS+ server entry name.",
				ForceNew:    true,
				Required:    true,
			},
			"port": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),

				Description: "Port number of the TACACS+ server.",
				Optional:    true,
				Computed:    true,
			},
			"secondary_key": {
				Type: schema.TypeString,

				Description: "Key to access the secondary server.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"secondary_server": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Secondary TACACS+ server CN domain name or IP address.",
				Optional:    true,
				Computed:    true,
			},
			"server": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Primary TACACS+ server CN domain name or IP address.",
				Optional:    true,
				Computed:    true,
			},
			"source_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Source IP address for communications to TACACS+ server.",
				Optional:    true,
				Computed:    true,
			},
			"tertiary_key": {
				Type: schema.TypeString,

				Description: "Key to access the tertiary server.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"tertiary_server": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Tertiary TACACS+ server CN domain name or IP address.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceUserTacacsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating UserTacacs resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectUserTacacs(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateUserTacacs(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("UserTacacs")
	}

	return resourceUserTacacsRead(ctx, d, meta)
}

func resourceUserTacacsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectUserTacacs(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateUserTacacs(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating UserTacacs resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("UserTacacs")
	}

	return resourceUserTacacsRead(ctx, d, meta)
}

func resourceUserTacacsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteUserTacacs(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting UserTacacs resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserTacacsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadUserTacacs(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading UserTacacs resource: %v", err)
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

	diags := refreshObjectUserTacacs(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectUserTacacs(d *schema.ResourceData, o *models.UserTacacs, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AuthenType != nil {
		v := *o.AuthenType

		if err = d.Set("authen_type", v); err != nil {
			return diag.Errorf("error reading authen_type: %v", err)
		}
	}

	if o.Authorization != nil {
		v := *o.Authorization

		if err = d.Set("authorization", v); err != nil {
			return diag.Errorf("error reading authorization: %v", err)
		}
	}

	if o.Interface != nil {
		v := *o.Interface

		if err = d.Set("interface", v); err != nil {
			return diag.Errorf("error reading interface: %v", err)
		}
	}

	if o.InterfaceSelectMethod != nil {
		v := *o.InterfaceSelectMethod

		if err = d.Set("interface_select_method", v); err != nil {
			return diag.Errorf("error reading interface_select_method: %v", err)
		}
	}

	if o.Key != nil {
		v := *o.Key

		if v == "ENC XXXX" {
		} else if err = d.Set("key", v); err != nil {
			return diag.Errorf("error reading key: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.Port != nil {
		v := *o.Port

		if err = d.Set("port", v); err != nil {
			return diag.Errorf("error reading port: %v", err)
		}
	}

	if o.SecondaryKey != nil {
		v := *o.SecondaryKey

		if v == "ENC XXXX" {
		} else if err = d.Set("secondary_key", v); err != nil {
			return diag.Errorf("error reading secondary_key: %v", err)
		}
	}

	if o.SecondaryServer != nil {
		v := *o.SecondaryServer

		if err = d.Set("secondary_server", v); err != nil {
			return diag.Errorf("error reading secondary_server: %v", err)
		}
	}

	if o.Server != nil {
		v := *o.Server

		if err = d.Set("server", v); err != nil {
			return diag.Errorf("error reading server: %v", err)
		}
	}

	if o.SourceIp != nil {
		v := *o.SourceIp

		if err = d.Set("source_ip", v); err != nil {
			return diag.Errorf("error reading source_ip: %v", err)
		}
	}

	if o.TertiaryKey != nil {
		v := *o.TertiaryKey

		if v == "ENC XXXX" {
		} else if err = d.Set("tertiary_key", v); err != nil {
			return diag.Errorf("error reading tertiary_key: %v", err)
		}
	}

	if o.TertiaryServer != nil {
		v := *o.TertiaryServer

		if err = d.Set("tertiary_server", v); err != nil {
			return diag.Errorf("error reading tertiary_server: %v", err)
		}
	}

	return nil
}

func getObjectUserTacacs(d *schema.ResourceData, sv string) (*models.UserTacacs, diag.Diagnostics) {
	obj := models.UserTacacs{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("authen_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("authen_type", sv)
				diags = append(diags, e)
			}
			obj.AuthenType = &v2
		}
	}
	if v1, ok := d.GetOk("authorization"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("authorization", sv)
				diags = append(diags, e)
			}
			obj.Authorization = &v2
		}
	}
	if v1, ok := d.GetOk("interface"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.0") {
				e := utils.AttributeVersionWarning("interface", sv)
				diags = append(diags, e)
			}
			obj.Interface = &v2
		}
	}
	if v1, ok := d.GetOk("interface_select_method"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.0") {
				e := utils.AttributeVersionWarning("interface_select_method", sv)
				diags = append(diags, e)
			}
			obj.InterfaceSelectMethod = &v2
		}
	}
	if v1, ok := d.GetOk("key"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("key", sv)
				diags = append(diags, e)
			}
			obj.Key = &v2
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
	if v1, ok := d.GetOk("port"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("port", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Port = &tmp
		}
	}
	if v1, ok := d.GetOk("secondary_key"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("secondary_key", sv)
				diags = append(diags, e)
			}
			obj.SecondaryKey = &v2
		}
	}
	if v1, ok := d.GetOk("secondary_server"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("secondary_server", sv)
				diags = append(diags, e)
			}
			obj.SecondaryServer = &v2
		}
	}
	if v1, ok := d.GetOk("server"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("server", sv)
				diags = append(diags, e)
			}
			obj.Server = &v2
		}
	}
	if v1, ok := d.GetOk("source_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("source_ip", sv)
				diags = append(diags, e)
			}
			obj.SourceIp = &v2
		}
	}
	if v1, ok := d.GetOk("tertiary_key"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("tertiary_key", sv)
				diags = append(diags, e)
			}
			obj.TertiaryKey = &v2
		}
	}
	if v1, ok := d.GetOk("tertiary_server"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("tertiary_server", sv)
				diags = append(diags, e)
			}
			obj.TertiaryServer = &v2
		}
	}
	return &obj, diags
}
