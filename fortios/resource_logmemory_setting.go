// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Settings for memory buffer.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceLogMemorySetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogMemorySettingUpdate,
		Read:   resourceLogMemorySettingRead,
		Update: resourceLogMemorySettingUpdate,
		Delete: resourceLogMemorySettingDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"diskfull": {
				Type:     schema.TypeString,
				Optional: true,
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

func resourceLogMemorySettingUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectLogMemorySetting(d, c.Fv)
	if err != nil {
		return fmt.Errorf("error updating LogMemorySetting resource while getting object: %v", err)
	}

	o, err := c.UpdateLogMemorySetting(obj, mkey, vdomparam, urlparams, batchid)
	if err != nil {
		return fmt.Errorf("error updating LogMemorySetting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("LogMemorySetting")
	}

	return resourceLogMemorySettingRead(d, m)
}

func resourceLogMemorySettingDelete(d *schema.ResourceData, m interface{}) error {
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

	err := c.DeleteLogMemorySetting(mkey, vdomparam, batchid)
	if err != nil {
		return fmt.Errorf("error deleting LogMemorySetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogMemorySettingRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadLogMemorySetting(mkey, vdomparam, urlparams, batchid)
	if err != nil {
		return fmt.Errorf("error reading LogMemorySetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLogMemorySetting(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("error reading LogMemorySetting resource from API: %v", err)
	}
	return nil
}

func flattenLogMemorySettingStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogMemorySettingDiskfull(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLogMemorySetting(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("status", flattenLogMemorySettingStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("error reading status: %v", err)
		}
	}

	if err = d.Set("diskfull", flattenLogMemorySettingDiskfull(o["diskfull"], d, "diskfull", sv)); err != nil {
		if !fortiAPIPatch(o["diskfull"]) {
			return fmt.Errorf("error reading diskfull: %v", err)
		}
	}

	return nil
}

func flattenLogMemorySettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandLogMemorySettingStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogMemorySettingDiskfull(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLogMemorySetting(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {

		t, err := expandLogMemorySettingStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("diskfull"); ok {

		t, err := expandLogMemorySettingDiskfull(d, v, "diskfull", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diskfull"] = t
		}
	}

	return &obj, nil
}
