// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure IPv4 and IPv6 central SNAT policies.

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

func resourceFirewallCentralSnatMap() *schema.Resource {
	return &schema.Resource{
		Description: "Configure IPv4 and IPv6 central SNAT policies.",

		CreateContext: resourceFirewallCentralSnatMapCreate,
		ReadContext:   resourceFirewallCentralSnatMapRead,
		UpdateContext: resourceFirewallCentralSnatMapUpdate,
		DeleteContext: resourceFirewallCentralSnatMapDelete,

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
			"comments": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 1023),

				Description: "Comment.",
				Optional:    true,
				Computed:    true,
			},
			"dst_addr": {
				Type:        schema.TypeList,
				Description: "IPv4 Destination address.",
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
			"dst_addr6": {
				Type:        schema.TypeList,
				Description: "IPv6 Destination address.",
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
				Type:        schema.TypeList,
				Description: "Destination interface name from available interfaces.",
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
			"nat": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable source NAT.",
				Optional:    true,
				Computed:    true,
			},
			"nat_ippool": {
				Type:        schema.TypeList,
				Description: "Name of the IP pools to be used to translate addresses from available IP Pools.",
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
			"nat_ippool6": {
				Type:        schema.TypeList,
				Description: "IPv6 pools to be used for source NAT.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "IPv6 pool name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"nat_port": {
				Type: schema.TypeString,

				Description: "Translated port or port range (1 to 65535, 0 means any port).",
				Optional:    true,
				Computed:    true,
			},
			"nat46": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable NAT46.",
				Optional:    true,
				Computed:    true,
			},
			"nat64": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable NAT64.",
				Optional:    true,
				Computed:    true,
			},
			"orig_addr": {
				Type:        schema.TypeList,
				Description: "IPv4 Original address.",
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
			"orig_addr6": {
				Type:        schema.TypeList,
				Description: "IPv6 Original address.",
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
			"orig_port": {
				Type: schema.TypeString,

				Description: "Original TCP port (1 to 65535, 0 means any port).",
				Optional:    true,
				Computed:    true,
			},
			"policyid": {
				Type: schema.TypeInt,

				Description: "Policy ID.",
				ForceNew:    true,
				Required:    true,
			},
			"protocol": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),

				Description: "Integer value for the protocol type (0 - 255).",
				Optional:    true,
				Computed:    true,
			},
			"srcintf": {
				Type:        schema.TypeList,
				Description: "Source interface name from available interfaces.",
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
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable the active status of this policy.",
				Optional:    true,
				Computed:    true,
			},
			"type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"ipv4", "ipv6"}, false),

				Description: "IPv4/IPv6 source NAT.",
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

func resourceFirewallCentralSnatMapCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating FirewallCentralSnatMap resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectFirewallCentralSnatMap(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateFirewallCentralSnatMap(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallCentralSnatMap")
	}

	return resourceFirewallCentralSnatMapRead(ctx, d, meta)
}

func resourceFirewallCentralSnatMapUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectFirewallCentralSnatMap(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateFirewallCentralSnatMap(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating FirewallCentralSnatMap resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallCentralSnatMap")
	}

	return resourceFirewallCentralSnatMapRead(ctx, d, meta)
}

func resourceFirewallCentralSnatMapDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteFirewallCentralSnatMap(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting FirewallCentralSnatMap resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallCentralSnatMapRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallCentralSnatMap(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallCentralSnatMap resource: %v", err)
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

	diags := refreshObjectFirewallCentralSnatMap(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenFirewallCentralSnatMapDstAddr(v *[]models.FirewallCentralSnatMapDstAddr, sort bool) interface{} {
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

func flattenFirewallCentralSnatMapDstAddr6(v *[]models.FirewallCentralSnatMapDstAddr6, sort bool) interface{} {
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

func flattenFirewallCentralSnatMapDstintf(v *[]models.FirewallCentralSnatMapDstintf, sort bool) interface{} {
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

func flattenFirewallCentralSnatMapNatIppool(v *[]models.FirewallCentralSnatMapNatIppool, sort bool) interface{} {
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

func flattenFirewallCentralSnatMapNatIppool6(v *[]models.FirewallCentralSnatMapNatIppool6, sort bool) interface{} {
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

func flattenFirewallCentralSnatMapOrigAddr(v *[]models.FirewallCentralSnatMapOrigAddr, sort bool) interface{} {
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

func flattenFirewallCentralSnatMapOrigAddr6(v *[]models.FirewallCentralSnatMapOrigAddr6, sort bool) interface{} {
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

func flattenFirewallCentralSnatMapSrcintf(v *[]models.FirewallCentralSnatMapSrcintf, sort bool) interface{} {
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

func refreshObjectFirewallCentralSnatMap(d *schema.ResourceData, o *models.FirewallCentralSnatMap, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Comments != nil {
		v := *o.Comments

		if err = d.Set("comments", v); err != nil {
			return diag.Errorf("error reading comments: %v", err)
		}
	}

	if o.DstAddr != nil {
		if err = d.Set("dst_addr", flattenFirewallCentralSnatMapDstAddr(o.DstAddr, sort)); err != nil {
			return diag.Errorf("error reading dst_addr: %v", err)
		}
	}

	if o.DstAddr6 != nil {
		if err = d.Set("dst_addr6", flattenFirewallCentralSnatMapDstAddr6(o.DstAddr6, sort)); err != nil {
			return diag.Errorf("error reading dst_addr6: %v", err)
		}
	}

	if o.Dstintf != nil {
		if err = d.Set("dstintf", flattenFirewallCentralSnatMapDstintf(o.Dstintf, sort)); err != nil {
			return diag.Errorf("error reading dstintf: %v", err)
		}
	}

	if o.Nat != nil {
		v := *o.Nat

		if err = d.Set("nat", v); err != nil {
			return diag.Errorf("error reading nat: %v", err)
		}
	}

	if o.NatIppool != nil {
		if err = d.Set("nat_ippool", flattenFirewallCentralSnatMapNatIppool(o.NatIppool, sort)); err != nil {
			return diag.Errorf("error reading nat_ippool: %v", err)
		}
	}

	if o.NatIppool6 != nil {
		if err = d.Set("nat_ippool6", flattenFirewallCentralSnatMapNatIppool6(o.NatIppool6, sort)); err != nil {
			return diag.Errorf("error reading nat_ippool6: %v", err)
		}
	}

	if o.NatPort != nil {
		v := *o.NatPort

		if err = d.Set("nat_port", v); err != nil {
			return diag.Errorf("error reading nat_port: %v", err)
		}
	}

	if o.Nat46 != nil {
		v := *o.Nat46

		if err = d.Set("nat46", v); err != nil {
			return diag.Errorf("error reading nat46: %v", err)
		}
	}

	if o.Nat64 != nil {
		v := *o.Nat64

		if err = d.Set("nat64", v); err != nil {
			return diag.Errorf("error reading nat64: %v", err)
		}
	}

	if o.OrigAddr != nil {
		if err = d.Set("orig_addr", flattenFirewallCentralSnatMapOrigAddr(o.OrigAddr, sort)); err != nil {
			return diag.Errorf("error reading orig_addr: %v", err)
		}
	}

	if o.OrigAddr6 != nil {
		if err = d.Set("orig_addr6", flattenFirewallCentralSnatMapOrigAddr6(o.OrigAddr6, sort)); err != nil {
			return diag.Errorf("error reading orig_addr6: %v", err)
		}
	}

	if o.OrigPort != nil {
		v := *o.OrigPort

		if err = d.Set("orig_port", v); err != nil {
			return diag.Errorf("error reading orig_port: %v", err)
		}
	}

	if o.Policyid != nil {
		v := *o.Policyid

		if err = d.Set("policyid", v); err != nil {
			return diag.Errorf("error reading policyid: %v", err)
		}
	}

	if o.Protocol != nil {
		v := *o.Protocol

		if err = d.Set("protocol", v); err != nil {
			return diag.Errorf("error reading protocol: %v", err)
		}
	}

	if o.Srcintf != nil {
		if err = d.Set("srcintf", flattenFirewallCentralSnatMapSrcintf(o.Srcintf, sort)); err != nil {
			return diag.Errorf("error reading srcintf: %v", err)
		}
	}

	if o.Status != nil {
		v := *o.Status

		if err = d.Set("status", v); err != nil {
			return diag.Errorf("error reading status: %v", err)
		}
	}

	if o.Type != nil {
		v := *o.Type

		if err = d.Set("type", v); err != nil {
			return diag.Errorf("error reading type: %v", err)
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

func expandFirewallCentralSnatMapDstAddr(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallCentralSnatMapDstAddr, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallCentralSnatMapDstAddr

	for i := range l {
		tmp := models.FirewallCentralSnatMapDstAddr{}
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

func expandFirewallCentralSnatMapDstAddr6(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallCentralSnatMapDstAddr6, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallCentralSnatMapDstAddr6

	for i := range l {
		tmp := models.FirewallCentralSnatMapDstAddr6{}
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

func expandFirewallCentralSnatMapDstintf(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallCentralSnatMapDstintf, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallCentralSnatMapDstintf

	for i := range l {
		tmp := models.FirewallCentralSnatMapDstintf{}
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

func expandFirewallCentralSnatMapNatIppool(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallCentralSnatMapNatIppool, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallCentralSnatMapNatIppool

	for i := range l {
		tmp := models.FirewallCentralSnatMapNatIppool{}
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

func expandFirewallCentralSnatMapNatIppool6(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallCentralSnatMapNatIppool6, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallCentralSnatMapNatIppool6

	for i := range l {
		tmp := models.FirewallCentralSnatMapNatIppool6{}
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

func expandFirewallCentralSnatMapOrigAddr(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallCentralSnatMapOrigAddr, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallCentralSnatMapOrigAddr

	for i := range l {
		tmp := models.FirewallCentralSnatMapOrigAddr{}
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

func expandFirewallCentralSnatMapOrigAddr6(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallCentralSnatMapOrigAddr6, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallCentralSnatMapOrigAddr6

	for i := range l {
		tmp := models.FirewallCentralSnatMapOrigAddr6{}
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

func expandFirewallCentralSnatMapSrcintf(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallCentralSnatMapSrcintf, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallCentralSnatMapSrcintf

	for i := range l {
		tmp := models.FirewallCentralSnatMapSrcintf{}
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

func getObjectFirewallCentralSnatMap(d *schema.ResourceData, sv string) (*models.FirewallCentralSnatMap, diag.Diagnostics) {
	obj := models.FirewallCentralSnatMap{}
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
	if v, ok := d.GetOk("dst_addr"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("dst_addr", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallCentralSnatMapDstAddr(d, v, "dst_addr", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.DstAddr = t
		}
	} else if d.HasChange("dst_addr") {
		old, new := d.GetChange("dst_addr")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.DstAddr = &[]models.FirewallCentralSnatMapDstAddr{}
		}
	}
	if v, ok := d.GetOk("dst_addr6"); ok {
		if !utils.CheckVer(sv, "v6.4.0", "") {
			e := utils.AttributeVersionWarning("dst_addr6", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallCentralSnatMapDstAddr6(d, v, "dst_addr6", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.DstAddr6 = t
		}
	} else if d.HasChange("dst_addr6") {
		old, new := d.GetChange("dst_addr6")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.DstAddr6 = &[]models.FirewallCentralSnatMapDstAddr6{}
		}
	}
	if v, ok := d.GetOk("dstintf"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("dstintf", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallCentralSnatMapDstintf(d, v, "dstintf", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Dstintf = t
		}
	} else if d.HasChange("dstintf") {
		old, new := d.GetChange("dstintf")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Dstintf = &[]models.FirewallCentralSnatMapDstintf{}
		}
	}
	if v1, ok := d.GetOk("nat"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("nat", sv)
				diags = append(diags, e)
			}
			obj.Nat = &v2
		}
	}
	if v, ok := d.GetOk("nat_ippool"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("nat_ippool", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallCentralSnatMapNatIppool(d, v, "nat_ippool", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.NatIppool = t
		}
	} else if d.HasChange("nat_ippool") {
		old, new := d.GetChange("nat_ippool")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.NatIppool = &[]models.FirewallCentralSnatMapNatIppool{}
		}
	}
	if v, ok := d.GetOk("nat_ippool6"); ok {
		if !utils.CheckVer(sv, "v6.4.0", "") {
			e := utils.AttributeVersionWarning("nat_ippool6", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallCentralSnatMapNatIppool6(d, v, "nat_ippool6", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.NatIppool6 = t
		}
	} else if d.HasChange("nat_ippool6") {
		old, new := d.GetChange("nat_ippool6")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.NatIppool6 = &[]models.FirewallCentralSnatMapNatIppool6{}
		}
	}
	if v1, ok := d.GetOk("nat_port"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("nat_port", sv)
				diags = append(diags, e)
			}
			obj.NatPort = &v2
		}
	}
	if v1, ok := d.GetOk("nat46"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.1", "") {
				e := utils.AttributeVersionWarning("nat46", sv)
				diags = append(diags, e)
			}
			obj.Nat46 = &v2
		}
	}
	if v1, ok := d.GetOk("nat64"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.1", "") {
				e := utils.AttributeVersionWarning("nat64", sv)
				diags = append(diags, e)
			}
			obj.Nat64 = &v2
		}
	}
	if v, ok := d.GetOk("orig_addr"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("orig_addr", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallCentralSnatMapOrigAddr(d, v, "orig_addr", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.OrigAddr = t
		}
	} else if d.HasChange("orig_addr") {
		old, new := d.GetChange("orig_addr")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.OrigAddr = &[]models.FirewallCentralSnatMapOrigAddr{}
		}
	}
	if v, ok := d.GetOk("orig_addr6"); ok {
		if !utils.CheckVer(sv, "v6.4.0", "") {
			e := utils.AttributeVersionWarning("orig_addr6", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallCentralSnatMapOrigAddr6(d, v, "orig_addr6", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.OrigAddr6 = t
		}
	} else if d.HasChange("orig_addr6") {
		old, new := d.GetChange("orig_addr6")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.OrigAddr6 = &[]models.FirewallCentralSnatMapOrigAddr6{}
		}
	}
	if v1, ok := d.GetOk("orig_port"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("orig_port", sv)
				diags = append(diags, e)
			}
			obj.OrigPort = &v2
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
	if v, ok := d.GetOk("srcintf"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("srcintf", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallCentralSnatMapSrcintf(d, v, "srcintf", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Srcintf = t
		}
	} else if d.HasChange("srcintf") {
		old, new := d.GetChange("srcintf")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Srcintf = &[]models.FirewallCentralSnatMapSrcintf{}
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
	if v1, ok := d.GetOk("type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.0", "") {
				e := utils.AttributeVersionWarning("type", sv)
				diags = append(diags, e)
			}
			obj.Type = &v2
		}
	}
	if v1, ok := d.GetOk("uuid"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.0", "") {
				e := utils.AttributeVersionWarning("uuid", sv)
				diags = append(diags, e)
			}
			obj.Uuid = &v2
		}
	}
	return &obj, diags
}
