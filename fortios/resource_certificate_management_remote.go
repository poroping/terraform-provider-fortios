// Author: Justin Roberts (@poroping)
// Documentation:
// Justin Roberts (@poroping),

// Description: Remote certificate management.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	forticlient "github.com/poroping/forti-sdk-go/fortios/sdkcore"
)

func resourceCertificateManagementRemote() *schema.Resource {
	return &schema.Resource{
		Create: resourceCertificateManagementRemoteCreate,
		Read:   resourceCertificateManagementRemoteRead,
		Update: resourceCertificateManagementRemoteUpdate,
		Delete: resourceCertificateManagementRemoteDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
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
				DiffSuppressFunc: diffSuppCertificates,
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
			"batchid": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
		},
	}
}

func resourceCertificateManagementRemoteCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	batchid := 0

	if v, ok := d.GetOk("batchid"); ok {
		if i, ok := v.(int); ok {
			batchid = i
		}
	}

	i := &forticlient.JSONSystemVpnCertificateRemoteImport{
		Certificate: d.Get("certificate").(string),
		Scope:       d.Get("scope").(string),
	}

	mkey, err := c.CreateSystemVpnCertificateRemoteImport(i, vdomparam, batchid)

	if err != nil {
		return fmt.Errorf("error creating CertificateManagementRemoteImport resource: %v", err)
	}

	d.SetId(mkey)

	obj, err := getObjectCertificateManagementRemote(d, c.Fv)
	if err != nil {
		return fmt.Errorf("error creating CertificateManagementRemote resource while getting object: %v", err)
	}

	urlparams := make(map[string][]string)

	scope := d.Get("scope").(string)
	var o map[string]interface{}
	if scope == "global" {
		o, err = c.UpdateCertificateRemote(obj, mkey, vdomparam, urlparams, batchid)
	} else {
		o, err = c.UpdateVpnCertificateRemote(obj, mkey, vdomparam, urlparams, batchid)
	}

	if err != nil {
		return fmt.Errorf("error creating CertificateManagementRemote resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("CertificateManagementRemote")
	}

	return resourceCertificateManagementRemoteRead(d, m)
}

func resourceCertificateManagementRemoteUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	batchid := 0

	if v, ok := d.GetOk("batchid"); ok {
		if i, ok := v.(int); ok {
			batchid = i
		}
	}

	urlparams := make(map[string][]string)

	obj, err := getObjectCertificateManagementRemote(d, c.Fv)
	if err != nil {
		return fmt.Errorf("error updating CertificateManagementRemote resource while getting object: %v", err)
	}

	scope := d.Get("scope").(string)
	var o map[string]interface{}
	if scope == "global" {
		o, err = c.UpdateCertificateRemote(obj, mkey, vdomparam, urlparams, batchid)
	} else {
		o, err = c.UpdateVpnCertificateRemote(obj, mkey, vdomparam, urlparams, batchid)
	}

	if err != nil {
		return fmt.Errorf("error updating CertificateManagementRemote resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("CertificateManagementRemote")
	}

	return resourceCertificateManagementRemoteRead(d, m)
}

func resourceCertificateManagementRemoteDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	batchid := 0

	if v, ok := d.GetOk("batchid"); ok {
		if i, ok := v.(int); ok {
			batchid = i
		}
	}

	scope := d.Get("scope").(string)
	var err error
	if scope == "global" {
		err = c.DeleteCertificateRemote(mkey, vdomparam, batchid)
	} else {
		err = c.DeleteVpnCertificateRemote(mkey, vdomparam, batchid)
	}

	if err != nil {
		return fmt.Errorf("error deleting CertificateManagementRemote resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceCertificateManagementRemoteRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	batchid := 0

	if v, ok := d.GetOk("batchid"); ok {
		if i, ok := v.(int); ok {
			batchid = i
		}
	}

	urlparams := make(map[string][]string)

	// get certificate from monitor endpoint

	i := &forticlient.JSONSystemCertificateDownload{
		Mkey: mkey,
		Type: "remote-cer",
	}

	cert, err := c.ReadSystemCertificateDownload(i, vdomparam)
	if err != nil {
		cert = ""
		// return fmt.Errorf("error reading VpnCertificateRemoteImport resource: %v", err)
	}

	if cert == "" {
		log.Printf("[WARN] certificate (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	// get cert attributes from cmdb endpoint

	scope := d.Get("scope").(string)

	var o map[string]interface{}

	if scope == "global" {
		o, err = c.ReadCertificateRemote(mkey, vdomparam, urlparams, batchid)
	} else {
		o, err = c.ReadVpnCertificateRemote(mkey, vdomparam, urlparams, batchid)
	}

	if err != nil {
		return fmt.Errorf("error reading CertificateManagementRemote resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectCertificateManagementRemote(d, o, cert, c.Fv)
	if err != nil {
		return fmt.Errorf("error reading CertificateManagementRemote resource from API: %v", err)
	}
	return nil
}

func flattenCertificateManagementRemoteName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCertificateManagementRemoteCertificate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCertificateManagementRemoteScope(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectCertificateManagementRemote(d *schema.ResourceData, o map[string]interface{}, cert string, sv string) error {
	var err error

	parsed_cert, err := parseDownloadedPemCertificate(cert)

	if err != nil {
		return fmt.Errorf("%v", err)
	}

	var certificate_details []interface{}
	certificate_details = append(certificate_details, parsed_cert)

	if err = d.Set("certificate_details", certificate_details); err != nil {
		return fmt.Errorf("error reading certificate details: %v", err)
	}

	if err = d.Set("name", flattenCertificateManagementRemoteName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("error reading name: %v", err)
		}
	}

	if err = d.Set("certificate", cert); err != nil {
		if !fortiAPIPatch(cert) {
			return fmt.Errorf("error reading certificate: %v", err)
		}
	}

	if err = d.Set("scope", flattenCertificateManagementRemoteScope(o["range"], d, "scope", sv)); err != nil {
		if !fortiAPIPatch(o["range"]) {
			return fmt.Errorf("error reading scope: %v", err)
		}
	}

	return nil
}

func flattenCertificateManagementRemoteFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandCertificateManagementRemoteName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateManagementRemoteCertificate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateManagementRemoteScope(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectCertificateManagementRemote(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandCertificateManagementRemoteName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
