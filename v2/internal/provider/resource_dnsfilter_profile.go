// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure DNS domain filter profile.

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

func resourceDnsfilterProfile() *schema.Resource {
	return &schema.Resource{
		Description: "Configure DNS domain filter profile.",

		CreateContext: resourceDnsfilterProfileCreate,
		ReadContext:   resourceDnsfilterProfileRead,
		UpdateContext: resourceDnsfilterProfileUpdate,
		DeleteContext: resourceDnsfilterProfileDelete,

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
				Description: "If set to true allows provider to overwrite existing resources instead of erroring. Useful for brownfield implementations. Use with caution!",
				Optional:    true,
				Default:     false,
			},
			"dynamic_sort_subtable": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"block_action": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"block", "redirect", "block-sevrfail"}, false),

				Description: "Action to take for blocked domains.",
				Optional:    true,
				Computed:    true,
			},
			"block_botnet": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable blocking botnet C&C DNS lookups.",
				Optional:    true,
				Computed:    true,
			},
			"comment": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Comment.",
				Optional:    true,
				Computed:    true,
			},
			"dns_translation": {
				Type:        schema.TypeList,
				Description: "DNS translation settings.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"addr_type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"ipv4", "ipv6"}, false),

							Description: "DNS translation type (IPv4 or IPv6).",
							Optional:    true,
							Computed:    true,
						},
						"dst": {
							Type:         schema.TypeString,
							ValidateFunc: validation.IsIPv4Address,

							Description: "IPv4 address or subnet on the external network to substitute for the resolved address in DNS query replies. Can be single IP address or subnet on the external network, but number of addresses must equal number of mapped IP addresses in src.",
							Optional:    true,
							Computed:    true,
						},
						"dst6": {
							Type:             schema.TypeString,
							ValidateFunc:     validation.IsIPv6Address,
							DiffSuppressFunc: suppressors.DiffIPEqual,
							Description:      "IPv6 address or subnet on the external network to substitute for the resolved address in DNS query replies. Can be single IP address or subnet on the external network, but number of addresses must equal number of mapped IP addresses in src6.",
							Optional:         true,
							Computed:         true,
						},
						"id": {
							Type: schema.TypeInt,

							Description: "ID.",
							Optional:    true,
							Computed:    true,
						},
						"netmask": {
							Type:         schema.TypeString,
							ValidateFunc: validators.FortiValidateIPv4Netmask,

							Description: "If src and dst are subnets rather than single IP addresses, enter the netmask for both src and dst.",
							Optional:    true,
							Computed:    true,
						},
						"prefix": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 128),

							Description: "If src6 and dst6 are subnets rather than single IP addresses, enter the prefix for both src6 and dst6 (1 - 128, default = 128).",
							Optional:    true,
							Computed:    true,
						},
						"src": {
							Type:         schema.TypeString,
							ValidateFunc: validation.IsIPv4Address,

							Description: "IPv4 address or subnet on the internal network to compare with the resolved address in DNS query replies. If the resolved address matches, the resolved address is substituted with dst.",
							Optional:    true,
							Computed:    true,
						},
						"src6": {
							Type:             schema.TypeString,
							ValidateFunc:     validation.IsIPv6Address,
							DiffSuppressFunc: suppressors.DiffIPEqual,
							Description:      "IPv6 address or subnet on the internal network to compare with the resolved address in DNS query replies. If the resolved address matches, the resolved address is substituted with dst6.",
							Optional:         true,
							Computed:         true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable this DNS translation entry.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"domain_filter": {
				Type:        schema.TypeList,
				Description: "Domain filter settings.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"domain_filter_table": {
							Type: schema.TypeInt,

							Description: "DNS domain filter table ID.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"external_ip_blocklist": {
				Type:        schema.TypeList,
				Description: "One or more external IP block lists.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),

							Description: "External domain block list name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"ftgd_dns": {
				Type:        schema.TypeList,
				Description: "FortiGuard DNS Filter settings.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"filters": {
							Type:        schema.TypeList,
							Description: "FortiGuard DNS domain filters.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"block", "monitor"}, false),

										Description: "Action to take for DNS requests matching the category.",
										Optional:    true,
										Computed:    true,
									},
									"category": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 255),

										Description: "Category number.",
										Optional:    true,
										Computed:    true,
									},
									"id": {
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 255),

										Description: "ID number.",
										Optional:    true,
										Computed:    true,
									},
									"log": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

										Description: "Enable/disable DNS filter logging for this DNS profile.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
						"options": {
							Type: schema.TypeString,

							DiffSuppressFunc: suppressors.DiffFakeListEqual,
							Description:      "FortiGuard DNS filter options.",
							Optional:         true,
							Computed:         true,
						},
					},
				},
			},
			"log_all_domain": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable logging of all domains visited (detailed DNS logging).",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Profile name.",
				ForceNew:    true,
				Required:    true,
			},
			"redirect_portal": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "IPv4 address of the SDNS redirect portal.",
				Optional:    true,
				Computed:    true,
			},
			"redirect_portal6": {
				Type:             schema.TypeString,
				ValidateFunc:     validation.IsIPv6Address,
				DiffSuppressFunc: suppressors.DiffIPEqual,
				Description:      "IPv6 address of the SDNS redirect portal.",
				Optional:         true,
				Computed:         true,
			},
			"safe_search": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable Google, Bing, YouTube, Qwant, DuckDuckGo safe search.",
				Optional:    true,
				Computed:    true,
			},
			"sdns_domain_log": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable domain filtering and botnet domain logging.",
				Optional:    true,
				Computed:    true,
			},
			"sdns_ftgd_err_log": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable FortiGuard SDNS rating error logging.",
				Optional:    true,
				Computed:    true,
			},
			"youtube_restrict": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"strict", "moderate"}, false),

				Description: "Set safe search for YouTube restriction level.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceDnsfilterProfileCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
	key := "name"
	if v, ok := d.GetOk(key); ok {
		mkey = utils.ParseMkey(v)
		if mkey == "" && allow_append {
			return diag.Errorf("error creating DnsfilterProfile resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectDnsfilterProfile(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateDnsfilterProfile(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("DnsfilterProfile")
	}

	return resourceDnsfilterProfileRead(ctx, d, meta)
}

func resourceDnsfilterProfileUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectDnsfilterProfile(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateDnsfilterProfile(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating DnsfilterProfile resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("DnsfilterProfile")
	}

	return resourceDnsfilterProfileRead(ctx, d, meta)
}

func resourceDnsfilterProfileDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteDnsfilterProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting DnsfilterProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceDnsfilterProfileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadDnsfilterProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading DnsfilterProfile resource: %v", err)
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

	diags := refreshObjectDnsfilterProfile(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenDnsfilterProfileDnsTranslation(v *[]models.DnsfilterProfileDnsTranslation, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.AddrType; tmp != nil {
				v["addr_type"] = *tmp
			}

			if tmp := cfg.Dst; tmp != nil {
				v["dst"] = *tmp
			}

			if tmp := cfg.Dst6; tmp != nil {
				v["dst6"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Netmask; tmp != nil {
				v["netmask"] = *tmp
			}

			if tmp := cfg.Prefix; tmp != nil {
				v["prefix"] = *tmp
			}

			if tmp := cfg.Src; tmp != nil {
				v["src"] = *tmp
			}

			if tmp := cfg.Src6; tmp != nil {
				v["src6"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func flattenDnsfilterProfileDomainFilter(v *[]models.DnsfilterProfileDomainFilter, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.DomainFilterTable; tmp != nil {
				v["domain_filter_table"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenDnsfilterProfileExternalIpBlocklist(v *[]models.DnsfilterProfileExternalIpBlocklist, sort bool) interface{} {
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

func flattenDnsfilterProfileFtgdDns(v *[]models.DnsfilterProfileFtgdDns, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Filters; tmp != nil {
				v["filters"] = flattenDnsfilterProfileFtgdDnsFilters(tmp, sort)
			}

			if tmp := cfg.Options; tmp != nil {
				v["options"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	return flat
}

func flattenDnsfilterProfileFtgdDnsFilters(v *[]models.DnsfilterProfileFtgdDnsFilters, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Action; tmp != nil {
				v["action"] = *tmp
			}

			if tmp := cfg.Category; tmp != nil {
				v["category"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Log; tmp != nil {
				v["log"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func refreshObjectDnsfilterProfile(d *schema.ResourceData, o *models.DnsfilterProfile, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.BlockAction != nil {
		v := *o.BlockAction

		if err = d.Set("block_action", v); err != nil {
			return diag.Errorf("error reading block_action: %v", err)
		}
	}

	if o.BlockBotnet != nil {
		v := *o.BlockBotnet

		if err = d.Set("block_botnet", v); err != nil {
			return diag.Errorf("error reading block_botnet: %v", err)
		}
	}

	if o.Comment != nil {
		v := *o.Comment

		if err = d.Set("comment", v); err != nil {
			return diag.Errorf("error reading comment: %v", err)
		}
	}

	if o.DnsTranslation != nil {
		if err = d.Set("dns_translation", flattenDnsfilterProfileDnsTranslation(o.DnsTranslation, sort)); err != nil {
			return diag.Errorf("error reading dns_translation: %v", err)
		}
	}

	if o.DomainFilter != nil {
		if err = d.Set("domain_filter", flattenDnsfilterProfileDomainFilter(o.DomainFilter, sort)); err != nil {
			return diag.Errorf("error reading domain_filter: %v", err)
		}
	}

	if o.ExternalIpBlocklist != nil {
		if err = d.Set("external_ip_blocklist", flattenDnsfilterProfileExternalIpBlocklist(o.ExternalIpBlocklist, sort)); err != nil {
			return diag.Errorf("error reading external_ip_blocklist: %v", err)
		}
	}

	if o.FtgdDns != nil {
		if err = d.Set("ftgd_dns", flattenDnsfilterProfileFtgdDns(o.FtgdDns, sort)); err != nil {
			return diag.Errorf("error reading ftgd_dns: %v", err)
		}
	}

	if o.LogAllDomain != nil {
		v := *o.LogAllDomain

		if err = d.Set("log_all_domain", v); err != nil {
			return diag.Errorf("error reading log_all_domain: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.RedirectPortal != nil {
		v := *o.RedirectPortal

		if err = d.Set("redirect_portal", v); err != nil {
			return diag.Errorf("error reading redirect_portal: %v", err)
		}
	}

	if o.RedirectPortal6 != nil {
		v := *o.RedirectPortal6

		if err = d.Set("redirect_portal6", v); err != nil {
			return diag.Errorf("error reading redirect_portal6: %v", err)
		}
	}

	if o.SafeSearch != nil {
		v := *o.SafeSearch

		if err = d.Set("safe_search", v); err != nil {
			return diag.Errorf("error reading safe_search: %v", err)
		}
	}

	if o.SdnsDomainLog != nil {
		v := *o.SdnsDomainLog

		if err = d.Set("sdns_domain_log", v); err != nil {
			return diag.Errorf("error reading sdns_domain_log: %v", err)
		}
	}

	if o.SdnsFtgdErrLog != nil {
		v := *o.SdnsFtgdErrLog

		if err = d.Set("sdns_ftgd_err_log", v); err != nil {
			return diag.Errorf("error reading sdns_ftgd_err_log: %v", err)
		}
	}

	if o.YoutubeRestrict != nil {
		v := *o.YoutubeRestrict

		if err = d.Set("youtube_restrict", v); err != nil {
			return diag.Errorf("error reading youtube_restrict: %v", err)
		}
	}

	return nil
}

func expandDnsfilterProfileDnsTranslation(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.DnsfilterProfileDnsTranslation, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.DnsfilterProfileDnsTranslation

	for i := range l {
		tmp := models.DnsfilterProfileDnsTranslation{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.addr_type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.AddrType = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dst", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dst = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.dst6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Dst6 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.netmask", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Netmask = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.prefix", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Prefix = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.src", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Src = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.src6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Src6 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandDnsfilterProfileDomainFilter(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.DnsfilterProfileDomainFilter, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.DnsfilterProfileDomainFilter

	for i := range l {
		tmp := models.DnsfilterProfileDomainFilter{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.domain_filter_table", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.DomainFilterTable = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandDnsfilterProfileExternalIpBlocklist(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.DnsfilterProfileExternalIpBlocklist, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.DnsfilterProfileExternalIpBlocklist

	for i := range l {
		tmp := models.DnsfilterProfileExternalIpBlocklist{}
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

func expandDnsfilterProfileFtgdDns(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.DnsfilterProfileFtgdDns, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.DnsfilterProfileFtgdDns

	for i := range l {
		tmp := models.DnsfilterProfileFtgdDns{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.filters", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandDnsfilterProfileFtgdDnsFilters(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.DnsfilterProfileFtgdDnsFilters
			// 	}
			tmp.Filters = v2

		}

		pre_append = fmt.Sprintf("%s.%d.options", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Options = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandDnsfilterProfileFtgdDnsFilters(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.DnsfilterProfileFtgdDnsFilters, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.DnsfilterProfileFtgdDnsFilters

	for i := range l {
		tmp := models.DnsfilterProfileFtgdDnsFilters{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.action", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Action = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.category", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Category = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.log", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Log = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectDnsfilterProfile(d *schema.ResourceData, sv string) (*models.DnsfilterProfile, diag.Diagnostics) {
	obj := models.DnsfilterProfile{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("block_action"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("block_action", sv)
				diags = append(diags, e)
			}
			obj.BlockAction = &v2
		}
	}
	if v1, ok := d.GetOk("block_botnet"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("block_botnet", sv)
				diags = append(diags, e)
			}
			obj.BlockBotnet = &v2
		}
	}
	if v1, ok := d.GetOk("comment"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("comment", sv)
				diags = append(diags, e)
			}
			obj.Comment = &v2
		}
	}
	if v, ok := d.GetOk("dns_translation"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("dns_translation", sv)
			diags = append(diags, e)
		}
		t, err := expandDnsfilterProfileDnsTranslation(d, v, "dns_translation", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.DnsTranslation = t
		}
	} else if d.HasChange("dns_translation") {
		old, new := d.GetChange("dns_translation")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.DnsTranslation = &[]models.DnsfilterProfileDnsTranslation{}
		}
	}
	if v, ok := d.GetOk("domain_filter"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("domain_filter", sv)
			diags = append(diags, e)
		}
		t, err := expandDnsfilterProfileDomainFilter(d, v, "domain_filter", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.DomainFilter = t
		}
	} else if d.HasChange("domain_filter") {
		old, new := d.GetChange("domain_filter")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.DomainFilter = &[]models.DnsfilterProfileDomainFilter{}
		}
	}
	if v, ok := d.GetOk("external_ip_blocklist"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("external_ip_blocklist", sv)
			diags = append(diags, e)
		}
		t, err := expandDnsfilterProfileExternalIpBlocklist(d, v, "external_ip_blocklist", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.ExternalIpBlocklist = t
		}
	} else if d.HasChange("external_ip_blocklist") {
		old, new := d.GetChange("external_ip_blocklist")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.ExternalIpBlocklist = &[]models.DnsfilterProfileExternalIpBlocklist{}
		}
	}
	if v, ok := d.GetOk("ftgd_dns"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("ftgd_dns", sv)
			diags = append(diags, e)
		}
		t, err := expandDnsfilterProfileFtgdDns(d, v, "ftgd_dns", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.FtgdDns = t
		}
	} else if d.HasChange("ftgd_dns") {
		old, new := d.GetChange("ftgd_dns")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.FtgdDns = &[]models.DnsfilterProfileFtgdDns{}
		}
	}
	if v1, ok := d.GetOk("log_all_domain"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("log_all_domain", sv)
				diags = append(diags, e)
			}
			obj.LogAllDomain = &v2
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
	if v1, ok := d.GetOk("redirect_portal"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("redirect_portal", sv)
				diags = append(diags, e)
			}
			obj.RedirectPortal = &v2
		}
	}
	if v1, ok := d.GetOk("redirect_portal6"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("redirect_portal6", sv)
				diags = append(diags, e)
			}
			obj.RedirectPortal6 = &v2
		}
	}
	if v1, ok := d.GetOk("safe_search"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("safe_search", sv)
				diags = append(diags, e)
			}
			obj.SafeSearch = &v2
		}
	}
	if v1, ok := d.GetOk("sdns_domain_log"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("sdns_domain_log", sv)
				diags = append(diags, e)
			}
			obj.SdnsDomainLog = &v2
		}
	}
	if v1, ok := d.GetOk("sdns_ftgd_err_log"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("sdns_ftgd_err_log", sv)
				diags = append(diags, e)
			}
			obj.SdnsFtgdErrLog = &v2
		}
	}
	if v1, ok := d.GetOk("youtube_restrict"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("youtube_restrict", sv)
				diags = append(diags, e)
			}
			obj.YoutubeRestrict = &v2
		}
	}
	return &obj, diags
}
