// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure IPv6 firewall addresses.

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

func resourceFirewallAddress6() *schema.Resource {
	return &schema.Resource{
		Description: "Configure IPv6 firewall addresses.",

		CreateContext: resourceFirewallAddress6Create,
		ReadContext:   resourceFirewallAddress6Read,
		UpdateContext: resourceFirewallAddress6Update,
		DeleteContext: resourceFirewallAddress6Delete,

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
			"cache_ttl": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 86400),

				Description: "Minimal TTL of individual IPv6 addresses in FQDN cache.",
				Optional:    true,
				Computed:    true,
			},
			"color": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 32),

				Description: "Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).",
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
			"country": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 2),

				Description: "IPv6 addresses associated to a specific country.",
				Optional:    true,
				Computed:    true,
			},
			"end_ip": {
				Type:             schema.TypeString,
				ValidateFunc:     validation.IsIPv6Address,
				DiffSuppressFunc: suppressors.DiffIPEqual,
				Description:      "Final IP address (inclusive) in the range for the address (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).",
				Optional:         true,
				Computed:         true,
			},
			"end_mac": {
				Type: schema.TypeString,

				Description: "Last MAC address in the range.",
				Optional:    true,
				Computed:    true,
			},
			"fabric_object": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Security Fabric global object setting.",
				Optional:    true,
				Computed:    true,
			},
			"fqdn": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Fully qualified domain name.",
				Optional:    true,
				Computed:    true,
			},
			"host": {
				Type:             schema.TypeString,
				ValidateFunc:     validation.IsIPv6Address,
				DiffSuppressFunc: suppressors.DiffIPEqual,
				Description:      "Host Address.",
				Optional:         true,
				Computed:         true,
			},
			"host_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"any", "specific"}, false),

				Description: "Host type.",
				Optional:    true,
				Computed:    true,
			},
			"ip6": {
				Type:             schema.TypeString,
				ValidateFunc:     validators.FortiValidateIPv6Network,
				DiffSuppressFunc: suppressors.DiffCidrEqual,
				Description:      "IPv6 address prefix (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx).",
				Optional:         true,
				Computed:         true,
			},
			"list": {
				Type:        schema.TypeList,
				Description: "IP address list.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 89),

							Description: "IP.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"macaddr": {
				Type:        schema.TypeList,
				Description: "Multiple MAC address ranges.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"macaddr": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 127),

							Description: "MAC address ranges <start>[-<end>] separated by space.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),

				Description: "Address name.",
				ForceNew:    true,
				Required:    true,
			},
			"obj_id": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Object ID for NSX.",
				Optional:    true,
				Computed:    true,
			},
			"sdn": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "SDN.",
				Optional:    true,
				Computed:    true,
			},
			"start_ip": {
				Type:             schema.TypeString,
				ValidateFunc:     validation.IsIPv6Address,
				DiffSuppressFunc: suppressors.DiffIPEqual,
				Description:      "First IP address (inclusive) in the range for the address (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).",
				Optional:         true,
				Computed:         true,
			},
			"start_mac": {
				Type: schema.TypeString,

				Description: "First MAC address in the range.",
				Optional:    true,
				Computed:    true,
			},
			"subnet_segment": {
				Type:        schema.TypeList,
				Description: "IPv6 subnet segments.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "Name.",
							Optional:    true,
							Computed:    true,
						},
						"type": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"any", "specific"}, false),

							Description: "Subnet segment type.",
							Optional:    true,
							Computed:    true,
						},
						"value": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

							Description: "Subnet segment value.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"tagging": {
				Type:        schema.TypeList,
				Description: "Config object tagging.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"category": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "Tag category.",
							Optional:    true,
							Computed:    true,
						},
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),

							Description: "Tagging entry name.",
							Optional:    true,
							Computed:    true,
						},
						"tags": {
							Type:        schema.TypeList,
							Description: "Tags.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),

										Description: "Tag name.",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
					},
				},
			},
			"template": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "IPv6 address template.",
				Optional:    true,
				Computed:    true,
			},
			"type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"ipprefix", "iprange", "fqdn", "geography", "dynamic", "template", "mac"}, false),

				Description: "Type of IPv6 address object (default = ipprefix).",
				Optional:    true,
				Computed:    true,
			},
			"uuid": {
				Type: schema.TypeString,

				Description: "Universally Unique Identifier (UUID; automatically assigned but can be manually reset).",
				Optional:    true,
				Computed:    true,
			},
			"visibility": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable the visibility of the object in the GUI.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceFirewallAddress6Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating FirewallAddress6 resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectFirewallAddress6(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateFirewallAddress6(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallAddress6")
	}

	return resourceFirewallAddress6Read(ctx, d, meta)
}

