package fortios

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSystemBatchStart() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemBatchStartCreate,
		Read:   resourceSystemBatchStartRead,
		Update: resourceSystemBatchStartUpdate,
		Delete: resourceSystemBatchStartDelete,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  30,
			},
			"force_recreate": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"batchid": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemBatchStartCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	timeout, _ := d.Get("timeout").(int)

	//Call process by sdk
	batchid, err := c.StartBatch(timeout, vdomparam)
	if err != nil {
		return fmt.Errorf("error starting batch transaction: %s", err)
	}

	//Set index for d
	d.SetId("BatchTransaction" + uuid.New().String())
	d.Set("batchid", batchid)

	return nil
}

func resourceSystemBatchStartRead(d *schema.ResourceData, m interface{}) error {
	// no API for this
	return nil
}

func resourceSystemBatchStartUpdate(d *schema.ResourceData, m interface{}) error {
	// no API for this
	return nil
}

func resourceSystemBatchStartDelete(d *schema.ResourceData, m interface{}) error {
	// no API for this
	return nil
}
