// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure WAN optimization peers.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceWanoptPeer() *schema.Resource {
	return &schema.Resource{
		Create: resourceWanoptPeerCreate,
		Read:   resourceWanoptPeerRead,
		Update: resourceWanoptPeerUpdate,
		Delete: resourceWanoptPeerDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"peer_host_id": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"ip": {
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

func resourceWanoptPeerCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWanoptPeer(d, c.Fv)
	if err != nil {
		return fmt.Errorf("error creating WanoptPeer resource while getting object: %v", err)
	}

	o, err := c.CreateWanoptPeer(obj, vdomparam, urlparams, batchid)

	if err != nil {
		return fmt.Errorf("error creating WanoptPeer resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WanoptPeer")
	}

	return resourceWanoptPeerRead(d, m)
}

func resourceWanoptPeerUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWanoptPeer(d, c.Fv)
	if err != nil {
		return fmt.Errorf("error updating WanoptPeer resource while getting object: %v", err)
	}

	o, err := c.UpdateWanoptPeer(obj, mkey, vdomparam, urlparams, batchid)
	if err != nil {
		return fmt.Errorf("error updating WanoptPeer resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WanoptPeer")
	}

	return resourceWanoptPeerRead(d, m)
}

func resourceWanoptPeerDelete(d *schema.ResourceData, m interface{}) error {
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

	err := c.DeleteWanoptPeer(mkey, vdomparam, batchid)
	if err != nil {
		return fmt.Errorf("error deleting WanoptPeer resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWanoptPeerRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWanoptPeer(mkey, vdomparam, urlparams, batchid)
	if err != nil {
		return fmt.Errorf("error reading WanoptPeer resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWanoptPeer(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("error reading WanoptPeer resource from API: %v", err)
	}
	return nil
}

func flattenWanoptPeerPeerHostId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWanoptPeerIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWanoptPeer(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("peer_host_id", flattenWanoptPeerPeerHostId(o["peer-host-id"], d, "peer_host_id", sv)); err != nil {
		if !fortiAPIPatch(o["peer-host-id"]) {
			return fmt.Errorf("error reading peer_host_id: %v", err)
		}
	}

	if err = d.Set("ip", flattenWanoptPeerIp(o["ip"], d, "ip", sv)); err != nil {
		if !fortiAPIPatch(o["ip"]) {
			return fmt.Errorf("error reading ip: %v", err)
		}
	}

	return nil
}

func flattenWanoptPeerFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandWanoptPeerPeerHostId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWanoptPeerIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWanoptPeer(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("peer_host_id"); ok {

		t, err := expandWanoptPeerPeerHostId(d, v, "peer_host_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["peer-host-id"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok {

		t, err := expandWanoptPeerIp(d, v, "ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	return &obj, nil
}
