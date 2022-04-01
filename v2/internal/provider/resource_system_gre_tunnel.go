// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure GRE tunnel.

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

func resourceSystemGreTunnel() *schema.Resource {
	return &schema.Resource{
		Description: "Configure GRE tunnel.",

		CreateContext: resourceSystemGreTunnelCreate,
		ReadContext:   resourceSystemGreTunnelRead,
		UpdateContext: resourceSystemGreTunnelUpdate,
		DeleteContext: resourceSystemGreTunnelDelete,

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
			"checksum_reception": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable validating checksums in received GRE packets.",
				Optional:    true,
				Computed:    true,
			},
			"checksum_transmission": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable including checksums in transmitted GRE packets.",
				Optional:    true,
				Computed:    true,
			},
			"diffservcode": {
				Type: schema.TypeString,

				Description: "DiffServ setting to be applied to GRE tunnel outer IP header.",
				Optional:    true,
				Computed:    true,
			},
			"dscp_copying": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable DSCP copying.",
				Optional:    true,
				Computed:    true,
			},
			"interface": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "Interface name.",
				Optional:    true,
				Computed:    true,
			},
			"ip_version": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"4", "6"}, false),

				Description: "IP version to use for VPN interface.",
				Optional:    true,
				Computed:    true,
			},
			"keepalive_failtimes": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 255),

				Description: "Number of consecutive unreturned keepalive messages before a GRE connection is considered down (1 - 255).",
				Optional:    true,
				Computed:    true,
			},
			"keepalive_interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 32767),

				Description: "Keepalive message interval (0 - 32767, 0 = disabled).",
				Optional:    true,
				Computed:    true,
			},
			"key_inbound": {
				Type: schema.TypeInt,

				Description: "Require received GRE packets contain this key (0 - 4294967295).",
				Optional:    true,
				Computed:    true,
			},
			"key_outbound": {
				Type: schema.TypeInt,

				Description: "Include this key in transmitted GRE packets (0 - 4294967295).",
				Optional:    true,
				Computed:    true,
			},
			"local_gw": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "IP address of the local gateway.",
				Optional:    true,
				Computed:    true,
			},
			"local_gw6": {
				Type:             schema.TypeString,
				ValidateFunc:     validation.IsIPv6Address,
				DiffSuppressFunc: suppressors.DiffIPEqual,
				Description:      "IPv6 address of the local gateway.",
				Optional:         true,
				Computed:         true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "Tunnel name.",
				ForceNew:    true,
				Required:    true,
			},
			"remote_gw": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "IP address of the remote gateway.",
				Optional:    true,
				Computed:    true,
			},
			"remote_gw6": {
				Type:             schema.TypeString,
				ValidateFunc:     validation.IsIPv6Address,
				DiffSuppressFunc: suppressors.DiffIPEqual,
				Description:      "IPv6 address of the remote gateway.",
				Optional:         true,
				Computed:         true,
			},
			"sequence_number_reception": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable validating sequence numbers in received GRE packets.",
				Optional:    true,
				Computed:    true,
			},
			"sequence_number_transmission": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable including of sequence numbers in transmitted GRE packets.",
				Optional:    true,
				Computed:    true,
			},
			"use_sdwan": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable use of SD-WAN to reach remote gateway.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSystemGreTunnelCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating SystemGreTunnel resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectSystemGreTunnel(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemGreTunnel(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemGreTunnel")
	}

	return resourceSystemGreTunnelRead(ctx, d, meta)
}

func resourceSystemGreTunnelUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemGreTunnel(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemGreTunnel(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemGreTunnel resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemGreTunnel")
	}

	return resourceSystemGreTunnelRead(ctx, d, meta)
}

func resourceSystemGreTunnelDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteSystemGreTunnel(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemGreTunnel resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemGreTunnelRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemGreTunnel(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemGreTunnel resource: %v", err)
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

	diags := refreshObjectSystemGreTunnel(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectSystemGreTunnel(d *schema.ResourceData, o *models.SystemGreTunnel, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.ChecksumReception != nil {
		v := *o.ChecksumReception

		if err = d.Set("checksum_reception", v); err != nil {
			return diag.Errorf("error reading checksum_reception: %v", err)
		}
	}

	if o.ChecksumTransmission != nil {
		v := *o.ChecksumTransmission

		if err = d.Set("checksum_transmission", v); err != nil {
			return diag.Errorf("error reading checksum_transmission: %v", err)
		}
	}

	if o.Diffservcode != nil {
		v := *o.Diffservcode

		if err = d.Set("diffservcode", v); err != nil {
			return diag.Errorf("error reading diffservcode: %v", err)
		}
	}

	if o.DscpCopying != nil {
		v := *o.DscpCopying

		if err = d.Set("dscp_copying", v); err != nil {
			return diag.Errorf("error reading dscp_copying: %v", err)
		}
	}

	if o.Interface != nil {
		v := *o.Interface

		if err = d.Set("interface", v); err != nil {
			return diag.Errorf("error reading interface: %v", err)
		}
	}

	if o.IpVersion != nil {
		v := *o.IpVersion

		if err = d.Set("ip_version", v); err != nil {
			return diag.Errorf("error reading ip_version: %v", err)
		}
	}

	if o.KeepaliveFailtimes != nil {
		v := *o.KeepaliveFailtimes

		if err = d.Set("keepalive_failtimes", v); err != nil {
			return diag.Errorf("error reading keepalive_failtimes: %v", err)
		}
	}

	if o.KeepaliveInterval != nil {
		v := *o.KeepaliveInterval

		if err = d.Set("keepalive_interval", v); err != nil {
			return diag.Errorf("error reading keepalive_interval: %v", err)
		}
	}

	if o.KeyInbound != nil {
		v := *o.KeyInbound

		if err = d.Set("key_inbound", v); err != nil {
			return diag.Errorf("error reading key_inbound: %v", err)
		}
	}

	if o.KeyOutbound != nil {
		v := *o.KeyOutbound

		if err = d.Set("key_outbound", v); err != nil {
			return diag.Errorf("error reading key_outbound: %v", err)
		}
	}

	if o.LocalGw != nil {
		v := *o.LocalGw

		if err = d.Set("local_gw", v); err != nil {
			return diag.Errorf("error reading local_gw: %v", err)
		}
	}

	if o.LocalGw6 != nil {
		v := *o.LocalGw6

		if err = d.Set("local_gw6", v); err != nil {
			return diag.Errorf("error reading local_gw6: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.RemoteGw != nil {
		v := *o.RemoteGw

		if err = d.Set("remote_gw", v); err != nil {
			return diag.Errorf("error reading remote_gw: %v", err)
		}
	}

	if o.RemoteGw6 != nil {
		v := *o.RemoteGw6

		if err = d.Set("remote_gw6", v); err != nil {
			return diag.Errorf("error reading remote_gw6: %v", err)
		}
	}

	if o.SequenceNumberReception != nil {
		v := *o.SequenceNumberReception

		if err = d.Set("sequence_number_reception", v); err != nil {
			return diag.Errorf("error reading sequence_number_reception: %v", err)
		}
	}

	if o.SequenceNumberTransmission != nil {
		v := *o.SequenceNumberTransmission

		if err = d.Set("sequence_number_transmission", v); err != nil {
			return diag.Errorf("error reading sequence_number_transmission: %v", err)
		}
	}

	if o.UseSdwan != nil {
		v := *o.UseSdwan

		if err = d.Set("use_sdwan", v); err != nil {
			return diag.Errorf("error reading use_sdwan: %v", err)
		}
	}

	return nil
}

func getObjectSystemGreTunnel(d *schema.ResourceData, sv string) (*models.SystemGreTunnel, diag.Diagnostics) {
	obj := models.SystemGreTunnel{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("checksum_reception"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("checksum_reception", sv)
				diags = append(diags, e)
			}
			obj.ChecksumReception = &v2
		}
	}
	if v1, ok := d.GetOk("checksum_transmission"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("checksum_transmission", sv)
				diags = append(diags, e)
			}
			obj.ChecksumTransmission = &v2
		}
	}
	if v1, ok := d.GetOk("diffservcode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("diffservcode", sv)
				diags = append(diags, e)
			}
			obj.Diffservcode = &v2
		}
	}
	if v1, ok := d.GetOk("dscp_copying"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("dscp_copying", sv)
				diags = append(diags, e)
			}
			obj.DscpCopying = &v2
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
	if v1, ok := d.GetOk("ip_version"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ip_version", sv)
				diags = append(diags, e)
			}
			obj.IpVersion = &v2
		}
	}
	if v1, ok := d.GetOk("keepalive_failtimes"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("keepalive_failtimes", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.KeepaliveFailtimes = &tmp
		}
	}
	if v1, ok := d.GetOk("keepalive_interval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("keepalive_interval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.KeepaliveInterval = &tmp
		}
	}
	if v1, ok := d.GetOk("key_inbound"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("key_inbound", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.KeyInbound = &tmp
		}
	}
	if v1, ok := d.GetOk("key_outbound"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("key_outbound", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.KeyOutbound = &tmp
		}
	}
	if v1, ok := d.GetOk("local_gw"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("local_gw", sv)
				diags = append(diags, e)
			}
			obj.LocalGw = &v2
		}
	}
	if v1, ok := d.GetOk("local_gw6"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("local_gw6", sv)
				diags = append(diags, e)
			}
			obj.LocalGw6 = &v2
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
	if v1, ok := d.GetOk("remote_gw"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("remote_gw", sv)
				diags = append(diags, e)
			}
			obj.RemoteGw = &v2
		}
	}
	if v1, ok := d.GetOk("remote_gw6"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("remote_gw6", sv)
				diags = append(diags, e)
			}
			obj.RemoteGw6 = &v2
		}
	}
	if v1, ok := d.GetOk("sequence_number_reception"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("sequence_number_reception", sv)
				diags = append(diags, e)
			}
			obj.SequenceNumberReception = &v2
		}
	}
	if v1, ok := d.GetOk("sequence_number_transmission"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("sequence_number_transmission", sv)
				diags = append(diags, e)
			}
			obj.SequenceNumberTransmission = &v2
		}
	}
	if v1, ok := d.GetOk("use_sdwan"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("use_sdwan", sv)
				diags = append(diags, e)
			}
			obj.UseSdwan = &v2
		}
	}
	return &obj, diags
}
