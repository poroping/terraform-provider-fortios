package fortios

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	forticlient "github.com/poroping/forti-sdk-go/fortios/sdkcore"
)

func resourceJSONGenericAPI() *schema.Resource {
	return &schema.Resource{
		Create: resourceJSONGenericAPICreateUpdate,
		Read:   resourceJSONGenericAPIRead,
		Update: resourceJSONGenericAPICreateUpdate,
		Delete: resourceJSONGenericAPIDelete,

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
			"method": {
				Type:     schema.TypeString,
				Required: true,
			},
			"specialparams": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"force_recreate": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"json": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"response": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"batchid": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
		},
	}
}

func resourceJSONGenericAPICreateUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	//Build input data by sdk
	i := &forticlient.JSONJSONGenericAPI{
		Path:          d.Get("path").(string),
		Method:        d.Get("method").(string),
		Specialparams: d.Get("specialparams").(string),
		Json:          d.Get("json").(string),
	}

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

	//Call process by sdk
	res, err := c.CreateJSONGenericAPI(i, vdomparam, batchid)
	if err != nil {
		return fmt.Errorf("error creating json generic api: %v", err)
	}

	//Set index for d
	d.SetId("JsonGenericApi" + uuid.New().String())
	d.Set("response", res)

	return nil
}

func resourceJSONGenericAPIDelete(d *schema.ResourceData, m interface{}) error {
	// no API for this
	return nil
}

func resourceJSONGenericAPIRead(d *schema.ResourceData, m interface{}) error {
	return nil
}
