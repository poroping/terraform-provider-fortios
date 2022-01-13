// Author: Justin Roberts (@poroping)
// Documentation:
// Justin Roberts (@poroping),

// Description: Local certificate management.

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

func resourceCertificateManagementLocal() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCertificateManagementLocalCreate,
		ReadContext:   resourceCertificateManagementLocalRead,
		UpdateContext: resourceCertificateManagementLocalUpdate,
		DeleteContext: resourceCertificateManagementLocalDelete,

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
			"password": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				ForceNew:     true,
				Optional:     true,
				Sensitive:    true,
			},
			"comments": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 511),
				Optional:     true,
				Computed:     true,
			},
			"private_key": {
				Type:      schema.TypeString,
				ForceNew:  true,
				Optional:  true,
				Sensitive: true,
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
			"ike_localid": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"ike_localid_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},
			"batchid": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
		},
	}
}

func resourceCertificateManagementLocalCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	mkey := d.Get("name").(string)

	payload := &monitor.MonitorSystemVpnCertificateLocalImport{
		Name:        d.Get("name").(string),
		Type:        d.Get("type").(string),
		Certificate: d.Get("certificate").(string),
		Password:    d.Get("password").(string),
		PrivateKey:  d.Get("private_key").(string),
		Scope:       d.Get("scope").(string),
	}

	_, err = c.Monitor.MonitorSystemVpnCertificateLocalImport(payload, urlparams)

	if err != nil {
		return diag.FromErr(err)
	}

	o := &models.CmdbResponse{}
	scope := d.Get("scope").(string)
	if scope == "global" {
		obj, diags := getObjectCertificateManagementLocal(d, c.Config.Fv)
		if diags.HasError() {
			return diags
		}
		o, err = c.Cmdb.UpdateCertificateLocal(mkey, obj, urlparams)
	} else {
		obj, diags := getObjectVpnCertificateManagementLocal(d, c.Config.Fv)
		if diags.HasError() {
			return diags
		}
		o, err = c.Cmdb.UpdateVpnCertificateLocal(mkey, obj, urlparams)
	}

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("CertificateManagementLocal")
	}

	return resourceCertificateManagementLocalRead(ctx, d, meta)
}

func resourceCertificateManagementLocalUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
		obj, diags := getObjectCertificateManagementLocal(d, c.Config.Fv)
		if diags.HasError() {
			return diags
		}
		o, err = c.Cmdb.UpdateCertificateLocal(mkey, obj, urlparams)
		if err != nil {
			return diag.FromErr(err)
		}
	} else {
		obj, diags := getObjectVpnCertificateManagementLocal(d, c.Config.Fv)
		if diags.HasError() {
			return diags
		}
		o, err = c.Cmdb.UpdateVpnCertificateLocal(mkey, obj, urlparams)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	if o.Mkey != nil {
		d.SetId(utils.ParseMkey(o.Mkey))
	} else {
		d.SetId("CertificateManagementLocal")
	}

	return resourceCertificateManagementLocalRead(ctx, d, meta)
}

func resourceCertificateManagementLocalDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
		err = c.Cmdb.DeleteCertificateLocal(mkey, urlparams)
	} else {
		err = c.Cmdb.DeleteVpnCertificateLocal(mkey, urlparams)
	}

	if err != nil {
		e := diag.FromErr(err)
		return append(diags, e...)
	}

	d.SetId("")

	return nil
}

func resourceCertificateManagementLocalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
	urlparams.Type = "local-cer"
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
		o, err := c.Cmdb.ReadCertificateLocal(mkey, urlparams)
		if err != nil {
			e := diag.FromErr(err)
			return append(diags, e...)
		}
		if o == nil {
			log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
			d.SetId("")
			return nil
		}
		diags := refreshObjectCertificateManagementLocal(d, o, c.Config.Fv, *cert)
		if diags.HasError() {
			return diags
		}
	} else {
		o, err := c.Cmdb.ReadVpnCertificateLocal(mkey, urlparams)
		if err != nil {
			e := diag.FromErr(err)
			return append(diags, e...)
		}
		if o == nil {
			log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
			d.SetId("")
			return nil
		}
		diags := refreshObjectVpnCertificateManagementLocal(d, o, c.Config.Fv, *cert)
		if diags.HasError() {
			return diags
		}
	}

	return nil
}

