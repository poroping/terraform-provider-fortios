// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure access point status (rogue | accepted | suppressed).

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceWirelessControllerApStatus() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerApStatusCreate,
		Read:   resourceWirelessControllerApStatusRead,
		Update: resourceWirelessControllerApStatusUpdate,
		Delete: resourceWirelessControllerApStatusDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"fosid": {
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"bssid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssid": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 32),
				Optional:     true,
				Computed:     true,
			},
			"status": {
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

func resourceWirelessControllerApStatusCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWirelessControllerApStatus(d, c.Fv)
	if err != nil {
		return fmt.Errorf("error creating WirelessControllerApStatus resource while getting object: %v", err)
	}

	o, err := c.CreateWirelessControllerApStatus(obj, vdomparam, urlparams, batchid)

	if err != nil {
		return fmt.Errorf("error creating WirelessControllerApStatus resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("WirelessControllerApStatus")
	}

	return resourceWirelessControllerApStatusRead(d, m)
}

func resourceWirelessControllerApStatusUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWirelessControllerApStatus(d, c.Fv)
	if err != nil {
		return fmt.Errorf("error updating WirelessControllerApStatus resource while getting object: %v", err)
	}

	o, err := c.UpdateWirelessControllerApStatus(obj, mkey, vdomparam, urlparams, batchid)
	if err != nil {
		return fmt.Errorf("error updating WirelessControllerApStatus resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("WirelessControllerApStatus")
	}

	return resourceWirelessControllerApStatusRead(d, m)
}

func resourceWirelessControllerApStatusDelete(d *schema.ResourceData, m interface{}) error {
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

	err := c.DeleteWirelessControllerApStatus(mkey, vdomparam, batchid)
	if err != nil {
		return fmt.Errorf("error deleting WirelessControllerApStatus resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerApStatusRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWirelessControllerApStatus(mkey, vdomparam, urlparams, batchid)
	if err != nil {
		return fmt.Errorf("error reading WirelessControllerApStatus resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerApStatus(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("error reading WirelessControllerApStatus resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerApStatusId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerApStatusBssid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerApStatusSsid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerApStatusStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWirelessControllerApStatus(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("fosid", flattenWirelessControllerApStatusId(o["id"], d, "fosid", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("error reading fosid: %v", err)
		}
	}

	if err = d.Set("bssid", flattenWirelessControllerApStatusBssid(o["bssid"], d, "bssid", sv)); err != nil {
		if !fortiAPIPatch(o["bssid"]) {
			return fmt.Errorf("error reading bssid: %v", err)
		}
	}

	if err = d.Set("ssid", flattenWirelessControllerApStatusSsid(o["ssid"], d, "ssid", sv)); err != nil {
		if !fortiAPIPatch(o["ssid"]) {
			return fmt.Errorf("error reading ssid: %v", err)
		}
	}

	if err = d.Set("status", flattenWirelessControllerApStatusStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("error reading status: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerApStatusFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandWirelessControllerApStatusId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerApStatusBssid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerApStatusSsid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerApStatusStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerApStatus(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOkExists("fosid"); ok {

		t, err := expandWirelessControllerApStatusId(d, v, "fosid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("bssid"); ok {

		t, err := expandWirelessControllerApStatusBssid(d, v, "bssid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bssid"] = t
		}
	}

	if v, ok := d.GetOk("ssid"); ok {

		t, err := expandWirelessControllerApStatusSsid(d, v, "ssid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssid"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {

		t, err := expandWirelessControllerApStatusStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
