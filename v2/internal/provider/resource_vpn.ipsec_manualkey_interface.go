// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure IPsec manual keys.

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

func resourceVpnipsecManualkeyInterface() *schema.Resource {
	return &schema.Resource{
		Description: "Configure IPsec manual keys.",

		CreateContext: resourceVpnipsecManualkeyInterfaceCreate,
		ReadContext:   resourceVpnipsecManualkeyInterfaceRead,
		UpdateContext: resourceVpnipsecManualkeyInterfaceUpdate,
		DeleteContext: resourceVpnipsecManualkeyInterfaceDelete,

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
			"addr_type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"4", "6"}, false),

				Description: "IP version to use for IP packets.",
				Optional:    true,
				Computed:    true,
			},
			"auth_alg": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"null", "md5", "sha1", "sha256", "sha384", "sha512"}, false),

				Description: "Authentication algorithm. Must be the same for both ends of the tunnel.",
				Optional:    true,
				Computed:    true,
			},
			"auth_key": {
				Type: schema.TypeString,

				Description: "Hexadecimal authentication key in 16-digit (8-byte) segments separated by hyphens.",
				Optional:    true,
				Computed:    true,
			},
			"enc_alg": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"null", "des", "3des", "aes128", "aes192", "aes256", "aria128", "aria192", "aria256", "seed"}, false),

				Description: "Encryption algorithm. Must be the same for both ends of the tunnel.",
				Optional:    true,
				Computed:    true,
			},
			"enc_key": {
				Type: schema.TypeString,

				Description: "Hexadecimal encryption key in 16-digit (8-byte) segments separated by hyphens.",
				Optional:    true,
				Computed:    true,
			},
			"interface": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "Name of the physical, aggregate, or VLAN interface.",
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
			"local_gw": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "IPv4 address of the local gateway's external interface.",
				Optional:    true,
				Computed:    true,
			},
			"local_gw6": {
				Type:             schema.TypeString,
				ValidateFunc:     validation.IsIPv6Address,
				DiffSuppressFunc: suppressors.DiffIPEqual,
				Description:      "Local IPv6 address of VPN gateway.",
				Optional:         true,
				Computed:         true,
			},
			"local_spi": {
				Type: schema.TypeString,

				Description: "Local SPI, a hexadecimal 8-digit (4-byte) tag. Discerns between two traffic streams with different encryption rules.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "IPsec tunnel name.",
				ForceNew:    true,
				Required:    true,
			},
			"npu_offload": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable offloading IPsec VPN manual key sessions to NPUs.",
				Optional:    true,
				Computed:    true,
			},
			"remote_gw": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "IPv4 address of the remote gateway's external interface.",
				Optional:    true,
				Computed:    true,
			},
			"remote_gw6": {
				Type:             schema.TypeString,
				ValidateFunc:     validation.IsIPv6Address,
				DiffSuppressFunc: suppressors.DiffIPEqual,
				Description:      "Remote IPv6 address of VPN gateway.",
				Optional:         true,
				Computed:         true,
			},
			"remote_spi": {
				Type: schema.TypeString,

				Description: "Remote SPI, a hexadecimal 8-digit (4-byte) tag. Discerns between two traffic streams with different encryption rules.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceVpnipsecManualkeyInterfaceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating VpnipsecManualkeyInterface resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectVpnipsecManualkeyInterface(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateVpnipsecManualkeyInterface(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("VpnipsecManualkeyInterface")
	}

	return resourceVpnipsecManualkeyInterfaceRead(ctx, d, meta)
}

func resourceVpnipsecManualkeyInterfaceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectVpnipsecManualkeyInterface(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateVpnipsecManualkeyInterface(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating VpnipsecManualkeyInterface resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("VpnipsecManualkeyInterface")
	}

	return resourceVpnipsecManualkeyInterfaceRead(ctx, d, meta)
}

func resourceVpnipsecManualkeyInterfaceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteVpnipsecManualkeyInterface(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting VpnipsecManualkeyInterface resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnipsecManualkeyInterfaceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadVpnipsecManualkeyInterface(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading VpnipsecManualkeyInterface resource: %v", err)
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

	diags := refreshObjectVpnipsecManualkeyInterface(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectVpnipsecManualkeyInterface(d *schema.ResourceData, o *models.VpnipsecManualkeyInterface, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AddrType != nil {
		v := *o.AddrType

		if err = d.Set("addr_type", v); err != nil {
			return diag.Errorf("error reading addr_type: %v", err)
		}
	}

	if o.AuthAlg != nil {
		v := *o.AuthAlg

		if err = d.Set("auth_alg", v); err != nil {
			return diag.Errorf("error reading auth_alg: %v", err)
		}
	}

	if o.AuthKey != nil {
		v := *o.AuthKey

		if err = d.Set("auth_key", v); err != nil {
			return diag.Errorf("error reading auth_key: %v", err)
		}
	}

	if o.EncAlg != nil {
		v := *o.EncAlg

		if err = d.Set("enc_alg", v); err != nil {
			return diag.Errorf("error reading enc_alg: %v", err)
		}
	}

	if o.EncKey != nil {
		v := *o.EncKey

		if err = d.Set("enc_key", v); err != nil {
			return diag.Errorf("error reading enc_key: %v", err)
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

	if o.LocalSpi != nil {
		v := *o.LocalSpi

		if err = d.Set("local_spi", v); err != nil {
			return diag.Errorf("error reading local_spi: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.NpuOffload != nil {
		v := *o.NpuOffload

		if err = d.Set("npu_offload", v); err != nil {
			return diag.Errorf("error reading npu_offload: %v", err)
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

	if o.RemoteSpi != nil {
		v := *o.RemoteSpi

		if err = d.Set("remote_spi", v); err != nil {
			return diag.Errorf("error reading remote_spi: %v", err)
		}
	}

	return nil
}

func getObjectVpnipsecManualkeyInterface(d *schema.ResourceData, sv string) (*models.VpnipsecManualkeyInterface, diag.Diagnostics) {
	obj := models.VpnipsecManualkeyInterface{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("addr_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("addr_type", sv)
				diags = append(diags, e)
			}
			obj.AddrType = &v2
		}
	}
	if v1, ok := d.GetOk("auth_alg"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auth_alg", sv)
				diags = append(diags, e)
			}
			obj.AuthAlg = &v2
		}
	}
	if v1, ok := d.GetOk("auth_key"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auth_key", sv)
				diags = append(diags, e)
			}
			obj.AuthKey = &v2
		}
	}
	if v1, ok := d.GetOk("enc_alg"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("enc_alg", sv)
				diags = append(diags, e)
			}
			obj.EncAlg = &v2
		}
	}
	if v1, ok := d.GetOk("enc_key"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("enc_key", sv)
				diags = append(diags, e)
			}
			obj.EncKey = &v2
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
	if v1, ok := d.GetOk("local_spi"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("local_spi", sv)
				diags = append(diags, e)
			}
			obj.LocalSpi = &v2
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
	if v1, ok := d.GetOk("npu_offload"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "v6.4.2") {
				e := utils.AttributeVersionWarning("npu_offload", sv)
				diags = append(diags, e)
			}
			obj.NpuOffload = &v2
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
	if v1, ok := d.GetOk("remote_spi"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("remote_spi", sv)
				diags = append(diags, e)
			}
			obj.RemoteSpi = &v2
		}
	}
	return &obj, diags
}
