// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: OSPF6 interface configuration.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/suppressors"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func resourceRouterOspf6Ospf6Interface() *schema.Resource {
	return &schema.Resource{
		Description: "OSPF6 interface configuration.",

		CreateContext: resourceRouterOspf6Ospf6InterfaceCreate,
		ReadContext:   resourceRouterOspf6Ospf6InterfaceRead,
		UpdateContext: resourceRouterOspf6Ospf6InterfaceUpdate,
		DeleteContext: resourceRouterOspf6Ospf6InterfaceDelete,

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
			"area_id": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "A.B.C.D, in IPv4 address format.",
				Optional:    true,
				Computed:    true,
			},
			"authentication": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"none", "ah", "esp", "area"}, false),

				Description: "Authentication mode.",
				Optional:    true,
				Computed:    true,
			},
			"bfd": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"global", "enable", "disable"}, false),

				Description: "Enable/disable Bidirectional Forwarding Detection (BFD).",
				Optional:    true,
				Computed:    true,
			},
			"cost": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "Cost of the interface, value range from 0 to 65535, 0 means auto-cost.",
				Optional:    true,
				Computed:    true,
			},
			"dead_interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),

				Description: "Dead interval.",
				Optional:    true,
				Computed:    true,
			},
			"hello_interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),

				Description: "Hello interval.",
				Optional:    true,
				Computed:    true,
			},
			"interface": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "Configuration interface name.",
				Optional:    true,
				Computed:    true,
			},
			"ipsec_auth_alg": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"md5", "sha1", "sha256", "sha384", "sha512"}, false),

				Description: "Authentication algorithm.",
				Optional:    true,
				Computed:    true,
			},
			"ipsec_enc_alg": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"null", "des", "3des", "aes128", "aes192", "aes256"}, false),

				Description: "Encryption algorithm.",
				Optional:    true,
				Computed:    true,
			},
			"ipsec_keys": {
				Type:        schema.TypeList,
				Description: "IPsec authentication and encryption keys.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auth_key": {
							Type: schema.TypeString,

							Description: "Authentication key.",
							Optional:    true,
							Computed:    true,
							Sensitive:   true,
						},
						"enc_key": {
							Type: schema.TypeString,

							Description: "Encryption key.",
							Optional:    true,
							Computed:    true,
							Sensitive:   true,
						},
						"spi": {
							Type: schema.TypeInt,

							Description: "Security Parameters Index.",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"key_rollover_interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(300, 216000),

				Description: "Key roll-over interval.",
				Optional:    true,
				Computed:    true,
			},
			"mtu": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(576, 65535),

				Description: "MTU for OSPFv3 packets.",
				Optional:    true,
				Computed:    true,
			},
			"mtu_ignore": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable ignoring MTU field in DBD packets.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Interface entry name.",
				ForceNew:    true,
				Required:    true,
			},
			"neighbor": {
				Type:        schema.TypeList,
				Description: "OSPFv3 neighbors are used when OSPFv3 runs on non-broadcast media",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cost": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),

							Description: "Cost of the interface, value range from 0 to 65535, 0 means auto-cost.",
							Optional:    true,
							Computed:    true,
						},
						"ip6": {
							Type:             schema.TypeString,
							ValidateFunc:     validation.IsIPv6Address,
							DiffSuppressFunc: suppressors.DiffIPEqual,
							Description:      "IPv6 link local address of the neighbor.",
							Optional:         true,
							Computed:         true,
						},
						"poll_interval": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),

							Description: "Poll interval time in seconds.",
							Optional:    true,
							Computed:    true,
						},
						"priority": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),

							Description: "priority",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"network_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"broadcast", "point-to-point", "non-broadcast", "point-to-multipoint", "point-to-multipoint-non-broadcast"}, false),

				Description: "Network type.",
				Optional:    true,
				Computed:    true,
			},
			"priority": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),

				Description: "priority",
				Optional:    true,
				Computed:    true,
			},
			"retransmit_interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),

				Description: "Retransmit interval.",
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable OSPF6 routing on this interface.",
				Optional:    true,
				Computed:    true,
			},
			"transmit_delay": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),

				Description: "Transmit delay.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceRouterOspf6Ospf6InterfaceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating RouterOspf6Ospf6Interface resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectRouterOspf6Ospf6Interface(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateRouterOspf6Ospf6Interface(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("RouterOspf6Ospf6Interface")
	}

	return resourceRouterOspf6Ospf6InterfaceRead(ctx, d, meta)
}

