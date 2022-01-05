// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure IPv6 DoS policies.

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

func resourceFirewallDoSPolicy6() *schema.Resource {
	return &schema.Resource{
		Description: "Configure IPv6 DoS policies.",

		CreateContext: resourceFirewallDoSPolicy6Create,
		ReadContext:   resourceFirewallDoSPolicy6Read,
		UpdateContext: resourceFirewallDoSPolicy6Update,
		DeleteContext: resourceFirewallDoSPolicy6Delete,

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
			"dynamic_sort_subtable": {
				Type:        schema.TypeBool,
				Description: "If set will sort table response by mkey",
				Optional:    true,
				Default:     false,
			},
			"anomaly": {
				Type:        schema.TypeList,
				Description: "Anomaly name.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"pass", "block"}, false),

							Description: "Action taken when the threshold is reached.",
							Optional:    true,
							Computed:    true,
						},
						"log": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable anomaly logging.",
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "Anomaly name.",
							Optional:    true,
							Computed:    true,
						},
						"quarantine": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"none", "attacker"}, false),

							Description: "Quarantine method.",
							Optional:    true,
							Computed:    true,
						},
						"quarantine_expiry": {
							Type: schema.TypeString,

							Description: "Duration of quarantine. (Format ###d##h##m, minimum 1m, maximum 364d23h59m, default = 5m). Requires quarantine set to attacker.",
							Optional:    true,
							Computed:    true,
						},
						"quarantine_log": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable quarantine logging.",
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

							Description: "Enable/disable this anomaly.",
							Optional:    true,
							Computed:    true,
						},
						"threshold": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 2147483647),

							Description: "Anomaly threshold. Number of detected instances per minute that triggers the anomaly action.",
							Optional:    true,
							Computed:    true,
						},
						"thresholddefault": {
							Type: schema.TypeInt,

							Description: "Number of detected instances per minute which triggers action (1 - 2147483647, default = 1000). Note that each anomaly has a different threshold value assigned to it.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"comments": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 1023),

				Description: "Comment.",
				Optional:    true,
				Computed:    true,
			},
			"dstaddr": {
				Type:        schema.TypeList,
				Description: "Destination address name from available addresses.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Address name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"interface": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Incoming interface name from available interfaces.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Policy name.",
				Optional:    true,
				Computed:    true,
			},
			"policyid": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 9999),

				Description: "Policy ID.",
				ForceNew:    true,
				Required:    true,
			},
			"service": {
				Type:        schema.TypeList,
				Description: "Service object from available options.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Service name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"srcaddr": {
				Type:        schema.TypeList,
				Description: "Source address name from available addresses.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "Address name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable this policy.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceFirewallDoSPolicy6Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
	key := "policyid"
	if v, ok := d.GetOk(key); ok {
		mkey = utils.ParseMkey(v)
		if mkey == "" && allow_append {
			return diag.Errorf("error creating FirewallDoSPolicy6 resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectFirewallDoSPolicy6(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateFirewallDoSPolicy6(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallDoSPolicy6")
	}

	return resourceFirewallDoSPolicy6Read(ctx, d, meta)
}

func resourceFirewallDoSPolicy6Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectFirewallDoSPolicy6(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateFirewallDoSPolicy6(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating FirewallDoSPolicy6 resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallDoSPolicy6")
	}

	return resourceFirewallDoSPolicy6Read(ctx, d, meta)
}

func resourceFirewallDoSPolicy6Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteFirewallDoSPolicy6(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting FirewallDoSPolicy6 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallDoSPolicy6Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallDoSPolicy6(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallDoSPolicy6 resource: %v", err)
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

	diags := refreshObjectFirewallDoSPolicy6(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenFirewallDoSPolicy6Anomaly(v *[]models.FirewallDoSPolicy6Anomaly, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Action; tmp != nil {
				v["action"] = *tmp
			}

			if tmp := cfg.Log; tmp != nil {
				v["log"] = *tmp
			}

			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.Quarantine; tmp != nil {
				v["quarantine"] = *tmp
			}

			if tmp := cfg.QuarantineExpiry; tmp != nil {
				v["quarantine_expiry"] = *tmp
			}

			if tmp := cfg.QuarantineLog; tmp != nil {
				v["quarantine_log"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			if tmp := cfg.Threshold; tmp != nil {
				v["threshold"] = *tmp
			}

			if tmp := cfg.Thresholddefault; tmp != nil {
				v["thresholddefault"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenFirewallDoSPolicy6Dstaddr(v *[]models.FirewallDoSPolicy6Dstaddr, sort bool) interface{} {
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

func flattenFirewallDoSPolicy6Service(v *[]models.FirewallDoSPolicy6Service, sort bool) interface{} {
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

func flattenFirewallDoSPolicy6Srcaddr(v *[]models.FirewallDoSPolicy6Srcaddr, sort bool) interface{} {
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

func refreshObjectFirewallDoSPolicy6(d *schema.ResourceData, o *models.FirewallDoSPolicy6, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Anomaly != nil {
		if err = d.Set("anomaly", flattenFirewallDoSPolicy6Anomaly(o.Anomaly, sort)); err != nil {
			return diag.Errorf("error reading anomaly: %v", err)
		}
	}

	if o.Comments != nil {
		v := *o.Comments

		if err = d.Set("comments", v); err != nil {
			return diag.Errorf("error reading comments: %v", err)
		}
	}

	if o.Dstaddr != nil {
		if err = d.Set("dstaddr", flattenFirewallDoSPolicy6Dstaddr(o.Dstaddr, sort)); err != nil {
			return diag.Errorf("error reading dstaddr: %v", err)
		}
	}

	if o.Interface != nil {
		v := *o.Interface

		if err = d.Set("interface", v); err != nil {
			return diag.Errorf("error reading interface: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.Policyid != nil {
		v := *o.Policyid

		if err = d.Set("policyid", v); err != nil {
			return diag.Errorf("error reading policyid: %v", err)
		}
	}

	if o.Service != nil {
		if err = d.Set("service", flattenFirewallDoSPolicy6Service(o.Service, sort)); err != nil {
			return diag.Errorf("error reading service: %v", err)
		}
	}

	if o.Srcaddr != nil {
		if err = d.Set("srcaddr", flattenFirewallDoSPolicy6Srcaddr(o.Srcaddr, sort)); err != nil {
			return diag.Errorf("error reading srcaddr: %v", err)
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

func expandFirewallDoSPolicy6Anomaly(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallDoSPolicy6Anomaly, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallDoSPolicy6Anomaly

	for i := range l {
		tmp := models.FirewallDoSPolicy6Anomaly{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.action", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Action = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.log", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Log = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.quarantine", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Quarantine = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.quarantine_expiry", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.QuarantineExpiry = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.quarantine_log", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.QuarantineLog = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.threshold", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Threshold = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.thresholddefault", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Thresholddefault = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandFirewallDoSPolicy6Dstaddr(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallDoSPolicy6Dstaddr, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallDoSPolicy6Dstaddr

	for i := range l {
		tmp := models.FirewallDoSPolicy6Dstaddr{}
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

func expandFirewallDoSPolicy6Service(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallDoSPolicy6Service, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallDoSPolicy6Service

	for i := range l {
		tmp := models.FirewallDoSPolicy6Service{}
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

func expandFirewallDoSPolicy6Srcaddr(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallDoSPolicy6Srcaddr, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallDoSPolicy6Srcaddr

	for i := range l {
		tmp := models.FirewallDoSPolicy6Srcaddr{}
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

func getObjectFirewallDoSPolicy6(d *schema.ResourceData, sv string) (*models.FirewallDoSPolicy6, diag.Diagnostics) {
	obj := models.FirewallDoSPolicy6{}
	diags := diag.Diagnostics{}

	if v, ok := d.GetOk("anomaly"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("anomaly", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallDoSPolicy6Anomaly(d, v, "anomaly", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Anomaly = t
		}
	} else if d.HasChange("anomaly") {
		old, new := d.GetChange("anomaly")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Anomaly = &[]models.FirewallDoSPolicy6Anomaly{}
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
	if v, ok := d.GetOk("dstaddr"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("dstaddr", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallDoSPolicy6Dstaddr(d, v, "dstaddr", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Dstaddr = t
		}
	} else if d.HasChange("dstaddr") {
		old, new := d.GetChange("dstaddr")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Dstaddr = &[]models.FirewallDoSPolicy6Dstaddr{}
		}
	}
	if v1, ok := d.GetOk("interface"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("interface", sv)
				diags = append(diags, e)
			}
			obj.Interface = &v2
		}
	}
	if v1, ok := d.GetOk("name"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.2", "") {
				e := utils.AttributeVersionWarning("name", sv)
				diags = append(diags, e)
			}
			obj.Name = &v2
		}
	}
	if v1, ok := d.GetOk("policyid"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("policyid", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Policyid = &tmp
		}
	}
	if v, ok := d.GetOk("service"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("service", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallDoSPolicy6Service(d, v, "service", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Service = t
		}
	} else if d.HasChange("service") {
		old, new := d.GetChange("service")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Service = &[]models.FirewallDoSPolicy6Service{}
		}
	}
	if v, ok := d.GetOk("srcaddr"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("srcaddr", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallDoSPolicy6Srcaddr(d, v, "srcaddr", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Srcaddr = t
		}
	} else if d.HasChange("srcaddr") {
		old, new := d.GetChange("srcaddr")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Srcaddr = &[]models.FirewallDoSPolicy6Srcaddr{}
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
