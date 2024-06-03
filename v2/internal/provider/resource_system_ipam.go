// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure IP address management services.

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
	"github.com/poroping/terraform-provider-fortios/v2/validators"
)

func resourceSystemIpam() *schema.Resource {
	return &schema.Resource{
		Description: "Configure IP address management services.",

		CreateContext: resourceSystemIpamCreate,
		ReadContext:   resourceSystemIpamRead,
		UpdateContext: resourceSystemIpamUpdate,
		DeleteContext: resourceSystemIpamDelete,

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
			"dynamic_sort_subtable": {
				Type:        schema.TypeBool,
				Description: "If set will sort table response by mkey",
				Optional:    true,
				Default:     false,
			},
			"pool_subnet": {
				Type:         schema.TypeString,
				ValidateFunc: validators.FortiValidateIPv4Classnet,

				Description: "Configure IPAM pool subnet, Class A - Class B subnet.",
				Optional:    true,
				Computed:    true,
			},
			"pools": {
				Type:        schema.TypeList,
				Description: "Configure IPAM pools.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"description": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 127),

							Description: "Description.",
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "IPAM pool name.",
							Optional:    true,
							Computed:    true,
						},
						"subnet": {
							Type:         schema.TypeString,
							ValidateFunc: validators.FortiValidateIPv4Classnet,

							Description: "Configure IPAM pool subnet, Class A - Class B subnet.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"rules": {
				Type:        schema.TypeList,
				Description: "Configure IPAM allocation rules.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"description": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 127),

							Description: "Description.",
							Optional:    true,
							Computed:    true,
						},
						"device": {
							Type:        schema.TypeList,
							Description: "Configure serial number or wildcard of FortiGate to match.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),

										Description: "FortiGate serial number or wildcard.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"dhcp": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable DHCP server for matching IPAM interfaces.",
							Optional:    true,
							Computed:    true,
						},
						"interface": {
							Type:        schema.TypeList,
							Description: "Configure name or wildcard of interface to match.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),

										Description: "Interface name or wildcard.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "IPAM rule name.",
							Optional:    true,
							Computed:    true,
						},
						"pool": {
							Type:        schema.TypeList,
							Description: "Configure name of IPAM pool to use.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),

										Description: "Ipam pool name.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"role": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"any", "lan", "wan", "dmz", "undefined"}, false),

							Description: "Configure role of interface to match.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"server_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"fabric-root"}, false),

				Description: "Configure the type of IPAM server to use.",
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable IP address management services.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSystemIpamCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemIpam(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemIpam(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemIpam")
	}

	return resourceSystemIpamRead(ctx, d, meta)
}

func resourceSystemIpamUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemIpam(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemIpam(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemIpam resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemIpam")
	}

	return resourceSystemIpamRead(ctx, d, meta)
}

func resourceSystemIpamDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectSystemIpam(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateSystemIpam(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemIpam resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemIpamRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemIpam(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemIpam resource: %v", err)
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

	diags := refreshObjectSystemIpam(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenSystemIpamPools(d *schema.ResourceData, v *[]models.SystemIpamPools, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Description; tmp != nil {
				v["description"] = *tmp
			}

			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.Subnet; tmp != nil {
				if tmp = utils.Ipv4Read(d, fmt.Sprintf("%s.%d.subnet", prefix, i), *tmp); tmp != nil {
					v["subnet"] = *tmp
				}
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenSystemIpamRules(d *schema.ResourceData, v *[]models.SystemIpamRules, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Description; tmp != nil {
				v["description"] = *tmp
			}

			if tmp := cfg.Device; tmp != nil {
				v["device"] = flattenSystemIpamRulesDevice(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "device"), sort)
			}

			if tmp := cfg.Dhcp; tmp != nil {
				v["dhcp"] = *tmp
			}

			if tmp := cfg.Interface; tmp != nil {
				v["interface"] = flattenSystemIpamRulesInterface(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "interface"), sort)
			}

			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.Pool; tmp != nil {
				v["pool"] = flattenSystemIpamRulesPool(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "pool"), sort)
			}

			if tmp := cfg.Role; tmp != nil {
				v["role"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenSystemIpamRulesDevice(d *schema.ResourceData, v *[]models.SystemIpamRulesDevice, prefix string, sort bool) interface{} {
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

func flattenSystemIpamRulesInterface(d *schema.ResourceData, v *[]models.SystemIpamRulesInterface, prefix string, sort bool) interface{} {
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

func flattenSystemIpamRulesPool(d *schema.ResourceData, v *[]models.SystemIpamRulesPool, prefix string, sort bool) interface{} {
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

func refreshObjectSystemIpam(d *schema.ResourceData, o *models.SystemIpam, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.PoolSubnet != nil {
		v := *o.PoolSubnet
		if current, ok := d.GetOk("pool_subnet"); ok {
			if s, ok := current.(string); ok {
				v = utils.ValidateConvIPMask2CIDR(s, v)
			}
		}

		if err = d.Set("pool_subnet", v); err != nil {
			return diag.Errorf("error reading pool_subnet: %v", err)
		}
	}

	if o.Pools != nil {
		if err = d.Set("pools", flattenSystemIpamPools(d, o.Pools, "pools", sort)); err != nil {
			return diag.Errorf("error reading pools: %v", err)
		}
	}

	if o.Rules != nil {
		if err = d.Set("rules", flattenSystemIpamRules(d, o.Rules, "rules", sort)); err != nil {
			return diag.Errorf("error reading rules: %v", err)
		}
	}

	if o.ServerType != nil {
		v := *o.ServerType

		if err = d.Set("server_type", v); err != nil {
			return diag.Errorf("error reading server_type: %v", err)
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

func expandSystemIpamPools(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemIpamPools, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemIpamPools

	for i := range l {
		tmp := models.SystemIpamPools{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.description", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Description = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

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

func expandSystemIpamRules(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemIpamRules, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemIpamRules

	for i := range l {
		tmp := models.SystemIpamRules{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.description", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Description = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.device", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSystemIpamRulesDevice(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemIpamRulesDevice
			// 	}
			tmp.Device = v2

		}

		pre_append = fmt.Sprintf("%s.%d.dhcp", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dhcp = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.interface", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSystemIpamRulesInterface(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemIpamRulesInterface
			// 	}
			tmp.Interface = v2

		}

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.pool", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandSystemIpamRulesPool(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.SystemIpamRulesPool
			// 	}
			tmp.Pool = v2

		}

		pre_append = fmt.Sprintf("%s.%d.role", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Role = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemIpamRulesDevice(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemIpamRulesDevice, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemIpamRulesDevice

	for i := range l {
		tmp := models.SystemIpamRulesDevice{}
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

func expandSystemIpamRulesInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemIpamRulesInterface, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemIpamRulesInterface

	for i := range l {
		tmp := models.SystemIpamRulesInterface{}
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

func expandSystemIpamRulesPool(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemIpamRulesPool, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemIpamRulesPool

	for i := range l {
		tmp := models.SystemIpamRulesPool{}
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

func getObjectSystemIpam(d *schema.ResourceData, sv string) (*models.SystemIpam, diag.Diagnostics) {
	obj := models.SystemIpam{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("pool_subnet"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v7.2.1") {
				e := utils.AttributeVersionWarning("pool_subnet", sv)
				diags = append(diags, e)
			}
			obj.PoolSubnet = &v2
		}
	}
	if v, ok := d.GetOk("pools"); ok {
		if !utils.CheckVer(sv, "v7.2.1", "") {
			e := utils.AttributeVersionWarning("pools", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemIpamPools(d, v, "pools", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Pools = t
		}
	} else if d.HasChange("pools") {
		old, new := d.GetChange("pools")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Pools = &[]models.SystemIpamPools{}
		}
	}
	if v, ok := d.GetOk("rules"); ok {
		if !utils.CheckVer(sv, "v7.2.1", "") {
			e := utils.AttributeVersionWarning("rules", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemIpamRules(d, v, "rules", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Rules = t
		}
	} else if d.HasChange("rules") {
		old, new := d.GetChange("rules")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Rules = &[]models.SystemIpamRules{}
		}
	}
	if v1, ok := d.GetOk("server_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("server_type", sv)
				diags = append(diags, e)
			}
			obj.ServerType = &v2
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
func getEmptyObjectSystemIpam(d *schema.ResourceData, sv string) (*models.SystemIpam, diag.Diagnostics) {
	obj := models.SystemIpam{}
	diags := diag.Diagnostics{}

	obj.Pools = &[]models.SystemIpamPools{}
	obj.Rules = &[]models.SystemIpamRules{}

	return &obj, diags
}