func resourceFirewallAddress6Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectFirewallAddress6(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateFirewallAddress6(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating FirewallAddress6 resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallAddress6")
	}

	return resourceFirewallAddress6Read(ctx, d, meta)
}

func resourceFirewallAddress6Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteFirewallAddress6(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting FirewallAddress6 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallAddress6Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallAddress6(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallAddress6 resource: %v", err)
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

	diags := refreshObjectFirewallAddress6(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenFirewallAddress6List(d *schema.ResourceData, v *[]models.FirewallAddress6List, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Ip; tmp != nil {
				v["ip"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "ip")
	}

	return flat
}

func flattenFirewallAddress6Macaddr(d *schema.ResourceData, v *[]models.FirewallAddress6Macaddr, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Macaddr; tmp != nil {
				v["macaddr"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "macaddr")
	}

	return flat
}

func flattenFirewallAddress6SubnetSegment(d *schema.ResourceData, v *[]models.FirewallAddress6SubnetSegment, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.Type; tmp != nil {
				v["type"] = *tmp
			}

			if tmp := cfg.Value; tmp != nil {
				v["value"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenFirewallAddress6Tagging(d *schema.ResourceData, v *[]models.FirewallAddress6Tagging, prefix string, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for i, cfg := range *v {
			_ = i
			v := make(map[string]interface{})
			if tmp := cfg.Category; tmp != nil {
				v["category"] = *tmp
			}

			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.Tags; tmp != nil {
				v["tags"] = flattenFirewallAddress6TaggingTags(d, tmp, fmt.Sprintf("%s.%d.%s", prefix, i, "tags"), sort)
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenFirewallAddress6TaggingTags(d *schema.ResourceData, v *[]models.FirewallAddress6TaggingTags, prefix string, sort bool) interface{} {
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

func refreshObjectFirewallAddress6(d *schema.ResourceData, o *models.FirewallAddress6, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.CacheTtl != nil {
		v := *o.CacheTtl

		if err = d.Set("cache_ttl", v); err != nil {
			return diag.Errorf("error reading cache_ttl: %v", err)
		}
	}

	if o.Color != nil {
		v := *o.Color

		if err = d.Set("color", v); err != nil {
			return diag.Errorf("error reading color: %v", err)
		}
	}

	if o.Comment != nil {
		v := *o.Comment

		if err = d.Set("comment", v); err != nil {
			return diag.Errorf("error reading comment: %v", err)
		}
	}

	if o.Country != nil {
		v := *o.Country

		if err = d.Set("country", v); err != nil {
			return diag.Errorf("error reading country: %v", err)
		}
	}

	if o.EndIp != nil {
		v := *o.EndIp

		if err = d.Set("end_ip", v); err != nil {
			return diag.Errorf("error reading end_ip: %v", err)
		}
	}

	if o.EndMac != nil {
		v := *o.EndMac

		if err = d.Set("end_mac", v); err != nil {
			return diag.Errorf("error reading end_mac: %v", err)
		}
	}

	if o.FabricObject != nil {
		v := *o.FabricObject

		if err = d.Set("fabric_object", v); err != nil {
			return diag.Errorf("error reading fabric_object: %v", err)
		}
	}

	if o.Fqdn != nil {
		v := *o.Fqdn

		if err = d.Set("fqdn", v); err != nil {
			return diag.Errorf("error reading fqdn: %v", err)
		}
	}

	if o.Host != nil {
		v := *o.Host

		if err = d.Set("host", v); err != nil {
			return diag.Errorf("error reading host: %v", err)
		}
	}

	if o.HostType != nil {
		v := *o.HostType

		if err = d.Set("host_type", v); err != nil {
			return diag.Errorf("error reading host_type: %v", err)
		}
	}

	if o.Ip6 != nil {
		v := *o.Ip6

		if err = d.Set("ip6", v); err != nil {
			return diag.Errorf("error reading ip6: %v", err)
		}
	}

	if o.List != nil {
		if err = d.Set("list", flattenFirewallAddress6List(d, o.List, "list", sort)); err != nil {
			return diag.Errorf("error reading list: %v", err)
		}
	}

	if o.Macaddr != nil {
		if err = d.Set("macaddr", flattenFirewallAddress6Macaddr(d, o.Macaddr, "macaddr", sort)); err != nil {
			return diag.Errorf("error reading macaddr: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.ObjId != nil {
		v := *o.ObjId

		if err = d.Set("obj_id", v); err != nil {
			return diag.Errorf("error reading obj_id: %v", err)
		}
	}

	if o.Sdn != nil {
		v := *o.Sdn

		if err = d.Set("sdn", v); err != nil {
			return diag.Errorf("error reading sdn: %v", err)
		}
	}

	if o.StartIp != nil {
		v := *o.StartIp

		if err = d.Set("start_ip", v); err != nil {
			return diag.Errorf("error reading start_ip: %v", err)
		}
	}

	if o.StartMac != nil {
		v := *o.StartMac

		if err = d.Set("start_mac", v); err != nil {
			return diag.Errorf("error reading start_mac: %v", err)
		}
	}

	if o.SubnetSegment != nil {
		if err = d.Set("subnet_segment", flattenFirewallAddress6SubnetSegment(d, o.SubnetSegment, "subnet_segment", sort)); err != nil {
			return diag.Errorf("error reading subnet_segment: %v", err)
		}
	}

	if o.Tagging != nil {
		if err = d.Set("tagging", flattenFirewallAddress6Tagging(d, o.Tagging, "tagging", sort)); err != nil {
			return diag.Errorf("error reading tagging: %v", err)
		}
	}

	if o.Template != nil {
		v := *o.Template

		if err = d.Set("template", v); err != nil {
			return diag.Errorf("error reading template: %v", err)
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

	if o.Visibility != nil {
		v := *o.Visibility

		if err = d.Set("visibility", v); err != nil {
			return diag.Errorf("error reading visibility: %v", err)
		}
	}

	return nil
}

func expandFirewallAddress6List(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallAddress6List, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallAddress6List

	for i := range l {
		tmp := models.FirewallAddress6List{}
		var pre_append string

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

func expandFirewallAddress6Macaddr(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallAddress6Macaddr, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallAddress6Macaddr

	for i := range l {
		tmp := models.FirewallAddress6Macaddr{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.macaddr", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Macaddr = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandFirewallAddress6SubnetSegment(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallAddress6SubnetSegment, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallAddress6SubnetSegment

	for i := range l {
		tmp := models.FirewallAddress6SubnetSegment{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.type", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Type = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.value", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Value = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandFirewallAddress6Tagging(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallAddress6Tagging, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallAddress6Tagging

	for i := range l {
		tmp := models.FirewallAddress6Tagging{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.category", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Category = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.name", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Name = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.tags", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			v2, _ := expandFirewallAddress6TaggingTags(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.FirewallAddress6TaggingTags
			// 	}
			tmp.Tags = v2

		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandFirewallAddress6TaggingTags(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallAddress6TaggingTags, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallAddress6TaggingTags

	for i := range l {
		tmp := models.FirewallAddress6TaggingTags{}
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

func getObjectFirewallAddress6(d *schema.ResourceData, sv string) (*models.FirewallAddress6, diag.Diagnostics) {
	obj := models.FirewallAddress6{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("cache_ttl"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("cache_ttl", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.CacheTtl = &tmp
		}
	}
	if v1, ok := d.GetOk("color"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("color", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Color = &tmp
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
	if v1, ok := d.GetOk("country"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.0", "") {
				e := utils.AttributeVersionWarning("country", sv)
				diags = append(diags, e)
			}
			obj.Country = &v2
		}
	}
	if v1, ok := d.GetOk("end_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("end_ip", sv)
				diags = append(diags, e)
			}
			obj.EndIp = &v2
		}
	}
	if v1, ok := d.GetOk("end_mac"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.0") {
				e := utils.AttributeVersionWarning("end_mac", sv)
				diags = append(diags, e)
			}
			obj.EndMac = &v2
		}
	}
	if v1, ok := d.GetOk("fabric_object"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.5", "") {
				e := utils.AttributeVersionWarning("fabric_object", sv)
				diags = append(diags, e)
			}
			obj.FabricObject = &v2
		}
	}
	if v1, ok := d.GetOk("fqdn"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("fqdn", sv)
				diags = append(diags, e)
			}
			obj.Fqdn = &v2
		}
	}
	if v1, ok := d.GetOk("host"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("host", sv)
				diags = append(diags, e)
			}
			obj.Host = &v2
		}
	}
	if v1, ok := d.GetOk("host_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("host_type", sv)
				diags = append(diags, e)
			}
			obj.HostType = &v2
		}
	}
	if v1, ok := d.GetOk("ip6"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ip6", sv)
				diags = append(diags, e)
			}
			obj.Ip6 = &v2
		}
	}
	if v, ok := d.GetOk("list"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("list", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallAddress6List(d, v, "list", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.List = t
		}
	} else if d.HasChange("list") {
		old, new := d.GetChange("list")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.List = &[]models.FirewallAddress6List{}
		}
	}
	if v, ok := d.GetOk("macaddr"); ok {
		if !utils.CheckVer(sv, "v7.0.0", "") {
			e := utils.AttributeVersionWarning("macaddr", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallAddress6Macaddr(d, v, "macaddr", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Macaddr = t
		}
	} else if d.HasChange("macaddr") {
		old, new := d.GetChange("macaddr")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Macaddr = &[]models.FirewallAddress6Macaddr{}
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
	if v1, ok := d.GetOk("obj_id"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("obj_id", sv)
				diags = append(diags, e)
			}
			obj.ObjId = &v2
		}
	}
	if v1, ok := d.GetOk("sdn"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("sdn", sv)
				diags = append(diags, e)
			}
			obj.Sdn = &v2
		}
	}
	if v1, ok := d.GetOk("start_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("start_ip", sv)
				diags = append(diags, e)
			}
			obj.StartIp = &v2
		}
	}
	if v1, ok := d.GetOk("start_mac"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.0") {
				e := utils.AttributeVersionWarning("start_mac", sv)
				diags = append(diags, e)
			}
			obj.StartMac = &v2
		}
	}
	if v, ok := d.GetOk("subnet_segment"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("subnet_segment", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallAddress6SubnetSegment(d, v, "subnet_segment", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.SubnetSegment = t
		}
	} else if d.HasChange("subnet_segment") {
		old, new := d.GetChange("subnet_segment")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.SubnetSegment = &[]models.FirewallAddress6SubnetSegment{}
		}
	}
	if v, ok := d.GetOk("tagging"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("tagging", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallAddress6Tagging(d, v, "tagging", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Tagging = t
		}
	} else if d.HasChange("tagging") {
		old, new := d.GetChange("tagging")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Tagging = &[]models.FirewallAddress6Tagging{}
		}
	}
	if v1, ok := d.GetOk("template"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("template", sv)
				diags = append(diags, e)
			}
			obj.Template = &v2
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
	if v1, ok := d.GetOk("uuid"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("uuid", sv)
				diags = append(diags, e)
			}
			obj.Uuid = &v2
		}
	}
	if v1, ok := d.GetOk("visibility"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.0") {
				e := utils.AttributeVersionWarning("visibility", sv)
				diags = append(diags, e)
			}
			obj.Visibility = &v2
		}
	}
	return &obj, diags
}
