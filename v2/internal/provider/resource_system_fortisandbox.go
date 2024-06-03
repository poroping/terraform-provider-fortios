// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3,v7.0.4,v7.0.5,v7.0.6,v7.2.0,v7.2.1,v7.2.8 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: Configure FortiSandbox.

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

func resourceSystemFortisandbox() *schema.Resource {
	return &schema.Resource{
		Description: "Configure FortiSandbox.",

		CreateContext: resourceSystemFortisandboxCreate,
		ReadContext:   resourceSystemFortisandboxRead,
		UpdateContext: resourceSystemFortisandboxUpdate,
		DeleteContext: resourceSystemFortisandboxDelete,

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
			"email": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Notifier email address.",
				Optional:    true,
				Computed:    true,
			},
			"enc_algorithm": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"default", "high", "low"}, false),

				Description: "Configure the level of SSL protection for secure communication with FortiSandbox.",
				Optional:    true,
				Computed:    true,
			},
			"forticloud": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable FortiSandbox Cloud.",
				Optional:    true,
				Computed:    true,
			},
			"inline_scan": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable FortiSandbox inline scan.",
				Optional:    true,
				Computed:    true,
			},
			"interface": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),

				Description: "Specify outgoing interface to reach server.",
				Optional:    true,
				Computed:    true,
			},
			"interface_select_method": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"auto", "sdwan", "specify"}, false),

				Description: "Specify how to select outgoing interface to reach server.",
				Optional:    true,
				Computed:    true,
			},
			"server": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Server address of the remote FortiSandbox.",
				Optional:    true,
				Computed:    true,
			},
			"source_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),

				Description: "Source IP address for communications to FortiSandbox.",
				Optional:    true,
				Computed:    true,
			},
			"ssl_min_proto_version": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"default", "SSLv3", "TLSv1", "TLSv1-1", "TLSv1-2"}, false),

				Description: "Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).",
				Optional:    true,
				Computed:    true,
			},
			"status": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable FortiSandbox.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceSystemFortisandboxCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemFortisandbox(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateSystemFortisandbox(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemFortisandbox")
	}

	return resourceSystemFortisandboxRead(ctx, d, meta)
}

func resourceSystemFortisandboxUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectSystemFortisandbox(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateSystemFortisandbox(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating SystemFortisandbox resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("SystemFortisandbox")
	}

	return resourceSystemFortisandboxRead(ctx, d, meta)
}

func resourceSystemFortisandboxDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getEmptyObjectSystemFortisandbox(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	_, err := c.Cmdb.UpdateSystemFortisandbox(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error deleting SystemFortisandbox resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemFortisandboxRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadSystemFortisandbox(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading SystemFortisandbox resource: %v", err)
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

	diags := refreshObjectSystemFortisandbox(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectSystemFortisandbox(d *schema.ResourceData, o *models.SystemFortisandbox, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Email != nil {
		v := *o.Email

		if err = d.Set("email", v); err != nil {
			return diag.Errorf("error reading email: %v", err)
		}
	}

	if o.EncAlgorithm != nil {
		v := *o.EncAlgorithm

		if err = d.Set("enc_algorithm", v); err != nil {
			return diag.Errorf("error reading enc_algorithm: %v", err)
		}
	}

	if o.Forticloud != nil {
		v := *o.Forticloud

		if err = d.Set("forticloud", v); err != nil {
			return diag.Errorf("error reading forticloud: %v", err)
		}
	}

	if o.InlineScan != nil {
		v := *o.InlineScan

		if err = d.Set("inline_scan", v); err != nil {
			return diag.Errorf("error reading inline_scan: %v", err)
		}
	}

	if o.Interface != nil {
		v := *o.Interface

		if err = d.Set("interface", v); err != nil {
			return diag.Errorf("error reading interface: %v", err)
		}
	}

	if o.InterfaceSelectMethod != nil {
		v := *o.InterfaceSelectMethod

		if err = d.Set("interface_select_method", v); err != nil {
			return diag.Errorf("error reading interface_select_method: %v", err)
		}
	}

	if o.Server != nil {
		v := *o.Server

		if err = d.Set("server", v); err != nil {
			return diag.Errorf("error reading server: %v", err)
		}
	}

	if o.SourceIp != nil {
		v := *o.SourceIp

		if err = d.Set("source_ip", v); err != nil {
			return diag.Errorf("error reading source_ip: %v", err)
		}
	}

	if o.SslMinProtoVersion != nil {
		v := *o.SslMinProtoVersion

		if err = d.Set("ssl_min_proto_version", v); err != nil {
			return diag.Errorf("error reading ssl_min_proto_version: %v", err)
		}
	}

	if o.Status != nil {
		v := *o.Status

		if err = d.Set("status", v); err != nil {
			return diag.Errorf("error reading status: %v", err)
		}
	}

	return nil
}

func getObjectSystemFortisandbox(d *schema.ResourceData, sv string) (*models.SystemFortisandbox, diag.Diagnostics) {
	obj := models.SystemFortisandbox{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("email"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("email", sv)
				diags = append(diags, e)
			}
			obj.Email = &v2
		}
	}
	if v1, ok := d.GetOk("enc_algorithm"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("enc_algorithm", sv)
				diags = append(diags, e)
			}
			obj.EncAlgorithm = &v2
		}
	}
	if v1, ok := d.GetOk("forticloud"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.0", "") {
				e := utils.AttributeVersionWarning("forticloud", sv)
				diags = append(diags, e)
			}
			obj.Forticloud = &v2
		}
	}
	if v1, ok := d.GetOk("inline_scan"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.2.0", "") {
				e := utils.AttributeVersionWarning("inline_scan", sv)
				diags = append(diags, e)
			}
			obj.InlineScan = &v2
		}
	}
	if v1, ok := d.GetOk("interface"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.2", "") {
				e := utils.AttributeVersionWarning("interface", sv)
				diags = append(diags, e)
			}
			obj.Interface = &v2
		}
	}
	if v1, ok := d.GetOk("interface_select_method"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v6.4.2", "") {
				e := utils.AttributeVersionWarning("interface_select_method", sv)
				diags = append(diags, e)
			}
			obj.InterfaceSelectMethod = &v2
		}
	}
	if v1, ok := d.GetOk("server"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("server", sv)
				diags = append(diags, e)
			}
			obj.Server = &v2
		}
	}
	if v1, ok := d.GetOk("source_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("source_ip", sv)
				diags = append(diags, e)
			}
			obj.SourceIp = &v2
		}
	}
	if v1, ok := d.GetOk("ssl_min_proto_version"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssl_min_proto_version", sv)
				diags = append(diags, e)
			}
			obj.SslMinProtoVersion = &v2
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
	return &obj, diags
}

// Return an object with explicitly empty objects for tables that have been set.
func getEmptyObjectSystemFortisandbox(d *schema.ResourceData, sv string) (*models.SystemFortisandbox, diag.Diagnostics) {
	obj := models.SystemFortisandbox{}
	diags := diag.Diagnostics{}

	return &obj, diags
}
