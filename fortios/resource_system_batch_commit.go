package fortios

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSystemBatchCommit() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemBatchCommitCreate,
		Read:   resourceSystemBatchCommitRead,
		Delete: resourceSystemBatchCommitDelete,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"batchid": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceSystemBatchCommitCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	batchid, _ := d.Get("batchid").(int)

	//Call process by sdk
	err := c.CommitBatch(vdomparam, batchid)
	if err != nil {
		return fmt.Errorf("error starting batch transaction: %s", err)
	}

	//Set index for d
	d.SetId("BatchCommit" + uuid.New().String())

	return nil
}

func resourceSystemBatchCommitRead(d *schema.ResourceData, m interface{}) error {
	// no API for this
	return nil
}

func resourceSystemBatchCommitUpdate(d *schema.ResourceData, m interface{}) error {
	// no API for this
	return nil
}

func resourceSystemBatchCommitDelete(d *schema.ResourceData, m interface{}) error {
	// no API for this
	return nil
}
