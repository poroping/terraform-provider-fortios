// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure wireless access gateway (WAG) profiles used for tunnels on AP.

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

func resourceWirelessControllerWagProfile() *schema.Resource {
	return &schema.Resource{
		Description: "Configure wireless access gateway (WAG) profiles used for tunnels on AP.",

		CreateContext: resourceWirelessControllerWagProfileCreate,
		ReadContext:   resourceWirelessControllerWagProfileRead,
		UpdateContext: resourceWirelessControllerWagProfileUpdate,
		DeleteContext: resourceWirelessControllerWagProfileDelete,

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
			"comment": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "Comment.",
				Optional:    true,
				Computed:    true,
			},
			"dhcp_ip_addr": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "IP address of the monitoring DHCP request packet sent through the tunnel.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Tunnel profile name.",
				ForceNew:    true,
				Required:    true,
			},
			"ping_interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),

				Description: "Interval between two tunnel monitoring echo packets (1 - 65535 sec, default = 1).",
				Optional:    true,
				Computed:    true,
			},
			"ping_number": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),

				Description: "Number of the tunnel monitoring echo packets (1 - 65535, default = 5).",
				Optional:    true,
				Computed:    true,
			},
			"return_packet_timeout": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),

				Description: "Window of time for the return packets from the tunnel's remote end (1 - 65535 sec, default = 160).",
				Optional:    true,
				Computed:    true,
			},
			"tunnel_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"l2tpv3", "gre"}, false),

				Description: "Tunnel type.",
				Optional:    true,
				Computed:    true,
			},
			"wag_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "IP Address of the wireless access gateway.",
				Optional:    true,
				Computed:    true,
			},
			"wag_port": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),

				Description: "UDP port of the wireless access gateway.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceWirelessControllerWagProfileCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating WirelessControllerWagProfile resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectWirelessControllerWagProfile(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateWirelessControllerWagProfile(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WirelessControllerWagProfile")
	}

	return resourceWirelessControllerWagProfileRead(ctx, d, meta)
}

func resourceWirelessControllerWagProfileUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectWirelessControllerWagProfile(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateWirelessControllerWagProfile(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating WirelessControllerWagProfile resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WirelessControllerWagProfile")
	}

	return resourceWirelessControllerWagProfileRead(ctx, d, meta)
}

func resourceWirelessControllerWagProfileDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteWirelessControllerWagProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting WirelessControllerWagProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerWagProfileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadWirelessControllerWagProfile(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WirelessControllerWagProfile resource: %v", err)
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

	diags := refreshObjectWirelessControllerWagProfile(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectWirelessControllerWagProfile(d *schema.ResourceData, o *models.WirelessControllerWagProfile, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Comment != nil {
		v := *o.Comment

		if err = d.Set("comment", v); err != nil {
			return diag.Errorf("error reading comment: %v", err)
		}
	}

	if o.DhcpIpAddr != nil {
		v := *o.DhcpIpAddr

		if err = d.Set("dhcp_ip_addr", v); err != nil {
			return diag.Errorf("error reading dhcp_ip_addr: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.PingInterval != nil {
		v := *o.PingInterval

		if err = d.Set("ping_interval", v); err != nil {
			return diag.Errorf("error reading ping_interval: %v", err)
		}
	}

	if o.PingNumber != nil {
		v := *o.PingNumber

		if err = d.Set("ping_number", v); err != nil {
			return diag.Errorf("error reading ping_number: %v", err)
		}
	}

	if o.ReturnPacketTimeout != nil {
		v := *o.ReturnPacketTimeout

		if err = d.Set("return_packet_timeout", v); err != nil {
			return diag.Errorf("error reading return_packet_timeout: %v", err)
		}
	}

	if o.TunnelType != nil {
		v := *o.TunnelType

		if err = d.Set("tunnel_type", v); err != nil {
			return diag.Errorf("error reading tunnel_type: %v", err)
		}
	}

	if o.WagIp != nil {
		v := *o.WagIp

		if err = d.Set("wag_ip", v); err != nil {
			return diag.Errorf("error reading wag_ip: %v", err)
		}
	}

	if o.WagPort != nil {
		v := *o.WagPort

		if err = d.Set("wag_port", v); err != nil {
			return diag.Errorf("error reading wag_port: %v", err)
		}
	}

	return nil
}

func getObjectWirelessControllerWagProfile(d *schema.ResourceData, sv string) (*models.WirelessControllerWagProfile, diag.Diagnostics) {
	obj := models.WirelessControllerWagProfile{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("comment"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("comment", sv)
				diags = append(diags, e)
			}
			obj.Comment = &v2
		}
	}
	if v1, ok := d.GetOk("dhcp_ip_addr"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dhcp_ip_addr", sv)
				diags = append(diags, e)
			}
			obj.DhcpIpAddr = &v2
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
	if v1, ok := d.GetOk("ping_interval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ping_interval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.PingInterval = &tmp
		}
	}
	if v1, ok := d.GetOk("ping_number"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ping_number", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.PingNumber = &tmp
		}
	}
	if v1, ok := d.GetOk("return_packet_timeout"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("return_packet_timeout", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.ReturnPacketTimeout = &tmp
		}
	}
	if v1, ok := d.GetOk("tunnel_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("tunnel_type", sv)
				diags = append(diags, e)
			}
			obj.TunnelType = &v2
		}
	}
	if v1, ok := d.GetOk("wag_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("wag_ip", sv)
				diags = append(diags, e)
			}
			obj.WagIp = &v2
		}
	}
	if v1, ok := d.GetOk("wag_port"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("wag_port", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.WagPort = &tmp
		}
	}
	return &obj, diags
}
