// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure DHCPv6 servers.

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

func resourceSystemdhcp6Server() *schema.Resource {
	return &schema.Resource{
		Description: "Configure DHCPv6 servers.",

		CreateContext: resourceSystemdhcp6ServerCreate,
		ReadContext:   resourceSystemdhcp6ServerRead,
		UpdateContext: resourceSystemdhcp6ServerUpdate,
		DeleteContext: resourceSystemdhcp6ServerDelete,

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
				Description:  "If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution!",
				Optional:     true,
				Default:      false,
				RequiredWith: []string{"fosid"},
			},
			"dynamic_sort_subtable": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"delegated_prefix_iaid": {
				Type: schema.TypeInt,

				Description: "IAID of obtained delegated-prefix from the upstream interface.",
				Optional:    true,
				Computed:    true,
			},
			"dns_search_list": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"delegated", "specify"}, false),

				Description: "DNS search list options.",
				Optional:    true,
				Computed:    true,
			},
			"dns_server1": {
				Type:             schema.TypeString,
				ValidateFunc:     validation.IsIPv6Address,
				DiffSuppressFunc: suppressors.DiffIPEqual,
				Description:      "DNS server 1.",
				Optional:         true,
				Computed:         true,
			},
			"dns_server2": {
				Type:             schema.TypeString,
				ValidateFunc:     validation.IsIPv6Address,
				DiffSuppressFunc: suppressors.DiffIPEqual,
				Description:      "DNS server 2.",
				Optional:         true,
				Computed:         true,
			},
			"dns_server3": {
				Type:             schema.TypeString,
				ValidateFunc:     validation.IsIPv6Address,
				DiffSuppressFunc: suppressors.DiffIPEqual,
				Description:      "DNS server 3.",
				Optional:         true,
				Computed:         true,
			},
			"dns_server4": {
				Type:             schema.TypeString,
				ValidateFunc:     validation.IsIPv6Address,
				DiffSuppressFunc: suppressors.DiffIPEqual,
				Description:      "DNS server 4.",
				Optional:         true,
				Computed:         true,
			},
			"dns_service": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"delegated", "default", "specify"}, false),

				Description: " Options for assigning DNS servers to DHCPv6 clients.",
				Optional:    true,
				Computed:    true,
			},
			"domain": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Domain name suffix for the IP addresses that the DHCP server assigns to clients.",
				Optional:    true,
				Computed:    true,
			},
			"fosid": {
				Type: schema.TypeInt,

				Description: "ID.",
				ForceNew:    true,
				Optional:    true,
				Computed:    true,
			},
			"interface": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "DHCP server can assign IP configurations to clients connected to this interface.",
				Optional:    true,
				Computed:    true,
			},
			"ip_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"range", "delegated"}, false),

				Description: "Method used to assign client IP.",
				Optional:    true,
				Computed:    true,
			},
			"ip_range": {
				Type:        schema.TypeList,
				Description: "DHCP IP range configuration.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"end_ip": {
							Type:             schema.TypeString,
							ValidateFunc:     validation.IsIPv6Address,
							DiffSuppressFunc: suppressors.DiffIPEqual,
							Description:      "End of IP range.",
							Optional:         true,
							Computed:         true,
						},
						"id": {
							Type: schema.TypeInt,

							Description: "ID.",
							Optional:    true,
							Computed:    true,
						},
						"start_ip": {
							Type:             schema.TypeString,
							ValidateFunc:     validation.IsIPv6Address,
							DiffSuppressFunc: suppressors.DiffIPEqual,
							Description:      "Start of IP range.",
							Optional:         true,
							Computed:         true,
						},
					},
				},
			},
			"lease_time": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(300, 8640000),

				Description: "Lease time in seconds, 0 means unlimited.",
				Optional:    true,
				Computed:    true,
			},
			"option1": {
				Type: schema.TypeString,

				Description: "Option 1.",
				Optional:    true,
				Computed:    true,
			},
			"option2": {
				Type: schema.TypeString,

				Description: "Option 2.",
				Optional:    true,
				Computed:    true,
			},
			"option3": {
				Type: schema.TypeString,

				Description: "Option 3.",
				Optional:    true,
				Computed:    true,
			},
			"prefix_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"dhcp6", "ra"}, false),

				Description: "Assigning a prefix from a DHCPv6 client or RA.",
				Optional:    true,
				Computed:    true,
			},
			"prefix_range": {
				Type:        schema.TypeList,
				Description: "DHCP prefix configuration.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"end_prefix": {
							Type:             schema.TypeString,
							ValidateFunc:     validation.IsIPv6Address,
							DiffSuppressFunc: suppressors.DiffIPEqual,
							Description:      "End of prefix range.",
							Optional:         true,
							Computed:         true,
						},
						"id": {
							Type: schema.TypeInt,

							Description: "ID.",
							Optional:    true,
							Computed:    true,
						},
						"prefix_length": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 128),

							Description: "Prefix length.",
							Optional:    true,
							Computed:    true,
						},
						"start_prefix": {
							Type:             schema.TypeString,
							ValidateFunc:     validation.IsIPv6Address,
							DiffSuppressFunc: suppressors.DiffIPEqual,
							Description:      "Start of prefix range.",
							Optional:         true,
							Computed:         true,
						},
					},
				},
			},
			"rapid_commit": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable allow/disallow rapid commit.",
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable this DHCPv6 configuration.",
				Optional:    true,
				Computed:    true,
			},
			"subnet": {
				Type:             schema.TypeString,
				ValidateFunc:     validators.FortiValidateIPv6Prefix,
				DiffSuppressFunc: suppressors.DiffCidrEqual,
				Description:      "Subnet or subnet-id if the IP mode is delegated.",
				Optional:         true,
				Computed:         true,
			},
			"upstream_interface": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "Interface name from where delegated information is provided.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSystemdhcp6ServerCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating Systemdhcp6Server resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectSystemdhcp6Server(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemdhcp6Server(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("Systemdhcp6Server")
	}

	return resourceSystemdhcp6ServerRead(ctx, d, meta)
}

