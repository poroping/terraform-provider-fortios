// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure IP in IP Tunneling.

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

func resourceSystemIpipTunnel() *schema.Resource {
	return &schema.Resource{
		Description: "Configure IP in IP Tunneling.",

		CreateContext: resourceSystemIpipTunnelCreate,
		ReadContext:   resourceSystemIpipTunnelRead,
		UpdateContext: resourceSystemIpipTunnelUpdate,
		DeleteContext: resourceSystemIpipTunnelDelete,

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
			"auto_asic_offload": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable tunnel ASIC offloading.",
				Optional:    true,
				Computed:    true,
			},
			"interface": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "Interface name that is associated with the incoming traffic from available options.",
				Optional:    true,
				Computed:    true,
			},
			"local_gw": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "IPv4 address for the local gateway.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "IPIP Tunnel name.",
				ForceNew:    true,
				Required:    true,
			},
			"remote_gw": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "IPv4 address for the remote gateway.",
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

func resourceSystemIpipTunnelCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating SystemIpipTunnel resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectSystemIpipTunnel(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemIpipTunnel(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemIpipTunnel")
	}

	return resourceSystemIpipTunnelRead(ctx, d, meta)
}

func resourceSystemIpipTunnelUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemIpipTunnel(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemIpipTunnel(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemIpipTunnel resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemIpipTunnel")
	}

	return resourceSystemIpipTunnelRead(ctx, d, meta)
}

func resourceSystemIpipTunnelDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteSystemIpipTunnel(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemIpipTunnel resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemIpipTunnelRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemIpipTunnel(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemIpipTunnel resource: %v", err)
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

	diags := refreshObjectSystemIpipTunnel(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectSystemIpipTunnel(d *schema.ResourceData, o *models.SystemIpipTunnel, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AutoAsicOffload != nil {
		v := *o.AutoAsicOffload

		if err = d.Set("auto_asic_offload", v); err != nil {
			return diag.Errorf("error reading auto_asic_offload: %v", err)
		}
	}

	if o.Interface != nil {
		v := *o.Interface

		if err = d.Set("interface", v); err != nil {
			return diag.Errorf("error reading interface: %v", err)
		}
	}

	if o.LocalGw != nil {
		v := *o.LocalGw

		if err = d.Set("local_gw", v); err != nil {
			return diag.Errorf("error reading local_gw: %v", err)
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

	if o.UseSdwan != nil {
		v := *o.UseSdwan

		if err = d.Set("use_sdwan", v); err != nil {
			return diag.Errorf("error reading use_sdwan: %v", err)
		}
	}

	return nil
}

func getObjectSystemIpipTunnel(d *schema.ResourceData, sv string) (*models.SystemIpipTunnel, diag.Diagnostics) {
	obj := models.SystemIpipTunnel{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("auto_asic_offload"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("auto_asic_offload", sv)
				diags = append(diags, e)
			}
			obj.AutoAsicOffload = &v2
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
	if v1, ok := d.GetOk("local_gw"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("local_gw", sv)
				diags = append(diags, e)
			}
			obj.LocalGw = &v2
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
