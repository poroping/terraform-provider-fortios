// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure console.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemConsole() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemConsoleUpdate,
		Read:   resourceSystemConsoleRead,
		Update: resourceSystemConsoleUpdate,
		Delete: resourceSystemConsoleDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"mode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"baudrate": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"output": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"login": {
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

func resourceSystemConsoleUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemConsole(d, c.Fv)
	if err != nil {
		return fmt.Errorf("error updating SystemConsole resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemConsole(obj, mkey, vdomparam, urlparams, batchid)
	if err != nil {
		return fmt.Errorf("error updating SystemConsole resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemConsole")
	}

	return resourceSystemConsoleRead(d, m)
}

func resourceSystemConsoleDelete(d *schema.ResourceData, m interface{}) error {
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

	err := c.DeleteSystemConsole(mkey, vdomparam, batchid)
	if err != nil {
		return fmt.Errorf("error deleting SystemConsole resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemConsoleRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemConsole(mkey, vdomparam, urlparams, batchid)
	if err != nil {
		return fmt.Errorf("error reading SystemConsole resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemConsole(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("error reading SystemConsole resource from API: %v", err)
	}
	return nil
}

func flattenSystemConsoleMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemConsoleBaudrate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemConsoleOutput(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemConsoleLogin(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemConsole(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mode", flattenSystemConsoleMode(o["mode"], d, "mode", sv)); err != nil {
		if !fortiAPIPatch(o["mode"]) {
			return fmt.Errorf("error reading mode: %v", err)
		}
	}

	if err = d.Set("baudrate", flattenSystemConsoleBaudrate(o["baudrate"], d, "baudrate", sv)); err != nil {
		if !fortiAPIPatch(o["baudrate"]) {
			return fmt.Errorf("error reading baudrate: %v", err)
		}
	}

	if err = d.Set("output", flattenSystemConsoleOutput(o["output"], d, "output", sv)); err != nil {
		if !fortiAPIPatch(o["output"]) {
			return fmt.Errorf("error reading output: %v", err)
		}
	}

	if err = d.Set("login", flattenSystemConsoleLogin(o["login"], d, "login", sv)); err != nil {
		if !fortiAPIPatch(o["login"]) {
			return fmt.Errorf("error reading login: %v", err)
		}
	}

	return nil
}

func flattenSystemConsoleFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemConsoleMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemConsoleBaudrate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemConsoleOutput(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemConsoleLogin(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemConsole(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mode"); ok {

		t, err := expandSystemConsoleMode(d, v, "mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode"] = t
		}
	}

	if v, ok := d.GetOk("baudrate"); ok {

		t, err := expandSystemConsoleBaudrate(d, v, "baudrate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["baudrate"] = t
		}
	}

	if v, ok := d.GetOk("output"); ok {

		t, err := expandSystemConsoleOutput(d, v, "output", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["output"] = t
		}
	}

	if v, ok := d.GetOk("login"); ok {

		t, err := expandSystemConsoleLogin(d, v, "login", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["login"] = t
		}
	}

	return &obj, nil
}
