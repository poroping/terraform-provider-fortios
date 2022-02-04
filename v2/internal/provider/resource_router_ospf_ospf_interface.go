// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: OSPF interface configuration.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func resourceRouterOspfOspfInterface() *schema.Resource {
	return &schema.Resource{
		Description: "OSPF interface configuration.",

		CreateContext: resourceRouterOspfOspfInterfaceCreate,
		ReadContext:   resourceRouterOspfOspfInterfaceRead,
		UpdateContext: resourceRouterOspfOspfInterfaceUpdate,
		DeleteContext: resourceRouterOspfOspfInterfaceDelete,

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
			"authentication": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"none", "text", "message-digest"}, false),

				Description: "Authentication type.",
				Optional:    true,
				Computed:    true,
			},
			"authentication_key": {
				Type: schema.TypeString,

				Description: "Authentication key.",
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
			},
			"bfd": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"global", "enable", "disable"}, false),

				Description: "Bidirectional Forwarding Detection (BFD).",
				Optional:    true,
				Computed:    true,
			},
			"comments": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Comment.",
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
			"database_filter_out": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable control of flooding out LSAs.",
				Optional:    true,
				Computed:    true,
			},
			"dead_interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "Dead interval.",
				Optional:    true,
				Computed:    true,
			},
			"hello_interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "Hello interval.",
				Optional:    true,
				Computed:    true,
			},
			"hello_multiplier": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(3, 10),

				Description: "Number of hello packets within dead interval.",
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
			"ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "IP address.",
				Optional:    true,
				Computed:    true,
			},
			"keychain": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Message-digest key-chain name.",
				Optional:    true,
				Computed:    true,
			},
			"md5_keychain": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Authentication MD5 key-chain name.",
				Optional:    true,
				Computed:    true,
			},
			"md5_keys": {
				Type:        schema.TypeList,
				Description: "MD5 key.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 255),

							Description: "Key ID (1 - 255).",
							Optional:    true,
							Computed:    true,
						},
						"key_string": {
							Type: schema.TypeString,

							Description: "Password for the key.",
							Optional:    true,
							Computed:    true,
							Sensitive:   true,
						},
					},
				},
			},
			"mtu": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(576, 65535),

				Description: "MTU for database description packets.",
				Optional:    true,
				Computed:    true,
			},
			"mtu_ignore": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable ignore MTU.",
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
			"network_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"broadcast", "non-broadcast", "point-to-point", "point-to-multipoint", "point-to-multipoint-non-broadcast"}, false),

				Description: "Network type.",
				Optional:    true,
				Computed:    true,
			},
			"prefix_length": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 32),

				Description: "Prefix length.",
				Optional:    true,
				Computed:    true,
			},
			"priority": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),

				Description: "Priority.",
				Optional:    true,
				Computed:    true,
			},
			"resync_timeout": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 3600),

				Description: "Graceful restart neighbor resynchronization timeout.",
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

				Description: "Enable/disable status.",
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

func resourceRouterOspfOspfInterfaceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating RouterOspfOspfInterface resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectRouterOspfOspfInterface(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateRouterOspfOspfInterface(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("RouterOspfOspfInterface")
	}

	return resourceRouterOspfOspfInterfaceRead(ctx, d, meta)
}

func resourceRouterOspfOspfInterfaceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectRouterOspfOspfInterface(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateRouterOspfOspfInterface(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating RouterOspfOspfInterface resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("RouterOspfOspfInterface")
	}

	return resourceRouterOspfOspfInterfaceRead(ctx, d, meta)
}

func resourceRouterOspfOspfInterfaceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteRouterOspfOspfInterface(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting RouterOspfOspfInterface resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterOspfOspfInterfaceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadRouterOspfOspfInterface(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading RouterOspfOspfInterface resource: %v", err)
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

	diags := refreshObjectRouterOspfOspfInterface(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectRouterOspfOspfInterface(d *schema.ResourceData, o *models.RouterOspfOspfInterface, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Authentication != nil {
		v := *o.Authentication

		if err = d.Set("authentication", v); err != nil {
			return diag.Errorf("error reading authentication: %v", err)
		}
	}

	if o.AuthenticationKey != nil {
		v := *o.AuthenticationKey

		if v == "ENC XXXX" {
		} else if err = d.Set("authentication_key", v); err != nil {
			return diag.Errorf("error reading authentication_key: %v", err)
		}
	}

	if o.Bfd != nil {
		v := *o.Bfd

		if err = d.Set("bfd", v); err != nil {
			return diag.Errorf("error reading bfd: %v", err)
		}
	}

	if o.Comments != nil {
		v := *o.Comments

		if err = d.Set("comments", v); err != nil {
			return diag.Errorf("error reading comments: %v", err)
		}
	}

	if o.Cost != nil {
		v := *o.Cost

		if err = d.Set("cost", v); err != nil {
			return diag.Errorf("error reading cost: %v", err)
		}
	}

	if o.DatabaseFilterOut != nil {
		v := *o.DatabaseFilterOut

		if err = d.Set("database_filter_out", v); err != nil {
			return diag.Errorf("error reading database_filter_out: %v", err)
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

	if o.HelloMultiplier != nil {
		v := *o.HelloMultiplier

		if err = d.Set("hello_multiplier", v); err != nil {
			return diag.Errorf("error reading hello_multiplier: %v", err)
		}
	}

	if o.Interface != nil {
		v := *o.Interface

		if err = d.Set("interface", v); err != nil {
			return diag.Errorf("error reading interface: %v", err)
		}
	}

	if o.Ip != nil {
		v := *o.Ip

		if err = d.Set("ip", v); err != nil {
			return diag.Errorf("error reading ip: %v", err)
		}
	}

	if o.Keychain != nil {
		v := *o.Keychain

		if err = d.Set("keychain", v); err != nil {
			return diag.Errorf("error reading keychain: %v", err)
		}
	}

	if o.Md5Keychain != nil {
		v := *o.Md5Keychain

		if err = d.Set("md5_keychain", v); err != nil {
			return diag.Errorf("error reading md5_keychain: %v", err)
		}
	}

	if o.Md5Keys != nil {
		if err = d.Set("md5_keys", flattenRouterOspfOspfInterfaceMd5Keys(o.Md5Keys, sort)); err != nil {
			return diag.Errorf("error reading md5_keys: %v", err)
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

	if o.NetworkType != nil {
		v := *o.NetworkType

		if err = d.Set("network_type", v); err != nil {
			return diag.Errorf("error reading network_type: %v", err)
		}
	}

	if o.PrefixLength != nil {
		v := *o.PrefixLength

		if err = d.Set("prefix_length", v); err != nil {
			return diag.Errorf("error reading prefix_length: %v", err)
		}
	}

	if o.Priority != nil {
		v := *o.Priority

		if err = d.Set("priority", v); err != nil {
			return diag.Errorf("error reading priority: %v", err)
		}
	}

	if o.ResyncTimeout != nil {
		v := *o.ResyncTimeout

		if err = d.Set("resync_timeout", v); err != nil {
			return diag.Errorf("error reading resync_timeout: %v", err)
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

func getObjectRouterOspfOspfInterface(d *schema.ResourceData, sv string) (*models.RouterOspfOspfInterface, diag.Diagnostics) {
	obj := models.RouterOspfOspfInterface{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("authentication"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("authentication", sv)
				diags = append(diags, e)
			}
			obj.Authentication = &v2
		}
	}
	if v1, ok := d.GetOk("authentication_key"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("authentication_key", sv)
				diags = append(diags, e)
			}
			obj.AuthenticationKey = &v2
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
	if v1, ok := d.GetOk("comments"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("comments", sv)
				diags = append(diags, e)
			}
			obj.Comments = &v2
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
	if v1, ok := d.GetOk("database_filter_out"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("database_filter_out", sv)
				diags = append(diags, e)
			}
			obj.DatabaseFilterOut = &v2
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
	if v1, ok := d.GetOk("hello_multiplier"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("hello_multiplier", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.HelloMultiplier = &tmp
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
	if v1, ok := d.GetOk("ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ip", sv)
				diags = append(diags, e)
			}
			obj.Ip = &v2
		}
	}
	if v1, ok := d.GetOk("keychain"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.1", "") {
				e := utils.AttributeVersionWarning("keychain", sv)
				diags = append(diags, e)
			}
			obj.Keychain = &v2
		}
	}
	if v1, ok := d.GetOk("md5_keychain"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v7.0.1") {
				e := utils.AttributeVersionWarning("md5_keychain", sv)
				diags = append(diags, e)
			}
			obj.Md5Keychain = &v2
		}
	}
	if v, ok := d.GetOk("md5_keys"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("md5_keys", sv)
			diags = append(diags, e)
		}
		t, err := expandRouterOspfOspfInterfaceMd5Keys(d, v, "md5_keys", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Md5Keys = t
		}
	} else if d.HasChange("md5_keys") {
		old, new := d.GetChange("md5_keys")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Md5Keys = &[]models.RouterOspfOspfInterfaceMd5Keys{}
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
	if v1, ok := d.GetOk("network_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("network_type", sv)
				diags = append(diags, e)
			}
			obj.NetworkType = &v2
		}
	}
	if v1, ok := d.GetOk("prefix_length"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("prefix_length", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.PrefixLength = &tmp
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
	if v1, ok := d.GetOk("resync_timeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("resync_timeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.ResyncTimeout = &tmp
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
