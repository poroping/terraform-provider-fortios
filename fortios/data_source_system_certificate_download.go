// Author: Justin Roberts (@poroping)
// Documentation:
// Justin Roberts (@poroping),

// Description: Download system certificates.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	forticlient "github.com/poroping/forti-sdk-go/fortios/sdkcore"
)

func dataSourceSystemCertificateDownload() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemCertificateDownloadRead,
		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"certificate": {
				Type:     schema.TypeString,
				Computed: true,
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
		},
	}
}

func dataSourceSystemCertificateDownloadRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	i := &forticlient.JSONSystemCertificateDownload{
		Mkey: d.Get("name").(string),
		Type: d.Get("type").(string),
	}

	o, err := c.ReadSystemCertificateDownload(i, vdomparam)
	if err != nil {
		return fmt.Errorf("error reading certificate from API: %v", err)
	}

	if o == "" {
		log.Printf("[WARN] certificate (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemCertificateDownload(d, o)
	if err != nil {
		return fmt.Errorf("error describing SystemCertificateDownload from API: %v", err)
	}

	d.SetId(i.Mkey)

	return nil
}

func dataSourceRefreshObjectSystemCertificateDownload(d *schema.ResourceData, o string) error {

	parsed_cert, err := parseDownloadedPemCertificate(o)

	if err != nil {
		return fmt.Errorf("%v", err)
	}

	var certificate_details []interface{}
	certificate_details = append(certificate_details, parsed_cert)

	if err = d.Set("certificate_details", certificate_details); err != nil {
		return fmt.Errorf("error reading certificate details: %v", err)
	}

	if err = d.Set("certificate", o); err != nil {
		return fmt.Errorf("error reading certificate: %v", err)
	}

	return nil
}

func dataSourceFlattenSystemCertificateDownloadFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}
