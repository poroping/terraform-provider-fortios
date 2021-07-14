package fortios

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceFirewallProxypolicySort() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallProxypolicySortCreateUpdate,
		Read:   resourceFirewallProxypolicySortRead,
		Update: resourceFirewallProxypolicySortCreateUpdate,
		Delete: schema.Noop,

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"sortby": {
				Type:     schema.TypeString,
				Required: true,
			},
			"sortdirection": {
				Type:     schema.TypeString,
				Required: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"force_recreate": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"batchid": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
		},
	}
}

func resourceFirewallProxypolicySortCreateUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client

	if c == nil {
		return fmt.Errorf("FortiOS connection did not initialize successfully!")
	}

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

	sortby := d.Get("sortby").(string)
	sortdirection := d.Get("sortdirection").(string)

	if sortby != "policyid" {
		return fmt.Errorf("Unsupported sort type: " + sortby)
	}

	if sortdirection != "ascending" && sortdirection != "descending" {
		return fmt.Errorf("Unsupported sort direction: " + sortdirection)
	}

	err := c.CreateUpdateFirewallProxypolicySort(sortby, sortdirection, vdomparam, batchid)
	if err != nil {
		return fmt.Errorf("error sorting FirewallProxypolicy: %s", err)
	}

	d.SetId(sortby + sortdirection)

	return resourceFirewallProxypolicySortRead(d, m)
}

func resourceFirewallProxypolicySortRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client

	if c == nil {
		return fmt.Errorf("FortiOS connection did not initialize successfully!")
	}

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

	sortby := d.Get("sortby").(string)
	sortdirection := d.Get("sortdirection").(string)

	if sortby != "policyid" {
		return fmt.Errorf("Unsupported sort type: " + sortby)
	}

	if sortdirection != "ascending" && sortdirection != "descending" {
		return fmt.Errorf("Unsupported sort direction: " + sortdirection)
	}

	sorted, err := c.ReadFirewallProxypolicySort(sortby, sortdirection, vdomparam, batchid)
	if err != nil {
		return fmt.Errorf("error reading FirewallProxypolicy sort status: %s %s", err, mkey)
	}

	if sorted == false {
		d.Set("status", "unsorted")
	} else {
		d.Set("status", "")
	}

	return nil
}
