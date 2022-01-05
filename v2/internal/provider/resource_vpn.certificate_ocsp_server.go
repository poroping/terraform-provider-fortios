// Unofficial Fortinet Terraform Provider
// Generated from templates using FortiOS v6.2.7,v6.4.0,v6.4.2,v6.4.3,v6.4.5,v6.4.6,v6.4.7,v6.4.8,v7.0.0,v7.0.1,v7.0.2,v7.0.3 schemas
// Maintainers:
// Justin Roberts (@poroping)

// Description: OCSP server configuration.

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

func resourceVpnCertificateOcspServer() *schema.Resource {
	return &schema.Resource{
		Description: "OCSP server configuration.",

		CreateContext: resourceVpnCertificateOcspServerCreate,
		ReadContext:   resourceVpnCertificateOcspServerRead,
		UpdateContext: resourceVpnCertificateOcspServerUpdate,
		DeleteContext: resourceVpnCertificateOcspServerDelete,

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
			"cert": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),

				Description: "OCSP server certificate.",
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),

				Description: "OCSP server entry name.",
				ForceNew:    true,
				Required:    true,
			},
			"secondary_cert": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),

				Description: "Secondary OCSP server certificate.",
				Optional:    true,
				Computed:    true,
			},
			"secondary_url": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),

				Description: "Secondary OCSP server URL.",
				Optional:    true,
				Computed:    true,
			},
			"source_ip": {
				Type:         schema.TypeString,
				ValidateFunc: validation.IsIPv4Address,

				Description: "Source IP address for communications to the OCSP server.",
				Optional:    true,
				Computed:    true,
			},
			"unavail_action": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"revoke", "ignore"}, false),

				Description: "Action when server is unavailable (revoke the certificate or ignore the result of the check).",
				Optional:    true,
				Computed:    true,
			},
			"url": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),

				Description: "OCSP server URL.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

func resourceVpnCertificateOcspServerCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("error creating VpnCertificateOcspServer resource: %q must be set if \"allow_append\" is true", key)
		}
	}

	obj, diags := getObjectVpnCertificateOcspServer(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.CreateVpnCertificateOcspServer(obj, urlparams)

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("VpnCertificateOcspServer")
	}

	return resourceVpnCertificateOcspServerRead(ctx, d, meta)
}

func resourceVpnCertificateOcspServerUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	obj, diags := getObjectVpnCertificateOcspServer(d, c.Config.Fv)
	if diags.HasError() {
		return diags
	}

	o, err := c.Cmdb.UpdateVpnCertificateOcspServer(mkey, obj, urlparams)
	if err != nil {
		return diag.Errorf("error updating VpnCertificateOcspServer resource: %v", err)
	}

	// log.Printf(strconv.Itoa(c.Retries))
	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("VpnCertificateOcspServer")
	}

	return resourceVpnCertificateOcspServerRead(ctx, d, meta)
}

func resourceVpnCertificateOcspServerDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	err := c.Cmdb.DeleteVpnCertificateOcspServer(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error deleting VpnCertificateOcspServer resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnCertificateOcspServerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	o, err := c.Cmdb.ReadVpnCertificateOcspServer(mkey, urlparams)
	if err != nil {
		return diag.Errorf("error reading VpnCertificateOcspServer resource: %v", err)
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

	diags := refreshObjectVpnCertificateOcspServer(d, o, c.Config.Fv, sort)
	if diags.HasError() {
		return diags
	}
	return nil
}

func refreshObjectVpnCertificateOcspServer(d *schema.ResourceData, o *models.VpnCertificateOcspServer, sv string, sort bool) diag.Diagnostics {
	var err error

	if o.Cert != nil {
		v := *o.Cert

		if err = d.Set("cert", v); err != nil {
			return diag.Errorf("error reading cert: %v", err)
		}
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.SecondaryCert != nil {
		v := *o.SecondaryCert

		if err = d.Set("secondary_cert", v); err != nil {
			return diag.Errorf("error reading secondary_cert: %v", err)
		}
	}

	if o.SecondaryUrl != nil {
		v := *o.SecondaryUrl

		if err = d.Set("secondary_url", v); err != nil {
			return diag.Errorf("error reading secondary_url: %v", err)
		}
	}

	if o.SourceIp != nil {
		v := *o.SourceIp

		if err = d.Set("source_ip", v); err != nil {
			return diag.Errorf("error reading source_ip: %v", err)
		}
	}

	if o.UnavailAction != nil {
		v := *o.UnavailAction

		if err = d.Set("unavail_action", v); err != nil {
			return diag.Errorf("error reading unavail_action: %v", err)
		}
	}

	if o.Url != nil {
		v := *o.Url

		if err = d.Set("url", v); err != nil {
			return diag.Errorf("error reading url: %v", err)
		}
	}

	return nil
}

func getObjectVpnCertificateOcspServer(d *schema.ResourceData, sv string) (*models.VpnCertificateOcspServer, diag.Diagnostics) {
	obj := models.VpnCertificateOcspServer{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("cert"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("cert", sv)
				diags = append(diags, e)
			}
			obj.Cert = &v2
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
	if v1, ok := d.GetOk("secondary_cert"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("secondary_cert", sv)
				diags = append(diags, e)
			}
			obj.SecondaryCert = &v2
		}
	}
	if v1, ok := d.GetOk("secondary_url"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("secondary_url", sv)
				diags = append(diags, e)
			}
			obj.SecondaryUrl = &v2
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
	if v1, ok := d.GetOk("unavail_action"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("unavail_action", sv)
				diags = append(diags, e)
			}
			obj.UnavailAction = &v2
		}
	}
	if v1, ok := d.GetOk("url"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("url", sv)
				diags = append(diags, e)
			}
			obj.Url = &v2
		}
	}
	return &obj, diags
}
