// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure connection capability.

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

func resourceWirelessControllerHotspot20H2qpConnCapability() *schema.Resource {
	return &schema.Resource{
		Description: "Configure connection capability.",

		CreateContext: resourceWirelessControllerHotspot20H2qpConnCapabilityCreate,
		ReadContext:   resourceWirelessControllerHotspot20H2qpConnCapabilityRead,
		UpdateContext: resourceWirelessControllerHotspot20H2qpConnCapabilityUpdate,
		DeleteContext: resourceWirelessControllerHotspot20H2qpConnCapabilityDelete,

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
			"esp_port": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"closed", "open", "unknown"}, false),

				Description: "Set ESP port service (used by IPsec VPNs) status.",
				Optional:    true,
				Computed:    true,
			},
			"ftp_port": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"closed", "open", "unknown"}, false),

				Description: "Set FTP port service status.",
				Optional:    true,
				Computed:    true,
			},
			"http_port": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"closed", "open", "unknown"}, false),

				Description: "Set HTTP port service status.",
				Optional:    true,
				Computed:    true,
			},
			"icmp_port": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"closed", "open", "unknown"}, false),

				Description: "Set ICMP port service status.",
				Optional:    true,
				Computed:    true,
			},
			"ikev2_port": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"closed", "open", "unknown"}, false),

				Description: "Set IKEv2 port service for IPsec VPN status.",
				Optional:    true,
				Computed:    true,
			},
			"ikev2_xx_port": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"closed", "open", "unknown"}, false),

				Description: "Set UDP port 4500 (which may be used by IKEv2 for IPsec VPN) service status.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Connection capability name.",
				ForceNew:    true,
				Required:    true,
			},
			"pptp_vpn_port": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"closed", "open", "unknown"}, false),

				Description: "Set Point to Point Tunneling Protocol (PPTP) VPN port service status.",
				Optional:    true,
				Computed:    true,
			},
			"ssh_port": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"closed", "open", "unknown"}, false),

				Description: "Set SSH port service status.",
				Optional:    true,
				Computed:    true,
			},
			"tls_port": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"closed", "open", "unknown"}, false),

				Description: "Set TLS VPN (HTTPS) port service status.",
				Optional:    true,
				Computed:    true,
			},
			"voip_tcp_port": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"closed", "open", "unknown"}, false),

				Description: "Set VoIP TCP port service status.",
				Optional:    true,
				Computed:    true,
			},
			"voip_udp_port": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"closed", "open", "unknown"}, false),

				Description: "Set VoIP UDP port service status.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceWirelessControllerHotspot20H2qpConnCapabilityCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating WirelessControllerHotspot20H2qpConnCapability resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectWirelessControllerHotspot20H2qpConnCapability(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateWirelessControllerHotspot20H2qpConnCapability(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WirelessControllerHotspot20H2qpConnCapability")
	}

	return resourceWirelessControllerHotspot20H2qpConnCapabilityRead(ctx, d, meta)
}

func resourceWirelessControllerHotspot20H2qpConnCapabilityUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectWirelessControllerHotspot20H2qpConnCapability(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateWirelessControllerHotspot20H2qpConnCapability(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating WirelessControllerHotspot20H2qpConnCapability resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("WirelessControllerHotspot20H2qpConnCapability")
	}

	return resourceWirelessControllerHotspot20H2qpConnCapabilityRead(ctx, d, meta)
}