func refreshObjectCertificateManagementLocal(d *schema.ResourceData, o *models.CertificateLocal, sv string, cert string) diag.Diagnostics {
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

	if o.Comments != nil {
		v := *o.Comments

		if err = d.Set("comments", v); err != nil {
			return diag.Errorf("error reading comments: %v", err)
		}
	}

	if o.Range != nil {
		v := *o.Range

		if err = d.Set("scope", v); err != nil {
			return diag.Errorf("error reading scope: %v", err)
		}
	}

	if o.IkeLocalid != nil {
		v := *o.IkeLocalid

		if err = d.Set("ike_localid", v); err != nil {
			return diag.Errorf("error reading ike_localid: %v", err)
		}
	}

	if o.IkeLocalidType != nil {
		v := *o.IkeLocalidType

		if err = d.Set("ike_localid_type", v); err != nil {
			return diag.Errorf("error reading ike_localid_type: %v", err)
		}
	}

	if cert != "" {
		if err = d.Set("certificate", cert); err != nil {
			return diag.Errorf("error reading remote: %v", err)
		}
	}

	return nil
}

func refreshObjectVpnCertificateManagementLocal(d *schema.ResourceData, o *models.VpnCertificateLocal, sv string, cert string) diag.Diagnostics {
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

	if o.Comments != nil {
		v := *o.Comments

		if err = d.Set("comments", v); err != nil {
			return diag.Errorf("error reading comments: %v", err)
		}
	}

	if o.Range != nil {
		v := *o.Range

		if err = d.Set("scope", v); err != nil {
			return diag.Errorf("error reading scope: %v", err)
		}
	}

	if o.IkeLocalid != nil {
		v := *o.IkeLocalid

		if err = d.Set("ike_localid", v); err != nil {
			return diag.Errorf("error reading ike_localid: %v", err)
		}
	}

	if o.IkeLocalidType != nil {
		v := *o.IkeLocalidType

		if err = d.Set("ike_localid_type", v); err != nil {
			return diag.Errorf("error reading ike_localid_type: %v", err)
		}
	}

	if cert != "" {
		if err = d.Set("certificate", cert); err != nil {
			return diag.Errorf("error reading remote: %v", err)
		}
	}

	return nil
}

func getObjectCertificateManagementLocal(d *schema.ResourceData, sv string) (*models.CertificateLocal, diag.Diagnostics) {
	obj := models.CertificateLocal{}
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

	if v1, ok := d.GetOk("comments"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("comments", sv)
				diags = append(diags, e)
			}
			obj.Comments = &v2
		}
	}

	if v1, ok := d.GetOk("ike_localid"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ike_localid", sv)
				diags = append(diags, e)
			}
			obj.IkeLocalid = &v2
		}
	}

	if v1, ok := d.GetOk("ike_localid_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ike_localid_type", sv)
				diags = append(diags, e)
			}
			obj.IkeLocalidType = &v2
		}
	}

	return &obj, diags
}

func getObjectVpnCertificateManagementLocal(d *schema.ResourceData, sv string) (*models.VpnCertificateLocal, diag.Diagnostics) {
	obj := models.VpnCertificateLocal{}
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

	if v1, ok := d.GetOk("comments"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("comments", sv)
				diags = append(diags, e)
			}
			obj.Comments = &v2
		}
	}

	if v1, ok := d.GetOk("ike_localid"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ike_localid", sv)
				diags = append(diags, e)
			}
			obj.IkeLocalid = &v2
		}
	}

	if v1, ok := d.GetOk("ike_localid_type"); ok {
		if v2, ok := v1.(string); ok {
			if !utils.CheckVer(sv, "", "") {
				e := utils.AttributeVersionWarning("ike_localid_type", sv)
				diags = append(diags, e)
			}
			obj.IkeLocalidType = &v2
		}
	}

	return &obj, diags
}
