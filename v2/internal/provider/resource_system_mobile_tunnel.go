// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure Mobile tunnels, an implementation of Network Mobility (NEMO) extensions for Mobile IPv4 RFC5177.

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

func resourceSystemMobileTunnel() *schema.Resource {
	return &schema.Resource{
		Description: "Configure Mobile tunnels, an implementation of Network Mobility (NEMO) extensions for Mobile IPv4 RFC5177.",

		CreateContext: resourceSystemMobileTunnelCreate,
		ReadContext:   resourceSystemMobileTunnelRead,
		UpdateContext: resourceSystemMobileTunnelUpdate,
		DeleteContext: resourceSystemMobileTunnelDelete,

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
			"hash_algorithm": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"hmac-md5"}, false),

				Description: "Hash Algorithm (Keyed MD5).",
				Optional:    true,
				Computed:    true,
			},
			"home_address": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "Home IP address (Format: xxx.xxx.xxx.xxx).",
				Optional:    true,
				Computed:    true,
			},
			"home_agent": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "IPv4 address of the NEMO HA (Format: xxx.xxx.xxx.xxx).",
				Optional:    true,
				Computed:    true,
			},
			"lifetime": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(180, 65535),

				Description: "NMMO HA registration request lifetime (180 - 65535 sec, default = 65535).",
				Optional:    true,
				Computed:    true,
			},
			"n_mhae_key": {
				Type: schema.TypeString,

				Description: "NEMO authentication key.",
				Optional:    true,
				Computed:    true,
			},
			"n_mhae_key_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"ascii", "base64"}, false),

				Description: "NEMO authentication key type (ascii or base64).",
				Optional:    true,
				Computed:    true,
			},
			"n_mhae_spi": {
				Type: schema.TypeInt,

				Description: "NEMO authentication SPI (default: 256).",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "Tunnel name.",
				ForceNew:    true,
				Required:    true,
			},
			"network": {
				Type:        schema.TypeList,
				Description: "NEMO network configuration.",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type: schema.TypeInt,

							Description: "Network entry ID.",
							Optional:    true,
							Computed:    true,
						},
						"interface": {
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),

							Description: "Select the associated interface name from available options.",
							Optional:    true,
							Computed:    true,
						},
						"prefix": {
							Type:         schema.TypeString,
							ValidateFunc: validators.FortiValidateIPv4Classnet,

							Description: "Class IP and Netmask with correction (Format:xxx.xxx.xxx.xxx xxx.xxx.xxx.xxx or xxx.xxx.xxx.xxx/x).",
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"reg_interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(5, 300),

				Description: "NMMO HA registration interval (5 - 300, default = 5).",
				Optional:    true,
				Computed:    true,
			},
			"reg_retry": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 30),

				Description: "Maximum number of NMMO HA registration retries (1 to 30, default = 3).",
				Optional:    true,
				Computed:    true,
			},
			"renew_interval": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(5, 60),

				Description: "Time before lifetime expiraton to send NMMO HA re-registration (5 - 60, default = 60).",
				Optional:    true,
				Computed:    true,
			},
			"roaming_interface": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "Select the associated interface name from available options.",
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"disable", "enable"}, false),

				Description: "Enable/disable this mobile tunnel.",
				Optional:    true,
				Computed:    true,
			},
			"tunnel_mode": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"gre"}, false),

				Description: "NEMO tunnnel mode (GRE tunnel).",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSystemMobileTunnelCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating SystemMobileTunnel resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectSystemMobileTunnel(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemMobileTunnel(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemMobileTunnel")
	}

	return resourceSystemMobileTunnelRead(ctx, d, meta)
}

func resourceSystemMobileTunnelUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemMobileTunnel(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemMobileTunnel(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemMobileTunnel resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemMobileTunnel")
	}

	return resourceSystemMobileTunnelRead(ctx, d, meta)
}

func resourceSystemMobileTunnelDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteSystemMobileTunnel(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemMobileTunnel resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemMobileTunnelRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemMobileTunnel(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemMobileTunnel resource: %v", err)
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

	diags := refreshObjectSystemMobileTunnel(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func flattenSystemMobileTunnelNetwork(v *[]models.SystemMobileTunnelNetwork, sort bool) interface{} {
	flat := make([]map[string]interface{}, 0)

	if v != nil {
		for _, cfg := range *v {
			v := make(map[string]interface{})
			if tmp := cfg.Id; tmp != nil {
				v["id"] = *tmp
			}

			if tmp := cfg.Interface; tmp != nil {
				v["interface"] = *tmp
			}

			if tmp := cfg.Prefix; tmp != nil {
				v["prefix"] = *tmp
			}

			flat = append(flat, v)
		}
	}

	if sort {
		utils.SortSubtable(flat, "id")
	}

	return flat
}

func refreshObjectSystemMobileTunnel(d *schema.ResourceData, o *models.SystemMobileTunnel, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.HashAlgorithm != nil {
		v := *o.HashAlgorithm

		if err = d.Set("hash_algorithm", v); err != nil {
			return diag.Errorf("error reading hash_algorithm: %v", err)
		}
	}

	if o.HomeAddress != nil {
		v := *o.HomeAddress

		if err = d.Set("home_address", v); err != nil {
			return diag.Errorf("error reading home_address: %v", err)
		}
	}

	if o.HomeAgent != nil {
		v := *o.HomeAgent

		if err = d.Set("home_agent", v); err != nil {
			return diag.Errorf("error reading home_agent: %v", err)
		}
	}

	if o.Lifetime != nil {
		v := *o.Lifetime

		if err = d.Set("lifetime", v); err != nil {
			return diag.Errorf("error reading lifetime: %v", err)
		}
	}

	if o.NMhaeKey != nil {
		v := *o.NMhaeKey

		if err = d.Set("n_mhae_key", v); err != nil {
			return diag.Errorf("error reading n_mhae_key: %v", err)
		}
	}

	if o.NMhaeKeyType != nil {
		v := *o.NMhaeKeyType

		if err = d.Set("n_mhae_key_type", v); err != nil {
			return diag.Errorf("error reading n_mhae_key_type: %v", err)
		}
	}

	if o.NMhaeSpi != nil {
		v := *o.NMhaeSpi

		if err = d.Set("n_mhae_spi", v); err != nil {
			return diag.Errorf("error reading n_mhae_spi: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.Network != nil {
		if err = d.Set("network", flattenSystemMobileTunnelNetwork(o.Network, sort)); err != nil {
			return diag.Errorf("error reading network: %v", err)
		}
	}

	if o.RegInterval != nil {
		v := *o.RegInterval

		if err = d.Set("reg_interval", v); err != nil {
			return diag.Errorf("error reading reg_interval: %v", err)
		}
	}

	if o.RegRetry != nil {
		v := *o.RegRetry

		if err = d.Set("reg_retry", v); err != nil {
			return diag.Errorf("error reading reg_retry: %v", err)
		}
	}

	if o.RenewInterval != nil {
		v := *o.RenewInterval

		if err = d.Set("renew_interval", v); err != nil {
			return diag.Errorf("error reading renew_interval: %v", err)
		}
	}

	if o.RoamingInterface != nil {
		v := *o.RoamingInterface

		if err = d.Set("roaming_interface", v); err != nil {
			return diag.Errorf("error reading roaming_interface: %v", err)
		}
	}

	if o.Status != nil {
		v := *o.Status

		if err = d.Set("status", v); err != nil {
			return diag.Errorf("error reading status: %v", err)
		}
	}

	if o.TunnelMode != nil {
		v := *o.TunnelMode

		if err = d.Set("tunnel_mode", v); err != nil {
			return diag.Errorf("error reading tunnel_mode: %v", err)
		}
	}

	return nil
}

func expandSystemMobileTunnelNetwork(d *schema.ResourceData, v interface{}, pre string, sv string) (*[]models.SystemMobileTunnelNetwork, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	var result []models.SystemMobileTunnelNetwork

	for i := range l {
		tmp := models.SystemMobileTunnelNetwork{}
		var pre_append string

		pre_append = fmt.Sprintf("%s.%d.id", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(int64); ok {
				tmp.Id = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.interface", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Interface = &v2
			}
		}

		pre_append = fmt.Sprintf("%s.%d.prefix", pre, i)
		if v1, ok := d.GetOk(pre_append); ok {
			if v2, ok := v1.(string); ok {
				tmp.Prefix = &v2
			}
		}

		result = append(result, tmp)
	}
	return &result, nil
}

func getObjectSystemMobileTunnel(d *schema.ResourceData, sv string) (*models.SystemMobileTunnel, diag.Diagnostics) {
	obj := models.SystemMobileTunnel{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("hash_algorithm"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("hash_algorithm", sv)
				diags = append(diags, e)
			}
			obj.HashAlgorithm = &v2
		}
	}
	if v1, ok := d.GetOk("home_address"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("home_address", sv)
				diags = append(diags, e)
			}
			obj.HomeAddress = &v2
		}
	}
	if v1, ok := d.GetOk("home_agent"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("home_agent", sv)
				diags = append(diags, e)
			}
			obj.HomeAgent = &v2
		}
	}
	if v1, ok := d.GetOk("lifetime"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("lifetime", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.Lifetime = &tmp
		}
	}
	if v1, ok := d.GetOk("n_mhae_key"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("n_mhae_key", sv)
				diags = append(diags, e)
			}
			obj.NMhaeKey = &v2
		}
	}
	if v1, ok := d.GetOk("n_mhae_key_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("n_mhae_key_type", sv)
				diags = append(diags, e)
			}
			obj.NMhaeKeyType = &v2
		}
	}
	if v1, ok := d.GetOk("n_mhae_spi"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("n_mhae_spi", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.NMhaeSpi = &tmp
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
	if v, ok := d.GetOk("network"); ok {
		if !utils.CheckVer(sv, "", "") {
			e := utils.AttributeVersionWarning("network", sv)
			diags = append(diags, e)
		}
		t, err := expandSystemMobileTunnelNetwork(d, v, "network", sv)
		if err != nil {
			return &obj, diag.FromErr(err)
		} else if t != nil {
			obj.Network = t
		}
	} else if d.HasChange("network") {
		old, new := d.GetChange("network")
		if len(old.([]interface{})) > 0 && len(new.([]interface{})) == 0 {
			obj.Network = &[]models.SystemMobileTunnelNetwork{}
		}
	}
	if v1, ok := d.GetOk("reg_interval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("reg_interval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.RegInterval = &tmp
		}
	}
	if v1, ok := d.GetOk("reg_retry"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("reg_retry", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.RegRetry = &tmp
		}
	}
	if v1, ok := d.GetOk("renew_interval"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("renew_interval", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.RenewInterval = &tmp
		}
	}
	if v1, ok := d.GetOk("roaming_interface"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("roaming_interface", sv)
				diags = append(diags, e)
			}
			obj.RoamingInterface = &v2
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
	if v1, ok := d.GetOk("tunnel_mode"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("tunnel_mode", sv)
				diags = append(diags, e)
			}
			obj.TunnelMode = &v2
		}
	}
	return &obj, diags
}
