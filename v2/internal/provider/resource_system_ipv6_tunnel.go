// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure IPv6/IPv4 in IPv6 tunnel.

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

func resourceSystemIpv6Tunnel() *schema.Resource {
	return &schema.Resource{
		Description: "Configure IPv6/IPv4 in IPv6 tunnel.",

		CreateContext: resourceSystemIpv6TunnelCreate,
		ReadContext:   resourceSystemIpv6TunnelRead,
		UpdateContext: resourceSystemIpv6TunnelUpdate,
		DeleteContext: resourceSystemIpv6TunnelDelete,

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
			"auto_asic_offload": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable tunnel ASIC offloading.",
				Optional:    true,
				Computed:    true,
			},
			"destination": {
				Type:             schema.TypeString,
				ValidateFunc:     validation.IsIPv6Address,
				DiffSuppressFunc: suppressors.DiffIPEqual,
				Description:      "Remote IPv6 address of the tunnel.",
				Optional:         true,
				Computed:         true,
			},
			"interface": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "Interface name.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "IPv6 tunnel name.",
				ForceNew:    true,
				Required:    true,
			},
			"source": {
				Type:             schema.TypeString,
				ValidateFunc:     validation.IsIPv6Address,
				DiffSuppressFunc: suppressors.DiffIPEqual,
				Description:      "Local IPv6 address of the tunnel.",
				Optional:         true,
				Computed:         true,
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

func resourceSystemIpv6TunnelCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating SystemIpv6Tunnel resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectSystemIpv6Tunnel(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemIpv6Tunnel(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemIpv6Tunnel")
	}

	return resourceSystemIpv6TunnelRead(ctx, d, meta)
}

func resourceSystemIpv6TunnelUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemIpv6Tunnel(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemIpv6Tunnel(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemIpv6Tunnel resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemIpv6Tunnel")
	}

	return resourceSystemIpv6TunnelRead(ctx, d, meta)
}

func resourceSystemIpv6TunnelDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteSystemIpv6Tunnel(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemIpv6Tunnel resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemIpv6TunnelRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemIpv6Tunnel(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemIpv6Tunnel resource: %v", err)
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

	diags := refreshObjectSystemIpv6Tunnel(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectSystemIpv6Tunnel(d *schema.ResourceData, o *models.SystemIpv6Tunnel, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AutoAsicOffload != nil {
		v := *o.AutoAsicOffload

		if err = d.Set("auto_asic_offload", v); err != nil {
			return diag.Errorf("error reading auto_asic_offload: %v", err)
		}
	}

	if o.Destination != nil {
		v := *o.Destination

		if err = d.Set("destination", v); err != nil {
			return diag.Errorf("error reading destination: %v", err)
		}
	}

	if o.Interface != nil {
		v := *o.Interface

		if err = d.Set("interface", v); err != nil {
			return diag.Errorf("error reading interface: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.Source != nil {
		v := *o.Source

		if err = d.Set("source", v); err != nil {
			return diag.Errorf("error reading source: %v", err)
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

func getObjectSystemIpv6Tunnel(d *schema.ResourceData, sv string) (*models.SystemIpv6Tunnel, diag.Diagnostics) {
	obj := models.SystemIpv6Tunnel{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("auto_asic_offload"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.0") {
				e := utils.AttributeVersionWarning("auto_asic_offload", sv)
				diags = append(diags, e)
			}
			obj.AutoAsicOffload = &v2
		}
	}
	if v1, ok := d.GetOk("destination"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("destination", sv)
				diags = append(diags, e)
			}
			obj.Destination = &v2
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
	if v1, ok := d.GetOk("name"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("name", sv)
				diags = append(diags, e)
			}
			obj.Name = &v2
		}
	}
	if v1, ok := d.GetOk("source"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("source", sv)
				diags = append(diags, e)
			}
			obj.Source = &v2
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
