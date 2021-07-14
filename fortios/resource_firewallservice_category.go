// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure service categories.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceFirewallServiceCategory() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallServiceCategoryCreate,
		Read:   resourceFirewallServiceCategoryRead,
		Update: resourceFirewallServiceCategoryUpdate,
		Delete: resourceFirewallServiceCategoryDelete,

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
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"comment": {
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

func resourceFirewallServiceCategoryCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallServiceCategory(d, c.Fv)
	if err != nil {
		return fmt.Errorf("error creating FirewallServiceCategory resource while getting object: %v", err)
	}

	if mkey == "" && allow_append {
		return fmt.Errorf("error creating FirewallServiceCategory resource: %q must be set if \"allow_append\" is true", key)
	}

	o := make(map[string]interface{})
	if mkey != "" && allow_append {
		o, err = c.UpdateFirewallServiceCategory(obj, mkey, vdomparam, urlparams, batchid)
	} else {
		o, err = c.CreateFirewallServiceCategory(obj, vdomparam, urlparams, batchid)
	}

	if err != nil {
		return fmt.Errorf("error creating FirewallServiceCategory resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallServiceCategory")
	}

	return resourceFirewallServiceCategoryRead(d, m)
}

func resourceFirewallServiceCategoryUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallServiceCategory(d, c.Fv)
	if err != nil {
		return fmt.Errorf("error updating FirewallServiceCategory resource while getting object: %v", err)
	}

	o, err := c.UpdateFirewallServiceCategory(obj, mkey, vdomparam, urlparams, batchid)
	if err != nil {
		return fmt.Errorf("error updating FirewallServiceCategory resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallServiceCategory")
	}

	return resourceFirewallServiceCategoryRead(d, m)
}

func resourceFirewallServiceCategoryDelete(d *schema.ResourceData, m interface{}) error {
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

	err := c.DeleteFirewallServiceCategory(mkey, vdomparam, batchid)
	if err != nil {
		return fmt.Errorf("error deleting FirewallServiceCategory resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallServiceCategoryRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadFirewallServiceCategory(mkey, vdomparam, urlparams, batchid)
	if err != nil {
		return fmt.Errorf("error reading FirewallServiceCategory resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallServiceCategory(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("error reading FirewallServiceCategory resource from API: %v", err)
	}
	return nil
}

func flattenFirewallServiceCategoryName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallServiceCategoryComment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectFirewallServiceCategory(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenFirewallServiceCategoryName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("error reading name: %v", err)
		}
	}

	if err = d.Set("comment", flattenFirewallServiceCategoryComment(o["comment"], d, "comment", sv)); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("error reading comment: %v", err)
		}
	}

	return nil
}

func flattenFirewallServiceCategoryFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandFirewallServiceCategoryName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallServiceCategoryComment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallServiceCategory(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandFirewallServiceCategoryName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {

		t, err := expandFirewallServiceCategoryComment(d, v, "comment", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	return &obj, nil
}
