// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure IPv6 to IPv4 policies.

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

func resourceFirewallPolicy64() *schema.Resource {
	return &schema.Resource{
		Description: "Configure IPv6 to IPv4 policies.",

		CreateContext: resourceFirewallPolicy64Create,
		ReadContext:   resourceFirewallPolicy64Read,
		UpdateContext: resourceFirewallPolicy64Update,
		DeleteContext: resourceFirewallPolicy64Delete,

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
			"action": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"accept", "deny"}, false),

				Description: "Policy action.",
				Optional:    true,
				Computed:    true,
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
				Description: "Destination address name.",
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
			"dstintf": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Destination interface name.",
				Optional:    true,
				Computed:    true,
			},
			"fixedport": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable policy fixed port.",
				Optional:    true,
				Computed:    true,
			},
			"ippool": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable policy64 IP pool.",
				Optional:    true,
				Computed:    true,
			},
			"logtraffic": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable policy log traffic.",
				Optional:    true,
				Computed:    true,
			},
			"logtraffic_start": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Record logs when a session starts and ends.",
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
			"per_ip_shaper": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Per-IP traffic shaper.",
				Optional:    true,
				Computed:    true,
			},
			"permit_any_host": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable permit any host in.",
				Optional:    true,
				Computed:    true,
			},
			"policyid": {
				Type: schema.TypeInt,

				Description: "Policy ID (0 - 4294967294).",
				ForceNew:    true,
				Optional:    true,
				Computed:    true,
			},
			"poolname": {
				Type:        schema.TypeList,
				Description: "Policy IP pool names.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "IP pool name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"schedule": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Schedule name.",
				Optional:    true,
				Computed:    true,
			},
			"service": {
				Type:        schema.TypeList,
				Description: "Service name.",
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
			"srcaddr": {
				Type:        schema.TypeList,
				Description: "Source address name.",
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
			"srcintf": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Source interface name.",
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable policy status.",
				Optional:    true,
				Computed:    true,
			},
			"tcp_mss_receiver": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "TCP MSS value of receiver.",
				Optional:    true,
				Computed:    true,
			},
			"tcp_mss_sender": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "TCP MSS value of sender.",
				Optional:    true,
				Computed:    true,
			},
			"traffic_shaper": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Traffic shaper.",
				Optional:    true,
				Computed:    true,
			},
			"traffic_shaper_reverse": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Reverse traffic shaper.",
				Optional:    true,
				Computed:    true,
			},
			"uuid": {
				Type: schema.TypeString,

				Description: "Universally Unique Identifier (UUID; automatically assigned but can be manually reset).",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceFirewallPolicy64Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating FirewallPolicy64 resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectFirewallPolicy64(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateFirewallPolicy64(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallPolicy64")
	}

	return resourceFirewallPolicy64Read(ctx, d, meta)
}

func resourceFirewallPolicy64Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectFirewallPolicy64(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateFirewallPolicy64(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating FirewallPolicy64 resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallPolicy64")
	}

	return resourceFirewallPolicy64Read(ctx, d, meta)
}

func resourceFirewallPolicy64Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteFirewallPolicy64(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting FirewallPolicy64 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallPolicy64Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallPolicy64(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallPolicy64 resource: %v", err)
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

	diags := refreshObjectFirewallPolicy64(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenFirewallPolicy64Dstaddr(d *schema.ResourceData, v *[]models.FirewallPolicy64Dstaddr, prefix string, sort bool) interface{} {
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

func flattenFirewallPolicy64Poolname(d *schema.ResourceData, v *[]models.FirewallPolicy64Poolname, prefix string, sort bool) interface{} {
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

func flattenFirewallPolicy64Service(d *schema.ResourceData, v *[]models.FirewallPolicy64Service, prefix string, sort bool) interface{} {
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

func flattenFirewallPolicy64Srcaddr(d *schema.ResourceData, v *[]models.FirewallPolicy64Srcaddr, prefix string, sort bool) interface{} {
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

func refreshObjectFirewallPolicy64(d *schema.ResourceData, o *models.FirewallPolicy64, sv string, sort bool) diag.Diagnostics {
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

	if o.Dstaddr != nil {
		if err = d.Set("dstaddr", flattenFirewallPolicy64Dstaddr(d, o.Dstaddr, "dstaddr", sort)); err != nil {
			return diag.Errorf("error reading dstaddr: %v", err)
		}
	}

	if o.Dstintf != nil {
		v := *o.Dstintf

		if err = d.Set("dstintf", v); err != nil {
			return diag.Errorf("error reading dstintf: %v", err)
		}
	}

	if o.Fixedport != nil {
		v := *o.Fixedport

		if err = d.Set("fixedport", v); err != nil {
			return diag.Errorf("error reading fixedport: %v", err)
		}
	}

	if o.Ippool != nil {
		v := *o.Ippool

		if err = d.Set("ippool", v); err != nil {
			return diag.Errorf("error reading ippool: %v", err)
		}
	}

	if o.Logtraffic != nil {
		v := *o.Logtraffic

		if err = d.Set("logtraffic", v); err != nil {
			return diag.Errorf("error reading logtraffic: %v", err)
		}
	}

	if o.LogtrafficStart != nil {
		v := *o.LogtrafficStart

		if err = d.Set("logtraffic_start", v); err != nil {
			return diag.Errorf("error reading logtraffic_start: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.PerIpShaper != nil {
		v := *o.PerIpShaper

		if err = d.Set("per_ip_shaper", v); err != nil {
			return diag.Errorf("error reading per_ip_shaper: %v", err)
		}
	}

	if o.PermitAnyHost != nil {
		v := *o.PermitAnyHost

		if err = d.Set("permit_any_host", v); err != nil {
			return diag.Errorf("error reading permit_any_host: %v", err)
		}
	}

	if o.Policyid != nil {
		v := *o.Policyid

		if err = d.Set("policyid", v); err != nil {
			return diag.Errorf("error reading policyid: %v", err)
		}
	}

	if o.Poolname != nil {
		if err = d.Set("poolname", flattenFirewallPolicy64Poolname(d, o.Poolname, "poolname", sort)); err != nil {
			return diag.Errorf("error reading poolname: %v", err)
		}
	}

	if o.Schedule != nil {
		v := *o.Schedule

		if err = d.Set("schedule", v); err != nil {
			return diag.Errorf("error reading schedule: %v", err)
		}
	}

	if o.Service != nil {
		if err = d.Set("service", flattenFirewallPolicy64Service(d, o.Service, "service", sort)); err != nil {
			return diag.Errorf("error reading service: %v", err)
		}
	}

	if o.Srcaddr != nil {
		if err = d.Set("srcaddr", flattenFirewallPolicy64Srcaddr(d, o.Srcaddr, "srcaddr", sort)); err != nil {
			return diag.Errorf("error reading srcaddr: %v", err)
		}
	}

	if o.Srcintf != nil {
		v := *o.Srcintf

		if err = d.Set("srcintf", v); err != nil {
			return diag.Errorf("error reading srcintf: %v", err)
		}
	}

	if o.Status != nil {
		v := *o.Status

		if err = d.Set("status", v); err != nil {
			return diag.Errorf("error reading status: %v", err)
		}
	}

	if o.TcpMssReceiver != nil {
		v := *o.TcpMssReceiver

		if err = d.Set("tcp_mss_receiver", v); err != nil {
			return diag.Errorf("error reading tcp_mss_receiver: %v", err)
		}
	}

	if o.TcpMssSender != nil {
		v := *o.TcpMssSender

		if err = d.Set("tcp_mss_sender", v); err != nil {
			return diag.Errorf("error reading tcp_mss_sender: %v", err)
		}
	}

	if o.TrafficShaper != nil {
		v := *o.TrafficShaper

		if err = d.Set("traffic_shaper", v); err != nil {
			return diag.Errorf("error reading traffic_shaper: %v", err)
		}
	}

	if o.TrafficShaperReverse != nil {
		v := *o.TrafficShaperReverse

		if err = d.Set("traffic_shaper_reverse", v); err != nil {
			return diag.Errorf("error reading traffic_shaper_reverse: %v", err)
		}
	}

	if o.Uuid != nil {
		v := *o.Uuid

		if err = d.Set("uuid", v); err != nil {
			return diag.Errorf("error reading uuid: %v", err)
		}
	}

	return nil
}

func expandFirewallPolicy64Dstaddr(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallPolicy64Dstaddr, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallPolicy64Dstaddr

	for i := range l {
		tmp := models.FirewallPolicy64Dstaddr{}
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

func expandFirewallPolicy64Poolname(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallPolicy64Poolname, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallPolicy64Poolname

	for i := range l {
		tmp := models.FirewallPolicy64Poolname{}
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

func expandFirewallPolicy64Service(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallPolicy64Service, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallPolicy64Service

	for i := range l {
		tmp := models.FirewallPolicy64Service{}
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

func expandFirewallPolicy64Srcaddr(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallPolicy64Srcaddr, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallPolicy64Srcaddr

	for i := range l {
		tmp := models.FirewallPolicy64Srcaddr{}
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

func getObjectFirewallPolicy64(d *schema.ResourceData, sv string) (*models.FirewallPolicy64, diag.Diagnostics) {
	obj := models.FirewallPolicy64{}
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
	if v, ok := d.GetOk("dstaddr"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("dstaddr", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallPolicy64Dstaddr(d, v, "dstaddr", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Dstaddr = t
		}
	} else if d.HasChange("dstaddr") {
		old, new := d.GetChange("dstaddr")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Dstaddr = &[]models.FirewallPolicy64Dstaddr{}
		}
	}
	if v1, ok := d.GetOk("dstintf"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dstintf", sv)
				diags = append(diags, e)
			}
			obj.Dstintf = &v2
		}
	}
	if v1, ok := d.GetOk("fixedport"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fixedport", sv)
				diags = append(diags, e)
			}
			obj.Fixedport = &v2
		}
	}
	if v1, ok := d.GetOk("ippool"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ippool", sv)
				diags = append(diags, e)
			}
			obj.Ippool = &v2
		}
	}
	if v1, ok := d.GetOk("logtraffic"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("logtraffic", sv)
				diags = append(diags, e)
			}
			obj.Logtraffic = &v2
		}
	}
	if v1, ok := d.GetOk("logtraffic_start"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("logtraffic_start", sv)
				diags = append(diags, e)
			}
			obj.LogtrafficStart = &v2
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
	if v1, ok := d.GetOk("per_ip_shaper"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("per_ip_shaper", sv)
				diags = append(diags, e)
			}
			obj.PerIpShaper = &v2
		}
	}
	if v1, ok := d.GetOk("permit_any_host"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("permit_any_host", sv)
				diags = append(diags, e)
			}
			obj.PermitAnyHost = &v2
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
	if v, ok := d.GetOk("poolname"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("poolname", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallPolicy64Poolname(d, v, "poolname", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Poolname = t
		}
	} else if d.HasChange("poolname") {
		old, new := d.GetChange("poolname")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Poolname = &[]models.FirewallPolicy64Poolname{}
		}
	}
	if v1, ok := d.GetOk("schedule"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("schedule", sv)
				diags = append(diags, e)
			}
			obj.Schedule = &v2
		}
	}
	if v, ok := d.GetOk("service"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("service", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallPolicy64Service(d, v, "service", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Service = t
		}
	} else if d.HasChange("service") {
		old, new := d.GetChange("service")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Service = &[]models.FirewallPolicy64Service{}
		}
	}
	if v, ok := d.GetOk("srcaddr"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("srcaddr", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallPolicy64Srcaddr(d, v, "srcaddr", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Srcaddr = t
		}
	} else if d.HasChange("srcaddr") {
		old, new := d.GetChange("srcaddr")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Srcaddr = &[]models.FirewallPolicy64Srcaddr{}
		}
	}
	if v1, ok := d.GetOk("srcintf"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("srcintf", sv)
				diags = append(diags, e)
			}
			obj.Srcintf = &v2
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
	if v1, ok := d.GetOk("tcp_mss_receiver"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("tcp_mss_receiver", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.TcpMssReceiver = &tmp
		}
	}
	if v1, ok := d.GetOk("tcp_mss_sender"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("tcp_mss_sender", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.TcpMssSender = &tmp
		}
	}
	if v1, ok := d.GetOk("traffic_shaper"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("traffic_shaper", sv)
				diags = append(diags, e)
			}
			obj.TrafficShaper = &v2
		}
	}
	if v1, ok := d.GetOk("traffic_shaper_reverse"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("traffic_shaper_reverse", sv)
				diags = append(diags, e)
			}
			obj.TrafficShaperReverse = &v2
		}
	}
	if v1, ok := d.GetOk("uuid"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("uuid", sv)
				diags = append(diags, e)
			}
			obj.Uuid = &v2
		}
	}
	return &obj, diags
}
