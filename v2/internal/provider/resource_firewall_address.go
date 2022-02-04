// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure IPv4 addresses.

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

func resourceFirewallAddress() *schema.Resource {
	return &schema.Resource{
		Description: "Configure IPv4 addresses.",

		CreateContext: resourceFirewallAddressCreate,
		ReadContext:   resourceFirewallAddressRead,
		UpdateContext: resourceFirewallAddressUpdate,
		DeleteContext: resourceFirewallAddressDelete,

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
			"allow_routing": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable use of this address in the static route configuration.",
				Optional:    true,
				Computed:    true,
			},
			"associated_interface": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Network interface associated with address.",
				Optional:    true,
				Computed:    true,
			},
			"cache_ttl": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 86400),

				Description: "Defines the minimal TTL of individual IP addresses in FQDN cache measured in seconds.",
				Optional:    true,
				Computed:    true,
			},
			"clearpass_spt": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"unknown", "healthy", "quarantine", "checkup", "transient", "infected"}, false),

				Description: "SPT (System Posture Token) value.",
				Optional:    true,
				Computed:    true,
			},
			"color": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 32),

				Description: "Color of icon on the GUI.",
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

				Description: "IP addresses associated to a specific country.",
				Optional:    true,
				Computed:    true,
			},
			"end_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "Final IP address (inclusive) in the range for the address.",
				Optional:    true,
				Computed:    true,
			},
			"end_mac": {
				Type: schema.TypeString,

				Description: "Last MAC address in the range.",
				Optional:    true,
				Computed:    true,
			},
			"epg_name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Endpoint group name.",
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
			"filter": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 2047),

				Description: "Match criteria filter.",
				Optional:    true,
				Computed:    true,
			},
			"fqdn": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Fully Qualified Domain Name address.",
				Optional:    true,
				Computed:    true,
			},
			"fsso_group": {
				Type:        schema.TypeList,
				Description: "FSSO group(s).",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 511),

							Description: "FSSO group name.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"interface": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Name of interface whose IP address is to be used.",
				Optional:    true,
				Computed:    true,
			},
			"list": {
				Type:        schema.TypeList,
				Description: "IP address list.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),

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
			"node_ip_only": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable collection of node addresses only in Kubernetes.",
				Optional:    true,
				Computed:    true,
			},
			"obj_id": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Object ID for NSX.",
				Optional:    true,
				Computed:    true,
			},
			"obj_tag": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Tag of dynamic address object.",
				Optional:    true,
				Computed:    true,
			},
			"obj_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"ip", "mac"}, false),

				Description: "Object type.",
				Optional:    true,
				Computed:    true,
			},
			"organization": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Organization domain name (Syntax: organization/domain).",
				Optional:    true,
				Computed:    true,
			},
			"policy_group": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "Policy group name.",
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
			"sdn_addr_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"private", "public", "all"}, false),

				Description: "Type of addresses to collect.",
				Optional:    true,
				Computed:    true,
			},
			"sdn_tag": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "SDN Tag.",
				Optional:    true,
				Computed:    true,
			},
			"start_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "First IP address (inclusive) in the range for the address.",
				Optional:    true,
				Computed:    true,
			},
			"start_mac": {
				Type: schema.TypeString,

				Description: "First MAC address in the range.",
				Optional:    true,
				Computed:    true,
			},
			"sub_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"sdn", "clearpass-spt", "fsso", "ems-tag", "fortivoice-tag", "fortinac-tag", "swc-tag"}, false),

				Description: "Sub-type of address.",
				Optional:    true,
				Computed:    true,
			},
			"subnet": {
				Type:         schema.TypeString,
				ValidateFunc: validators.FortiValidateIPv4ClassnetAny,

				Description: "IP address and subnet mask of address.",
				Optional:    true,
				Computed:    true,
			},
			"subnet_name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Subnet name.",
				Optional:    true,
				Computed:    true,
			},
			"tag_detection_level": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "Tag detection level of dynamic address object.",
				Optional:    true,
				Computed:    true,
			},
			"tag_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Tag type of dynamic address object.",
				Optional:    true,
				Computed:    true,
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
			"tenant": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Tenant.",
				Optional:    true,
				Computed:    true,
			},
			"type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"ipmask", "iprange", "fqdn", "geography", "wildcard", "dynamic", "interface-subnet", "mac"}, false),

				Description: "Type of address.",
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

				Description: "Enable/disable address visibility in the GUI.",
				Optional:    true,
				Computed:    true,
			},
			"wildcard": {
				Type:         schema.TypeString,
				ValidateFunc: validators.FortiValidateIPv4ClassnetAny,

				Description: "IP address and wildcard netmask.",
				Optional:    true,
				Computed:    true,
			},
			"wildcard_fqdn": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Fully Qualified Domain Name with wildcard characters.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceFirewallAddressCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating FirewallAddress resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectFirewallAddress(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateFirewallAddress(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallAddress")
	}

	return resourceFirewallAddressRead(ctx, d, meta)
}

func resourceFirewallAddressUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectFirewallAddress(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateFirewallAddress(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating FirewallAddress resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallAddress")
	}

	return resourceFirewallAddressRead(ctx, d, meta)
}

func resourceFirewallAddressDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteFirewallAddress(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting FirewallAddress resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallAddressRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallAddress(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallAddress resource: %v", err)
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

	diags := refreshObjectFirewallAddress(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenFirewallAddressFssoGroup(v *[]models.FirewallAddressFssoGroup, sort bool) interface{} {
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

func flattenFirewallAddressList(v *[]models.FirewallAddressList, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
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

func flattenFirewallAddressMacaddr(v *[]models.FirewallAddressMacaddr, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
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

func flattenFirewallAddressTagging(v *[]models.FirewallAddressTagging, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Category; tmp != nil {
				v["category"] = *tmp
			}

			if tmp := cfg.Name; tmp != nil {
				v["name"] = *tmp
			}

			if tmp := cfg.Tags; tmp != nil {
				v["tags"] = flattenFirewallAddressTaggingTags(tmp, sort)
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "name")
	}

	return flat
}

func flattenFirewallAddressTaggingTags(v *[]models.FirewallAddressTaggingTags, sort bool) interface{} {
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

func refreshObjectFirewallAddress(d *schema.ResourceData, o *models.FirewallAddress, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AllowRouting != nil {
		v := *o.AllowRouting

		if err = d.Set("allow_routing", v); err != nil {
			return diag.Errorf("error reading allow_routing: %v", err)
		}
	}

	if o.AssociatedInterface != nil {
		v := *o.AssociatedInterface

		if err = d.Set("associated_interface", v); err != nil {
			return diag.Errorf("error reading associated_interface: %v", err)
		}
	}

	if o.CacheTtl != nil {
		v := *o.CacheTtl

		if err = d.Set("cache_ttl", v); err != nil {
			return diag.Errorf("error reading cache_ttl: %v", err)
		}
	}

	if o.ClearpassSpt != nil {
		v := *o.ClearpassSpt

		if err = d.Set("clearpass_spt", v); err != nil {
			return diag.Errorf("error reading clearpass_spt: %v", err)
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

	if o.EpgName != nil {
		v := *o.EpgName

		if err = d.Set("epg_name", v); err != nil {
			return diag.Errorf("error reading epg_name: %v", err)
		}
	}

	if o.FabricObject != nil {
		v := *o.FabricObject

		if err = d.Set("fabric_object", v); err != nil {
			return diag.Errorf("error reading fabric_object: %v", err)
		}
	}

	if o.Filter != nil {
		v := *o.Filter

		if err = d.Set("filter", v); err != nil {
			return diag.Errorf("error reading filter: %v", err)
		}
	}

	if o.Fqdn != nil {
		v := *o.Fqdn

		if err = d.Set("fqdn", v); err != nil {
			return diag.Errorf("error reading fqdn: %v", err)
		}
	}

	if o.FssoGroup != nil {
		if err = d.Set("fsso_group", flattenFirewallAddressFssoGroup(o.FssoGroup, sort)); err != nil {
			return diag.Errorf("error reading fsso_group: %v", err)
		}
	}

	if o.Interface != nil {
		v := *o.Interface

		if err = d.Set("interface", v); err != nil {
			return diag.Errorf("error reading interface: %v", err)
		}
	}

	if o.List != nil {
		if err = d.Set("list", flattenFirewallAddressList(o.List, sort)); err != nil {
			return diag.Errorf("error reading list: %v", err)
		}
	}

	if o.Macaddr != nil {
		if err = d.Set("macaddr", flattenFirewallAddressMacaddr(o.Macaddr, sort)); err != nil {
			return diag.Errorf("error reading macaddr: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.NodeIpOnly != nil {
		v := *o.NodeIpOnly

		if err = d.Set("node_ip_only", v); err != nil {
			return diag.Errorf("error reading node_ip_only: %v", err)
		}
	}

	if o.ObjId != nil {
		v := *o.ObjId

		if err = d.Set("obj_id", v); err != nil {
			return diag.Errorf("error reading obj_id: %v", err)
		}
	}

	if o.ObjTag != nil {
		v := *o.ObjTag

		if err = d.Set("obj_tag", v); err != nil {
			return diag.Errorf("error reading obj_tag: %v", err)
		}
	}

	if o.ObjType != nil {
		v := *o.ObjType

		if err = d.Set("obj_type", v); err != nil {
			return diag.Errorf("error reading obj_type: %v", err)
		}
	}

	if o.Organization != nil {
		v := *o.Organization

		if err = d.Set("organization", v); err != nil {
			return diag.Errorf("error reading organization: %v", err)
		}
	}

	if o.PolicyGroup != nil {
		v := *o.PolicyGroup

		if err = d.Set("policy_group", v); err != nil {
			return diag.Errorf("error reading policy_group: %v", err)
		}
	}

	if o.Sdn != nil {
		v := *o.Sdn

		if err = d.Set("sdn", v); err != nil {
			return diag.Errorf("error reading sdn: %v", err)
		}
	}

	if o.SdnAddrType != nil {
		v := *o.SdnAddrType

		if err = d.Set("sdn_addr_type", v); err != nil {
			return diag.Errorf("error reading sdn_addr_type: %v", err)
		}
	}

	if o.SdnTag != nil {
		v := *o.SdnTag

		if err = d.Set("sdn_tag", v); err != nil {
			return diag.Errorf("error reading sdn_tag: %v", err)
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

	if o.SubType != nil {
		v := *o.SubType

		if err = d.Set("sub_type", v); err != nil {
			return diag.Errorf("error reading sub_type: %v", err)
		}
	}

	if o.Subnet != nil {
		v := *o.Subnet
		if current, ok := d.GetOk("subnet"); ok {
			if s, ok := current.(string); ok {
				v = utils.ValidateConvIPMask2CIDR(s, v)
			}
		}

		if err = d.Set("subnet", v); err != nil {
			return diag.Errorf("error reading subnet: %v", err)
		}
	}

	if o.SubnetName != nil {
		v := *o.SubnetName

		if err = d.Set("subnet_name", v); err != nil {
			return diag.Errorf("error reading subnet_name: %v", err)
		}
	}

	if o.TagDetectionLevel != nil {
		v := *o.TagDetectionLevel

		if err = d.Set("tag_detection_level", v); err != nil {
			return diag.Errorf("error reading tag_detection_level: %v", err)
		}
	}

	if o.TagType != nil {
		v := *o.TagType

		if err = d.Set("tag_type", v); err != nil {
			return diag.Errorf("error reading tag_type: %v", err)
		}
	}

	if o.Tagging != nil {
		if err = d.Set("tagging", flattenFirewallAddressTagging(o.Tagging, sort)); err != nil {
			return diag.Errorf("error reading tagging: %v", err)
		}
	}

	if o.Tenant != nil {
		v := *o.Tenant

		if err = d.Set("tenant", v); err != nil {
			return diag.Errorf("error reading tenant: %v", err)
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

	if o.Wildcard != nil {
		v := *o.Wildcard
		if current, ok := d.GetOk("wildcard"); ok {
			if s, ok := current.(string); ok {
				v = utils.ValidateConvIPMask2CIDR(s, v)
			}
		}

		if err = d.Set("wildcard", v); err != nil {
			return diag.Errorf("error reading wildcard: %v", err)
		}
	}

	if o.WildcardFqdn != nil {
		v := *o.WildcardFqdn

		if err = d.Set("wildcard_fqdn", v); err != nil {
			return diag.Errorf("error reading wildcard_fqdn: %v", err)
		}
	}

	return nil
}

func expandFirewallAddressFssoGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallAddressFssoGroup, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallAddressFssoGroup

	for i := range l {
		tmp := models.FirewallAddressFssoGroup{}
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

func expandFirewallAddressList(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallAddressList, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallAddressList

	for i := range l {
		tmp := models.FirewallAddressList{}
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

func expandFirewallAddressMacaddr(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallAddressMacaddr, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallAddressMacaddr

	for i := range l {
		tmp := models.FirewallAddressMacaddr{}
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

func expandFirewallAddressTagging(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallAddressTagging, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallAddressTagging

	for i := range l {
		tmp := models.FirewallAddressTagging{}
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
			v2, _ := expandFirewallAddressTaggingTags(d, v1, pre_append, sv)
			// if err != nil {
			// 	v2 := &[]models.FirewallAddressTaggingTags
			// 	}
			tmp.Tags = v2

		}

		result = append(result, tmp)
	}
	return &result, nil
}

func expandFirewallAddressTaggingTags(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.FirewallAddressTaggingTags, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.FirewallAddressTaggingTags

	for i := range l {
		tmp := models.FirewallAddressTaggingTags{}
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

func getObjectFirewallAddress(d *schema.ResourceData, sv string) (*models.FirewallAddress, diag.Diagnostics) {
	obj := models.FirewallAddress{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("allow_routing"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("allow_routing", sv)
				diags = append(diags, e)
			}
			obj.AllowRouting = &v2
		}
	}
	if v1, ok := d.GetOk("associated_interface"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("associated_interface", sv)
				diags = append(diags, e)
			}
			obj.AssociatedInterface = &v2
		}
	}
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
	if v1, ok := d.GetOk("clearpass_spt"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("clearpass_spt", sv)
				diags = append(diags, e)
			}
			obj.ClearpassSpt = &v2
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
			if !utils.CheckVer(sv, "", "") {
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
			if !utils.CheckVer(sv, "", "v7.0.0") {
				e := utils.AttributeVersionWarning("end_mac", sv)
				diags = append(diags, e)
			}
			obj.EndMac = &v2
		}
	}
	if v1, ok := d.GetOk("epg_name"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("epg_name", sv)
				diags = append(diags, e)
			}
			obj.EpgName = &v2
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
	if v1, ok := d.GetOk("filter"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("filter", sv)
				diags = append(diags, e)
			}
			obj.Filter = &v2
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
	if v, ok := d.GetOk("fsso_group"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("fsso_group", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallAddressFssoGroup(d, v, "fsso_group", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.FssoGroup = t
		}
	} else if d.HasChange("fsso_group") {
		old, new := d.GetChange("fsso_group")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.FssoGroup = &[]models.FirewallAddressFssoGroup{}
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
	if v, ok := d.GetOk("list"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("list", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallAddressList(d, v, "list", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.List = t
		}
	} else if d.HasChange("list") {
		old, new := d.GetChange("list")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.List = &[]models.FirewallAddressList{}
		}
	}
	if v, ok := d.GetOk("macaddr"); ok {
		if !utils.CheckVer(sv, "v7.0.0", "") {
			e := utils.AttributeVersionWarning("macaddr", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallAddressMacaddr(d, v, "macaddr", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Macaddr = t
		}
	} else if d.HasChange("macaddr") {
		old, new := d.GetChange("macaddr")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Macaddr = &[]models.FirewallAddressMacaddr{}
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
	if v1, ok := d.GetOk("node_ip_only"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("node_ip_only", sv)
				diags = append(diags, e)
			}
			obj.NodeIpOnly = &v2
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
	if v1, ok := d.GetOk("obj_tag"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.2", "") {
				e := utils.AttributeVersionWarning("obj_tag", sv)
				diags = append(diags, e)
			}
			obj.ObjTag = &v2
		}
	}
	if v1, ok := d.GetOk("obj_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.2", "") {
				e := utils.AttributeVersionWarning("obj_type", sv)
				diags = append(diags, e)
			}
			obj.ObjType = &v2
		}
	}
	if v1, ok := d.GetOk("organization"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("organization", sv)
				diags = append(diags, e)
			}
			obj.Organization = &v2
		}
	}
	if v1, ok := d.GetOk("policy_group"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("policy_group", sv)
				diags = append(diags, e)
			}
			obj.PolicyGroup = &v2
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
	if v1, ok := d.GetOk("sdn_addr_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("sdn_addr_type", sv)
				diags = append(diags, e)
			}
			obj.SdnAddrType = &v2
		}
	}
	if v1, ok := d.GetOk("sdn_tag"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("sdn_tag", sv)
				diags = append(diags, e)
			}
			obj.SdnTag = &v2
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
			if !utils.CheckVer(sv, "", "v7.0.0") {
				e := utils.AttributeVersionWarning("start_mac", sv)
				diags = append(diags, e)
			}
			obj.StartMac = &v2
		}
	}
	if v1, ok := d.GetOk("sub_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("sub_type", sv)
				diags = append(diags, e)
			}
			obj.SubType = &v2
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
	if v1, ok := d.GetOk("subnet_name"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("subnet_name", sv)
				diags = append(diags, e)
			}
			obj.SubnetName = &v2
		}
	}
	if v1, ok := d.GetOk("tag_detection_level"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.4", "") {
				e := utils.AttributeVersionWarning("tag_detection_level", sv)
				diags = append(diags, e)
			}
			obj.TagDetectionLevel = &v2
		}
	}
	if v1, ok := d.GetOk("tag_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.4", "") {
				e := utils.AttributeVersionWarning("tag_type", sv)
				diags = append(diags, e)
			}
			obj.TagType = &v2
		}
	}
	if v, ok := d.GetOk("tagging"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("tagging", sv)
			diags = append(diags, e)
		}
		t, err := expandFirewallAddressTagging(d, v, "tagging", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Tagging = t
		}
	} else if d.HasChange("tagging") {
		old, new := d.GetChange("tagging")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Tagging = &[]models.FirewallAddressTagging{}
		}
	}
	if v1, ok := d.GetOk("tenant"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("tenant", sv)
				diags = append(diags, e)
			}
			obj.Tenant = &v2
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
	if v1, ok := d.GetOk("wildcard"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("wildcard", sv)
				diags = append(diags, e)
			}
			obj.Wildcard = &v2
		}
	}
	if v1, ok := d.GetOk("wildcard_fqdn"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("wildcard_fqdn", sv)
				diags = append(diags, e)
			}
			obj.WildcardFqdn = &v2
		}
	}
	return &obj, diags
}
