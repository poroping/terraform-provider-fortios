// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: SSH proxy settings.

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

func resourceFirewallSshSetting() *schema.Resource {
	return &schema.Resource{
		Description: "SSH proxy settings.",

		CreateContext: resourceFirewallSshSettingCreate,
		ReadContext:   resourceFirewallSshSettingRead,
		UpdateContext: resourceFirewallSshSettingUpdate,
		DeleteContext: resourceFirewallSshSettingDelete,

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
			"caname": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "CA certificate used by SSH Inspection.",
				Optional:    true,
				Computed:    true,
			},
			"host_trusted_checking": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable host trusted checking.",
				Optional:    true,
				Computed:    true,
			},
			"hostkey_dsa1024": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "DSA certificate used by SSH proxy.",
				Optional:    true,
				Computed:    true,
			},
			"hostkey_ecdsa256": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "ECDSA nid256 certificate used by SSH proxy.",
				Optional:    true,
				Computed:    true,
			},
			"hostkey_ecdsa384": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "ECDSA nid384 certificate used by SSH proxy.",
				Optional:    true,
				Computed:    true,
			},
			"hostkey_ecdsa521": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "ECDSA nid384 certificate used by SSH proxy.",
				Optional:    true,
				Computed:    true,
			},
			"hostkey_ed25519": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "ED25519 hostkey used by SSH proxy.",
				Optional:    true,
				Computed:    true,
			},
			"hostkey_rsa2048": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "RSA certificate used by SSH proxy.",
				Optional:    true,
				Computed:    true,
			},
			"untrusted_caname": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "Untrusted CA certificate used by SSH Inspection.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceFirewallSshSettingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectFirewallSshSetting(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateFirewallSshSetting(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallSshSetting")
	}

	return resourceFirewallSshSettingRead(ctx, d, meta)
}

func resourceFirewallSshSettingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectFirewallSshSetting(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateFirewallSshSetting(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating FirewallSshSetting resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("FirewallSshSetting")
	}

	return resourceFirewallSshSettingRead(ctx, d, meta)
}

func resourceFirewallSshSettingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectFirewallSshSetting(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateFirewallSshSetting(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting FirewallSshSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallSshSettingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadFirewallSshSetting(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading FirewallSshSetting resource: %v", err)
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

	diags := refreshObjectFirewallSshSetting(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectFirewallSshSetting(d *schema.ResourceData, o *models.FirewallSshSetting, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Caname != nil {
		v := *o.Caname

		if err = d.Set("caname", v); err != nil {
			return diag.Errorf("error reading caname: %v", err)
		}
	}

	if o.HostTrustedChecking != nil {
		v := *o.HostTrustedChecking

		if err = d.Set("host_trusted_checking", v); err != nil {
			return diag.Errorf("error reading host_trusted_checking: %v", err)
		}
	}

	if o.HostkeyDsa1024 != nil {
		v := *o.HostkeyDsa1024

		if err = d.Set("hostkey_dsa1024", v); err != nil {
			return diag.Errorf("error reading hostkey_dsa1024: %v", err)
		}
	}

	if o.HostkeyEcdsa256 != nil {
		v := *o.HostkeyEcdsa256

		if err = d.Set("hostkey_ecdsa256", v); err != nil {
			return diag.Errorf("error reading hostkey_ecdsa256: %v", err)
		}
	}

	if o.HostkeyEcdsa384 != nil {
		v := *o.HostkeyEcdsa384

		if err = d.Set("hostkey_ecdsa384", v); err != nil {
			return diag.Errorf("error reading hostkey_ecdsa384: %v", err)
		}
	}

	if o.HostkeyEcdsa521 != nil {
		v := *o.HostkeyEcdsa521

		if err = d.Set("hostkey_ecdsa521", v); err != nil {
			return diag.Errorf("error reading hostkey_ecdsa521: %v", err)
		}
	}

	if o.HostkeyEd25519 != nil {
		v := *o.HostkeyEd25519

		if err = d.Set("hostkey_ed25519", v); err != nil {
			return diag.Errorf("error reading hostkey_ed25519: %v", err)
		}
	}

	if o.HostkeyRsa2048 != nil {
		v := *o.HostkeyRsa2048

		if err = d.Set("hostkey_rsa2048", v); err != nil {
			return diag.Errorf("error reading hostkey_rsa2048: %v", err)
		}
	}

	if o.UntrustedCaname != nil {
		v := *o.UntrustedCaname

		if err = d.Set("untrusted_caname", v); err != nil {
			return diag.Errorf("error reading untrusted_caname: %v", err)
		}
	}

	return nil
}

func getObjectFirewallSshSetting(d *schema.ResourceData, sv string) (*models.FirewallSshSetting, diag.Diagnostics) {
	obj := models.FirewallSshSetting{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("caname"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("caname", sv)
				diags = append(diags, e)
			}
			obj.Caname = &v2
		}
	}
	if v1, ok := d.GetOk("host_trusted_checking"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("host_trusted_checking", sv)
				diags = append(diags, e)
			}
			obj.HostTrustedChecking = &v2
		}
	}
	if v1, ok := d.GetOk("hostkey_dsa1024"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("hostkey_dsa1024", sv)
				diags = append(diags, e)
			}
			obj.HostkeyDsa1024 = &v2
		}
	}
	if v1, ok := d.GetOk("hostkey_ecdsa256"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("hostkey_ecdsa256", sv)
				diags = append(diags, e)
			}
			obj.HostkeyEcdsa256 = &v2
		}
	}
	if v1, ok := d.GetOk("hostkey_ecdsa384"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("hostkey_ecdsa384", sv)
				diags = append(diags, e)
			}
			obj.HostkeyEcdsa384 = &v2
		}
	}
	if v1, ok := d.GetOk("hostkey_ecdsa521"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("hostkey_ecdsa521", sv)
				diags = append(diags, e)
			}
			obj.HostkeyEcdsa521 = &v2
		}
	}
	if v1, ok := d.GetOk("hostkey_ed25519"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("hostkey_ed25519", sv)
				diags = append(diags, e)
			}
			obj.HostkeyEd25519 = &v2
		}
	}
	if v1, ok := d.GetOk("hostkey_rsa2048"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("hostkey_rsa2048", sv)
				diags = append(diags, e)
			}
			obj.HostkeyRsa2048 = &v2
		}
	}
	if v1, ok := d.GetOk("untrusted_caname"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("untrusted_caname", sv)
				diags = append(diags, e)
			}
			obj.UntrustedCaname = &v2
		}
	}
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectFirewallSshSetting(d *schema.ResourceData, sv string) (*models.FirewallSshSetting, diag.Diagnostics) {
	obj := models.FirewallSshSetting{}
	diags := diag.Diagnostics{}

	return &obj, diags
}
