// Author: Justin Roberts (@poroping)
// Documentation:
// Justin Roberts (@poroping),

// Description: Remote certificate management.

package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/poroping/forti-sdk-go/v2/models"
	"github.com/poroping/forti-sdk-go/v2/monitor"
	"github.com/poroping/terraform-provider-fortios/v2/suppressors"
	"github.com/poroping/terraform-provider-fortios/v2/utils"
)

func resourceCertificateManagementRemote() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCertificateManagementRemoteCreate,
		ReadContext:   resourceCertificateManagementRemoteRead,
		// UpdateContext: resourceCertificateManagementRemoteUpdate,
		DeleteContext: resourceCertificateManagementRemoteDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Required:     true,
			},
			"certificate": {
				Type:             schema.TypeString,
				ForceNew:         true,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: suppressors.DiffSuppCertificates,
				// parse and check equality cause PEM can be formatted differently
			},
			"certificate_details": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"signature_algorithm": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"public_key_algorithm": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"serial_number": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_ca": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_valid": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"version": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"issuer": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"subject": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"not_before": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"not_after": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"sha1_fingerprint": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"scope": {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},
		},
	}
}

func resourceCertificateManagementRemoteCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	payload := &monitor.MonitorSystemVpnCertificateRemoteImport{
		Certificate: d.Get("certificate").(string),
		Scope:       d.Get("scope").(string),
	}

	mkey, err := c.Monitor.MonitorSystemVpnCertificateRemoteImport(payload, urlparams)

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*mkey)

	o := &models.CmdbResponse{}
	scope := d.Get("scope").(string)
	if scope == "global" {
		obj, diags := getObjectCertificateManagementRemote(d, c.Config.Fv)
		if diags.HasError() {
			return diags
		}
		o, err = c.Cmdb.UpdateCertificateRemote(*mkey, obj, urlparams)
	} else {
		obj, diags := getObjectVpnCertificateManagementRemote(d, c.Config.Fv)
		if diags.HasError() {
			return diags
		}
		o, err = c.Cmdb.UpdateVpnCertificateRemote(*mkey, obj, urlparams)
	}

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("CertificateManagementRemote")
	}

	return resourceCertificateManagementRemoteRead(ctx, d, meta)
}

func resourceCertificateManagementRemoteUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	mkey := d.Id()
	c := meta.(*apiClient).Client
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

	o := &models.CmdbResponse{}
	scope := d.Get("scope").(string)
	if scope == "global" {
		obj, diags := getObjectCertificateManagementRemote(d, c.Config.Fv)
		if diags.HasError() {
			return diags
		}
		o, err = c.Cmdb.UpdateCertificateRemote(mkey, obj, urlparams)
		if err != nil {
			return diag.FromErr(err)
		}
	} else {
		obj, diags := getObjectVpnCertificateManagementRemote(d, c.Config.Fv)
		if diags.HasError() {
			return diags
		}
		o, err = c.Cmdb.UpdateVpnCertificateRemote(mkey, obj, urlparams)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("CertificateManagementRemote")
	}

	return resourceCertificateManagementRemoteRead(ctx, d, meta)
}

func resourceCertificateManagementRemoteDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	mkey := d.Id()
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

	scope := d.Get("scope").(string)
	if scope == "global" {
		err = c.Cmdb.DeleteCertificateRemote(mkey, urlparams)
	} else {
		err = c.Cmdb.DeleteVpnCertificateRemote(mkey, urlparams)
	}

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	d.SetId("")

	return nil
}

func resourceCertificateManagementRemoteRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	mkey := d.Id()
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

	urlparams.Mkey = mkey
	urlparams.Type = "remote-cer"
	// get certificate from monitor endpoint

	cert, err := c.Monitor.MonitorSystemCertificateDownload(urlparams)
	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if cert == nil || *cert == "" {
		log.Printf("[WARN] certificate (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	// get cert attributes from cmdb endpoint

	scope := d.Get("scope").(string)
	if scope == "global" {
		o, err := c.Cmdb.ReadCertificateRemote(mkey, urlparams)
		if err != nil {
			e := diag.FromErr(err)
			return append(diags, e...)
		}
		if o == nil {
			log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
			d.SetId("")
			return nil
		}
		diags := refreshObjectCertificateManagementRemote(d, o, c.Config.Fv, *cert)
		if diags.HasError() {
			return diags
		}
	} else {
		o, err := c.Cmdb.ReadVpnCertificateRemote(mkey, urlparams)
		if err != nil {
			e := diag.FromErr(err)
			return append(diags, e...)
		}
		if o == nil {
			log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
			d.SetId("")
			return nil
		}
		diags := refreshObjectVpnCertificateManagementRemote(d, o, c.Config.Fv, *cert)
		if diags.HasError() {
			return diags
		}
	}

	return nil
}

func refreshObjectCertificateManagementRemote(d *schema.ResourceData, o *models.CertificateRemote, sv string, cert string) diag.Diagnostics {
	parsed_cert, err := utils.ParseDownloadedPemCertificate(cert)

	if err != nil {
		return diag.Errorf("%v", err)
	}

	var certificate_details []interface{}
	certificate_details = append(certificate_details, parsed_cert)

	if err = d.Set("certificate_details", certificate_details); err != nil {
		return diag.Errorf("error reading certificate details: %v", err)
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.Range != nil {
		v := *o.Range

		if err = d.Set("scope", v); err != nil {
			return diag.Errorf("error reading scope: %v", err)
		}
	}

	if cert != "" {
		if err = d.Set("certificate", cert); err != nil {
			return diag.Errorf("error reading remote: %v", err)
		}
	}

	return nil
}

func refreshObjectVpnCertificateManagementRemote(d *schema.ResourceData, o *models.VpnCertificateRemote, sv string, cert string) diag.Diagnostics {
	parsed_cert, err := utils.ParseDownloadedPemCertificate(cert)

	if err != nil {
		return diag.Errorf("%v", err)
	}

	var certificate_details []interface{}
	certificate_details = append(certificate_details, parsed_cert)

	if err = d.Set("certificate_details", certificate_details); err != nil {
		return diag.Errorf("error reading certificate details: %v", err)
	}

	if o.Name != nil {
		v := *o.Name

		if err = d.Set("name", v); err != nil {
			return diag.Errorf("error reading name: %v", err)
		}
	}

	if o.Range != nil {
		v := *o.Range

		if err = d.Set("scope", v); err != nil {
			return diag.Errorf("error reading scope: %v", err)
		}
	}

	if cert != "" {
		if err = d.Set("certificate", cert); err != nil {
			return diag.Errorf("error reading remote: %v", err)
		}
	}

	return nil
}

func getObjectCertificateManagementRemote(d *schema.ResourceData, sv string) (*models.CertificateRemote, diag.Diagnostics) {
	obj := models.CertificateRemote{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("name"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("name", sv)
				diags = append(diags, e)
			}
			obj.Name = &v2
		}
	}
	return &obj, diags
}

func getObjectVpnCertificateManagementRemote(d *schema.ResourceData, sv string) (*models.VpnCertificateRemote, diag.Diagnostics) {
	obj := models.VpnCertificateRemote{}
	diags := diag.Diagnostics{}

	if v1, ok := d.GetOk("name"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("name", sv)
				diags = append(diags, e)
			}
			obj.Name = &v2
		}
	}
	return &obj, diags
}
