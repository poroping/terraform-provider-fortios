// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: SNMP community configuration.

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

func resourceSystemSnmpCommunity() *schema.Resource {
	return &schema.Resource{
		Description: "SNMP community configuration.",

		CreateContext: resourceSystemSnmpCommunityCreate,
		ReadContext:   resourceSystemSnmpCommunityRead,
		UpdateContext: resourceSystemSnmpCommunityUpdate,
		DeleteContext: resourceSystemSnmpCommunityDelete,

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
				RequiredWith: []string{"fosid"},
			},
			"dynamic_sort_subtable": {
				Type:        schema.TypeBool,
				Description: "If set will sort table response by mkey",
				Optional:    true,
				Default:     false,
			},
			"events": {
				Type: schema.TypeString,

				DiffSuppressFunc: suppressors.DiffFakeListEqual,
				Description:      "SNMP trap events.",
				Optional:         true,
				Computed:         true,
			},
			"hosts": {
				Type:        schema.TypeList,
				Description: "Configure IPv4 SNMP managers (hosts).",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ha_direct": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable direct management of HA cluster members.",
							Optional:    true,
							Computed:    true,
						},
						"host_type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"any", "query", "trap"}, false),

							Description: "Control whether the SNMP manager sends SNMP queries, receives SNMP traps, or both. No traps will be sent when IP type is subnet.",
							Optional:    true,
							Computed:    true,
						},
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
						"source_ip": {
							Type:         schema.TypeString,
							ValidateFunc: validation.IsIPv4Address,

							Description: "Source IPv4 address for SNMP traps.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"hosts6": {
				Type:        schema.TypeList,
				Description: "Configure IPv6 SNMP managers.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ha_direct": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable direct management of HA cluster members.",
							Optional:    true,
							Computed:    true,
						},
						"host_type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"any", "query", "trap"}, false),

							Description: "Control whether the SNMP manager sends SNMP queries, receives SNMP traps, or both.",
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Type: schema.TypeInt,

							Description: "Host6 entry ID.",
							Optional:    true,
							Computed:    true,
						},
						"ipv6": {
							Type:             schema.TypeString,
							ValidateFunc:     validators.FortiValidateIPv6Prefix,
							DiffSuppressFunc: suppressors.DiffCidrEqual,
							Description:      "SNMP manager IPv6 address prefix.",
							Optional:         true,
							Computed:         true,
						},
						"source_ipv6": {
							Type:             schema.TypeString,
							ValidateFunc:     validation.IsIPv6Address,
							DiffSuppressFunc: suppressors.DiffIPEqual,
							Description:      "Source IPv6 address for SNMP traps.",
							Optional:         true,
							Computed:         true,
						},
					},
				},
			},
			"fosid": {
				Type: schema.TypeInt,

				Description: "Community ID.",
				ForceNew:    true,
				Optional:    true,
				Computed:    true,
			},
			"mib_view": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 32),

				Description: "SNMP access control MIB view.",
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
			"query_v1_port": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),

				Description: "SNMP v1 query port (default = 161).",
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
			"query_v2c_port": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "SNMP v2c query port (default = 161).",
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
			"trap_v1_lport": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),

				Description: "SNMP v1 trap local port (default = 162).",
				Optional:    true,
				Computed:    true,
			},
			"trap_v1_rport": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),

				Description: "SNMP v1 trap remote port (default = 162).",
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
			"trap_v2c_lport": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),

				Description: "SNMP v2c trap local port (default = 162).",
				Optional:    true,
				Computed:    true,
			},
			"trap_v2c_rport": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),

				Description: "SNMP v2c trap remote port (default = 162).",
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
			"vdoms": {
				Type:        schema.TypeList,
				Description: "SNMP access control VDOMs.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "VDOM name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func resourceSystemSnmpCommunityCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
	key := "id"
	if v, ok := d.GetOk(key); ok {
		mkey = utils.ParseMkey(v)
		if mkey == "" && allow_append {
			return diag.Errorf("error creating SystemSnmpCommunity resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectSystemSnmpCommunity(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemSnmpCommunity(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemSnmpCommunity")
	}

	return resourceSystemSnmpCommunityRead(ctx, d, meta)
}

func resourceSystemSnmpCommunityUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemSnmpCommunity(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemSnmpCommunity(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemSnmpCommunity resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemSnmpCommunity")
	}

	return resourceSystemSnmpCommunityRead(ctx, d, meta)
}

func resourceSystemSnmpCommunityDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteSystemSnmpCommunity(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemSnmpCommunity resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSnmpCommunityRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemSnmpCommunity(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemSnmpCommunity resource: %v", err)
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

	diags := refreshObjectSystemSnmpCommunity(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenSystemSnmpCommunityHosts(d *schema.ResourceData, v *[]models.SystemSnmpCommunityHosts, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.HaDirect; tmp != nil {
				v["ha_direct"] = *tmp
			}

			if tmp := cfg.HostType; tmp != nil {
				v["host_type"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Ip; tmp != nil {
				v["ip"] = *tmp
			}

			if tmp := cfg.SourceIp; tmp != nil {
				v["source_ip"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenSystemSnmpCommunityHosts6(d *schema.ResourceData, v *[]models.SystemSnmpCommunityHosts6, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.HaDirect; tmp != nil {
				v["ha_direct"] = *tmp
			}

			if tmp := cfg.HostType; tmp != nil {
				v["host_type"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Ipv6; tmp != nil {
				v["ipv6"] = *tmp
			}

			if tmp := cfg.SourceIpv6; tmp != nil {
				v["source_ipv6"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenSystemSnmpCommunityVdoms(d *schema.ResourceData, v *[]models.SystemSnmpCommunityVdoms, prefix string, sort bool) interface{} {
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

func refreshObjectSystemSnmpCommunity(d *schema.ResourceData, o *models.SystemSnmpCommunity, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Events != nil {
		v := *o.Events

		if err = d.Set("events", v); err != nil {
			return diag.Errorf("error reading events: %v", err)
		}
	}

	if o.Hosts != nil {
		if err = d.Set("hosts", flattenSystemSnmpCommunityHosts(d, o.Hosts, "hosts", sort)); err != nil {
			return diag.Errorf("error reading hosts: %v", err)
		}
	}

	if o.Hosts6 != nil {
		if err = d.Set("hosts6", flattenSystemSnmpCommunityHosts6(d, o.Hosts6, "hosts6", sort)); err != nil {
			return diag.Errorf("error reading hosts6: %v", err)
		}
	}

	if o.Id != nil {
		v := *o.Id

		if err = d.Set("fosid", v); err != nil {
			return diag.Errorf("error reading fosid: %v", err)
		}
	}

	if o.MibView != nil {
		v := *o.MibView

		if err = d.Set("mib_view", v); err != nil {
			return diag.Errorf("error reading mib_view: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.QueryV1Port != nil {
		v := *o.QueryV1Port

		if err = d.Set("query_v1_port", v); err != nil {
			return diag.Errorf("error reading query_v1_port: %v", err)
		}
	}

	if o.QueryV1Status != nil {
		v := *o.QueryV1Status

		if err = d.Set("query_v1_status", v); err != nil {
			return diag.Errorf("error reading query_v1_status: %v", err)
		}
	}

	if o.QueryV2cPort != nil {
		v := *o.QueryV2cPort

		if err = d.Set("query_v2c_port", v); err != nil {
			return diag.Errorf("error reading query_v2c_port: %v", err)
		}
	}

	if o.QueryV2cStatus != nil {
		v := *o.QueryV2cStatus

		if err = d.Set("query_v2c_status", v); err != nil {
			return diag.Errorf("error reading query_v2c_status: %v", err)
		}
	}

	if o.Status != nil {
		v := *o.Status

		if err = d.Set("status", v); err != nil {
			return diag.Errorf("error reading status: %v", err)
		}
	}

	if o.TrapV1Lport != nil {
		v := *o.TrapV1Lport

		if err = d.Set("trap_v1_lport", v); err != nil {
			return diag.Errorf("error reading trap_v1_lport: %v", err)
		}
	}

	if o.TrapV1Rport != nil {
		v := *o.TrapV1Rport

		if err = d.Set("trap_v1_rport", v); err != nil {
			return diag.Errorf("error reading trap_v1_rport: %v", err)
		}
	}

	if o.TrapV1Status != nil {
		v := *o.TrapV1Status

		if err = d.Set("trap_v1_status", v); err != nil {
			return diag.Errorf("error reading trap_v1_status: %v", err)
		}
	}

	if o.TrapV2cLport != nil {
		v := *o.TrapV2cLport

		if err = d.Set("trap_v2c_lport", v); err != nil {
			return diag.Errorf("error reading trap_v2c_lport: %v", err)
		}
	}

	if o.TrapV2cRport != nil {
		v := *o.TrapV2cRport

		if err = d.Set("trap_v2c_rport", v); err != nil {
			return diag.Errorf("error reading trap_v2c_rport: %v", err)
		}
	}

	if o.TrapV2cStatus != nil {
		v := *o.TrapV2cStatus

		if err = d.Set("trap_v2c_status", v); err != nil {
			return diag.Errorf("error reading trap_v2c_status: %v", err)
		}
	}

	if o.Vdoms != nil {
		if err = d.Set("vdoms", flattenSystemSnmpCommunityVdoms(d, o.Vdoms, "vdoms", sort)); err != nil {
			return diag.Errorf("error reading vdoms: %v", err)
		}
	}

	return nil
}

func expandSystemSnmpCommunityHosts(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemSnmpCommunityHosts, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemSnmpCommunityHosts

	for i := range l {
		tmp := models.SystemSnmpCommunityHosts{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.ha_direct", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.HaDirect = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.host_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.HostType = &v2
			}
		}

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

		pre_append = fmt.Sprintf("%s.%d.source_ip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SourceIp = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemSnmpCommunityHosts6(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemSnmpCommunityHosts6, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemSnmpCommunityHosts6

	for i := range l {
		tmp := models.SystemSnmpCommunityHosts6{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.ha_direct", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.HaDirect = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.host_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.HostType = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ipv6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ipv6 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.source_ipv6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.SourceIpv6 = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemSnmpCommunityVdoms(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemSnmpCommunityVdoms, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemSnmpCommunityVdoms

	for i := range l {
		tmp := models.SystemSnmpCommunityVdoms{}
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

func getObjectSystemSnmpCommunity(d *schema.ResourceData, sv string) (*models.SystemSnmpCommunity, diag.Diagnostics) {
	obj := models.SystemSnmpCommunity{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("events"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("events", sv)
				diags = append(diags, e)
			}
			obj.Events = &v2
		}
	}
	if v, ok := d.GetOk("hosts"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("hosts", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemSnmpCommunityHosts(d, v, "hosts", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Hosts = t
		}
	} else if d.HasChange("hosts") {
		old, new := d.GetChange("hosts")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Hosts = &[]models.SystemSnmpCommunityHosts{}
		}
	}
	if v, ok := d.GetOk("hosts6"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("hosts6", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemSnmpCommunityHosts6(d, v, "hosts6", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Hosts6 = t
		}
	} else if d.HasChange("hosts6") {
		old, new := d.GetChange("hosts6")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Hosts6 = &[]models.SystemSnmpCommunityHosts6{}
		}
	}
	if v1, ok := d.GetOk("fosid"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fosid", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Id = &tmp
		}
	}
	if v1, ok := d.GetOk("mib_view"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.2.0", "") {
				e := utils.AttributeVersionWarning("mib_view", sv)
				diags = append(diags, e)
			}
			obj.MibView = &v2
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
	if v1, ok := d.GetOk("query_v1_port"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("query_v1_port", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.QueryV1Port = &tmp
		}
	}
	if v1, ok := d.GetOk("query_v1_status"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("query_v1_status", sv)
				diags = append(diags, e)
			}
			obj.QueryV1Status = &v2
		}
	}
	if v1, ok := d.GetOk("query_v2c_port"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("query_v2c_port", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.QueryV2cPort = &tmp
		}
	}
	if v1, ok := d.GetOk("query_v2c_status"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("query_v2c_status", sv)
				diags = append(diags, e)
			}
			obj.QueryV2cStatus = &v2
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
	if v1, ok := d.GetOk("trap_v1_lport"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("trap_v1_lport", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.TrapV1Lport = &tmp
		}
	}
	if v1, ok := d.GetOk("trap_v1_rport"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("trap_v1_rport", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.TrapV1Rport = &tmp
		}
	}
	if v1, ok := d.GetOk("trap_v1_status"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("trap_v1_status", sv)
				diags = append(diags, e)
			}
			obj.TrapV1Status = &v2
		}
	}
	if v1, ok := d.GetOk("trap_v2c_lport"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("trap_v2c_lport", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.TrapV2cLport = &tmp
		}
	}
	if v1, ok := d.GetOk("trap_v2c_rport"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("trap_v2c_rport", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.TrapV2cRport = &tmp
		}
	}
	if v1, ok := d.GetOk("trap_v2c_status"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("trap_v2c_status", sv)
				diags = append(diags, e)
			}
			obj.TrapV2cStatus = &v2
		}
	}
	if v, ok := d.GetOk("vdoms"); ok {
		if !utils.CheckVer(sv, "v7.2.0", "") {
			e := utils.AttributeVersionWarning("vdoms", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemSnmpCommunityVdoms(d, v, "vdoms", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Vdoms = t
		}
	} else if d.HasChange("vdoms") {
		old, new := d.GetChange("vdoms")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Vdoms = &[]models.SystemSnmpCommunityVdoms{}
		}
	}
	return &obj, diags
}
