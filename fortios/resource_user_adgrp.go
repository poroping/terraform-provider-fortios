// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure FSSO groups.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceUserAdgrp() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserAdgrpCreate,
		Read:   resourceUserAdgrpRead,
		Update: resourceUserAdgrpUpdate,
		Delete: resourceUserAdgrpDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 511),
				ForceNew:     true,
				Required:     true,
			},
			"server_name": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"connector_source": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"fosid": {
				Type:     schema.TypeInt,
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

func resourceUserAdgrpCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectUserAdgrp(d, c.Fv)
	if err != nil {
		return fmt.Errorf("error creating UserAdgrp resource while getting object: %v", err)
	}

	o, err := c.CreateUserAdgrp(obj, vdomparam, urlparams, batchid)

	if err != nil {
		return fmt.Errorf("error creating UserAdgrp resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("UserAdgrp")
	}

	return resourceUserAdgrpRead(d, m)
}

func resourceUserAdgrpUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectUserAdgrp(d, c.Fv)
	if err != nil {
		return fmt.Errorf("error updating UserAdgrp resource while getting object: %v", err)
	}

	o, err := c.UpdateUserAdgrp(obj, mkey, vdomparam, urlparams, batchid)
	if err != nil {
		return fmt.Errorf("error updating UserAdgrp resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("UserAdgrp")
	}

	return resourceUserAdgrpRead(d, m)
}

func resourceUserAdgrpDelete(d *schema.ResourceData, m interface{}) error {
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

	err := c.DeleteUserAdgrp(mkey, vdomparam, batchid)
	if err != nil {
		return fmt.Errorf("error deleting UserAdgrp resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserAdgrpRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadUserAdgrp(mkey, vdomparam, urlparams, batchid)
	if err != nil {
		return fmt.Errorf("error reading UserAdgrp resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectUserAdgrp(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("error reading UserAdgrp resource from API: %v", err)
	}
	return nil
}

func flattenUserAdgrpName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserAdgrpServerName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserAdgrpConnectorSource(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserAdgrpId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectUserAdgrp(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenUserAdgrpName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("error reading name: %v", err)
		}
	}

	if err = d.Set("server_name", flattenUserAdgrpServerName(o["server-name"], d, "server_name", sv)); err != nil {
		if !fortiAPIPatch(o["server-name"]) {
			return fmt.Errorf("error reading server_name: %v", err)
		}
	}

	if err = d.Set("connector_source", flattenUserAdgrpConnectorSource(o["connector-source"], d, "connector_source", sv)); err != nil {
		if !fortiAPIPatch(o["connector-source"]) {
			return fmt.Errorf("error reading connector_source: %v", err)
		}
	}

	if err = d.Set("fosid", flattenUserAdgrpId(o["id"], d, "fosid", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("error reading fosid: %v", err)
		}
	}

	return nil
}

func flattenUserAdgrpFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandUserAdgrpName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserAdgrpServerName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserAdgrpConnectorSource(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserAdgrpId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectUserAdgrp(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandUserAdgrpName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("server_name"); ok {

		t, err := expandUserAdgrpServerName(d, v, "server_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-name"] = t
		}
	}

	if v, ok := d.GetOk("connector_source"); ok {

		t, err := expandUserAdgrpConnectorSource(d, v, "connector_source", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["connector-source"] = t
		}
	}

	if v, ok := d.GetOkExists("fosid"); ok {

		t, err := expandUserAdgrpId(d, v, "fosid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	return &obj, nil
}
