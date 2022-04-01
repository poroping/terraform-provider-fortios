// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.2.0 schemas
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
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func resourceVpnIpsecManualkey() *schema.Resource {
	return &schema.Resource{
		Description: "Configure IPsec manual keys.",

		CreateContext: resourceVpnIpsecManualkeyCreate,
		ReadContext:   resourceVpnIpsecManualkeyRead,
		UpdateContext: resourceVpnIpsecManualkeyUpdate,
		DeleteContext: resourceVpnIpsecManualkeyDelete,

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
			"authentication": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"null", "md5", "sha1", "sha256", "sha384", "sha512"}, false),

				Description: "Authentication algorithm. Must be the same for both ends of the tunnel.",
				Optional:    true,
				Computed:    true,
			},
			"authkey": {
				Type: schema.TypeString,

				Description: "Hexadecimal authentication key in 16-digit (8-byte) segments separated by hyphens.",
				Optional:    true,
				Computed:    true,
			},
			"enckey": {
				Type: schema.TypeString,

				Description: "Hexadecimal encryption key in 16-digit (8-byte) segments separated by hyphens.",
				Optional:    true,
				Computed:    true,
			},
			"encryption": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"null", "des", "3des", "aes128", "aes192", "aes256", "aria128", "aria192", "aria256", "seed"}, false),

				Description: "Encryption algorithm. Must be the same for both ends of the tunnel.",
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
			"local_gw": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "Local gateway.",
				Optional:    true,
				Computed:    true,
			},
			"localspi": {
				Type: schema.TypeString,

				Description: "Local SPI, a hexadecimal 8-digit (4-byte) tag. Discerns between two traffic streams with different encryption rules.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "IPsec tunnel name.",
				ForceNew:    true,
				Required:    true,
			},
			"npu_offload": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable NPU offloading.",
				Optional:    true,
				Computed:    true,
			},
			"remote_gw": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "Peer gateway.",
				Optional:    true,
				Computed:    true,
			},
			"remotespi": {
				Type: schema.TypeString,

				Description: "Remote SPI, a hexadecimal 8-digit (4-byte) tag. Discerns between two traffic streams with different encryption rules.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceVpnIpsecManualkeyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating VpnIpsecManualkey resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectVpnIpsecManualkey(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateVpnIpsecManualkey(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("VpnIpsecManualkey")
	}

	return resourceVpnIpsecManualkeyRead(ctx, d, meta)
}

func resourceVpnIpsecManualkeyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectVpnIpsecManualkey(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateVpnIpsecManualkey(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating VpnIpsecManualkey resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("VpnIpsecManualkey")
	}

	return resourceVpnIpsecManualkeyRead(ctx, d, meta)
}

func resourceVpnIpsecManualkeyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteVpnIpsecManualkey(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting VpnIpsecManualkey resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnIpsecManualkeyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadVpnIpsecManualkey(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading VpnIpsecManualkey resource: %v", err)
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

	diags := refreshObjectVpnIpsecManualkey(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectVpnIpsecManualkey(d *schema.ResourceData, o *models.VpnIpsecManualkey, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Authentication != nil {
		v := *o.Authentication

		if err = d.Set("authentication", v); err != nil {
			return diag.Errorf("error reading authentication: %v", err)
		}
	}

	if o.Authkey != nil {
		v := *o.Authkey

		if err = d.Set("authkey", v); err != nil {
			return diag.Errorf("error reading authkey: %v", err)
		}
	}

	if o.Enckey != nil {
		v := *o.Enckey

		if err = d.Set("enckey", v); err != nil {
			return diag.Errorf("error reading enckey: %v", err)
		}
	}

	if o.Encryption != nil {
		v := *o.Encryption

		if err = d.Set("encryption", v); err != nil {
			return diag.Errorf("error reading encryption: %v", err)
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

	if o.Localspi != nil {
		v := *o.Localspi

		if err = d.Set("localspi", v); err != nil {
			return diag.Errorf("error reading localspi: %v", err)
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

	if o.Remotespi != nil {
		v := *o.Remotespi

		if err = d.Set("remotespi", v); err != nil {
			return diag.Errorf("error reading remotespi: %v", err)
		}
	}

	return nil
}

func getObjectVpnIpsecManualkey(d *schema.ResourceData, sv string) (*models.VpnIpsecManualkey, diag.Diagnostics) {
	obj := models.VpnIpsecManualkey{}
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
	if v1, ok := d.GetOk("authkey"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("authkey", sv)
				diags = append(diags, e)
			}
			obj.Authkey = &v2
		}
	}
	if v1, ok := d.GetOk("enckey"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("enckey", sv)
				diags = append(diags, e)
			}
			obj.Enckey = &v2
		}
	}
	if v1, ok := d.GetOk("encryption"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("encryption", sv)
				diags = append(diags, e)
			}
			obj.Encryption = &v2
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
	if v1, ok := d.GetOk("localspi"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("localspi", sv)
				diags = append(diags, e)
			}
			obj.Localspi = &v2
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
			if !utils.CheckVer(sv, "", "v6.4.0") {
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
	if v1, ok := d.GetOk("remotespi"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("remotespi", sv)
				diags = append(diags, e)
			}
			obj.Remotespi = &v2
		}
	}
	return &obj, diags
}