func resourceRouterOspf6Ospf6InterfaceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectRouterOspf6Ospf6Interface(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateRouterOspf6Ospf6Interface(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating RouterOspf6Ospf6Interface resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("RouterOspf6Ospf6Interface")
	}

	return resourceRouterOspf6Ospf6InterfaceRead(ctx, d, meta)
}

func resourceRouterOspf6Ospf6InterfaceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteRouterOspf6Ospf6Interface(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting RouterOspf6Ospf6Interface resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterOspf6Ospf6InterfaceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadRouterOspf6Ospf6Interface(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading RouterOspf6Ospf6Interface resource: %v", err)
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

	diags := refreshObjectRouterOspf6Ospf6Interface(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectRouterOspf6Ospf6Interface(d *schema.ResourceData, o *models.RouterOspf6Ospf6Interface, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AreaId != nil {
		v := *o.AreaId

		if err = d.Set("area_id", v); err != nil {
			return diag.Errorf("error reading area_id: %v", err)
		}
	}

	if o.Authentication != nil {
		v := *o.Authentication

		if err = d.Set("authentication", v); err != nil {
			return diag.Errorf("error reading authentication: %v", err)
		}
	}

	if o.Bfd != nil {
		v := *o.Bfd

		if err = d.Set("bfd", v); err != nil {
			return diag.Errorf("error reading bfd: %v", err)
		}
	}

	if o.Cost != nil {
		v := *o.Cost

		if err = d.Set("cost", v); err != nil {
			return diag.Errorf("error reading cost: %v", err)
		}
	}

	if o.DeadInterval != nil {
		v := *o.DeadInterval

		if err = d.Set("dead_interval", v); err != nil {
			return diag.Errorf("error reading dead_interval: %v", err)
		}
	}

	if o.HelloInterval != nil {
		v := *o.HelloInterval

		if err = d.Set("hello_interval", v); err != nil {
			return diag.Errorf("error reading hello_interval: %v", err)
		}
	}

	if o.Interface != nil {
		v := *o.Interface

		if err = d.Set("interface", v); err != nil {
			return diag.Errorf("error reading interface: %v", err)
		}
	}

	if o.IpsecAuthAlg != nil {
		v := *o.IpsecAuthAlg

		if err = d.Set("ipsec_auth_alg", v); err != nil {
			return diag.Errorf("error reading ipsec_auth_alg: %v", err)
		}
	}

	if o.IpsecEncAlg != nil {
		v := *o.IpsecEncAlg

		if err = d.Set("ipsec_enc_alg", v); err != nil {
			return diag.Errorf("error reading ipsec_enc_alg: %v", err)
		}
	}

	if o.IpsecKeys != nil {
		if err = d.Set("ipsec_keys", flattenRouterOspf6Ospf6InterfaceIpsecKeys(o.IpsecKeys, sort)); err != nil {
			return diag.Errorf("error reading ipsec_keys: %v", err)
		}
	}

	if o.KeyRolloverInterval != nil {
		v := *o.KeyRolloverInterval

		if err = d.Set("key_rollover_interval", v); err != nil {
			return diag.Errorf("error reading key_rollover_interval: %v", err)
		}
	}

	if o.Mtu != nil {
		v := *o.Mtu

		if err = d.Set("mtu", v); err != nil {
			return diag.Errorf("error reading mtu: %v", err)
		}
	}

	if o.MtuIgnore != nil {
		v := *o.MtuIgnore

		if err = d.Set("mtu_ignore", v); err != nil {
			return diag.Errorf("error reading mtu_ignore: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.Neighbor != nil {
		if err = d.Set("neighbor", flattenRouterOspf6Ospf6InterfaceNeighbor(o.Neighbor, sort)); err != nil {
			return diag.Errorf("error reading neighbor: %v", err)
		}
	}

	if o.NetworkType != nil {
		v := *o.NetworkType

		if err = d.Set("network_type", v); err != nil {
			return diag.Errorf("error reading network_type: %v", err)
		}
	}

	if o.Priority != nil {
		v := *o.Priority

		if err = d.Set("priority", v); err != nil {
			return diag.Errorf("error reading priority: %v", err)
		}
	}

	if o.RetransmitInterval != nil {
		v := *o.RetransmitInterval

		if err = d.Set("retransmit_interval", v); err != nil {
			return diag.Errorf("error reading retransmit_interval: %v", err)
		}
	}

	if o.Status != nil {
		v := *o.Status

		if err = d.Set("status", v); err != nil {
			return diag.Errorf("error reading status: %v", err)
		}
	}

	if o.TransmitDelay != nil {
		v := *o.TransmitDelay

		if err = d.Set("transmit_delay", v); err != nil {
			return diag.Errorf("error reading transmit_delay: %v", err)
		}
	}

	return nil
}

func getObjectRouterOspf6Ospf6Interface(d *schema.ResourceData, sv string) (*models.RouterOspf6Ospf6Interface, diag.Diagnostics) {
	obj := models.RouterOspf6Ospf6Interface{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("area_id"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("area_id", sv)
				diags = append(diags, e)
			}
			obj.AreaId = &v2
		}
	}
	if v1, ok := d.GetOk("authentication"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("authentication", sv)
				diags = append(diags, e)
			}
			obj.Authentication = &v2
		}
	}
	if v1, ok := d.GetOk("bfd"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("bfd", sv)
				diags = append(diags, e)
			}
			obj.Bfd = &v2
		}
	}
	if v1, ok := d.GetOk("cost"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("cost", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Cost = &tmp
		}
	}
	if v1, ok := d.GetOk("dead_interval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dead_interval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.DeadInterval = &tmp
		}
	}
	if v1, ok := d.GetOk("hello_interval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("hello_interval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.HelloInterval = &tmp
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
	if v1, ok := d.GetOk("ipsec_auth_alg"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ipsec_auth_alg", sv)
				diags = append(diags, e)
			}
			obj.IpsecAuthAlg = &v2
		}
	}
	if v1, ok := d.GetOk("ipsec_enc_alg"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ipsec_enc_alg", sv)
				diags = append(diags, e)
			}
			obj.IpsecEncAlg = &v2
		}
	}
	if v, ok := d.GetOk("ipsec_keys"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("ipsec_keys", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterOspf6Ospf6InterfaceIpsecKeys(d, v, "ipsec_keys", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.IpsecKeys = t
		}
	} else if d.HasChange("ipsec_keys") {
		old, new := d.GetChange("ipsec_keys")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.IpsecKeys = &[]models.RouterOspf6Ospf6InterfaceIpsecKeys{}
		}
	}
	if v1, ok := d.GetOk("key_rollover_interval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("key_rollover_interval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.KeyRolloverInterval = &tmp
		}
	}
	if v1, ok := d.GetOk("mtu"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("mtu", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Mtu = &tmp
		}
	}
	if v1, ok := d.GetOk("mtu_ignore"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("mtu_ignore", sv)
				diags = append(diags, e)
			}
			obj.MtuIgnore = &v2
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
	if v, ok := d.GetOk("neighbor"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("neighbor", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterOspf6Ospf6InterfaceNeighbor(d, v, "neighbor", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Neighbor = t
		}
	} else if d.HasChange("neighbor") {
		old, new := d.GetChange("neighbor")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Neighbor = &[]models.RouterOspf6Ospf6InterfaceNeighbor{}
		}
	}
	if v1, ok := d.GetOk("network_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("network_type", sv)
				diags = append(diags, e)
			}
			obj.NetworkType = &v2
		}
	}
	if v1, ok := d.GetOk("priority"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("priority", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Priority = &tmp
		}
	}
	if v1, ok := d.GetOk("retransmit_interval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("retransmit_interval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.RetransmitInterval = &tmp
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
	if v1, ok := d.GetOk("transmit_delay"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("transmit_delay", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.TransmitDelay = &tmp
		}
	}
	return &obj, diags
}
