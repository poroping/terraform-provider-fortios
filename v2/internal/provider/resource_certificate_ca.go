// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: CA certificate.

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

func resourceCertificateCa() *schema.Resource {
	return &schema.Resource{
		Description: "CA certificate.",

		CreateContext: resourceCertificateCaCreate,
		ReadContext:   resourceCertificateCaRead,
		UpdateContext: resourceCertificateCaUpdate,
		DeleteContext: resourceCertificateCaDelete,

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
			"auto_update_days": {
				Type: schema.TypeInt,

				Description: "Number of days to wait before requesting an updated CA certificate (0 - 4294967295, 0 = disabled).",
				Optional:    true,
				Computed:    true,
			},
			"auto_update_days_warning": {
				Type: schema.TypeInt,

				Description: "Number of days before an expiry-warning message is generated (0 - 4294967295, 0 = disabled).",
				Optional:    true,
				Computed:    true,
			},
			"ca": {
				Type: schema.TypeString,

				Description: "CA certificate as a PEM file.",
				Optional:    true,
				Computed:    true,
			},
			"ca_identifier": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "CA identifier of the SCEP server.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),

				Description: "Name.",
				ForceNew:    true,
				Required:    true,
			},
			"range": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"global", "vdom"}, false),

				Description: "Either global or VDOM IP address range for the CA certificate.",
				Optional:    true,
				Computed:    true,
			},
			"scep_url": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),

				Description: "URL of the SCEP server.",
				Optional:    true,
				Computed:    true,
			},
			"source": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"factory", "user", "bundle"}, false),

				Description: "CA certificate source type.",
				Optional:    true,
				Computed:    true,
			},
			"source_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "Source IP address for communications to the SCEP server.",
				Optional:    true,
				Computed:    true,
			},
			"ssl_inspection_trusted": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),

				Description: "Enable/disable this CA as a trusted CA for SSL inspection.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceCertificateCaCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating CertificateCa resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectCertificateCa(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateCertificateCa(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("CertificateCa")
	}

	return resourceCertificateCaRead(ctx, d, meta)
}

func resourceCertificateCaUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectCertificateCa(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateCertificateCa(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating CertificateCa resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("CertificateCa")
	}

	return resourceCertificateCaRead(ctx, d, meta)
}

func resourceCertificateCaDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteCertificateCa(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting CertificateCa resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceCertificateCaRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadCertificateCa(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading CertificateCa resource: %v", err)
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

	diags := refreshObjectCertificateCa(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectCertificateCa(d *schema.ResourceData, o *models.CertificateCa, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.AutoUpdateDays != nil {
		v := *o.AutoUpdateDays

		if err = d.Set("auto_update_days", v); err != nil {
			return diag.Errorf("error reading auto_update_days: %v", err)
		}
	}

	if o.AutoUpdateDaysWarning != nil {
		v := *o.AutoUpdateDaysWarning

		if err = d.Set("auto_update_days_warning", v); err != nil {
			return diag.Errorf("error reading auto_update_days_warning: %v", err)
		}
	}

	if o.Ca != nil {
		v := *o.Ca

		if err = d.Set("ca", v); err != nil {
			return diag.Errorf("error reading ca: %v", err)
		}
	}

	if o.CaIdentifier != nil {
		v := *o.CaIdentifier

		if err = d.Set("ca_identifier", v); err != nil {
			return diag.Errorf("error reading ca_identifier: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.Range != nil {
		v := *o.Range

		if err = d.Set("range", v); err != nil {
			return diag.Errorf("error reading range: %v", err)
		}
	}

	if o.ScepUrl != nil {
		v := *o.ScepUrl

		if err = d.Set("scep_url", v); err != nil {
			return diag.Errorf("error reading scep_url: %v", err)
		}
	}

	if o.Source != nil {
		v := *o.Source

		if err = d.Set("source", v); err != nil {
			return diag.Errorf("error reading source: %v", err)
		}
	}

	if o.SourceIp != nil {
		v := *o.SourceIp

		if err = d.Set("source_ip", v); err != nil {
			return diag.Errorf("error reading source_ip: %v", err)
		}
	}

	if o.SslInspectionTrusted != nil {
		v := *o.SslInspectionTrusted

		if err = d.Set("ssl_inspection_trusted", v); err != nil {
			return diag.Errorf("error reading ssl_inspection_trusted: %v", err)
		}
	}

	return nil
}

func getObjectCertificateCa(d *schema.ResourceData, sv string) (*models.CertificateCa, diag.Diagnostics) {
	obj := models.CertificateCa{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("auto_update_days"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auto_update_days", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.AutoUpdateDays = &tmp
		}
	}
	if v1, ok := d.GetOk("auto_update_days_warning"); ok {
		if v2, ok := v1.(int); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("auto_update_days_warning", sv)
				diags = append(diags, e)
			}
			tmp := int64(v2)
			obj.AutoUpdateDaysWarning = &tmp
		}
	}
	if v1, ok := d.GetOk("ca"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ca", sv)
				diags = append(diags, e)
			}
			obj.Ca = &v2
		}
	}
	if v1, ok := d.GetOk("ca_identifier"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "v7.0.2", "") {
				e := utils.AttributeVersionWarning("ca_identifier", sv)
				diags = append(diags, e)
			}
			obj.CaIdentifier = &v2
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
	if v1, ok := d.GetOk("range"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("range", sv)
				diags = append(diags, e)
			}
			obj.Range = &v2
		}
	}
	if v1, ok := d.GetOk("scep_url"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("scep_url", sv)
				diags = append(diags, e)
			}
			obj.ScepUrl = &v2
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
	if v1, ok := d.GetOk("source_ip"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("source_ip", sv)
				diags = append(diags, e)
			}
			obj.SourceIp = &v2
		}
	}
	if v1, ok := d.GetOk("ssl_inspection_trusted"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ssl_inspection_trusted", sv)
				diags = append(diags, e)
			}
			obj.SslInspectionTrusted = &v2
		}
	}
	return &obj, diags
}
