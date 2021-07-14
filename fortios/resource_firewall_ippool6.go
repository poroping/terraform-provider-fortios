// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure IPv6 IP pools.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceFirewallIppool6() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallIppool6Create,
		Read:   resourceFirewallIppool6Read,
		Update: resourceFirewallIppool6Update,
		Delete: resourceFirewallIppool6Delete,

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
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"startip": {
				Type:     schema.TypeString,
				Required: true,
			},
			"endip": {
				Type:     schema.TypeString,
				Required: true,
			},
			"comments": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"batchid": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"allow_append": {
				Type:         schema.TypeBool,
				Optional:     true,
				Default:      false,
				RequiredWith: []string{"name"},
			},
		},
	}
}

func resourceFirewallIppool6Create(d *schema.ResourceData, m interface{}) error {
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

	allow_append := false

	if v, ok := d.GetOk("allow_append"); ok {
		if b, ok := v.(bool); ok {
			allow_append = b
		}
	}

	urlparams := make(map[string][]string)
	urlparams["allow_append"] = []string{strconv.FormatBool(allow_append)}

	key := "name"
	mkey := ""
	if v, ok := d.GetOk(key); ok {
		if s, ok := v.(string); ok {
			mkey = s
		}
	}

	obj, err := getObjectFirewallIppool6(d, c.Fv)
	if err != nil {
		return fmt.Errorf("error creating FirewallIppool6 resource while getting object: %v", err)
	}

	if mkey == "" && allow_append {
		return fmt.Errorf("error creating FirewallIppool6 resource: %q must be set if \"allow_append\" is true", key)
	}

	o := make(map[string]interface{})
	if mkey != "" && allow_append {
		o, err = c.UpdateFirewallIppool6(obj, mkey, vdomparam, urlparams, batchid)
	} else {
		o, err = c.CreateFirewallIppool6(obj, vdomparam, urlparams, batchid)
	}

	if err != nil {
		return fmt.Errorf("error creating FirewallIppool6 resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallIppool6")
	}

	return resourceFirewallIppool6Read(d, m)
}

func resourceFirewallIppool6Update(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallIppool6(d, c.Fv)
	if err != nil {
		return fmt.Errorf("error updating FirewallIppool6 resource while getting object: %v", err)
	}

	o, err := c.UpdateFirewallIppool6(obj, mkey, vdomparam, urlparams, batchid)
	if err != nil {
		return fmt.Errorf("error updating FirewallIppool6 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallIppool6")
	}

	return resourceFirewallIppool6Read(d, m)
}

func resourceFirewallIppool6Delete(d *schema.ResourceData, m interface{}) error {
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

	err := c.DeleteFirewallIppool6(mkey, vdomparam, batchid)
	if err != nil {
		return fmt.Errorf("error deleting FirewallIppool6 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallIppool6Read(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadFirewallIppool6(mkey, vdomparam, urlparams, batchid)
	if err != nil {
		return fmt.Errorf("error reading FirewallIppool6 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallIppool6(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("error reading FirewallIppool6 resource from API: %v", err)
	}
	return nil
}

func flattenFirewallIppool6Name(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallIppool6Startip(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallIppool6Endip(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallIppool6Comments(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectFirewallIppool6(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenFirewallIppool6Name(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("error reading name: %v", err)
		}
	}

	if err = d.Set("startip", flattenFirewallIppool6Startip(o["startip"], d, "startip", sv)); err != nil {
		if !fortiAPIPatch(o["startip"]) {
			return fmt.Errorf("error reading startip: %v", err)
		}
	}

	if err = d.Set("endip", flattenFirewallIppool6Endip(o["endip"], d, "endip", sv)); err != nil {
		if !fortiAPIPatch(o["endip"]) {
			return fmt.Errorf("error reading endip: %v", err)
		}
	}

	if err = d.Set("comments", flattenFirewallIppool6Comments(o["comments"], d, "comments", sv)); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("error reading comments: %v", err)
		}
	}

	return nil
}

func flattenFirewallIppool6FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandFirewallIppool6Name(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppool6Startip(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppool6Endip(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppool6Comments(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallIppool6(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandFirewallIppool6Name(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("startip"); ok {

		t, err := expandFirewallIppool6Startip(d, v, "startip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["startip"] = t
		}
	}

	if v, ok := d.GetOk("endip"); ok {

		t, err := expandFirewallIppool6Endip(d, v, "endip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["endip"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok {

		t, err := expandFirewallIppool6Comments(d, v, "comments", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	return &obj, nil
}
