// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure IPv4 routing policies.

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

func resourceRouterPolicy() *schema.Resource {
	return &schema.Resource{
		Description: "Configure IPv4 routing policies.",

		CreateContext: resourceRouterPolicyCreate,
		ReadContext:   resourceRouterPolicyRead,
		UpdateContext: resourceRouterPolicyUpdate,
		DeleteContext: resourceRouterPolicyDelete,

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
			"action": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"deny", "permit"}, false),

				Description: "Action of the policy route.",
				Optional:    true,
				Computed:    true,
			},
			"comments": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Optional comments.",
				Optional:    true,
				Computed:    true,
			},
			"dst": {
				Type:        schema.TypeList,
				Description: "Destination IP and mask (x.x.x.x/x).",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"subnet": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "IP and mask.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"dst_negate": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable negating destination address match.",
				Optional:    true,
				Computed:    true,
			},
			"dstaddr": {
				Type:        schema.TypeList,
				Description: "Destination address name.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Address/group name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"end_port": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "End destination port number (0 - 65535).",
				Optional:    true,
				Computed:    true,
			},
			"end_source_port": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "End source port number (0 - 65535).",
				Optional:    true,
				Computed:    true,
			},
			"gateway": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "IP address of the gateway.",
				Optional:    true,
				Computed:    true,
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
			"input_device_negate": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable negation of input device match.",
				Optional:    true,
				Computed:    true,
			},
			"internet_service_custom": {
				Type:        schema.TypeList,
				Description: "Custom Destination Internet Service name.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Custom Destination Internet Service name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"internet_service_id": {
				Type:        schema.TypeList,
				Description: "Destination Internet Service ID.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type: schema.TypeInt,

							Description: "Destination Internet Service ID.",
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
				Type:        schema.TypeList,
				Description: "Source IP and mask (x.x.x.x/x).",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"subnet": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "IP and mask.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"src_negate": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable negating source address match.",
				Optional:    true,
				Computed:    true,
			},
			"srcaddr": {
				Type:        schema.TypeList,
				Description: "Source address name.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Address/group name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"start_port": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "Start destination port number (0 - 65535).",
				Optional:    true,
				Computed:    true,
			},
			"start_source_port": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "Start source port number (0 - 65535).",
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

func resourceRouterPolicyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating RouterPolicy resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectRouterPolicy(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateRouterPolicy(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("RouterPolicy")
	}

	return resourceRouterPolicyRead(ctx, d, meta)
}

func resourceRouterPolicyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectRouterPolicy(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateRouterPolicy(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating RouterPolicy resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("RouterPolicy")
	}

	return resourceRouterPolicyRead(ctx, d, meta)
}

func resourceRouterPolicyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteRouterPolicy(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting RouterPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterPolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadRouterPolicy(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading RouterPolicy resource: %v", err)
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

	diags := refreshObjectRouterPolicy(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenRouterPolicyDst(d *schema.ResourceData, v *[]models.RouterPolicyDst, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Subnet; tmp != nil {
				v["subnet"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "subnet")
	}

	return flat
}

func flattenRouterPolicyDstaddr(d *schema.ResourceData, v *[]models.RouterPolicyDstaddr, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
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

func flattenRouterPolicyInputDevice(d *schema.ResourceData, v *[]models.RouterPolicyInputDevice, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
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

func flattenRouterPolicyInternetServiceCustom(d *schema.ResourceData, v *[]models.RouterPolicyInternetServiceCustom, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
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

func flattenRouterPolicyInternetServiceId(d *schema.ResourceData, v *[]models.RouterPolicyInternetServiceId, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenRouterPolicySrc(d *schema.ResourceData, v *[]models.RouterPolicySrc, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Subnet; tmp != nil {
				v["subnet"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "subnet")
	}

	return flat
}

func flattenRouterPolicySrcaddr(d *schema.ResourceData, v *[]models.RouterPolicySrcaddr, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
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

func refreshObjectRouterPolicy(d *schema.ResourceData, o *models.RouterPolicy, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Action != nil {
		v := *o.Action

		if err = d.Set("action", v); err != nil {
			return diag.Errorf("error reading action: %v", err)
		}
	}

	if o.Comments != nil {
		v := *o.Comments

		if err = d.Set("comments", v); err != nil {
			return diag.Errorf("error reading comments: %v", err)
		}
	}

	if o.Dst != nil {
		if err = d.Set("dst", flattenRouterPolicyDst(d, o.Dst, "dst", sort)); err != nil {
			return diag.Errorf("error reading dst: %v", err)
		}
	}

	if o.DstNegate != nil {
		v := *o.DstNegate

		if err = d.Set("dst_negate", v); err != nil {
			return diag.Errorf("error reading dst_negate: %v", err)
		}
	}

	if o.Dstaddr != nil {
		if err = d.Set("dstaddr", flattenRouterPolicyDstaddr(d, o.Dstaddr, "dstaddr", sort)); err != nil {
			return diag.Errorf("error reading dstaddr: %v", err)
		}
	}

	if o.EndPort != nil {
		v := *o.EndPort

		if err = d.Set("end_port", v); err != nil {
			return diag.Errorf("error reading end_port: %v", err)
		}
	}

	if o.EndSourcePort != nil {
		v := *o.EndSourcePort

		if err = d.Set("end_source_port", v); err != nil {
			return diag.Errorf("error reading end_source_port: %v", err)
		}
	}

	if o.Gateway != nil {
		v := *o.Gateway

		if err = d.Set("gateway", v); err != nil {
			return diag.Errorf("error reading gateway: %v", err)
		}
	}

	if o.InputDevice != nil {
		if err = d.Set("input_device", flattenRouterPolicyInputDevice(d, o.InputDevice, "input_device", sort)); err != nil {
			return diag.Errorf("error reading input_device: %v", err)
		}
	}

	if o.InputDeviceNegate != nil {
		v := *o.InputDeviceNegate

		if err = d.Set("input_device_negate", v); err != nil {
			return diag.Errorf("error reading input_device_negate: %v", err)
		}
	}

	if o.InternetServiceCustom != nil {
		if err = d.Set("internet_service_custom", flattenRouterPolicyInternetServiceCustom(d, o.InternetServiceCustom, "internet_service_custom", sort)); err != nil {
			return diag.Errorf("error reading internet_service_custom: %v", err)
		}
	}

	if o.InternetServiceId != nil {
		if err = d.Set("internet_service_id", flattenRouterPolicyInternetServiceId(d, o.InternetServiceId, "internet_service_id", sort)); err != nil {
			return diag.Errorf("error reading internet_service_id: %v", err)
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
		if err = d.Set("src", flattenRouterPolicySrc(d, o.Src, "src", sort)); err != nil {
			return diag.Errorf("error reading src: %v", err)
		}
	}

	if o.SrcNegate != nil {
		v := *o.SrcNegate

		if err = d.Set("src_negate", v); err != nil {
			return diag.Errorf("error reading src_negate: %v", err)
		}
	}

	if o.Srcaddr != nil {
		if err = d.Set("srcaddr", flattenRouterPolicySrcaddr(d, o.Srcaddr, "srcaddr", sort)); err != nil {
			return diag.Errorf("error reading srcaddr: %v", err)
		}
	}

	if o.StartPort != nil {
		v := *o.StartPort

		if err = d.Set("start_port", v); err != nil {
			return diag.Errorf("error reading start_port: %v", err)
		}
	}

	if o.StartSourcePort != nil {
		v := *o.StartSourcePort

		if err = d.Set("start_source_port", v); err != nil {
			return diag.Errorf("error reading start_source_port: %v", err)
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

func expandRouterPolicyDst(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterPolicyDst, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterPolicyDst

	for i := range l {
		tmp := models.RouterPolicyDst{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.subnet", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Subnet = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandRouterPolicyDstaddr(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterPolicyDstaddr, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterPolicyDstaddr

	for i := range l {
		tmp := models.RouterPolicyDstaddr{}
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

func expandRouterPolicyInputDevice(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterPolicyInputDevice, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterPolicyInputDevice

	for i := range l {
		tmp := models.RouterPolicyInputDevice{}
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

func expandRouterPolicyInternetServiceCustom(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterPolicyInternetServiceCustom, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterPolicyInternetServiceCustom

	for i := range l {
		tmp := models.RouterPolicyInternetServiceCustom{}
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

func expandRouterPolicyInternetServiceId(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterPolicyInternetServiceId, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterPolicyInternetServiceId

	for i := range l {
		tmp := models.RouterPolicyInternetServiceId{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandRouterPolicySrc(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterPolicySrc, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterPolicySrc

	for i := range l {
		tmp := models.RouterPolicySrc{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.subnet", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Subnet = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandRouterPolicySrcaddr(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.RouterPolicySrcaddr, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.RouterPolicySrcaddr

	for i := range l {
		tmp := models.RouterPolicySrcaddr{}
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

func getObjectRouterPolicy(d *schema.ResourceData, sv string) (*models.RouterPolicy, diag.Diagnostics) {
	obj := models.RouterPolicy{}
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
	if v1, ok := d.GetOk("comments"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("comments", sv)
				diags = append(diags, e)
			}
			obj.Comments = &v2
		}
	}
	if v, ok := d.GetOk("dst"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("dst", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterPolicyDst(d, v, "dst", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Dst = t
		}
	} else if d.HasChange("dst") {
		old, new := d.GetChange("dst")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Dst = &[]models.RouterPolicyDst{}
		}
	}
	if v1, ok := d.GetOk("dst_negate"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dst_negate", sv)
				diags = append(diags, e)
			}
			obj.DstNegate = &v2
		}
	}
	if v, ok := d.GetOk("dstaddr"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("dstaddr", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterPolicyDstaddr(d, v, "dstaddr", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Dstaddr = t
		}
	} else if d.HasChange("dstaddr") {
		old, new := d.GetChange("dstaddr")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Dstaddr = &[]models.RouterPolicyDstaddr{}
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
	if v1, ok := d.GetOk("end_source_port"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("end_source_port", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.EndSourcePort = &tmp
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
		t, err := expandRouterPolicyInputDevice(d, v, "input_device", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.InputDevice = t
		}
	} else if d.HasChange("input_device") {
		old, new := d.GetChange("input_device")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.InputDevice = &[]models.RouterPolicyInputDevice{}
		}
	}
	if v1, ok := d.GetOk("input_device_negate"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("input_device_negate", sv)
				diags = append(diags, e)
			}
			obj.InputDeviceNegate = &v2
		}
	}
	if v, ok := d.GetOk("internet_service_custom"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("internet_service_custom", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterPolicyInternetServiceCustom(d, v, "internet_service_custom", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.InternetServiceCustom = t
		}
	} else if d.HasChange("internet_service_custom") {
		old, new := d.GetChange("internet_service_custom")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.InternetServiceCustom = &[]models.RouterPolicyInternetServiceCustom{}
		}
	}
	if v, ok := d.GetOk("internet_service_id"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("internet_service_id", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterPolicyInternetServiceId(d, v, "internet_service_id", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.InternetServiceId = t
		}
	} else if d.HasChange("internet_service_id") {
		old, new := d.GetChange("internet_service_id")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.InternetServiceId = &[]models.RouterPolicyInternetServiceId{}
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
	if v, ok := d.GetOk("src"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("src", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterPolicySrc(d, v, "src", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Src = t
		}
	} else if d.HasChange("src") {
		old, new := d.GetChange("src")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Src = &[]models.RouterPolicySrc{}
		}
	}
	if v1, ok := d.GetOk("src_negate"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("src_negate", sv)
				diags = append(diags, e)
			}
			obj.SrcNegate = &v2
		}
	}
	if v, ok := d.GetOk("srcaddr"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("srcaddr", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterPolicySrcaddr(d, v, "srcaddr", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Srcaddr = t
		}
	} else if d.HasChange("srcaddr") {
		old, new := d.GetChange("srcaddr")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Srcaddr = &[]models.RouterPolicySrcaddr{}
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
	if v1, ok := d.GetOk("start_source_port"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("start_source_port", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.StartSourcePort = &tmp
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
