// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure DNS databases.

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
)

func resourceSystemDnsDatabase() *schema.Resource {
	return &schema.Resource{
		Description: "Configure DNS databases.",

		CreateContext: resourceSystemDnsDatabaseCreate,
		ReadContext:   resourceSystemDnsDatabaseRead,
		UpdateContext: resourceSystemDnsDatabaseUpdate,
		DeleteContext: resourceSystemDnsDatabaseDelete,

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
			"allow_transfer": {
				Type: schema.TypeString,

				Description: "DNS zone transfer IP address list.",
				Optional:    true,
				Computed:    true,
			},
			"authoritative": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable authoritative zone.",
				Optional:    true,
				Computed:    true,
			},
			"contact": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Email address of the administrator for this zone.\n\t\tYou can specify only the username (e.g. admin) or full email address (e.g. admin@test.com) \n\t\tWhen using a simple username, the domain of the email will be this zone.",
				Optional:    true,
				Computed:    true,
			},
			"dns_entry": {
				Type:        schema.TypeList,
				Description: "DNS entry.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"canonical_name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),

							Description: "Canonical name of the host.",
							Optional:    true,
							Computed:    true,
						},
						"hostname": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),

							Description: "Name of the host.",
							Optional:    true,
							Computed:    true,
						},
						"id": {
							Type: schema.TypeInt,

							Description: "DNS entry ID.",
							Optional:    true,
							Computed:    true,
						},
						"ip": {
							Type:         schema.TypeString,
							ValidateFunc: validation.IsIPv4Address,

							Description: "IPv4 address of the host.",
							Optional:    true,
							Computed:    true,
						},
						"ipv6": {
							Type:             schema.TypeString,
							ValidateFunc:     validation.IsIPv6Address,
							DiffSuppressFunc: suppressors.DiffIPEqual,
							Description:      "IPv6 address of the host.",
							Optional:         true,
							Computed:         true,
						},
						"preference": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),

							Description: "DNS entry preference, 0 is the highest preference (0 - 65535, default = 10)",
							Optional:    true,
							Computed:    true,
						},
						"status": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

							Description: "Enable/disable resource record status.",
							Optional:    true,
							Computed:    true,
						},
						"ttl": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 2147483647),

							Description: "Time-to-live for this entry (0 to 2147483647 sec, default = 0).",
							Optional:    true,
							Computed:    true,
						},
						"type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"A", "NS", "CNAME", "MX", "AAAA", "PTR", "PTR_V6"}, false),

							Description: "Resource record type.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"domain": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Domain name.",
				Optional:    true,
				Computed:    true,
			},
			"forwarder": {
				Type: schema.TypeString,

				Description: "DNS zone forwarder IP address list.",
				Optional:    true,
				Computed:    true,
			},
			"ip_master": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "IP address of master DNS server. Entries in this master DNS server and imported into the DNS zone.",
				Optional:    true,
				Computed:    true,
			},
			"ip_primary": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "IP address of primary DNS server. Entries in this primary DNS server and imported into the DNS zone.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Zone name.",
				ForceNew:    true,
				Required:    true,
			},
			"primary_name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Domain name of the default DNS server for this zone.",
				Optional:    true,
				Computed:    true,
			},
			"rr_max": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(10, 65536),

				Description: "Maximum number of resource records (10 - 65536, 0 means infinite).",
				Optional:    true,
				Computed:    true,
			},
			"source_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "Source IP for forwarding to DNS server.",
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable this DNS zone.",
				Optional:    true,
				Computed:    true,
			},
			"ttl": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 2147483647),

				Description: "Default time-to-live value for the entries of this DNS zone (0 - 2147483647 sec, default = 86400).",
				Optional:    true,
				Computed:    true,
			},
			"type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"primary", "secondary"}, false),

				Description: "Zone type (primary to manage entries directly, secondary to import entries from other zones).",
				Optional:    true,
				Computed:    true,
			},
			"view": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"shadow", "public"}, false),

				Description: "Zone view (public to serve public clients, shadow to serve internal clients).",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSystemDnsDatabaseCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating SystemDnsDatabase resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectSystemDnsDatabase(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemDnsDatabase(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemDnsDatabase")
	}

	return resourceSystemDnsDatabaseRead(ctx, d, meta)
}

func resourceSystemDnsDatabaseUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemDnsDatabase(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemDnsDatabase(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemDnsDatabase resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemDnsDatabase")
	}

	return resourceSystemDnsDatabaseRead(ctx, d, meta)
}

func resourceSystemDnsDatabaseDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteSystemDnsDatabase(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemDnsDatabase resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemDnsDatabaseRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemDnsDatabase(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemDnsDatabase resource: %v", err)
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

	diags := refreshObjectSystemDnsDatabase(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenSystemDnsDatabaseDnsEntry(v *[]models.SystemDnsDatabaseDnsEntry, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.CanonicalName; tmp != nil {
				v["canonical_name"] = *tmp
			}

			if tmp := cfg.Hostname; tmp != nil {
				v["hostname"] = *tmp
			}

			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Ip; tmp != nil {
				v["ip"] = *tmp
			}

			if tmp := cfg.Ipv6; tmp != nil {
				v["ipv6"] = *tmp
			}

			if tmp := cfg.Preference; tmp != nil {
				v["preference"] = *tmp
			}

			if tmp := cfg.Status; tmp != nil {
				v["status"] = *tmp
			}

			if tmp := cfg.Ttl; tmp != nil {
				v["ttl"] = *tmp
			}

			if tmp := cfg.Type; tmp != nil {
				v["type"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func refreshObjectSystemDnsDatabase(d *schema.ResourceData, o *models.SystemDnsDatabase, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AllowTransfer != nil {
		v := *o.AllowTransfer

		if err = d.Set("allow_transfer", v); err != nil {
			return diag.Errorf("error reading allow_transfer: %v", err)
		}
	}

	if o.Authoritative != nil {
		v := *o.Authoritative

		if err = d.Set("authoritative", v); err != nil {
			return diag.Errorf("error reading authoritative: %v", err)
		}
	}

	if o.Contact != nil {
		v := *o.Contact

		if err = d.Set("contact", v); err != nil {
			return diag.Errorf("error reading contact: %v", err)
		}
	}

	if o.DnsEntry != nil {
		if err = d.Set("dns_entry", flattenSystemDnsDatabaseDnsEntry(o.DnsEntry, sort)); err != nil {
			return diag.Errorf("error reading dns_entry: %v", err)
		}
	}

	if o.Domain != nil {
		v := *o.Domain

		if err = d.Set("domain", v); err != nil {
			return diag.Errorf("error reading domain: %v", err)
		}
	}

	if o.Forwarder != nil {
		v := *o.Forwarder

		if err = d.Set("forwarder", v); err != nil {
			return diag.Errorf("error reading forwarder: %v", err)
		}
	}

	if o.IpMaster != nil {
		v := *o.IpMaster

		if err = d.Set("ip_master", v); err != nil {
			return diag.Errorf("error reading ip_master: %v", err)
		}
	}

	if o.IpPrimary != nil {
		v := *o.IpPrimary

		if err = d.Set("ip_primary", v); err != nil {
			return diag.Errorf("error reading ip_primary: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.PrimaryName != nil {
		v := *o.PrimaryName

		if err = d.Set("primary_name", v); err != nil {
			return diag.Errorf("error reading primary_name: %v", err)
		}
	}

	if o.RrMax != nil {
		v := *o.RrMax

		if err = d.Set("rr_max", v); err != nil {
			return diag.Errorf("error reading rr_max: %v", err)
		}
	}

	if o.SourceIp != nil {
		v := *o.SourceIp

		if err = d.Set("source_ip", v); err != nil {
			return diag.Errorf("error reading source_ip: %v", err)
		}
	}

	if o.Status != nil {
		v := *o.Status

		if err = d.Set("status", v); err != nil {
			return diag.Errorf("error reading status: %v", err)
		}
	}

	if o.Ttl != nil {
		v := *o.Ttl

		if err = d.Set("ttl", v); err != nil {
			return diag.Errorf("error reading ttl: %v", err)
		}
	}

	if o.Type != nil {
		v := *o.Type

		if err = d.Set("type", v); err != nil {
			return diag.Errorf("error reading type: %v", err)
		}
	}

	if o.View != nil {
		v := *o.View

		if err = d.Set("view", v); err != nil {
			return diag.Errorf("error reading view: %v", err)
		}
	}

	return nil
}

func expandSystemDnsDatabaseDnsEntry(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemDnsDatabaseDnsEntry, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemDnsDatabaseDnsEntry

	for i := range l {
		tmp := models.SystemDnsDatabaseDnsEntry{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.canonical_name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.CanonicalName = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.hostname", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Hostname = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ip", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ip = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ipv6", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Ipv6 = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.preference", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Preference = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.status", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Status = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.ttl", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Ttl = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Type = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectSystemDnsDatabase(d *schema.ResourceData, sv string) (*models.SystemDnsDatabase, diag.Diagnostics) {
	obj := models.SystemDnsDatabase{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("allow_transfer"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("allow_transfer", sv)
				diags = append(diags, e)
			}
			obj.AllowTransfer = &v2
		}
	}
	if v1, ok := d.GetOk("authoritative"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("authoritative", sv)
				diags = append(diags, e)
			}
			obj.Authoritative = &v2
		}
	}
	if v1, ok := d.GetOk("contact"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("contact", sv)
				diags = append(diags, e)
			}
			obj.Contact = &v2
		}
	}
	if v, ok := d.GetOk("dns_entry"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("dns_entry", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemDnsDatabaseDnsEntry(d, v, "dns_entry", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.DnsEntry = t
		}
	} else if d.HasChange("dns_entry") {
		old, new := d.GetChange("dns_entry")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.DnsEntry = &[]models.SystemDnsDatabaseDnsEntry{}
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
	if v1, ok := d.GetOk("forwarder"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("forwarder", sv)
				diags = append(diags, e)
			}
			obj.Forwarder = &v2
		}
	}
	if v1, ok := d.GetOk("ip_master"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v7.0.0") {
				e := utils.AttributeVersionWarning("ip_master", sv)
				diags = append(diags, e)
			}
			obj.IpMaster = &v2
		}
	}
	if v1, ok := d.GetOk("ip_primary"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("ip_primary", sv)
				diags = append(diags, e)
			}
			obj.IpPrimary = &v2
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
	if v1, ok := d.GetOk("primary_name"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("primary_name", sv)
				diags = append(diags, e)
			}
			obj.PrimaryName = &v2
		}
	}
	if v1, ok := d.GetOk("rr_max"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "v6.4.0", "") {
				e := utils.AttributeVersionWarning("rr_max", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.RrMax = &tmp
		}
	}
	if v1, ok := d.GetOk("source_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("source_ip", sv)
				diags = append(diags, e)
			}
			obj.SourceIp = &v2
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
	if v1, ok := d.GetOk("ttl"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ttl", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Ttl = &tmp
		}
	}
	if v1, ok := d.GetOk("type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("type", sv)
				diags = append(diags, e)
			}
			obj.Type = &v2
		}
	}
	if v1, ok := d.GetOk("view"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("view", sv)
				diags = append(diags, e)
			}
			obj.View = &v2
		}
	}
	return &obj, diags
}