func resourceWirelessControllerHotspot20H2qpConnCapabilityDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteWirelessControllerHotspot20H2qpConnCapability(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting WirelessControllerHotspot20H2qpConnCapability resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerHotspot20H2qpConnCapabilityRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadWirelessControllerHotspot20H2qpConnCapability(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading WirelessControllerHotspot20H2qpConnCapability resource: %v", err)
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

	diags := refreshObjectWirelessControllerHotspot20H2qpConnCapability(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectWirelessControllerHotspot20H2qpConnCapability(d *schema.ResourceData, o *models.WirelessControllerHotspot20H2qpConnCapability, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.EspPort != nil {
		v := *o.EspPort

		if err = d.Set("esp_port", v); err != nil {
			return diag.Errorf("error reading esp_port: %v", err)
		}
	}

	if o.FtpPort != nil {
		v := *o.FtpPort

		if err = d.Set("ftp_port", v); err != nil {
			return diag.Errorf("error reading ftp_port: %v", err)
		}
	}

	if o.HttpPort != nil {
		v := *o.HttpPort

		if err = d.Set("http_port", v); err != nil {
			return diag.Errorf("error reading http_port: %v", err)
		}
	}

	if o.IcmpPort != nil {
		v := *o.IcmpPort

		if err = d.Set("icmp_port", v); err != nil {
			return diag.Errorf("error reading icmp_port: %v", err)
		}
	}

	if o.Ikev2Port != nil {
		v := *o.Ikev2Port

		if err = d.Set("ikev2_port", v); err != nil {
			return diag.Errorf("error reading ikev2_port: %v", err)
		}
	}

	if o.Ikev2XxPort != nil {
		v := *o.Ikev2XxPort

		if err = d.Set("ikev2_xx_port", v); err != nil {
			return diag.Errorf("error reading ikev2_xx_port: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.PptpVpnPort != nil {
		v := *o.PptpVpnPort

		if err = d.Set("pptp_vpn_port", v); err != nil {
			return diag.Errorf("error reading pptp_vpn_port: %v", err)
		}
	}

	if o.SshPort != nil {
		v := *o.SshPort

		if err = d.Set("ssh_port", v); err != nil {
			return diag.Errorf("error reading ssh_port: %v", err)
		}
	}

	if o.TlsPort != nil {
		v := *o.TlsPort

		if err = d.Set("tls_port", v); err != nil {
			return diag.Errorf("error reading tls_port: %v", err)
		}
	}

	if o.VoipTcpPort != nil {
		v := *o.VoipTcpPort

		if err = d.Set("voip_tcp_port", v); err != nil {
			return diag.Errorf("error reading voip_tcp_port: %v", err)
		}
	}

	if o.VoipUdpPort != nil {
		v := *o.VoipUdpPort

		if err = d.Set("voip_udp_port", v); err != nil {
			return diag.Errorf("error reading voip_udp_port: %v", err)
		}
	}

	return nil
}

func getObjectWirelessControllerHotspot20H2qpConnCapability(d *schema.ResourceData, sv string) (*models.WirelessControllerHotspot20H2qpConnCapability, diag.Diagnostics) {
	obj := models.WirelessControllerHotspot20H2qpConnCapability{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("esp_port"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("esp_port", sv)
				diags = append(diags, e)
			}
			obj.EspPort = &v2
		}
	}
	if v1, ok := d.GetOk("ftp_port"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ftp_port", sv)
				diags = append(diags, e)
			}
			obj.FtpPort = &v2
		}
	}
	if v1, ok := d.GetOk("http_port"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("http_port", sv)
				diags = append(diags, e)
			}
			obj.HttpPort = &v2
		}
	}
	if v1, ok := d.GetOk("icmp_port"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("icmp_port", sv)
				diags = append(diags, e)
			}
			obj.IcmpPort = &v2
		}
	}
	if v1, ok := d.GetOk("ikev2_port"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ikev2_port", sv)
				diags = append(diags, e)
			}
			obj.Ikev2Port = &v2
		}
	}
	if v1, ok := d.GetOk("ikev2_xx_port"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ikev2_xx_port", sv)
				diags = append(diags, e)
			}
			obj.Ikev2XxPort = &v2
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
	if v1, ok := d.GetOk("pptp_vpn_port"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("pptp_vpn_port", sv)
				diags = append(diags, e)
			}
			obj.PptpVpnPort = &v2
		}
	}
	if v1, ok := d.GetOk("ssh_port"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssh_port", sv)
				diags = append(diags, e)
			}
			obj.SshPort = &v2
		}
	}
	if v1, ok := d.GetOk("tls_port"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("tls_port", sv)
				diags = append(diags, e)
			}
			obj.TlsPort = &v2
		}
	}
	if v1, ok := d.GetOk("voip_tcp_port"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("voip_tcp_port", sv)
				diags = append(diags, e)
			}
			obj.VoipTcpPort = &v2
		}
	}
	if v1, ok := d.GetOk("voip_udp_port"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("voip_udp_port", sv)
				diags = append(diags, e)
			}
			obj.VoipUdpPort = &v2
		}
	}
	return &obj, diags
}
