// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure SNMP.

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

func resourceWirelessControllerSnmp() *schema.Resource {
	return &schema.Resource{
		Description: "Configure SNMP.",

		CreateContext: resourceWirelessControllerSnmpCreate,
		ReadContext:   resourceWirelessControllerSnmpRead,
		UpdateContext: resourceWirelessControllerSnmpUpdate,
		DeleteContext: resourceWirelessControllerSnmpDelete,

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
			"community": {
				Type:        schema.TypeList,
				Description: "SNMP Community Configuration.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hosts": {
							Type:        schema.TypeList,
							Description: "Configure IPv4 SNMP managers (hosts).",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type: schema.TypeInt,

										Description: "Host entry ID.",
										Optional:    true,
										Computed:    true,
									},
									"ip": {
										Type: schema.TypeString,

										Description: "IPv4 address of the SNMP manager (host).",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"id": {
							Type: schema.TypeInt,

							Description: "Community ID.",
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Community name.",
							Optional:    true,
							Computed:    true,
						},
						"query_v1_status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable SNMP v1 queries.",
							Optional:    true,
							Computed:    true,
						},
						"query_v2c_status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable SNMP v2c queries.",
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable this SNMP community.",
							Optional:    true,
							Computed:    true,
						},
						"trap_v1_status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable SNMP v1 traps.",
							Optional:    true,
							Computed:    true,
						},
						"trap_v2c_status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable SNMP v2c traps.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"contact_info": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),

				Description: "Contact Information.",
				Optional:    true,
				Computed:    true,
			},
			"engine_id": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 23),

				Description: "AC SNMP engineID string (maximum 24 characters).",
				Optional:    true,
				Computed:    true,
			},
			"trap_high_cpu_threshold": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(10, 100),

				Description: "CPU usage when trap is sent.",
				Optional:    true,
				Computed:    true,
			},
			"trap_high_mem_threshold": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(10, 100),

				Description: "Memory usage when trap is sent.",
				Optional:    true,
				Computed:    true,
			},
			"user": {
				Type:        schema.TypeList,
				Description: "SNMP User Configuration.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auth_proto": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"md5", "sha"}, false),

							Description: "Authentication protocol.",
							Optional:    true,
							Computed:    true,
						},
						"auth_pwd": {
							Type: schema.TypeString,

							Description: "Password for authentication protocol.",
							Optional:    true,
							Computed:    true,
							Sensitive:   true,
						},
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 32),

							Description: "SNMP user name.",
							Optional:    true,
							Computed:    true,
						},
						"notify_hosts": {
							Type:         schema.TypeString,
							ValidateFunc: validation.IsIPv4Address,

							Description: "Configure SNMP User Notify Hosts.",
							Optional:    true,
							Computed:    true,
						},
						"priv_proto": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"aes", "des", "aes256", "aes256cisco"}, false),

							Description: "Privacy (encryption) protocol.",
							Optional:    true,
							Computed:    true,
						},
						"priv_pwd": {
							Type: schema.TypeString,

							Description: "Password for privacy (encryption) protocol.",
							Optional:    true,
							Computed:    true,
							Sensitive:   true,
						},
						"queries": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable SNMP queries for this user.",
							Optional:    true,
							Computed:    true,
						},
						"security_level": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"no-auth-no-priv", "auth-no-priv", "auth-priv"}, false),

							Description: "Security level for message authentication and encryption.",
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "SNMP user enable.",
							Optional:    true,
							Computed:    true,
						},
						"trap_status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable traps for this SNMP user.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func resourceWirelessControllerSnmpCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectWirelessControllerSnmp(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateWirelessControllerSnmp(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WirelessControllerSnmp")
	}

	return resourceWirelessControllerSnmpRead(ctx, d, meta)
}

func resourceWirelessControllerSnmpUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectWirelessControllerSnmp(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateWirelessControllerSnmp(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating WirelessControllerSnmp resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WirelessControllerSnmp")
	}

	return resourceWirelessControllerSnmpRead(ctx, d, meta)
}

func resourceWirelessControllerSnmpDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectWirelessControllerSnmp(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateWirelessControllerSnmp(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting WirelessControllerSnmp resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerSnmpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadWirelessControllerSnmp(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WirelessControllerSnmp resource: %v", err)
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

	diags := refreshObjectWirelessControllerSnmp(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenWirelessControllerSnmpCommunity(d *schema.ResourceData, v *[]models.WirelessControllerSnmpCommunity, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Hosts; tmp != nil {
				v["hosts"] = flattenWirelessControllerSnmpCommunityHosts(d, tmp, prefix+"hosts", sort)
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.QueryV1Status; tmp != nil {
				v["query_v1_status"] = *tmp
			}

			if tmp := cfg.QueryV2cStatus; tmp != nil {
				v["query_v2c_status"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			if tmp := cfg.TrapV1Status; tmp != nil {
				v["trap_v1_status"] = *tmp
			}

			if tmp := cfg.TrapV2cStatus; tmp != nil {
				v["trap_v2c_status"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenWirelessControllerSnmpCommunityHosts(d *schema.ResourceData, v *[]models.WirelessControllerSnmpCommunityHosts, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Ip; tmp != nil {
				v["ip"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenWirelessControllerSnmpUser(d *schema.ResourceData, v *[]models.WirelessControllerSnmpUser, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.AuthProto; tmp != nil {
				v["auth_proto"] = *tmp
			}

			if tmp := cfg.AuthPwd; tmp != nil {
				v["auth_pwd"] = *tmp
			}

			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.NotifyHosts; tmp != nil {
				v["notify_hosts"] = *tmp
			}

			if tmp := cfg.PrivProto; tmp != nil {
				v["priv_proto"] = *tmp
			}

			if tmp := cfg.PrivPwd; tmp != nil {
				v["priv_pwd"] = *tmp
			}

			if tmp := cfg.Queries; tmp != nil {
				v["queries"] = *tmp
			}

			if tmp := cfg.SecurityLevel; tmp != nil {
				v["security_level"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			if tmp := cfg.TrapStatus; tmp != nil {
				v["trap_status"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func refreshObjectWirelessControllerSnmp(d *schema.ResourceData, o *models.WirelessControllerSnmp, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Community != nil {
		if err = d.Set("community", flattenWirelessControllerSnmpCommunity(d, o.Community, "community", sort)); err != nil {
			return diag.Errorf("error reading community: %v", err)
		}
	}

	if o.ContactInfo != nil {
		v := *o.ContactInfo

		if err = d.Set("contact_info", v); err != nil {
			return diag.Errorf("error reading contact_info: %v", err)
		}
	}

	if o.EngineId != nil {
		v := *o.EngineId

		if err = d.Set("engine_id", v); err != nil {
			return diag.Errorf("error reading engine_id: %v", err)
		}
	}

	if o.TrapHighCpuThreshold != nil {
		v := *o.TrapHighCpuThreshold

		if err = d.Set("trap_high_cpu_threshold", v); err != nil {
			return diag.Errorf("error reading trap_high_cpu_threshold: %v", err)
		}
	}

	if o.TrapHighMemThreshold != nil {
		v := *o.TrapHighMemThreshold

		if err = d.Set("trap_high_mem_threshold", v); err != nil {
			return diag.Errorf("error reading trap_high_mem_threshold: %v", err)
		}
	}

	if o.User != nil {
		if err = d.Set("user", flattenWirelessControllerSnmpUser(d, o.User, "user", sort)); err != nil {
			return diag.Errorf("error reading user: %v", err)
		}
	}

	return nil
}

func expandWirelessControllerSnmpCommunity(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WirelessControllerSnmpCommunity, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WirelessControllerSnmpCommunity

	for i := range l {
		tmp := models.WirelessControllerSnmpCommunity{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.hosts", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandWirelessControllerSnmpCommunityHosts(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.WirelessControllerSnmpCommunityHosts
			// 	}
			tmp.Hosts = v2

		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.query_v1_status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.QueryV1Status = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.query_v2c_status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.QueryV2cStatus = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.trap_v1_status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.TrapV1Status = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.trap_v2c_status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.TrapV2cStatus = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandWirelessControllerSnmpCommunityHosts(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WirelessControllerSnmpCommunityHosts, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WirelessControllerSnmpCommunityHosts

	for i := range l {
		tmp := models.WirelessControllerSnmpCommunityHosts{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ip = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandWirelessControllerSnmpUser(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.WirelessControllerSnmpUser, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.WirelessControllerSnmpUser

	for i := range l {
		tmp := models.WirelessControllerSnmpUser{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.auth_proto", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AuthProto = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.auth_pwd", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AuthPwd = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.notify_hosts", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.NotifyHosts = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.priv_proto", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PrivProto = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.priv_pwd", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.PrivPwd = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.queries", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Queries = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.security_level", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SecurityLevel = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.trap_status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.TrapStatus = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectWirelessControllerSnmp(d *schema.ResourceData, sv string) (*models.WirelessControllerSnmp, diag.Diagnostics) {
	obj := models.WirelessControllerSnmp{}
	diags := diag.Diagnostics{}

	if v, ok := d.GetOk("community"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("community", sv)
			diags = append(diags, e)
		}
		t, err := expandWirelessControllerSnmpCommunity(d, v, "community", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Community = t
		}
	} else if d.HasChange("community") {
		old, new := d.GetChange("community")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Community = &[]models.WirelessControllerSnmpCommunity{}
		}
	}
	if v1, ok := d.GetOk("contact_info"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("contact_info", sv)
				diags = append(diags, e)
			}
			obj.ContactInfo = &v2
		}
	}
	if v1, ok := d.GetOk("engine_id"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("engine_id", sv)
				diags = append(diags, e)
			}
			obj.EngineId = &v2
		}
	}
	if v1, ok := d.GetOk("trap_high_cpu_threshold"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("trap_high_cpu_threshold", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.TrapHighCpuThreshold = &tmp
		}
	}
	if v1, ok := d.GetOk("trap_high_mem_threshold"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("trap_high_mem_threshold", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.TrapHighMemThreshold = &tmp
		}
	}
	if v, ok := d.GetOk("user"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("user", sv)
			diags = append(diags, e)
		}
		t, err := expandWirelessControllerSnmpUser(d, v, "user", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.User = t
		}
	} else if d.HasChange("user") {
		old, new := d.GetChange("user")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.User = &[]models.WirelessControllerSnmpUser{}
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectWirelessControllerSnmp(d *schema.ResourceData, sv string) (*models.WirelessControllerSnmp, diag.Diagnostics) {
	obj := models.WirelessControllerSnmp{}
	diags := diag.Diagnostics{}

	obj.Community = &[]models.WirelessControllerSnmpCommunity{}
	obj.User = &[]models.WirelessControllerSnmpUser{}

	return &obj, diags
}
