// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure AP local configuration profiles.

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

func resourceWirelessControllerApcfgProfile() *schema.Resource {
	return &schema.Resource{
		Description: "Configure AP local configuration profiles.",

		CreateContext: resourceWirelessControllerApcfgProfileCreate,
		ReadContext:   resourceWirelessControllerApcfgProfileRead,
		UpdateContext: resourceWirelessControllerApcfgProfileUpdate,
		DeleteContext: resourceWirelessControllerApcfgProfileDelete,

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
				Description: "If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution!",
				Optional:    true,
				Default:     false,
			},
			"dynamic_sort_subtable": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"ac_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "IP address of the validation controller that AP must be able to join after applying AP local configuration.",
				Optional:    true,
				Computed:    true,
			},
			"ac_port": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1024, 49150),

				Description: "Port of the validation controller that AP must be able to join after applying AP local configuration (1024 - 49150, default = 5246).",
				Optional:    true,
				Computed:    true,
			},
			"ac_timer": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(3, 30),

				Description: "Maximum waiting time for the AP to join the validation controller after applying AP local configuration (3 - 30 min, default = 10).",
				Optional:    true,
				Computed:    true,
			},
			"ac_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"default", "specify", "apcfg"}, false),

				Description: "Validation controller type (default = default).",
				Optional:    true,
				Computed:    true,
			},
			"ap_family": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"fap", "fap-u", "fap-c"}, false),

				Description: "FortiAP family type (default = fap).",
				Optional:    true,
				Computed:    true,
			},
			"command_list": {
				Type:        schema.TypeList,
				Description: "AP local configuration command list.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 255),

							Description: "Command ID.",
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "AP local configuration command name.",
							Optional:    true,
							Computed:    true,
						},
						"passwd_value": {
							Type: schema.TypeString,

							Description: "AP local configuration command password value.",
							Optional:    true,
							Computed:    true,
							Sensitive:   true,
						},
						"type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"non-password", "password"}, false),

							Description: "The command type (default = non-password).",
							Optional:    true,
							Computed:    true,
						},
						"value": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 127),

							Description: "AP local configuration command value.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"comment": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Comment.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "AP local configuration profile name.",
				ForceNew:    true,
				Required:    true,
			},
		},
	}
}

func resourceWirelessControllerApcfgProfileCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating WirelessControllerApcfgProfile resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectWirelessControllerApcfgProfile(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateWirelessControllerApcfgProfile(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WirelessControllerApcfgProfile")
	}

	return resourceWirelessControllerApcfgProfileRead(ctx, d, meta)
}

func resourceWirelessControllerApcfgProfileUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectWirelessControllerApcfgProfile(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateWirelessControllerApcfgProfile(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating WirelessControllerApcfgProfile resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WirelessControllerApcfgProfile")
	}

	return resourceWirelessControllerApcfgProfileRead(ctx, d, meta)
}

func resourceWirelessControllerApcfgProfileDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteWirelessControllerApcfgProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting WirelessControllerApcfgProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerApcfgProfileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadWirelessControllerApcfgProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WirelessControllerApcfgProfile resource: %v", err)
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

	diags := refreshObjectWirelessControllerApcfgProfile(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenWirelessControllerApcfgProfileCommandList(v *[]models.WirelessControllerApcfgProfileCommandList, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.PasswdValue; tmp != nil {
				v["passwd_value"] = *tmp
			}

			if tmp := cfg.Type; tmp != nil {
				v["type"] = *tmp
			}

			if tmp := cfg.Value; tmp != nil {
				v["value"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func refreshObjectWirelessControllerApcfgProfile(d *schema.ResourceData, o *models.WirelessControllerApcfgProfile, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AcIp != nil {
		v := *o.AcIp

		if err = d.Set("ac_ip", v); err != nil {
			return diag.Errorf("error reading ac_ip: %v", err)
		}
	}

	if o.AcPort != nil {
		v := *o.AcPort

		if err = d.Set("ac_port", v); err != nil {
			return diag.Errorf("error reading ac_port: %v", err)
		}
	}

	if o.AcTimer != nil {
		v := *o.AcTimer

		if err = d.Set("ac_timer", v); err != nil {
			return diag.Errorf("error reading ac_timer: %v", err)
		}
	}

	if o.AcType != nil {
		v := *o.AcType

		if err = d.Set("ac_type", v); err != nil {
			return diag.Errorf("error reading ac_type: %v", err)
		}
	}

	if o.ApFamily != nil {
		v := *o.ApFamily

		if err = d.Set("ap_family", v); err != nil {
			return diag.Errorf("error reading ap_family: %v", err)
		}
	}

	if o.CommandList != nil {
		if err = d.Set("command_list", flattenWirelessControllerApcfgProfileCommandList(o.CommandList, sort)); err != nil {
			return diag.Errorf("error reading command_list: %v", err)
		}
	}

	if o.Comment != nil {
		v := *o.Comment

		if err = d.Set("comment", v); err != nil {
			return diag.Errorf("error reading comment: %v", err)
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

func expandWirelessControllerApcfgProfileCommandList(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WirelessControllerApcfgProfileCommandList, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WirelessControllerApcfgProfileCommandList

	for i := range l {
		tmp := models.WirelessControllerApcfgProfileCommandList{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.passwd_value", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PasswdValue = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Type = &v2
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

func getObjectWirelessControllerApcfgProfile(d *schema.ResourceData, sv string) (*models.WirelessControllerApcfgProfile, diag.Diagnostics) {
	obj := models.WirelessControllerApcfgProfile{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("ac_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ac_ip", sv)
				diags = append(diags, e)
			}
			obj.AcIp = &v2
		}
	}
	if v1, ok := d.GetOk("ac_port"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ac_port", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.AcPort = &tmp
		}
	}
	if v1, ok := d.GetOk("ac_timer"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ac_timer", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.AcTimer = &tmp
		}
	}
	if v1, ok := d.GetOk("ac_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ac_type", sv)
				diags = append(diags, e)
			}
			obj.AcType = &v2
		}
	}
	if v1, ok := d.GetOk("ap_family"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.2", "") {
				e := utils.AttributeVersionWarning("ap_family", sv)
				diags = append(diags, e)
			}
			obj.ApFamily = &v2
		}
	}
	if v, ok := d.GetOk("command_list"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("command_list", sv)
			diags = append(diags, e)
		}
		t, err := expandWirelessControllerApcfgProfileCommandList(d, v, "command_list", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.CommandList = t
		}
	} else if d.HasChange("command_list") {
		old, new := d.GetChange("command_list")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.CommandList = &[]models.WirelessControllerApcfgProfileCommandList{}
		}
	}
	if v1, ok := d.GetOk("comment"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("comment", sv)
				diags = append(diags, e)
			}
			obj.Comment = &v2
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
