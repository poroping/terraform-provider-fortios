// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure debug URL addresses.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceWebProxyDebugUrl() *schema.Resource {
	return &schema.Resource{
		Create: resourceWebProxyDebugUrlCreate,
		Read:   resourceWebProxyDebugUrlRead,
		Update: resourceWebProxyDebugUrlUpdate,
		Delete: resourceWebProxyDebugUrlDelete,

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
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"url_pattern": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 511),
				Required:     true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"exact": {
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

func resourceWebProxyDebugUrlCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWebProxyDebugUrl(d, c.Fv)
	if err != nil {
		return fmt.Errorf("error creating WebProxyDebugUrl resource while getting object: %v", err)
	}

	o, err := c.CreateWebProxyDebugUrl(obj, vdomparam, urlparams, batchid)

	if err != nil {
		return fmt.Errorf("error creating WebProxyDebugUrl resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WebProxyDebugUrl")
	}

	return resourceWebProxyDebugUrlRead(d, m)
}

func resourceWebProxyDebugUrlUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWebProxyDebugUrl(d, c.Fv)
	if err != nil {
		return fmt.Errorf("error updating WebProxyDebugUrl resource while getting object: %v", err)
	}

	o, err := c.UpdateWebProxyDebugUrl(obj, mkey, vdomparam, urlparams, batchid)
	if err != nil {
		return fmt.Errorf("error updating WebProxyDebugUrl resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WebProxyDebugUrl")
	}

	return resourceWebProxyDebugUrlRead(d, m)
}

func resourceWebProxyDebugUrlDelete(d *schema.ResourceData, m interface{}) error {
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

	err := c.DeleteWebProxyDebugUrl(mkey, vdomparam, batchid)
	if err != nil {
		return fmt.Errorf("error deleting WebProxyDebugUrl resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWebProxyDebugUrlRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWebProxyDebugUrl(mkey, vdomparam, urlparams, batchid)
	if err != nil {
		return fmt.Errorf("error reading WebProxyDebugUrl resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWebProxyDebugUrl(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("error reading WebProxyDebugUrl resource from API: %v", err)
	}
	return nil
}

func flattenWebProxyDebugUrlName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyDebugUrlUrlPattern(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyDebugUrlStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyDebugUrlExact(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWebProxyDebugUrl(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenWebProxyDebugUrlName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("error reading name: %v", err)
		}
	}

	if err = d.Set("url_pattern", flattenWebProxyDebugUrlUrlPattern(o["url-pattern"], d, "url_pattern", sv)); err != nil {
		if !fortiAPIPatch(o["url-pattern"]) {
			return fmt.Errorf("error reading url_pattern: %v", err)
		}
	}

	if err = d.Set("status", flattenWebProxyDebugUrlStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("error reading status: %v", err)
		}
	}

	if err = d.Set("exact", flattenWebProxyDebugUrlExact(o["exact"], d, "exact", sv)); err != nil {
		if !fortiAPIPatch(o["exact"]) {
			return fmt.Errorf("error reading exact: %v", err)
		}
	}

	return nil
}

func flattenWebProxyDebugUrlFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandWebProxyDebugUrlName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyDebugUrlUrlPattern(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyDebugUrlStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyDebugUrlExact(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWebProxyDebugUrl(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandWebProxyDebugUrlName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("url_pattern"); ok {

		t, err := expandWebProxyDebugUrlUrlPattern(d, v, "url_pattern", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url-pattern"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {

		t, err := expandWebProxyDebugUrlStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("exact"); ok {

		t, err := expandWebProxyDebugUrlExact(d, v, "exact", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exact"] = t
		}
	}

	return &obj, nil
}
