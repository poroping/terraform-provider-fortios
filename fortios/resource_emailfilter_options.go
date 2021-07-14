// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure AntiSpam options.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceEmailfilterOptions() *schema.Resource {
	return &schema.Resource{
		Create: resourceEmailfilterOptionsUpdate,
		Read:   resourceEmailfilterOptionsRead,
		Update: resourceEmailfilterOptionsUpdate,
		Delete: resourceEmailfilterOptionsDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"dns_timeout": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 30),
				Optional:     true,
				Computed:     true,
			},
			"batchid": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
		},
	}
}

func resourceEmailfilterOptionsUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectEmailfilterOptions(d, c.Fv)
	if err != nil {
		return fmt.Errorf("error updating EmailfilterOptions resource while getting object: %v", err)
	}

	o, err := c.UpdateEmailfilterOptions(obj, mkey, vdomparam, urlparams, batchid)
	if err != nil {
		return fmt.Errorf("error updating EmailfilterOptions resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("EmailfilterOptions")
	}

	return resourceEmailfilterOptionsRead(d, m)
}

func resourceEmailfilterOptionsDelete(d *schema.ResourceData, m interface{}) error {
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

	err := c.DeleteEmailfilterOptions(mkey, vdomparam, batchid)
	if err != nil {
		return fmt.Errorf("error deleting EmailfilterOptions resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceEmailfilterOptionsRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadEmailfilterOptions(mkey, vdomparam, urlparams, batchid)
	if err != nil {
		return fmt.Errorf("error reading EmailfilterOptions resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectEmailfilterOptions(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("error reading EmailfilterOptions resource from API: %v", err)
	}
	return nil
}

func flattenEmailfilterOptionsDnsTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectEmailfilterOptions(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("dns_timeout", flattenEmailfilterOptionsDnsTimeout(o["dns-timeout"], d, "dns_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["dns-timeout"]) {
			return fmt.Errorf("error reading dns_timeout: %v", err)
		}
	}

	return nil
}

func flattenEmailfilterOptionsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandEmailfilterOptionsDnsTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectEmailfilterOptions(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("dns_timeout"); ok {

		t, err := expandEmailfilterOptionsDnsTimeout(d, v, "dns_timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-timeout"] = t
		}
	}

	return &obj, nil
}
