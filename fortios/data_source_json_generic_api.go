package fortios

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	forticlient "github.com/poroping/forti-sdk-go/fortios/sdkcore"
)

func dataSourceJSONGenericAPI() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceJSONGenericAPIRead,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"path": {
				Type:     schema.TypeString,
				Required: true,
			},
			"specialparams": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"response": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceJSONGenericAPIRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	i := &forticlient.JSONJSONGenericAPI{
		Path:          d.Get("path").(string),
		Method:        "GET",
		Specialparams: d.Get("specialparams").(string),
		Json:          "",
	}

	res, err := c.CreateJSONGenericAPI(i, vdomparam, 0)
	if err != nil {
		return fmt.Errorf("error reading json generic api: %v", err)
	}

	d.SetId("DataSourceJsonGenericApi" + uuid.New().String())
	d.Set("response", res)

	return nil
}
