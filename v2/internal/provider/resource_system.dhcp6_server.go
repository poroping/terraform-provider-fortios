// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
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

func resourceSystemDhcp6Server() *schema.Resource {
	return &schema.Resource{
		Description: "Configure DHCPv6 servers.",

		CreateContext: resourceSystemDhcp6ServerCreate,
		ReadContext:   resourceSystemDhcp6ServerRead,
		UpdateContext: resourceSystemDhcp6ServerUpdate,
		DeleteContext: resourceSystemDhcp6ServerDelete,

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

				Description: "Options for assigning DNS servers to DHCPv6 clients.",
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

func resourceSystemDhcp6ServerCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating SystemDhcp6Server resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectSystemDhcp6Server(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemDhcp6Server(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemDhcp6Server")
	}

	return resourceSystemDhcp6ServerRead(ctx, d, meta)
}

func resourceSystemDhcp6ServerUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemDhcp6Server(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemDhcp6Server(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemDhcp6Server resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemDhcp6Server")
	}

	return resourceSystemDhcp6ServerRead(ctx, d, meta)
}

func resourceSystemDhcp6ServerDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteSystemDhcp6Server(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemDhcp6Server resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemDhcp6ServerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemDhcp6Server(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemDhcp6Server resource: %v", err)
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

	diags := refreshObjectSystemDhcp6Server(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenSystemDhcp6ServerIpRange(d *schema.ResourceData, v *[]models.SystemDhcp6ServerIpRange, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
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

func flattenSystemDhcp6ServerPrefixRange(d *schema.ResourceData, v *[]models.SystemDhcp6ServerPrefixRange, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
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

func refreshObjectSystemDhcp6Server(d *schema.ResourceData, o *models.SystemDhcp6Server, sv string, sort bool) diag.Diagnostics {
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
		if err = d.Set("ip_range", flattenSystemDhcp6ServerIpRange(d, o.IpRange, "ip_range", sort)); err != nil {
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
		if err = d.Set("prefix_range", flattenSystemDhcp6ServerPrefixRange(d, o.PrefixRange, "prefix_range", sort)); err != nil {
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

func expandSystemDhcp6ServerIpRange(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemDhcp6ServerIpRange, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemDhcp6ServerIpRange

	for i := range l {
		tmp := models.SystemDhcp6ServerIpRange{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.end_ip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.EndIp = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
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

func expandSystemDhcp6ServerPrefixRange(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemDhcp6ServerPrefixRange, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemDhcp6ServerPrefixRange

	for i := range l {
		tmp := models.SystemDhcp6ServerPrefixRange{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.end_prefix", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.EndPrefix = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.Id = &v3
			}
		}

		pre_append = fmt.Sprintf("%s.%d.prefix_length", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int); ok {
				v3 := int64(v2)
				tmp.PrefixLength = &v3
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

func getObjectSystemDhcp6Server(d *schema.ResourceData, sv string) (*models.SystemDhcp6Server, diag.Diagnostics) {
	obj := models.SystemDhcp6Server{}
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
		t, err := expandSystemDhcp6ServerIpRange(d, v, "ip_range", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.IpRange = t
		}
	} else if d.HasChange("ip_range") {
		old, new := d.GetChange("ip_range")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.IpRange = &[]models.SystemDhcp6ServerIpRange{}
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
		t, err := expandSystemDhcp6ServerPrefixRange(d, v, "prefix_range", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.PrefixRange = t
		}
	} else if d.HasChange("prefix_range") {
		old, new := d.GetChange("prefix_range")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.PrefixRange = &[]models.SystemDhcp6ServerPrefixRange{}
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
