// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Remote certificate as a PEM file.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceVpnCertificateRemote() *schema.Resource {
	return &schema.Resource{
		Create: resourceVpnCertificateRemoteCreate,
		Read:   resourceVpnCertificateRemoteRead,
		Update: resourceVpnCertificateRemoteUpdate,
		Delete: resourceVpnCertificateRemoteDelete,

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
			"remote": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"range": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source": {
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

func resourceVpnCertificateRemoteCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectVpnCertificateRemote(d, c.Fv)
	if err != nil {
		return fmt.Errorf("error creating VpnCertificateRemote resource while getting object: %v", err)
	}

	o, err := c.CreateVpnCertificateRemote(obj, vdomparam, urlparams, batchid)

	if err != nil {
		return fmt.Errorf("error creating VpnCertificateRemote resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("VpnCertificateRemote")
	}

	return resourceVpnCertificateRemoteRead(d, m)
}

func resourceVpnCertificateRemoteUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectVpnCertificateRemote(d, c.Fv)
	if err != nil {
		return fmt.Errorf("error updating VpnCertificateRemote resource while getting object: %v", err)
	}

	o, err := c.UpdateVpnCertificateRemote(obj, mkey, vdomparam, urlparams, batchid)
	if err != nil {
		return fmt.Errorf("error updating VpnCertificateRemote resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("VpnCertificateRemote")
	}

	return resourceVpnCertificateRemoteRead(d, m)
}

func resourceVpnCertificateRemoteDelete(d *schema.ResourceData, m interface{}) error {
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

	err := c.DeleteVpnCertificateRemote(mkey, vdomparam, batchid)
	if err != nil {
		return fmt.Errorf("error deleting VpnCertificateRemote resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnCertificateRemoteRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadVpnCertificateRemote(mkey, vdomparam, urlparams, batchid)
	if err != nil {
		return fmt.Errorf("error reading VpnCertificateRemote resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVpnCertificateRemote(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("error reading VpnCertificateRemote resource from API: %v", err)
	}
	return nil
}

func flattenVpnCertificateRemoteName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateRemoteRemote(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateRemoteRange(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnCertificateRemoteSource(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectVpnCertificateRemote(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenVpnCertificateRemoteName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("error reading name: %v", err)
		}
	}

	if err = d.Set("remote", flattenVpnCertificateRemoteRemote(o["remote"], d, "remote", sv)); err != nil {
		if !fortiAPIPatch(o["remote"]) {
			return fmt.Errorf("error reading remote: %v", err)
		}
	}

	if err = d.Set("range", flattenVpnCertificateRemoteRange(o["range"], d, "range", sv)); err != nil {
		if !fortiAPIPatch(o["range"]) {
			return fmt.Errorf("error reading range: %v", err)
		}
	}

	if err = d.Set("source", flattenVpnCertificateRemoteSource(o["source"], d, "source", sv)); err != nil {
		if !fortiAPIPatch(o["source"]) {
			return fmt.Errorf("error reading source: %v", err)
		}
	}

	return nil
}

func flattenVpnCertificateRemoteFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandVpnCertificateRemoteName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateRemoteRemote(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateRemoteRange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateRemoteSource(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectVpnCertificateRemote(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandVpnCertificateRemoteName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("remote"); ok {

		t, err := expandVpnCertificateRemoteRemote(d, v, "remote", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote"] = t
		}
	}

	if v, ok := d.GetOk("range"); ok {

		t, err := expandVpnCertificateRemoteRange(d, v, "range", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["range"] = t
		}
	}

	if v, ok := d.GetOk("source"); ok {

		t, err := expandVpnCertificateRemoteSource(d, v, "source", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source"] = t
		}
	}

	return &obj, nil
}
