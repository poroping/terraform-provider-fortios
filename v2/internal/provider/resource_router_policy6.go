// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure IPv6 routing policies.

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
	"github.com/poroping/terraform-provider-fortios/v2/validators"
)

func resourceRouterPolicy6() *schema.Resource {
	return &schema.Resource{
		Description: "Configure IPv6 routing policies.",

		CreateContext: resourceRouterPolicy6Create,
		ReadContext:   resourceRouterPolicy6Read,
		UpdateContext: resourceRouterPolicy6Update,
		DeleteContext: resourceRouterPolicy6Delete,

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
				RequiredWith: []string{"seq_num"},
			},
			"dynamic_sort_subtable": {
				Type:        schema.TypeBool,
				Description: "If set will sort table response by mkey",
				Optional:    true,
				Default:     false,
			},
			"comments": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Optional comments.",
				Optional:    true,
				Computed:    true,
			},
			"dst": {
				Type:             schema.TypeString,
				ValidateFunc:     validators.FortiValidateIPv6Network,
				DiffSuppressFunc: suppressors.DiffCidrEqual,
				Description:      "Destination IPv6 prefix.",
				Optional:         true,
				Computed:         true,
			},
			"end_port": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),

				Description: "End destination port number (1 - 65535).",
				Optional:    true,
				Computed:    true,
			},
			"gateway": {
				Type:             schema.TypeString,
				ValidateFunc:     validation.IsIPv6Address,
				DiffSuppressFunc: suppressors.DiffIPEqual,
				Description:      "IPv6 address of the gateway.",
				Optional:         true,
				Computed:         true,
			},
			"input_device": {
				Type:        schema.TypeList,
				Description: "Incoming interface name.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Interface name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"output_device": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Outgoing interface name.",
				Optional:    true,
				Computed:    true,
			},
			"protocol": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),

				Description: "Protocol number (0 - 255).",
				Optional:    true,
				Computed:    true,
			},
			"seq_num": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),

				Description: "Sequence number(1-65535).",
				ForceNew:    true,
				Optional:    true,
				Computed:    true,
			},
			"src": {
				Type:             schema.TypeString,
				ValidateFunc:     validators.FortiValidateIPv6Network,
				DiffSuppressFunc: suppressors.DiffCidrEqual,
				Description:      "Source IPv6 prefix.",
				Optional:         true,
				Computed:         true,
			},
			"start_port": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),

				Description: "Start destination port number (1 - 65535).",
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable this policy route.",
				Optional:    true,
				Computed:    true,
			},
			"tos": {
				Type: schema.TypeString,

				Description: "Type of service bit pattern.",
				Optional:    true,
				Computed:    true,
			},
			"tos_mask": {
				Type: schema.TypeString,

				Description: "Type of service evaluated bits.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceRouterPolicy6Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
	key := "seq_num"
	if v, ok := d.GetOk(key); ok {
		mkey = utils.ParseMkey(v)
		if mkey == "" && allow_append {
			return diag.Errorf("error creating RouterPolicy6 resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectRouterPolicy6(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateRouterPolicy6(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("RouterPolicy6")
	}

	return resourceRouterPolicy6Read(ctx, d, meta)
}

func resourceRouterPolicy6Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectRouterPolicy6(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateRouterPolicy6(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating RouterPolicy6 resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("RouterPolicy6")
	}

	return resourceRouterPolicy6Read(ctx, d, meta)
}

func resourceRouterPolicy6Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteRouterPolicy6(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting RouterPolicy6 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterPolicy6Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadRouterPolicy6(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading RouterPolicy6 resource: %v", err)
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

	diags := refreshObjectRouterPolicy6(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenRouterPolicy6InputDevice(v *[]models.RouterPolicy6InputDevice, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func refreshObjectRouterPolicy6(d *schema.ResourceData, o *models.RouterPolicy6, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Comments != nil {
		v := *o.Comments

		if err = d.Set("comments", v); err != nil {
			return diag.Errorf("error reading comments: %v", err)
		}
	}

	if o.Dst != nil {
		v := *o.Dst

		if err = d.Set("dst", v); err != nil {
			return diag.Errorf("error reading dst: %v", err)
		}
	}

	if o.EndPort != nil {
		v := *o.EndPort

		if err = d.Set("end_port", v); err != nil {
			return diag.Errorf("error reading end_port: %v", err)
		}
	}

	if o.Gateway != nil {
		v := *o.Gateway

		if err = d.Set("gateway", v); err != nil {
			return diag.Errorf("error reading gateway: %v", err)
		}
	}

	if o.InputDevice != nil {
		if err = d.Set("input_device", flattenRouterPolicy6InputDevice(o.InputDevice, sort)); err != nil {
			return diag.Errorf("error reading input_device: %v", err)
		}
	}

	if o.OutputDevice != nil {
		v := *o.OutputDevice

		if err = d.Set("output_device", v); err != nil {
			return diag.Errorf("error reading output_device: %v", err)
		}
	}

	if o.Protocol != nil {
		v := *o.Protocol

		if err = d.Set("protocol", v); err != nil {
			return diag.Errorf("error reading protocol: %v", err)
		}
	}

	if o.SeqNum != nil {
		v := *o.SeqNum

		if err = d.Set("seq_num", v); err != nil {
			return diag.Errorf("error reading seq_num: %v", err)
		}
	}

	if o.Src != nil {
		v := *o.Src

		if err = d.Set("src", v); err != nil {
			return diag.Errorf("error reading src: %v", err)
		}
	}

	if o.StartPort != nil {
		v := *o.StartPort

		if err = d.Set("start_port", v); err != nil {
			return diag.Errorf("error reading start_port: %v", err)
		}
	}

	if o.Status != nil {
		v := *o.Status

		if err = d.Set("status", v); err != nil {
			return diag.Errorf("error reading status: %v", err)
		}
	}

	if o.Tos != nil {
		v := *o.Tos

		if err = d.Set("tos", v); err != nil {
			return diag.Errorf("error reading tos: %v", err)
		}
	}

	if o.TosMask != nil {
		v := *o.TosMask

		if err = d.Set("tos_mask", v); err != nil {
			return diag.Errorf("error reading tos_mask: %v", err)
		}
	}

	return nil
}

func expandRouterPolicy6InputDevice(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterPolicy6InputDevice, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterPolicy6InputDevice

	for i := range l {
		tmp := models.RouterPolicy6InputDevice{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectRouterPolicy6(d *schema.ResourceData, sv string) (*models.RouterPolicy6, diag.Diagnostics) {
	obj := models.RouterPolicy6{}
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
	if v1, ok := d.GetOk("dst"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dst", sv)
				diags = append(diags, e)
			}
			obj.Dst = &v2
		}
	}
	if v1, ok := d.GetOk("end_port"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("end_port", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.EndPort = &tmp
		}
	}
	if v1, ok := d.GetOk("gateway"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("gateway", sv)
				diags = append(diags, e)
			}
			obj.Gateway = &v2
		}
	}
	if v, ok := d.GetOk("input_device"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("input_device", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterPolicy6InputDevice(d, v, "input_device", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.InputDevice = t
		}
	} else if d.HasChange("input_device") {
		old, new := d.GetChange("input_device")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.InputDevice = &[]models.RouterPolicy6InputDevice{}
		}
	}
	if v1, ok := d.GetOk("output_device"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("output_device", sv)
				diags = append(diags, e)
			}
			obj.OutputDevice = &v2
		}
	}
	if v1, ok := d.GetOk("protocol"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("protocol", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Protocol = &tmp
		}
	}
	if v1, ok := d.GetOk("seq_num"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("seq_num", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.SeqNum = &tmp
		}
	}
	if v1, ok := d.GetOk("src"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("src", sv)
				diags = append(diags, e)
			}
			obj.Src = &v2
		}
	}
	if v1, ok := d.GetOk("start_port"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("start_port", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.StartPort = &tmp
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
	if v1, ok := d.GetOk("tos"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("tos", sv)
				diags = append(diags, e)
			}
			obj.Tos = &v2
		}
	}
	if v1, ok := d.GetOk("tos_mask"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("tos_mask", sv)
				diags = append(diags, e)
			}
			obj.TosMask = &v2
		}
	}
	return &obj, diags
}
