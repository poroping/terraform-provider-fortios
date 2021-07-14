// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure custom languages.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemCustomLanguage() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemCustomLanguageCreate,
		Read:   resourceSystemCustomLanguageRead,
		Update: resourceSystemCustomLanguageUpdate,
		Delete: resourceSystemCustomLanguageDelete,

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
				ForceNew:     true,
				Required:     true,
			},
			"filename": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Required:     true,
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
		},
	}
}

func resourceSystemCustomLanguageCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemCustomLanguage(d, c.Fv)
	if err != nil {
		return fmt.Errorf("error creating SystemCustomLanguage resource while getting object: %v", err)
	}

	o, err := c.CreateSystemCustomLanguage(obj, vdomparam, urlparams, batchid)

	if err != nil {
		return fmt.Errorf("error creating SystemCustomLanguage resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemCustomLanguage")
	}

	return resourceSystemCustomLanguageRead(d, m)
}

func resourceSystemCustomLanguageUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemCustomLanguage(d, c.Fv)
	if err != nil {
		return fmt.Errorf("error updating SystemCustomLanguage resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemCustomLanguage(obj, mkey, vdomparam, urlparams, batchid)
	if err != nil {
		return fmt.Errorf("error updating SystemCustomLanguage resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemCustomLanguage")
	}

	return resourceSystemCustomLanguageRead(d, m)
}

func resourceSystemCustomLanguageDelete(d *schema.ResourceData, m interface{}) error {
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

	err := c.DeleteSystemCustomLanguage(mkey, vdomparam, batchid)
	if err != nil {
		return fmt.Errorf("error deleting SystemCustomLanguage resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemCustomLanguageRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemCustomLanguage(mkey, vdomparam, urlparams, batchid)
	if err != nil {
		return fmt.Errorf("error reading SystemCustomLanguage resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemCustomLanguage(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("error reading SystemCustomLanguage resource from API: %v", err)
	}
	return nil
}

func flattenSystemCustomLanguageName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCustomLanguageFilename(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCustomLanguageComments(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemCustomLanguage(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenSystemCustomLanguageName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("error reading name: %v", err)
		}
	}

	if err = d.Set("filename", flattenSystemCustomLanguageFilename(o["filename"], d, "filename", sv)); err != nil {
		if !fortiAPIPatch(o["filename"]) {
			return fmt.Errorf("error reading filename: %v", err)
		}
	}

	if err = d.Set("comments", flattenSystemCustomLanguageComments(o["comments"], d, "comments", sv)); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("error reading comments: %v", err)
		}
	}

	return nil
}

func flattenSystemCustomLanguageFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemCustomLanguageName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCustomLanguageFilename(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCustomLanguageComments(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemCustomLanguage(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandSystemCustomLanguageName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("filename"); ok {

		t, err := expandSystemCustomLanguageFilename(d, v, "filename", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filename"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok {

		t, err := expandSystemCustomLanguageComments(d, v, "comments", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	return &obj, nil
}
