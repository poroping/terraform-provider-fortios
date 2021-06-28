package fortios

import (
	"fmt"

	fmgclient "github.com/fortinetdev/forti-sdk-go/fortimanager/sdkcore"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceFortimanagerSystemLicenseFortiCare() *schema.Resource {
	return &schema.Resource{
		Create:             addFMGSystemLicenseFortiCare,
		Read:               readFMGSystemLicenseFortiCare,
		Update:             updateFMGSystemLicenseFortiCare,
		Delete:             deleteFMGSystemLicenseFortiCare,
		DeprecationMessage: "This resource has been depecated. Please use the dedicated FortiManager provider. This resource will be removed in v2.0.0",

		Schema: map[string]*schema.Schema{
			"target": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"adom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "root",
			},
			"registration_code": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func addFMGSystemLicenseFortiCare(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("addFMGSystemLicenseFortiCare")()

	i := &fmgclient.SystemLicenseFortiCare{
		Target:           d.Get("target").(string),
		Adom:             d.Get("adom").(string),
		RegistrationCode: d.Get("registration_code").(string),
	}

	err := c.AddSystemLicenseFortiCare(i)
	if err != nil {
		return fmt.Errorf("Error adding License for %s: %s", i.Target, err)
	}

	d.SetId("fortimanager-license-forticare")

	return nil
}

func updateFMGSystemLicenseFortiCare(d *schema.ResourceData, m interface{}) error {
	return nil
}

func readFMGSystemLicenseFortiCare(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteFMGSystemLicenseFortiCare(d *schema.ResourceData, m interface{}) error {
	return nil
}