func resourceSystemdhcp6ServerUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemdhcp6Server(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemdhcp6Server(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating Systemdhcp6Server resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("Systemdhcp6Server")
	}

	return resourceSystemdhcp6ServerRead(ctx, d, meta)
}

func resourceSystemdhcp6ServerDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteSystemdhcp6Server(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting Systemdhcp6Server resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemdhcp6ServerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemdhcp6Server(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading Systemdhcp6Server resource: %v", err)
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

	diags := refreshObjectSystemdhcp6Server(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenSystemdhcp6ServerIpRange(v *[]models.Systemdhcp6ServerIpRange, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.EndIp; tmp != nil {
				v["end_ip"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.StartIp; tmp != nil {
				v["start_ip"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenSystemdhcp6ServerPrefixRange(v *[]models.Systemdhcp6ServerPrefixRange, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.EndPrefix; tmp != nil {
				v["end_prefix"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.PrefixLength; tmp != nil {
				v["prefix_length"] = *tmp
			}

			if tmp := cfg.StartPrefix; tmp != nil {
				v["start_prefix"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func refreshObjectSystemdhcp6Server(d *schema.ResourceData, o *models.Systemdhcp6Server, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.DelegatedPrefixIaid != nil {
		v := *o.DelegatedPrefixIaid

		if err = d.Set("delegated_prefix_iaid", v); err != nil {
			return diag.Errorf("error reading delegated_prefix_iaid: %v", err)
		}
	}

	if o.DnsSearchList != nil {
		v := *o.DnsSearchList

		if err = d.Set("dns_search_list", v); err != nil {
			return diag.Errorf("error reading dns_search_list: %v", err)
		}
	}

	if o.DnsServer1 != nil {
		v := *o.DnsServer1

		if err = d.Set("dns_server1", v); err != nil {
			return diag.Errorf("error reading dns_server1: %v", err)
		}
	}

	if o.DnsServer2 != nil {
		v := *o.DnsServer2

		if err = d.Set("dns_server2", v); err != nil {
			return diag.Errorf("error reading dns_server2: %v", err)
		}
	}

	if o.DnsServer3 != nil {
		v := *o.DnsServer3

		if err = d.Set("dns_server3", v); err != nil {
			return diag.Errorf("error reading dns_server3: %v", err)
		}
	}

	if o.DnsServer4 != nil {
		v := *o.DnsServer4

		if err = d.Set("dns_server4", v); err != nil {
			return diag.Errorf("error reading dns_server4: %v", err)
		}
	}

	if o.DnsService != nil {
		v := *o.DnsService

		if err = d.Set("dns_service", v); err != nil {
			return diag.Errorf("error reading dns_service: %v", err)
		}
	}

	if o.Domain != nil {
		v := *o.Domain

		if err = d.Set("domain", v); err != nil {
			return diag.Errorf("error reading domain: %v", err)
		}
	}

	if o.Id != nil {
		v := *o.Id

		if err = d.Set("fosid", v); err != nil {
			return diag.Errorf("error reading fosid: %v", err)
		}
	}

	if o.Interface != nil {
		v := *o.Interface

		if err = d.Set("interface", v); err != nil {
			return diag.Errorf("error reading interface: %v", err)
		}
	}

	if o.IpMode != nil {
		v := *o.IpMode

		if err = d.Set("ip_mode", v); err != nil {
			return diag.Errorf("error reading ip_mode: %v", err)
		}
	}

	if o.IpRange != nil {
		if err = d.Set("ip_range", flattenSystemdhcp6ServerIpRange(o.IpRange, sort)); err != nil {
			return diag.Errorf("error reading ip_range: %v", err)
		}
	}

	if o.LeaseTime != nil {
		v := *o.LeaseTime

		if err = d.Set("lease_time", v); err != nil {
			return diag.Errorf("error reading lease_time: %v", err)
		}
	}

	if o.Option1 != nil {
		v := *o.Option1

		if err = d.Set("option1", v); err != nil {
			return diag.Errorf("error reading option1: %v", err)
		}
	}

	if o.Option2 != nil {
		v := *o.Option2

		if err = d.Set("option2", v); err != nil {
			return diag.Errorf("error reading option2: %v", err)
		}
	}

	if o.Option3 != nil {
		v := *o.Option3

		if err = d.Set("option3", v); err != nil {
			return diag.Errorf("error reading option3: %v", err)
		}
	}

	if o.PrefixMode != nil {
		v := *o.PrefixMode

		if err = d.Set("prefix_mode", v); err != nil {
			return diag.Errorf("error reading prefix_mode: %v", err)
		}
	}

	if o.PrefixRange != nil {
		if err = d.Set("prefix_range", flattenSystemdhcp6ServerPrefixRange(o.PrefixRange, sort)); err != nil {
			return diag.Errorf("error reading prefix_range: %v", err)
		}
	}

	if o.RapidCommit != nil {
		v := *o.RapidCommit

		if err = d.Set("rapid_commit", v); err != nil {
			return diag.Errorf("error reading rapid_commit: %v", err)
		}
	}

	if o.Status != nil {
		v := *o.Status

		if err = d.Set("status", v); err != nil {
			return diag.Errorf("error reading status: %v", err)
		}
	}

	if o.Subnet != nil {
		v := *o.Subnet

		if err = d.Set("subnet", v); err != nil {
			return diag.Errorf("error reading subnet: %v", err)
		}
	}

	if o.UpstreamInterface != nil {
		v := *o.UpstreamInterface

		if err = d.Set("upstream_interface", v); err != nil {
			return diag.Errorf("error reading upstream_interface: %v", err)
		}
	}

	return nil
}

func expandSystemdhcp6ServerIpRange(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.Systemdhcp6ServerIpRange, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.Systemdhcp6ServerIpRange

	for i := range l {
		tmp := models.Systemdhcp6ServerIpRange{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.end_ip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.EndIp = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.start_ip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.StartIp = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandSystemdhcp6ServerPrefixRange(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.Systemdhcp6ServerPrefixRange, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.Systemdhcp6ServerPrefixRange

	for i := range l {
		tmp := models.Systemdhcp6ServerPrefixRange{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.end_prefix", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.EndPrefix = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.prefix_length", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.PrefixLength = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.start_prefix", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.StartPrefix = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectSystemdhcp6Server(d *schema.ResourceData, sv string) (*models.Systemdhcp6Server, diag.Diagnostics) {
	obj := models.Systemdhcp6Server{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("delegated_prefix_iaid"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v7.0.2", "") {
				e := utils.AttributeVersionWarning("delegated_prefix_iaid", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.DelegatedPrefixIaid = &tmp
		}
	}
	if v1, ok := d.GetOk("dns_search_list"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dns_search_list", sv)
				diags = append(diags, e)
			}
			obj.DnsSearchList = &v2
		}
	}
	if v1, ok := d.GetOk("dns_server1"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dns_server1", sv)
				diags = append(diags, e)
			}
			obj.DnsServer1 = &v2
		}
	}
	if v1, ok := d.GetOk("dns_server2"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dns_server2", sv)
				diags = append(diags, e)
			}
			obj.DnsServer2 = &v2
		}
	}
	if v1, ok := d.GetOk("dns_server3"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dns_server3", sv)
				diags = append(diags, e)
			}
			obj.DnsServer3 = &v2
		}
	}
	if v1, ok := d.GetOk("dns_server4"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dns_server4", sv)
				diags = append(diags, e)
			}
			obj.DnsServer4 = &v2
		}
	}
	if v1, ok := d.GetOk("dns_service"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dns_service", sv)
				diags = append(diags, e)
			}
			obj.DnsService = &v2
		}
	}
	if v1, ok := d.GetOk("domain"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("domain", sv)
				diags = append(diags, e)
			}
			obj.Domain = &v2
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
	if v1, ok := d.GetOk("interface"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("interface", sv)
				diags = append(diags, e)
			}
			obj.Interface = &v2
		}
	}
	if v1, ok := d.GetOk("ip_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ip_mode", sv)
				diags = append(diags, e)
			}
			obj.IpMode = &v2
		}
	}
	if v, ok := d.GetOk("ip_range"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("ip_range", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemdhcp6ServerIpRange(d, v, "ip_range", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.IpRange = t
		}
	} else if d.HasChange("ip_range") {
		old, new := d.GetChange("ip_range")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.IpRange = &[]models.Systemdhcp6ServerIpRange{}
		}
	}
	if v1, ok := d.GetOk("lease_time"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("lease_time", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.LeaseTime = &tmp
		}
	}
	if v1, ok := d.GetOk("option1"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("option1", sv)
				diags = append(diags, e)
			}
			obj.Option1 = &v2
		}
	}
	if v1, ok := d.GetOk("option2"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("option2", sv)
				diags = append(diags, e)
			}
			obj.Option2 = &v2
		}
	}
	if v1, ok := d.GetOk("option3"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("option3", sv)
				diags = append(diags, e)
			}
			obj.Option3 = &v2
		}
	}
	if v1, ok := d.GetOk("prefix_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("prefix_mode", sv)
				diags = append(diags, e)
			}
			obj.PrefixMode = &v2
		}
	}
	if v, ok := d.GetOk("prefix_range"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("prefix_range", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemdhcp6ServerPrefixRange(d, v, "prefix_range", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.PrefixRange = t
		}
	} else if d.HasChange("prefix_range") {
		old, new := d.GetChange("prefix_range")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.PrefixRange = &[]models.Systemdhcp6ServerPrefixRange{}
		}
	}
	if v1, ok := d.GetOk("rapid_commit"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("rapid_commit", sv)
				diags = append(diags, e)
			}
			obj.RapidCommit = &v2
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
	if v1, ok := d.GetOk("subnet"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("subnet", sv)
				diags = append(diags, e)
			}
			obj.Subnet = &v2
		}
	}
	if v1, ok := d.GetOk("upstream_interface"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("upstream_interface", sv)
				diags = append(diags, e)
			}
			obj.UpstreamInterface = &v2
		}
	}
	return &obj, diags
}
